package logger

import (
	"time"

	"github.com/rs/zerolog"
)

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
