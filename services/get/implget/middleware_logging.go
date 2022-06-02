package implget

import (
	"os"

	"github.com/golang-http-server/entities/logger"
	"github.com/golang-http-server/services/get"
	"github.com/rs/zerolog"
)

type loggingMiddleware struct {
	next   get.GetGateway
	logger *zerolog.Logger
}

var config = logger.ConfigLogger{
	Fields:       map[string]interface{}{"service": "get"},
	From:         "get",
	MinimumLevel: "trace",
	Output:       os.Stdout,
}

func WrapGetServiceWithLogging(next get.GetGateway) get.GetGateway {
	return &loggingMiddleware{
		next:   next,
		logger: logger.NewLogger(config),
	}
}

func (m *loggingMiddleware) Processor(request get.Request) (get.Response, error) {
	defer m.logger.Trace().Msg("Processing request")
	return m.next.Processor(request)
}
