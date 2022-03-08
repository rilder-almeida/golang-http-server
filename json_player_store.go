package main

const (
	JSONFILENAME = "players.json"
)

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type JsonPlayerStore struct {
	Store []Player `json:"store"`
}

func NewJsonPlayerStore() *JsonPlayerStore {
	return &JsonPlayerStore{}
}

func (j *JsonPlayerStore) RecordWin(name string) {
	err := FromJsonFile(JSONFILENAME, j)
	if err != nil {
		panic(err)
	}
	for i, p := range j.Store {
		if p.Name == name {
			j.Store[i].Score++
			ToJsonFile(JSONFILENAME, j)
			return
		}
	}
	j.Store = append(j.Store, Player{name, 1})
	ToJsonFile(JSONFILENAME, j)
}

func (j *JsonPlayerStore) GetPlayerScore(name string) int {
	err := FromJsonFile(JSONFILENAME, j)
	if err != nil {
		panic(err)
	}
	for i, p := range j.Store {
		if p.Name == name {
			return j.Store[i].Score
		}
	}
	return 0
}
