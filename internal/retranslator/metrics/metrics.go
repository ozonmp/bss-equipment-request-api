package metrics

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var numberOfEventsInRetranslator prometheus.Gauge

// InitMetrics - init retranslator service metrics
func InitMetrics(cfg config.Config) {
	numberOfEventsInRetranslator = promauto.NewGauge(prometheus.GaugeOpts{
		Subsystem: cfg.Project.ServiceName,
		Name:      "events_in_retranslator",
		Help:      "Number of events in retranslator",
	})
}

// AddEventsInRetranslator - add new events that should be processed
func AddEventsInRetranslator(c uint) {
	numberOfEventsInRetranslator.Add(float64(c))
}

// RemoveEventsFromRetranslator - remove already processed events
func RemoveEventsFromRetranslator(c uint) {
	numberOfEventsInRetranslator.Sub(float64(c))
}
