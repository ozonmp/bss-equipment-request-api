package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"net/http"
	"sync/atomic"
)

type statusCfg struct {
	Host          string
	Port          int
	LivenessPath  string
	ReadinessPath string
	VersionPath   string
}

type projectCfg struct {
	Name        string
	Debug       bool
	Environment string
	Version     string
	CommitHash  string
}

func createStatusServer(ctx context.Context, statusCfg *statusCfg, projectCfg *projectCfg, isReady *atomic.Value) *http.Server {
	statusAddr := fmt.Sprintf("%s:%v", statusCfg.Host, statusCfg.Port)

	mux := http.DefaultServeMux

	mux.HandleFunc(statusCfg.LivenessPath, livenessHandler)
	mux.HandleFunc(statusCfg.ReadinessPath, readinessHandler(isReady))
	mux.HandleFunc(statusCfg.VersionPath, versionHandler(ctx, projectCfg))

	statusServer := &http.Server{
		Addr:    statusAddr,
		Handler: mux,
	}

	return statusServer
}

func livenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)

			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func versionHandler(ctx context.Context, projectCfg *projectCfg) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		data := map[string]interface{}{
			"name":        projectCfg.Name,
			"debug":       projectCfg.Debug,
			"environment": projectCfg.Environment,
			"version":     projectCfg.Version,
			"commitHash":  projectCfg.CommitHash,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logger.ErrorKV(ctx, "status.versionHandler: encode failed", "err", err)
		}
	}
}
