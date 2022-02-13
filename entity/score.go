package entity

type ScoreRepository interface {
	AddScore()
	GetScore() int
	UpdateScore(int)
	ResetScore()
}

type Score struct {
	Value int
}
