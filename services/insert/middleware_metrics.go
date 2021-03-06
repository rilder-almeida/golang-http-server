package insert

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
				Name: "insert_response_successed_counter",
				Help: "Insert Response successed counter",
			},
			[]string{"method", "path"},
		),
		ResponseFailedCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "insert_response_failed_counter",
				Help: "Insert Response failed counter",
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

func (m *metricsMiddleware) Insert(request Request) (Response, error) {
	response, err := m.next.Insert(request)
	if err != nil {
		m.metrics.ResponseFailedCounter.WithLabelValues("INSERT", "/nfe/v1", err.(customerrors.Error).ErrorCode).Inc()
		return Response{}, err
	}
	m.metrics.ResponseSuccessedCounter.WithLabelValues("INSERT", "/nfe/v1").Inc()

	return response, nil
}
