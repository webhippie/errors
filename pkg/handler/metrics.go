package handler

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/webhippie/errors/pkg/version"
)

const (
	namespace = "default_http_backend"
	subsystem = "http"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_count_total",
			Help:      "counter of http requests made",
		},
		[]string{"proto"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_duration_milliseconds",
			Help:      "histogram of the time (in milliseconds) each request took",
			Buckets:   append([]float64{.001, .003}, prometheus.DefBuckets...),
		},
		[]string{"proto"},
	)
)

func init() {
	prometheus.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
		Namespace: namespace,
	}))

	prometheus.MustRegister(version.Collector(namespace))

	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}

func handleMetrics(start time.Time, major, minor int) {
	duration := time.Since(start).Seconds() * 1e3

	requestCounter.WithLabelValues(fmt.Sprintf("%d.%d", major, minor)).Inc()
	requestDuration.WithLabelValues(fmt.Sprintf("%d.%d", major, minor)).Observe(duration)
}
