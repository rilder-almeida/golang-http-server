package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	ZLogger *zerolog.Logger
	Config  ConfigLogger
}

type ConfigLogger struct {
	Fields       map[string]interface{}
	From         string
	DefaultLevel string
	Output       io.Writer
}
