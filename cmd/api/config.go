package main

import (
	"time"

	fklog "github.com/arquivei/foundationkit/log"
)

var config struct {
	Log fklog.Config

	HTTP struct {
		Port string `default:"8000"`
	}

	Shutdown struct {
		Graceperiod time.Duration `default:"3s"`
		Timeout     time.Duration `default:"5s"`
	}
}
