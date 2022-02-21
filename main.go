package main

import (
	"log"
	"net/http"

	repositoryMemory "github.com/rilder-almeida/golang-http-server/repository/memory"
	"github.com/rilder-almeida/golang-http-server/service"
)

func main() {
	scoreboard := service.NewScoreBoardRepository(repositoryMemory.NewPlayerRepositoryMemory())
	server := service.NewServer(scoreboard)
	log.Fatal(http.ListenAndServe(":5000", server))
}
