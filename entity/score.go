package entity

type Score struct {
	Value int
}

type ScoreRepository interface {
	AddScore(score *Score) error
	GetScore(name string) (*Score, error)
}
