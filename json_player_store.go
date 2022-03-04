package main

type Player struct {
	name  string
	score int
}

type JsonPlayerStore struct {
	store []*Player

	// ! WHY GOD? WHY?????
	// ? store []*Player != store []Player
	// ? &JsonPlayerStore{[]*Player{}} != &JsonPlayerStore{[]Player{}}
}

func NewJsonPlayerStore() *JsonPlayerStore {
	return &JsonPlayerStore{[]*Player{}}
}

func (j *JsonPlayerStore) RecordWin(name string) {
	for _, p := range j.store {
		if p.name == name {
			p.score++
			return
		}
	}
	j.store = append(j.store, &Player{name, 1})
}

func (j *JsonPlayerStore) GetPlayerScore(name string) int {
	for _, p := range j.store {
		if p.name == name {
			return p.score
		}
	}
	return 0
}
