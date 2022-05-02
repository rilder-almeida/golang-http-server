package main

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/nfe/impltnfe"
)

var config = impltnfe.Config{
	Repository:   "INFILE",
	JsonFilePath: "./data/nfe.json",
}

func NewRepository(config impltnfe.Config) nfe.Repository {
	return impltnfe.NewNfeRepository(config)
}
