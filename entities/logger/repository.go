package logger

import (
	"time"

	"github.com/rs/zerolog"
)

type ZLogger interface {
	Msg(string, string)
	Err(error, string)
	Fatal(error, string)
}

func NewLogger(config ConfigLogger) *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: config.Output, TimeFormat: time.RFC3339}

	levelLog, err := zerolog.ParseLevel(config.DefaultLevel)
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
