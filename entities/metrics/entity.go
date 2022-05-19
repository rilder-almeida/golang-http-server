package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
	TotalRequestsCounter     *prometheus.CounterVec
	RequestLatencyHistogram  *prometheus.HistogramVec
	ResponseSuccessedCounter *prometheus.CounterVec
	ResponseFailedCounter    *prometheus.CounterVec
}

func NewPrometheus() (*PrometheusMetrics, error) {
	prometheusMetrics := &PrometheusMetrics{
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
		),
		ResponseSuccessedCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "response_successed_counter",
				Help: "Response successed counter",
			},
			[]string{"method", "path"},
		),
		ResponseFailedCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "response_failed_counter",
				Help: "Response failed counter",
			},
			[]string{"method", "path", "error", "status_code"},
		),
	}
	err := prometheus.Register(prometheusMetrics.TotalRequestsCounter)
	if err != nil {
		return nil, errors.New("failed to register TotalRequestsCounter")
	}

	err = prometheus.Register(prometheusMetrics.RequestLatencyHistogram)
	if err != nil {
		return nil, errors.New("failed to register RequestLatencyHistogram")
	}

	err = prometheus.Register(prometheusMetrics.ResponseSuccessedCounter)
	if err != nil {
		return nil, errors.New("failed to register ResponseSuccessedCounter")
	}

	err = prometheus.Register(prometheusMetrics.ResponseFailedCounter)
	if err != nil {
		return nil, errors.New("failed to register ResponseFailedCounter")
	}

	return prometheusMetrics, nil
}
