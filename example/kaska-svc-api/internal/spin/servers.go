package spin

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	httpApi "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/api/http"
)

type (
	Servers struct {
		AppConfig          *config.AppConfig
		HealthCheckerLive  health.Checker
		HealthCheckerReady health.Checker

		Grpc   *myGrpcServer
		Http   *myHttpServer
		stopFn sync.Once
	}

	myHttpServer struct {
		http   *http.Server
		logger *zap.Logger
		mux    *gin.Engine
		config *config.AppConfig
	}

	myGrpcServer struct {
		grpc   *grpc.Server
		logger *zap.Logger
		mux    *runtime.ServeMux
		config *config.AppConfig
	}
)

func (s *Servers) Run(ctx context.Context) (err error) {
	var ec = make(chan error, 2)
	ctx, cancel := context.WithCancel(ctx)

	grpcGwMux := s.getGrpcMux(s.AppConfig)
	s.Grpc = &myGrpcServer{
		logger: s.AppConfig.AppLogger,
		mux:    grpcGwMux,
		config: s.AppConfig,
	}

	httpMux := s.getHttpMux(s.AppConfig)
	s.Http = &myHttpServer{
		logger: s.AppConfig.AppLogger,
		mux:    httpMux,
		config: s.AppConfig,
	}

	go func() {
		err := s.Grpc.Run(ctx, net.JoinHostPort(s.AppConfig.AppListenHostname, s.AppConfig.AppListenPortGrpc))
		if err != nil {
			err = fmt.Errorf("gRPC server stop : %w", err)
		}
		ec <- err
	}()

	go func() {
		err := s.Http.Run(ctx, net.JoinHostPort(s.AppConfig.AppListenHostname, s.AppConfig.AppListenPortHttp))
		if err != nil {
			err = fmt.Errorf("HTTP Server stop : %w", err)
		}
		ec <- err
	}()

	var es []string

	for i := 0; i < cap(ec); i++ {
		if err := <-ec; err != nil {
			es = append(es, err.Error())
			if ctx.Err() == nil {
				s.Shutdown(context.Background())
			}
		}
	}

	if len(es) > 0 {
		err = errors.New(strings.Join(es, ", "))
	}

	if len(es) > 0 {
		err = errors.New(strings.Join(es, ", "))
	}

	cancel()

	return err
}

func (s *Servers) Shutdown(ctx context.Context) {
	s.stopFn.Do(func() {
		s.Http.Shutdown(ctx)
		s.Grpc.Shutdown(ctx)
	})
}

func (s *Servers) getGrpcMux(c *config.AppConfig) *runtime.ServeMux {
	return runtime.NewServeMux()
}

func (s *Servers) getHttpMux(c *config.AppConfig) *gin.Engine {

	httpRoutingEngine := httpApi.InitHTTPRoutingEngine(c)
	httpApi.AddHttpEndpoints(httpRoutingEngine, c)

	httpRoutingEngine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error_info": "Requested resource was not found"})
	})

	return httpRoutingEngine
}

func (s *myHttpServer) Run(ctx context.Context, httpAddr string) error {
	httpListener, err := net.Listen("tcp", httpAddr)

	if err != nil {
		s.logger.Fatal("error on http address : " + httpAddr)
		os.Exit(1)
	}

	//TODO: add zap logger
	hs := &http.Server{
		//TODO: add these to main config
		Handler:        s.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.http = hs
	s.logger.Sugar().Infof("HTTP service listening at %s", httpAddr)
	return hs.Serve(httpListener)
}

func (s *myHttpServer) Shutdown(ctx context.Context) {

	s.logger.Sugar().Infof("HTTP service gracefully shutting down ")

	if s.http != nil {
		if err := s.http.Shutdown(ctx); err != nil {
			s.logger.Fatal("graceful shutdown of HTTP service failed ")
		}
	}
}

func (s *myGrpcServer) Run(ctx context.Context, grpcAddress string) error {

	var lc net.ListenConfig
	lis, err := lc.Listen(ctx, "tcp", grpcAddress)

	if err != nil {
		s.logger.Sugar().Fatalf("error on grpc address : %s", err)
	}

	s.grpc = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	reflection.Register(s.grpc)

	/*

		KaskaApiSvcImpl := grpcApi.NewKaskaApiSvcServerImpl(
			s.config,
		)

		v1.RegisterKaskaApiSvcServer(s.grpc, KaskaApiSvcImpl)

	*/

	s.logger.Sugar().Infof("gRPC service listening at %v", lis.Addr())
	return s.grpc.Serve(lis)
}

func (s *myGrpcServer) Shutdown(ctx context.Context) {
	s.logger.Sugar().Infof("gRPC service gracefully shutting down ")

	done := make(chan struct{}, 1)

	go func() {
		if s.grpc != nil {
			s.grpc.GracefulStop()
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-ctx.Done():
		if s.grpc != nil {
			s.grpc.Stop()
		}
		s.logger.Fatal("graceful shutdown of gRPC server failed. ")
	}
}
