package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
	TotalRequestsCounter     *prometheus.CounterVec
	RequestLatencyHistogram  *prometheus.HistogramVec
	ResponseSuccessedCounter *prometheus.CounterVec
	ResponseFailedCounter    *prometheus.CounterVec
}
