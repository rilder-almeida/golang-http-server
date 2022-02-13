package entity

type PlayerRepository interface {
	ShowAll() ([]*Player, error)
	GetPlayer(string) (*Player, error)
	PostPlayer(*Player) error
	PutPlayer(string, *Player) error
	DeletePlayer(string) error
	NewName(string) *Player
}

type Player struct {
	Name  string
	Score ScoreRepository
}
