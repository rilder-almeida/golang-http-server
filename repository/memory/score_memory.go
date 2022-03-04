package repository

import (
	"github.com/rilder-almeida/golang-http-server/entity"
)

type ScoreMemoryDb struct {
	Points *entity.Score
}

type ScoreRepositoryMemory struct {
	Db *ScoreMemoryDb
}

func NewScoreRepositoryMemory() *ScoreRepositoryMemory {
	return &ScoreRepositoryMemory{&ScoreMemoryDb{&entity.Score{Value: 0}}}
}

func (s *ScoreRepositoryMemory) AddScore() {
	s.Db.Points.Value += 1
}

func (s *ScoreRepositoryMemory) GetScore() int {
	return s.Db.Points.Value
}

func (s *ScoreRepositoryMemory) UpdateScore(newValue int) {
	s.Db.Points.Value = newValue
}

func (s *ScoreRepositoryMemory) ResetScore() {
	s.Db.Points.Value = 0
}

func (s *ScoreRepositoryMemory) NewScore() *entity.Score {
	return &entity.Score{Value: 0}
}
