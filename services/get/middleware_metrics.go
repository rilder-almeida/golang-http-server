package get

import (
	"log"

	customerrors "github.com/golang-http-server/entities/errors"
	"github.com/golang-http-server/entities/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type metricsMiddleware struct {
	next    Service
	metrics *metrics.PrometheusMetrics
}

func WrapServiceWithMetrics(next Service) Service {
	prometheusMetrics := &metrics.PrometheusMetrics{
		ResponseSuccessedCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "get_response_successed_counter",
				Help: "Get Response successed counter",
			},
			[]string{"method", "path"},
		),
		ResponseFailedCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "get_response_failed_counter",
				Help: "Get Response failed counter",
			},
			[]string{"method", "path", "error_code"},
		)}
	err := prometheus.Register(prometheusMetrics.ResponseSuccessedCounter)
	if err != nil {
		log.Fatalf("Failed to register metrics: %v", err)
	}

	err = prometheus.Register(prometheusMetrics.ResponseFailedCounter)
	if err != nil {
		log.Fatalf("Failed to register metrics: %v", err)
	}

	return &metricsMiddleware{
		next:    next,
		metrics: prometheusMetrics,
	}
}

func (m *metricsMiddleware) Get(request Request) (Response, error) {
	response, err := m.next.Get(request)
	if err != nil {
		m.metrics.ResponseFailedCounter.WithLabelValues("GET", "/nfe/get", err.(customerrors.Error).ErrorCode).Inc()
		return Response{}, err
	}
	m.metrics.ResponseSuccessedCounter.WithLabelValues("GET", "/nfe/get").Inc()

	return response, nil
}
