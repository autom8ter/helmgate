package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/kubego"
	"github.com/autom8ter/machine"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/autom8ter/meshpaas/internal/core"
	"github.com/autom8ter/meshpaas/internal/gql"
	"github.com/autom8ter/meshpaas/internal/helpers"
	"github.com/autom8ter/meshpaas/internal/logger"
	"github.com/autom8ter/meshpaas/internal/service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	jwksUri        string
	jwtIssuer      string
	outOfCluster   bool
	namespaceClaim string
)

func init() {
	godotenv.Load()
	pflag.CommandLine.BoolVar(&debug, "debug", helpers.BoolEnvOr("MESHPAAS_DEBUG", false), "enable debug logs (env: MESHPAAS_DEBUG)")
	pflag.CommandLine.Int64Var(&listenPort, "listen-port", int64(helpers.IntEnvOr("MESHPAAS_LISTEN_PORT", 8820)), "serve gRPC & graphQL on this port (env: MESHPAAS_LISTEN_PORT)")
	pflag.CommandLine.BoolVar(&outOfCluster, "out-of-cluster", helpers.BoolEnvOr("MESHPAAS_OUT_OF_CLUSTER", false), "enable out of cluster k8s config discovery (env: MESHPAAS_OUT_OF_CLUSTER)")
	pflag.CommandLine.StringVar(&jwksUri, "jwks-uri", helpers.EnvOr("MESHPAAS_JWKS_URI", ""), "remote json web key set uri for verifying authorization tokens (env: MESHPAAS_JWKS_URI)")
	pflag.CommandLine.StringVar(&jwtIssuer, "allow-jwt-issuer", helpers.EnvOr("MESHPAAS_ALLOW_ISSUER", ""), "allowed jwt.claim.iss issuer (env: MESHPAAS_ALLOW_ISSUER)")
	pflag.CommandLine.StringVar(&namespaceClaim, "namespace-claim", helpers.EnvOr("MESHPAAS_NAMESPACE_CLAIM", "aud"), "the jwt attribute on the id token that returns the namespace the user is allowed access to (required) (env: MESHPAAS_NAMESPACE_CLAIM)")
	pflag.Parse()
}

func main() {
	run(context.Background())
}

func run(ctx context.Context) {
	var lgger = logger.New(debug)
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
	var cli *core.Manager
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
		cli = core.NewManager(kclient, iclient, lgger, namespaceClaim)
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
		cli = core.NewManager(kclient, iclient, lgger, namespaceClaim)
	}
	a, err := auth.NewAuth(jwksUri, jwtIssuer, namespaceClaim, lgger)
	if err != nil {
		lgger.Error(err.Error())
		return
	}
	gopts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(lgger.Zap()),
			grpc_auth.UnaryServerInterceptor(a.Interceptor()),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(lgger.Zap()),
			grpc_auth.StreamServerInterceptor(a.Interceptor()),
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		),
	}
	gserver := grpc.NewServer(gopts...)
	meshpaaspb.RegisterMeshPaasServiceServer(gserver, service.NewMeshPaasService(cli))
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
	resolver := gql.NewResolver(meshpaaspb.NewMeshPaasServiceClient(conn), lgger)

	mux := http.NewServeMux()

	mux.Handle("/graphql", resolver.QueryHandler())
	mux.Handle("/", resolver.Playground())
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
