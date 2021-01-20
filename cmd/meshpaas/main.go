package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/kubego"
	"github.com/autom8ter/machine"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/client"
	"github.com/autom8ter/meshpaas/internal/gql"
	"github.com/autom8ter/meshpaas/internal/helpers"
	"github.com/autom8ter/meshpaas/internal/logger"
	"github.com/autom8ter/meshpaas/internal/service"
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
	allowedOrigins []string
	allowedHeaders []string
	allowedMethods []string
	outOfCluster   bool
)

func init() {
	godotenv.Load()
	pflag.CommandLine.BoolVar(&debug, "debug", helpers.BoolEnvOr("KDEPLOY_DEBUG", false), "enable debug logs (env: KDEPLOY_DEBUG)")
	pflag.CommandLine.Int64Var(&listenPort, "listen-port", int64(helpers.IntEnvOr("KDEPLOY_LISTEN_PORT", 8820)), "serve gRPC & graphQL on this port (env: KDEPLOY_LISTEN_PORT)")
	pflag.CommandLine.StringSliceVar(&allowedHeaders, "allow-headers", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_HEADERS", []string{"*"}), "cors allow headers (env: KDEPLOY_ALLOW_HEADERS)")
	pflag.CommandLine.StringSliceVar(&allowedOrigins, "allow-origins", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_ORIGINS", []string{"*"}), "cors allow origins (env: KDEPLOY_ALLOW_ORIGINS)")
	pflag.CommandLine.StringSliceVar(&allowedMethods, "allow-methods", helpers.StringSliceEnvOr("KDEPLOY_ALLOW_METHODS", []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"}), "cors allow methods (env: KDEPLOY_ALLOW_METHODS)")
	pflag.CommandLine.BoolVar(&outOfCluster, "out-of-cluster", helpers.BoolEnvOr("KDEPLOY_OUT_OF_CLUSTER", false), "enable out of cluster k8s config discovery (env: KDEPLOY_OUT_OF_CLUSTER)")
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
		cli = client.New(kclient, iclient, lgger)
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
		cli = client.New(kclient, iclient, lgger)
	}

	gopts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(lgger.Zap()),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(lgger.Zap()),
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
	resolver := gql.NewResolver(meshpaaspb.NewMeshPaasServiceClient(conn), cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
	}), lgger)

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
