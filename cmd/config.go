package main

import (
	"github.com/golang-http-server/entities/nfe/impltnfe"
)

var Config = impltnfe.Config{
	Repository:   "INFILE",
	JsonFilePath: "./data/nfe.json",
}
