package repository

import (
	"fmt"

	"github.com/rilder-almeida/golang-http-server/entity"
)

type PlayersMemoryDb struct {
	Players []*entity.Player
}

type PlayerRepositoryMemory struct {
	db *PlayersMemoryDb
}

func (p *PlayerRepositoryMemory) GetPlayer(name string) (*entity.Player, error) {
	for _, player := range p.db.Players {
		if player.Name == name {
			return player, nil
		}
	}
	return &entity.Player{}, fmt.Errorf("player %s not found", name)
}

func (p *PlayerRepositoryMemory) PostPlayer(player *entity.Player) error {
	p.db.Players = append(p.db.Players, player)
	return nil
}
