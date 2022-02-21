package repository

import (
	"fmt"

	"github.com/rilder-almeida/golang-http-server/entity"
)

type PlayersMemoryDb struct {
	Players []*entity.Player
}

type PlayerRepositoryMemory struct {
	Db *PlayersMemoryDb
}

func NewPlayerRepositoryMemory() *PlayerRepositoryMemory {
	return &PlayerRepositoryMemory{&PlayersMemoryDb{[]*entity.Player{}}}
}

func (p *PlayerRepositoryMemory) ShowAll() ([]*entity.Player, error) {
	if len(p.Db.Players) == 0 {
		return nil, fmt.Errorf("no players found")
	}
	return p.Db.Players, nil
}

// Retrieve player
func (p *PlayerRepositoryMemory) GetPlayer(name string) (*entity.Player, error) {
	for i, playerFromDb := range p.Db.Players {
		if playerFromDb.Name == name {
			return p.Db.Players[i], nil
		}
	}
	return &entity.Player{}, fmt.Errorf("player %s not found", name)
}

// Create player
func (p *PlayerRepositoryMemory) PostPlayer(player *entity.Player) error {
	_, err := p.GetPlayer(player.Name)
	if err == nil {
		return fmt.Errorf("player %s already exists", player.Name)
	}
	p.Db.Players = append(p.Db.Players, player)
	return nil
}

// Update player
func (p *PlayerRepositoryMemory) PutPlayer(name string, player *entity.Player) error {
	for i, playerFromDb := range p.Db.Players {
		if playerFromDb.Name == name {
			p.Db.Players[i].Name = player.Name
			return nil
		}
	}
	return fmt.Errorf("player %s not found", player.Name)
}

// Delete player
func (p *PlayerRepositoryMemory) DeletePlayer(name string) error {
	for i, playerFromDb := range p.Db.Players {
		if playerFromDb.Name == name {
			p.Db.Players = append(p.Db.Players[:i], p.Db.Players[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("player %s not found", name)
}

// New player
func (p *PlayerRepositoryMemory) NewPlayer(name string, scoreRepository entity.ScoreRepository) *entity.Player {
	return &entity.Player{Name: name, Score: scoreRepository}
}
