package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func NewServer() *PlayerServer {
	return &PlayerServer{
		store: getPlayerStore(),
	}
}

func getPlayerStore() PlayerStore {
	envVar := os.Getenv("PLAYER_STORE")
	switch envVar {
	case "JSON":
		return NewJsonPlayerStore()
	case "IN_MEMORY":
		return NewInMemoryPlayerStore()
	}
	panic(fmt.Sprintf("Bad environment variable value: %s", envVar))
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
