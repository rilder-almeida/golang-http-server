package main

import (
	"log"
	"net/http"

	repository "github.com/rilder-almeida/golang-http-server/repository/memory"
	"github.com/rilder-almeida/golang-http-server/service"
)

func main() {
	server := &service.NewScoreBoardRepository(repository.NewPlayerRepositoryMemory(), repository.NewScoreRepositoryMemory())
	log.Fatal(http.ListenAndServe(":5000", server))
}
