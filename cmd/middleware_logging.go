package main

import (
	"net/http"
	"os"

	"github.com/golang-http-server/entities/logger"
	"github.com/rs/zerolog"
)

type loggingMiddleware struct {
	next   http.Handler
	logger *zerolog.Logger
}

var config = logger.ConfigLogger{
	From:         "server",
	MinimumLevel: "trace",
	Output:       os.Stdout,
}

func WrapServerWithLogging(next http.Handler) http.Handler {
	return &loggingMiddleware{
		next:   next,
		logger: logger.NewLogger(config),
	}
}

func (m *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer m.logger.Trace().Msg("Request received")
	m.next.ServeHTTP(w, r)
}
