package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rilder-almeida/golang-http-server/entity"
)

type ScoreBoardRepository struct {
	storage entity.Storage
}

func NewScoreBoardRepository(playerRepo entity.PlayerRepository, scoreRepo entity.ScoreRepository) *ScoreBoardRepository {
	return &ScoreBoardRepository{entity.NewStorage(playerRepo, scoreRepo)}
}

func printPlayer(player *entity.Player) string {
	return fmt.Sprintf("Player:\t%v\n\tScore:\t%v\n\n", player.Name, player.Score.GetScore())
}

func (sb *ScoreBoardRepository) ShowAll(writer http.ResponseWriter) {
	players, err := sb.storage.PlayerStorage.ShowAll()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, player := range players {
		fmt.Fprint(writer, printPlayer(player))
	}
}

func (sb *ScoreBoardRepository) GetPlayer(writer http.ResponseWriter, name string) {
	player, err := sb.storage.PlayerStorage.GetPlayer(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	printPlayer(player)
}

func (sb *ScoreBoardRepository) PostPlayer(writer http.ResponseWriter, name string) {
	newPlayer := sb.storage.PlayerStorage.NewPlayer(name, sb.storage.ScoreStorage)
	err := sb.storage.PlayerStorage.PostPlayer(newPlayer)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusAccepted)
}

func (sb *ScoreBoardRepository) PutPlayer(writer http.ResponseWriter, oldName string, newName string) {
	oldPlayer, err := sb.storage.PlayerStorage.GetPlayer(oldName)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = sb.storage.PlayerStorage.PutPlayer(newName, oldPlayer)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusAccepted)
}
func (sb *ScoreBoardRepository) DeletePlayer(writer http.ResponseWriter, name string) {
	err := sb.storage.PlayerStorage.DeletePlayer(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusAccepted)
}

func (sb *ScoreBoardRepository) ProcessWin(writer http.ResponseWriter, name string) {
	player, err := sb.storage.PlayerStorage.GetPlayer(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	player.Score.AddScore()
	writer.WriteHeader(http.StatusAccepted)
}
func (sb *ScoreBoardRepository) ShowScore(writer http.ResponseWriter, name string) string {
	player, err := sb.storage.PlayerStorage.GetPlayer(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return ""
	}
	writer.WriteHeader(http.StatusAccepted)
	return strconv.Itoa(player.Score.GetScore())
}
func (sb *ScoreBoardRepository) UpdateScore(writer http.ResponseWriter, name string, score string) {
	player, err := sb.storage.PlayerStorage.GetPlayer(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	intScore, err := strconv.Atoi(score)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	player.Score.UpdateScore(intScore)
	writer.WriteHeader(http.StatusAccepted)
}
func (sb *ScoreBoardRepository) ResetScore(writer http.ResponseWriter, name string) {
	player, err := sb.storage.PlayerStorage.GetPlayer(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	player.Score.ResetScore()
	writer.WriteHeader(http.StatusAccepted)
}
