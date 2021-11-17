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
		logger.InfoKV(ctx, grpcServerStartLogTag+": gateway server is running on",
			"address", grpcAddr)

		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, grpcServerStartLogTag+": gatewayServer.ListenAndServe failed",
				"err", err)
			cancel()
		}
	}()

	metricsServer := createMetricsServer(cfg)

	go func() {
		logger.InfoKV(ctx, grpcServerStartLogTag+": metrics server is running on",
			"address", metricsAddr)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, grpcServerStartLogTag+": metricsServer.ListenAndServe() failed",
				"err", err)
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(ctx, cfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, grpcServerStartLogTag+": status server is running on",
			"address", statusAdrr)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, grpcServerStartLogTag+": statusServer.ListenAndServe() failed",
				"err", err)
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
		logger.InfoKV(ctx, grpcServerStartLogTag+": GRPC server is listening on",
			"address", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			logger.ErrorKV(ctx, grpcServerStartLogTag+": grpcServer.Serve failed",
				"err", err)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.InfoKV(ctx, grpcServerStartLogTag+": the service is ready to accept requests")
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, grpcServerStartLogTag+": signal.Notify", "quit", v)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, grpcServerStartLogTag+": ctx.Done", "done", done)
	}

	isReady.Store(false)

	if err := gatewayServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, grpcServerStartLogTag+": gatewayServer.Shutdown failed", "err", err)
	} else {
		logger.InfoKV(ctx, grpcServerStartLogTag+": gatewayServer shut down correctly")
	}

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, grpcServerStartLogTag+": statusServer.Shutdown failed", "err", err)
	} else {
		logger.InfoKV(ctx, grpcServerStartLogTag+": statusServer shut down correctly")
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, grpcServerStartLogTag+": metricsServer.Shutdown failed", "err", err)
	} else {
		logger.InfoKV(ctx, grpcServerStartLogTag+": metricsServer shut down correctly")
	}

	grpcServer.GracefulStop()
	logger.InfoKV(ctx, grpcServerStartLogTag+": grpcServer shut down correctly")

	return nil
}
