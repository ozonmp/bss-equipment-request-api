package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/retranslator"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"
)

const retranslatorServerStartLogTag = "RetranslatorServe.Start()"

// RetranslatorServer is retranslator server
type RetranslatorServer struct {
	retranslator retranslator.Retranslator
}

// NewRetranslatorServer returns retranslator
func NewRetranslatorServer(retranslator retranslator.Retranslator) *RetranslatorServer {
	return &RetranslatorServer{
		retranslator: retranslator,
	}
}

// Start method runs server
func (r *RetranslatorServer) Start(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)

	metricsServer := createMetricsServer(cfg)

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("%s: metrics server is running on", retranslatorServerStartLogTag),
			"address", metricsAddr)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: metricsServer.ListenAndServe failed", retranslatorServerStartLogTag),
				"err", err)
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(ctx, cfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, fmt.Sprintf("%s: status server is running on", retranslatorServerStartLogTag),
			"address", statusAdrr)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, fmt.Sprintf("%s: statusServer.ListenAndServe failed", retranslatorServerStartLogTag),
				"err", err)
		}
	}()

	r.retranslator.Start()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.Info(ctx, fmt.Sprintf("%s: the service is ready to accept requests", retranslatorServerStartLogTag))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, fmt.Sprintf("%s: signal.Notify", retranslatorServerStartLogTag),
			"quit", v,
		)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, fmt.Sprintf("%s: ctx.Done", retranslatorServerStartLogTag),
			"done", done,
		)
	}

	isReady.Store(false)

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: statusServer.Shutdown failed", retranslatorServerStartLogTag),
			"err", err,
		)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: statusServer shut down correctly", retranslatorServerStartLogTag))
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: metricsServer.Shutdown failed", retranslatorServerStartLogTag),
			"err", err,
		)
	} else {
		logger.Info(ctx, fmt.Sprintf("%s: metricsServer shut down correctly", retranslatorServerStartLogTag))
	}

	r.retranslator.Close()
	logger.Info(ctx, fmt.Sprintf("%s: retranslator closed", retranslatorServerStartLogTag))

	return nil
}
