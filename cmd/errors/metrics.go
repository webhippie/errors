package main

import (
	"github.com/prometheus/client_golang/prometheus"
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
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}
