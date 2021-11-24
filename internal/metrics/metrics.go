package metrics

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalEquipmentRequestNotFound prometheus.Counter
var totalCudOperations *prometheus.CounterVec

// InitMetrics - init equipment request service metrics
func InitMetrics(cfg config.Config) {
	totalEquipmentRequestNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: cfg.Project.ServiceName,
		Name:      "not_found_total",
		Help:      "Total number of equipment requests that were not found",
	})

	totalCudOperations = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: cfg.Project.ServiceName,
		Name:      "cud_total",
		Help:      "Total cud operations",
	}, []string{"type"})
}

// IncTotalEquipmentRequestNotFound - increment amount of total not founded equipment requests
func IncTotalEquipmentRequestNotFound() {
	totalEquipmentRequestNotFound.Inc()
}

// IncTotalCudOperations - increment amount of total CUD operations
func IncTotalCudOperations(eventType model.EventType) {
	totalCudOperations.WithLabelValues(eventType.String()).Inc()
}
