package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/kdeploy/gql"
	"github.com/autom8ter/kdeploy/helpers"
	"github.com/autom8ter/kdeploy/logger"
	"github.com/autom8ter/machine"
	"github.com/graphikDB/kubego"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	debug          bool
	listenPort     int64
	oidc           string
	allowedOrigins []string
	allowedHeaders []string
	allowedMethods []string
	clientID       string
	clientSecret   string
	redirect       string
	outOfCluster   bool
)

func init() {
	godotenv.Load()
	pflag.CommandLine.BoolVar(&debug, "debug", helpers.BoolEnvOr("KDEPLOY_DEBUG", false), "enable debug logs (env: KDEPLOY_DEBUG)")
	pflag.CommandLine.Int64Var(&listenPort, "listen-port", int64(helpers.IntEnvOr("KDEPLOY_LISTEN_PORT", 8820)), "serve gRPC & graphQL on this port (env: KDEPLOY_LISTEN_PORT)")
	pflag.CommandLine.StringVar(&oidc, "open-id", helpers.EnvOr("KDEPLOY_OPEN_ID", ""), "open id connect discovery uri ex: https://accounts.google.com/.well-known/openid-configuration (env: KDEPLOY_OPEN_ID) (required)")
	pflag.CommandLine.StringSliceVar(&allowedHeaders, "allow-headers", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_HEADERS", []string{"*"}), "cors allow headers (env: KDEPLOY_ALLOW_HEADERS)")
	pflag.CommandLine.StringSliceVar(&allowedOrigins, "allow-origins", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_ORIGINS", []string{"*"}), "cors allow origins (env: KDEPLOY_ALLOW_ORIGINS)")
	pflag.CommandLine.StringSliceVar(&allowedMethods, "allow-methods", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_METHODS", []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"}), "cors allow methods (env: KDEPLOY_ALLOW_METHODS)")
	pflag.CommandLine.StringVar(&clientID, "oauth-client-id", helpers.EnvOr("KDEPLOY_OAUTH_CLIENT_ID", ""), "playground oauth client id (env: KDEPLOY_OAUTH_CLIENT_ID) (required for graphQL playground)")
	pflag.CommandLine.StringVar(&clientSecret, "oauth-client-secret", helpers.EnvOr("KDEPLOY_OAUTH_CLIENT_SECRET", ""), "playground oauth client secret (env: KDEPLOY_OAUTH_CLIENT_SECRET) (required for graphQL playground)")
	pflag.CommandLine.StringVar(&redirect, "oauth-redirect", helpers.EnvOr("KDEPLOY_OAUTH_REDIRECT", ""), "playground oauth redirect (env: KDEPLOY_OAUTH_REDIRECT) (required for graphQL playground)")
	pflag.CommandLine.BoolVar(&outOfCluster, "out-of-cluster", helpers.BoolEnvOr("KDEPLOY_OUT_OF_CLUSTER", false), "enable out of cluster k8s config discovery (env: KDEPLOY_OUT_OF_CLUSTER)")
	pflag.Parse()
}

func main() {
	run(context.Background())
}

func run(ctx context.Context) {
	var lgger = logger.New(debug)
	if oidc == "" {
		lgger.Error("empty open-id connect discovery --open-id")
		return
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var (
		adminLis  net.Listener
		m         = machine.New(ctx)
		err       error
		interrupt = make(chan os.Signal, 1)
		apiLis    net.Listener
	)
	defer m.Close()
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	{
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", listenPort+1))
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
		adminLis, err = net.ListenTCP("tcp", addr)
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
	}
	defer adminLis.Close()
	adminMux := cmux.New(adminLis)
	var metricServer *http.Server
	{
		router := http.NewServeMux()
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		router.Handle("/metrics", promhttp.Handler())
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
		metricServer = &http.Server{Handler: router}
	}

	m.Go(func(routine machine.Routine) {
		hmatcher := adminMux.Match(cmux.HTTP1(), cmux.HTTP1Fast())
		defer hmatcher.Close()
		lgger.Info("starting metrics/admin server", zap.String("address", hmatcher.Addr().String()))
		if err := metricServer.Serve(hmatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("metrics server failure", zap.Error(err))
		}
	})
	m.Go(func(routine machine.Routine) {
		if err := adminMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
	})
	var config *oauth2.Config
	if clientID != "" {
		lgger.Debug("kdeploy graphQL playground enabled")
		resp, err := http.DefaultClient.Get(oidc)
		if err != nil {
			lgger.Error("failed to get oidc", zap.Error(err))
			return
		}
		defer resp.Body.Close()
		var openID = map[string]interface{}{}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			lgger.Error("failed to get oidc", zap.Error(err))
			return
		}
		if err := json.Unmarshal(bits, &openID); err != nil {
			lgger.Error("failed to get oidc", zap.Error(err))
			return
		}
		config = &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  openID["authorization_endpoint"].(string),
				TokenURL: openID["token_endpoint"].(string),
			},
			RedirectURL: redirect,
			Scopes:      []string{"openid", "email", "profile"},
		}
	}
	var client *kubego.Client
	if outOfCluster {
		client, err = kubego.NewOutOfClusterClient()
		if err != nil {
			lgger.Error("failed to create out of cluster k8s client", zap.Error(err))
			return
		}
	} else {
		client, err = kubego.NewInClusterClient()
		if err != nil {
			lgger.Error("failed to create in cluster k8s client", zap.Error(err))
			return
		}
	}

	resp, err := http.DefaultClient.Get(oidc)
	if err != nil {
		lgger.Error("failed to get oidc", zap.Error(err))
		return
	}
	defer resp.Body.Close()
	var openID = map[string]interface{}{}
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgger.Error("failed to get oidc", zap.Error(err))
		return
	}
	if err := json.Unmarshal(bits, &openID); err != nil {
		lgger.Error("failed to get oidc", zap.Error(err))
		return
	}
	resolver := gql.NewResolver(client, cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
	}), config, lgger, openID["userinfo_endpoint"].(string))
	mux := http.NewServeMux()
	mux.Handle("/", resolver.QueryHandler())
	if config != nil {
		mux.Handle("/playground", resolver.Playground())
		mux.Handle("/playground/callback", resolver.PlaygroundCallback("/playground"))
	}

	httpServer := &http.Server{
		Handler: mux,
	}

	apiLis, err = net.Listen("tcp", fmt.Sprintf(":%v", listenPort))
	if err != nil {
		lgger.Error("failed to create api server listener", zap.Error(err))
		return
	}
	defer apiLis.Close()
	apiMux := cmux.New(apiLis)
	m.Go(func(routine machine.Routine) {
		hmatcher := apiMux.Match(cmux.HTTP1())
		defer hmatcher.Close()
		lgger.Info("starting http server",
			zap.String("address", hmatcher.Addr().String()),
		)
		if err := httpServer.Serve(hmatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("http server failure", zap.Error(err))
		}
	})
	m.Go(func(routine machine.Routine) {
		if err := apiMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
	})
	select {
	case <-interrupt:
		m.Close()
		break
	case <-ctx.Done():
		m.Close()
		break
	}
	lgger.Warn("shutdown signal received")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	_ = httpServer.Shutdown(shutdownCtx)
	_ = metricServer.Shutdown(shutdownCtx)
	m.Wait()
	lgger.Debug("shutdown successful")
}
