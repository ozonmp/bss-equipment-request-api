package server

import (
	"context"
	"errors"
	"fmt"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/pkg/grps_logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/ozonmp/bss-equipment-request-api/internal/api"
	"github.com/ozonmp/bss-equipment-request-api/internal/config"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
)

const grpcServerStartLogTag = "GrpcServer.Start()"

// GrpcServer is gRPC server
type GrpcServer struct {
	equipmentRequestService equipment_request.ServiceInterface
}

// NewGrpcServer returns gRPC server with supporting of batch listing
func NewGrpcServer(equipmentRequestService equipment_request.ServiceInterface) *GrpcServer {
	return &GrpcServer{
		equipmentRequestService: equipmentRequestService,
	}
}

// Start method runs server
func (s *GrpcServer) Start(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gatewayAddr := fmt.Sprintf("%s:%v", cfg.Rest.Host, cfg.Rest.Port)
	grpcAddr := fmt.Sprintf("%s:%v", cfg.Grpc.Host, cfg.Grpc.Port)
	metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)

	gatewayServer := createGatewayServer(ctx, grpcAddr, gatewayAddr)

	go func() {

		logger.InfoKV(ctx, fmt.Sprintf("%s: gateway server is running on", grpcServerStartLogTag),
			"address", grpcAddr)

		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: gatewayServer.ListenAndServe failed", grpcServerStartLogTag),
				"err", err)
			cancel()
		}
	}()

	metricsServer := createMetricsServer(cfg.Metrics.Host, cfg.Metrics.Path, cfg.Metrics.Port)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("%s: metrics server is listening on", grpcServerStartLogTag),
			"address", metricsAddr,
		)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: metricsServer.ListenAndServe failed", grpcServerStartLogTag),
				"err", err)
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusCfg := &statusCfg{
		Port:          cfg.Status.Port,
		Host:          cfg.Status.Host,
		LivenessPath:  cfg.Status.LivenessPath,
		ReadinessPath: cfg.Status.ReadinessPath,
		VersionPath:   cfg.Status.VersionPath,
	}

	projectCfg := &projectCfg{
		Name:        cfg.Project.Name,
		Debug:       cfg.Project.Debug,
		Environment: cfg.Project.Environment,
		Version:     cfg.Project.Version,
		CommitHash:  cfg.Project.CommitHash,
	}
	statusServer := createStatusServer(ctx, statusCfg, projectCfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, fmt.Sprintf("%s: status server is listening on", grpcServerStartLogTag),
			"address", statusAdrr,
		)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: statusServer.ListenAndServe failed", grpcServerStartLogTag),
				"err", err,
			)
		}
	}()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	//nolint
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_opentracing.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
			grpc_zap.PayloadUnaryServerInterceptor(logger.Clone(ctx), grps_logger.ServerPayloadLoggingDecider()),
			grps_logger.UnaryServerInterceptor(),
		)),
	)

	pb.RegisterBssEquipmentRequestApiServiceServer(grpcServer, api.NewEquipmentRequestAPI(s.equipmentRequestService))
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("%s: GRPC server is listening on", grpcServerStartLogTag),
			"address", grpcAddr,
		)
		if err := grpcServer.Serve(l); err != nil {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: grpcServer.Serve failed", grpcServerStartLogTag),
				"err", err,
			)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.Info(ctx, fmt.Sprintf("%s: the service is ready to accept requests", grpcServerStartLogTag))
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, fmt.Sprintf("%s: signal.Notify", grpcServerStartLogTag), "quit", v)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, fmt.Sprintf("%s: ctx.Done", grpcServerStartLogTag), "done", done)
	}

	isReady.Store(false)

	if err := gatewayServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: gatewayServer.Shutdown failed", grpcServerStartLogTag), "err", err)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: gatewayServer shut down correctly", grpcServerStartLogTag))
	}

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: statusServer.Shutdown failed", grpcServerStartLogTag), "err", err)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: statusServer shut down correctly", grpcServerStartLogTag))
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: metricsServer.Shutdown failed", grpcServerStartLogTag), "err", err)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: metricsServer shut down correctly", grpcServerStartLogTag))
	}

	grpcServer.GracefulStop()
	logger.Info(ctx, fmt.Sprintf("%s: grpcServer shut down correctly", grpcServerStartLogTag))

	return nil
}
