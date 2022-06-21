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

	Repository struct {
		Type     string `default:"INFILE"`
		FilePath string `default:"./data/nfe.json"`
	}

	Postgresql struct {
		Host     string `default:"localhost",`
		User     string `default:"stark",`
		Password string `default:"stark",`
		Dbname   string `default:"winterfell",`
		Sslmode  string `default:"disable",`
		Port     string `default:"5432",`
	}
}
