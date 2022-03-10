package main

func NewInMemoryPlayerStore() PlayerStore {
	return &inMemoryPlayerStore{map[string]int{}}
}

type inMemoryPlayerStore struct {
	store map[string]int
}

func (i *inMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *inMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
