package main

import (
	"net/http"

	"github.com/golang-http-server/entities/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type metricsMiddleware struct {
	next    http.Handler
	metrics *metrics.PrometheusMetrics
}

func WrapHandlerWithMetrics(next http.Handler, metrics *metrics.PrometheusMetrics) http.Handler {
	return &metricsMiddleware{
		next:    next,
		metrics: metrics,
	}
}

func (m *metricsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.metrics.TotalRequestsCounter.WithLabelValues(r.Method, r.URL.Path).Inc()

	timer := prometheus.NewTimer(m.metrics.RequestLatencyHistogram.WithLabelValues(r.Method, r.URL.Path))
	m.metrics.RequestLatencyHistogram.WithLabelValues(r.Method, r.URL.Path)

	m.next.ServeHTTP(w, r)

	// TODO: if for ResponseSuccessedCounter or ResponseFailedCounter based on body

	m.metrics.ResponseSuccessedCounter.WithLabelValues(r.Method, r.URL.Path).Inc()
	m.metrics.ResponseFailedCounter.WithLabelValues(r.Method, r.URL.Path, "", "").Inc() //TODO: get status code and error from w context

	timer.ObserveDuration()

}
