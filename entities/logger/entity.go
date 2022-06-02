package logger

import (
	"io"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	ZLogger *zerolog.Logger
	Config  ConfigLogger
}

type ConfigLogger struct {
	Fields       map[string]interface{}
	From         string
	MinimumLevel string
	Output       io.Writer
}

func NewLogger(config ConfigLogger) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: config.Output, TimeFormat: time.RFC3339, NoColor: true}

	levelLog, err := zerolog.ParseLevel(config.MinimumLevel)
	if err != nil {
		levelLog = zerolog.InfoLevel
	}

	logger := zerolog.
		New(output).
		Level(zerolog.Level(levelLog)).
		With().
		Timestamp().
		Str("From", config.From).
		Fields(config.Fields).
		Caller().
		Logger()

	return &logger
}
