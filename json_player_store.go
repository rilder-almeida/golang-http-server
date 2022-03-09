package main

const (
	JSONFILENAME = "players.json"
)

type player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type jsonPlayerStore struct {
	Store    []player `json:"store"`
	jsonFile string   `json:"-"`
}

func NewJsonPlayerStore() PlayerStore {
	return &jsonPlayerStore{[]player{}, JSONFILENAME}
}

func (j *jsonPlayerStore) RecordWin(name string) {
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
	j.Store = append(j.Store, player{name, 1})
	ToJsonFile(JSONFILENAME, j)
}

func (j *jsonPlayerStore) GetPlayerScore(name string) int {
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
