package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	TotalRequestsCounter           *prometheus.CounterVec
	ResponseSuccessedTimeHistogram *prometheus.HistogramVec
	ResponseFailedTimeHistogram    *prometheus.HistogramVec
	ResponseSuccessedCounter       *prometheus.CounterVec
	ResponseFailedCounter          *prometheus.CounterVec
	RequestLatencyHistogram        *prometheus.HistogramVec
}
