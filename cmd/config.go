package main

import (
	"os"

	"github.com/golang-http-server/entities/nfe/impltnfe"
)

var Config = impltnfe.Config{
	RepositoryType:     os.Getenv("REPOSITORY_TYPE"),
	RepositoryFilePath: "./data/nfe.json",
}
