package implinsert

import (
	"os"

	"github.com/golang-http-server/entities/logger"
	"github.com/golang-http-server/services/insert"
	"github.com/rs/zerolog"
)

type loggingMiddleware struct {
	next   insert.InsertGateway
	logger *zerolog.Logger
}

var config = logger.ConfigLogger{
	Fields:       map[string]interface{}{"service": "insert"},
	From:         "insert",
	MinimumLevel: "trace",
	Output:       os.Stdout,
}

func WrapInsertServiceWithLogging(next insert.InsertGateway) insert.InsertGateway {
	return &loggingMiddleware{
		next:   next,
		logger: logger.NewLogger(config),
	}
}

func (m *loggingMiddleware) Processor(request insert.Request) (insert.Response, error) {
	defer m.logger.Trace().Msg("Processing request")
	return m.next.Processor(request)
}
