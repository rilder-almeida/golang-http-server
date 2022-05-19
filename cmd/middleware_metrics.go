package main

import (
	"net/http"
	"time"

	"github.com/golang-http-server/entities/metrics"
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
	m.metrics.RequestLatencyHistogram.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(time.Now()).Seconds())

	m.next.ServeHTTP(w, r)

	m.metrics.ResponseSuccessedCounter.WithLabelValues(r.Method, r.URL.Path).Inc()
	m.metrics.ResponseFailedCounter.WithLabelValues(r.Method, r.URL.Path, "", "").Inc() //TODO get status code and error from w context
}
