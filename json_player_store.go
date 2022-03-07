package main

import "fmt"

const (
	JSONFILENAME = "players.json"
)

type Player struct {
	name  string
	score int
}

type JsonPlayerStore struct {
	store []*Player

	// ! WHY GOD? WHY?????
	// ? store []*Player != store []Player
	// ? &JsonPlayerStore{[]*Player{}} != &JsonPlayerStore{[]Player{}}
	// ? &JsonPlayerStore{*[]Player{}} != &JsonPlayerStore{[]*Player{}}
}

func NewJsonPlayerStore() *JsonPlayerStore {
	return &JsonPlayerStore{[]*Player{}}
}

func (j *JsonPlayerStore) RecordWin(name string) {
	err := ReadJsonFromFile(JSONFILENAME, &j.store)
	if err != nil {
		fmt.Println("Error reading to file: ", err)
	}
	for _, p := range j.store {
		if p.name == name {
			p.score++
			err := WriteJsonToFile(JSONFILENAME, &j.store)
			if err != nil {
				fmt.Println("Error writing to file: ", err)
			}
			return
		}
	}
	j.store = append(j.store, &Player{name, 1})
	err = WriteJsonToFile(JSONFILENAME, &j.store)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
}

func (j *JsonPlayerStore) GetPlayerScore(name string) int {
	err := ReadJsonFromFile(JSONFILENAME, &j.store)
	if err != nil {
		fmt.Println("Error reading to file: ", err)
	}
	for _, p := range j.store {
		if p.name == name {
			return p.score
		}
	}
	return 0
}
