package entity

type Storage struct {
	PlayerStorage PlayerRepository
	ScoreStorage  ScoreRepository
}

func NewStorage(playerRepo PlayerRepository) Storage {
	return Storage{
		PlayerStorage: playerRepo,
	}
}
