package entity

type Storage struct {
	PlayerStorage PlayerRepository
	ScoreStorage  ScoreRepository
}

func NewStorage(playerRepo PlayerRepository, scoreRepo ScoreRepository) Storage {
	return Storage{
		PlayerStorage: playerRepo,
		ScoreStorage:  scoreRepo,
	}
}
