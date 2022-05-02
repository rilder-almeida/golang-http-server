package main

import (
	"os"

	"github.com/golang-http-server/entities/nfe/impltnfe"
)

var Config = impltnfe.Config{
	Repository:   os.Getenv("REPOSITORY"),
	JsonFilePath: "./data/nfe.json",
}
