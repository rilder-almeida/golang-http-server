package main

import (
	"log"
	"net/http"

	"github.com/golang-http-server/entities/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type metricsMiddleware struct {
	next    http.Handler
	metrics *metrics.PrometheusMetrics
}

func WrapHandlerWithMetrics(next http.Handler) http.Handler {
	prometheusMetrics := &metrics.PrometheusMetrics{
		TotalRequestsCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "total_requests_counter",
				Help: "Total requests counter",
			},
			[]string{"method", "path"},
		),
		RequestLatencyHistogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "request_latency_histogram",
				Help: "Request latency histogram",
			},
			[]string{"method", "path"},
		)}
	err := prometheus.Register(prometheusMetrics.TotalRequestsCounter)
	if err != nil {
		log.Fatalf("Failed to register metrics: %v", err)
	}

	err = prometheus.Register(prometheusMetrics.RequestLatencyHistogram)
	if err != nil {
		log.Fatalf("Failed to register metrics: %v", err)
	}

	return &metricsMiddleware{
		next:    next,
		metrics: prometheusMetrics,
	}
}

func (m *metricsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.metrics.TotalRequestsCounter.WithLabelValues(r.Method, r.URL.Path).Inc()

	timer := prometheus.NewTimer(m.metrics.RequestLatencyHistogram.WithLabelValues(r.Method, r.URL.Path))
	m.metrics.RequestLatencyHistogram.WithLabelValues(r.Method, r.URL.Path)

	m.next.ServeHTTP(w, r)

	timer.ObserveDuration()
}
