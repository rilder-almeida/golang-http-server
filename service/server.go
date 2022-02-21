package service

import (
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	scoreboard *ScoreBoardRepository
}

func NewServer(scoreboard *ScoreBoardRepository) *Server {
	return &Server{scoreboard}
}

func UrlParser(url string) (string, string, string, error) {
	splittedUrl := strings.Split(url, "/")
	switch len(splittedUrl) {
	case 2:
		infoData := splittedUrl[1]
		return infoData, "", "", nil
	case 3:
		infoData := splittedUrl[1]
		playerName := splittedUrl[2]
		return infoData, playerName, "", nil
	case 4:
		infoData := splittedUrl[1]
		playerName := splittedUrl[2]
		putArg := splittedUrl[3]
		return infoData, playerName, putArg, nil
	default:
		return "", "", "", fmt.Errorf("invalid url")
	}
}

func (p *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	infoData, playerName, putArg, err := UrlParser(request.URL.Path)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	switch infoData {
	case "scoreboard":
		switch request.Method {
		case http.MethodGet:
			p.scoreboard.ShowAll(writer)
		}
	case "player":
		if playerName == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		switch request.Method {
		case http.MethodGet:
			p.scoreboard.GetPlayer(writer, playerName)
		case http.MethodPost:
			p.scoreboard.PostPlayer(writer, playerName) // * new player
		case http.MethodPut:
			if putArg == "" {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			p.scoreboard.PutPlayer(writer, playerName, putArg) // * update player name | split new name "robert/bob"
		case http.MethodDelete:
			p.scoreboard.DeletePlayer(writer, playerName)
		}
	case "score":
		if playerName == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		switch request.Method {
		case http.MethodPost:
			p.scoreboard.ProcessWin(writer, playerName) // * score + 1
		case http.MethodGet:
			p.scoreboard.ShowScore(writer, playerName)
		case http.MethodPut:
			if putArg == "" {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			p.scoreboard.UpdateScore(writer, playerName, putArg) // * update player score | split new score "robert/10"
		case http.MethodDelete:
			p.scoreboard.ResetScore(writer, playerName)
		}
	}
}
