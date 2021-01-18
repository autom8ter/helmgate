package main

import (
	"context"
	"encoding/json"
	"fmt"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/autom8ter/kdeploy/internal/client"
	"github.com/autom8ter/kdeploy/internal/gql"
	"github.com/autom8ter/kdeploy/internal/helpers"
	"github.com/autom8ter/kdeploy/internal/logger"
	"github.com/autom8ter/kdeploy/internal/service"
	"github.com/autom8ter/kubego"
	"github.com/autom8ter/machine"
	"github.com/graphikDB/trigger"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	debug              bool
	listenPort         int64
	oidc               string
	allowedOrigins     []string
	allowedHeaders     []string
	allowedMethods     []string
	rootUsers          []string
	requestAuthorizers []string
	clientID           string
	clientSecret       string
	redirect           string
	outOfCluster       bool
)

func init() {
	godotenv.Load()
	pflag.CommandLine.BoolVar(&debug, "debug", helpers.BoolEnvOr("KDEPLOY_DEBUG", false), "enable debug logs (env: KDEPLOY_DEBUG)")
	pflag.CommandLine.Int64Var(&listenPort, "listen-port", int64(helpers.IntEnvOr("KDEPLOY_LISTEN_PORT", 8820)), "serve gRPC & graphQL on this port (env: KDEPLOY_LISTEN_PORT)")
	pflag.CommandLine.StringVar(&oidc, "open-id", helpers.EnvOr("KDEPLOY_OPEN_ID", ""), "open id connect discovery uri ex: https://accounts.google.com/.well-known/openid-configuration (env: KDEPLOY_OPEN_ID) (required)")
	pflag.CommandLine.StringSliceVar(&allowedHeaders, "allow-headers", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_HEADERS", []string{"*"}), "cors allow headers (env: KDEPLOY_ALLOW_HEADERS)")
	pflag.CommandLine.StringSliceVar(&allowedOrigins, "allow-origins", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_ORIGINS", []string{"*"}), "cors allow origins (env: KDEPLOY_ALLOW_ORIGINS)")
	pflag.CommandLine.StringSliceVar(&allowedMethods, "allow-methods", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_METHODS", []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"}), "cors allow methods (env: KDEPLOY_ALLOW_METHODS)")
	pflag.CommandLine.StringSliceVar(&rootUsers, "root-users", helpers.StringSliceEnvOr("KDEPLOY_ROOT_USERS", nil), "root users that bypass request authorizers (env: KDEPLOY_ROOT_USERS)")
	pflag.CommandLine.StringSliceVar(&requestAuthorizers, "request-authorizers", helpers.StringSliceEnvOr("KDEPLOY_REQUEST_AUTHORIZERS", nil), "request authorizer expressions (env: KDEPLOY_REQUEST_AUTHORIZERS)")
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
	var (
		metricsLis net.Listener
		m          = machine.New(ctx)
		err        error
		interrupt  = make(chan os.Signal, 1)
		apiLis     net.Listener
	)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	{
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", listenPort))
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
		apiLis, err = net.ListenTCP("tcp", addr)
		if err != nil {
			lgger.Error("failed to create api server listener", zap.Error(err))
			return
		}
	}
	defer apiLis.Close()
	apiMux := cmux.New(apiLis)
	apiMux.SetReadTimeout(1 * time.Second)
	grpcMatcher := apiMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	defer grpcMatcher.Close()
	gqlMatchermatcher := apiMux.Match(cmux.Any())
	defer gqlMatchermatcher.Close()
	m.Go(func(routine machine.Routine) {
		if err := apiMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
	})
	{
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", listenPort+1))
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
		metricsLis, err = net.ListenTCP("tcp", addr)
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
	}
	defer metricsLis.Close()
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
	lgger.Info("starting metrics server", zap.String("address", metricsLis.Addr().String()))
	m.Go(func(routine machine.Routine) {
		if err := metricServer.Serve(metricsLis); err != nil && err != http.ErrServerClosed {
			lgger.Error("metrics server failure", zap.Error(err))
		}
	})
	resp, err := http.DefaultClient.Get(oidc)
	if err != nil {
		lgger.Error("failed to get oidc", zap.Error(err))
		return
	}

	var openID = map[string]interface{}{}
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgger.Error("failed to get oidc", zap.Error(err))
		return
	}
	resp.Body.Close()
	if err := json.Unmarshal(bits, &openID); err != nil {
		lgger.Error("failed to get oidc", zap.Error(err))
		return
	}
	var authorizers []*trigger.Decision
	for _, a := range requestAuthorizers {
		decision, err := trigger.NewDecision(a)
		if err != nil {
			lgger.Error("failed to create authorizer", zap.Error(err))
			return
		}
		authorizers = append(authorizers, decision)
	}
	var config *oauth2.Config
	if clientID != "" {
		lgger.Debug("kdeploy graphQL playground enabled")
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

	var cli *client.Manager
	if outOfCluster {
		kclient, err := kubego.NewOutOfClusterKubeClient()
		if err != nil {
			lgger.Error(err.Error())
			return
		}
		iclient, err := kubego.NewOutOfClusterIstioClient()
		if err != nil {
			lgger.Error(err.Error())
			return
		}
		cli = client.New(kclient, iclient, lgger, rootUsers, openID["userinfo_endpoint"].(string), authorizers)
	} else {
		kclient, err := kubego.NewInClusterKubeClient()
		if err != nil {
			lgger.Error(err.Error())
			return
		}
		iclient, err := kubego.NewInClusterIstioClient()
		if err != nil {
			lgger.Error(err.Error())
			return
		}
		cli = client.New(kclient, iclient, lgger, rootUsers, openID["userinfo_endpoint"].(string), authorizers)
	}

	gopts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(lgger.Zap()),
			grpc_validator.UnaryServerInterceptor(),
			cli.UnaryInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(lgger.Zap()),
			grpc_validator.StreamServerInterceptor(),
			cli.StreamInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		),
	}
	gserver := grpc.NewServer(gopts...)
	kdeploypb.RegisterKdeployServiceServer(gserver, service.NewKdeployService(cli))
	reflection.Register(gserver)
	grpc_prometheus.Register(gserver)
	m.Go(func(routine machine.Routine) {
		lgger.Info("starting grpc server",
			zap.String("address", grpcMatcher.Addr().String()),
		)
		if err := gserver.Serve(grpcMatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("grpc server failure", zap.Error(err))
		}
	})
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("localhost:%v", listenPort),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		lgger.Error("failed to setup graphql server", zap.Error(err))
		return
	}
	defer conn.Close()
	resolver := gql.NewResolver(kdeploypb.NewKdeployServiceClient(conn), cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
	}), config, lgger, openID["userinfo_endpoint"].(string))

	mux := http.NewServeMux()

	mux.Handle("/graphql", resolver.QueryHandler())
	if config != nil {
		mux.Handle("/", resolver.Playground())
		mux.Handle("/oauth2/callback", resolver.PlaygroundCallback("/"))
	}

	graphQLServer := &http.Server{
		Handler: mux,
	}

	m.Go(func(routine machine.Routine) {
		lgger.Info("starting graphQL server", zap.String("address", gqlMatchermatcher.Addr().String()))
		if err := graphQLServer.Serve(gqlMatchermatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("http server failure", zap.Error(err))
		}
	})
	select {
	case <-interrupt:
		m.Cancel()
		break
	case <-ctx.Done():
		m.Cancel()
		break
	}
	lgger.Debug("shutdown signal received")
	go func() {
		shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer shutdowncancel()
		if err := graphQLServer.Shutdown(shutdownctx); err != nil {
			lgger.Error("graphQL server shutdown failure", zap.Error(err))
		} else {
			lgger.Debug("shutdown graphQL server")
		}
	}()
	go func() {
		shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer shutdowncancel()
		if err := metricServer.Shutdown(shutdownctx); err != nil {
			lgger.Error("metrics server shutdown failure", zap.Error(err))
		} else {
			lgger.Debug("shutdown metrics server")
		}

	}()
	go func() {
		stopped := make(chan struct{})
		go func() {
			gserver.GracefulStop()
			stopped <- struct{}{}
		}()
		select {
		case <-time.After(15 * time.Second):
			gserver.Stop()
		case <-stopped:
			close(stopped)
			break
		}
		lgger.Debug("shutdown gRPC server")
	}()
	m.Wait()
	lgger.Debug("shutdown successful")
}
