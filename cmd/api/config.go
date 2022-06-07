package main

import (
	"os"
	"time"

	fklog "github.com/arquivei/foundationkit/log"
	"github.com/golang-http-server/entities/nfe/impltnfe"
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

// TODO PASSAR O repositoryConfig PARA A config E SETAR AS VARIÁVEIS NO DOCKER COMPOSE

var repositoryConfig = impltnfe.Config{
	RepositoryType:     os.Getenv("REPOSITORY_TYPE"),
	RepositoryFilePath: "./data/nfe.json",
}
