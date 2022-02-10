package entity

type PlayerRepository interface {
	GetPlayer(name string) (*Player, error)
	PostPlayer(player *Player) error
}

type Player struct {
	Name  string
	Score Score
}
