package server

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func createMetricsServer(host, path string, port int) *http.Server {
	addr := fmt.Sprintf("%s:%d", host, port)

	mux := http.DefaultServeMux
	mux.Handle(path, promhttp.Handler())

	metricsServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return metricsServer
}
