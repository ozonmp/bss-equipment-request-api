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
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": metrics server is running on",
			"address", metricsAddr)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, retranslatorServerStartLogTag+": metricsServer.ListenAndServe() failed",
				"err", err)
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(ctx, cfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": status server is running on",
			"address", statusAdrr)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, retranslatorServerStartLogTag+": statusServer.ListenAndServe() failed",
				"err", err)
		}
	}()

	r.retranslator.Start()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": the service is ready to accept requests")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": signal.Notify", "quit", v)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": ctx.Done", "done", done)
	}

	isReady.Store(false)

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, retranslatorServerStartLogTag+": statusServer.Shutdown failed", "err", err)
	} else {
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": statusServer shut down correctly")
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, retranslatorServerStartLogTag+": metricsServer.Shutdown failed", "err", err)
	} else {
		logger.InfoKV(ctx, retranslatorServerStartLogTag+": metricsServer shut down correctly")
	}

	r.retranslator.Close()
	logger.InfoKV(ctx, retranslatorServerStartLogTag+": retranslator closed")

	return nil
}
