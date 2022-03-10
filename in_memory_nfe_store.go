package main

type inMemoryNfeStore struct {
	store map[int]XmlDocument
}

func NewInMemoryNfeStore() NfeStore {
	return &inMemoryNfeStore{}
}

func (i *inMemoryNfeStore) ReceiveNfe() {}

func (i *inMemoryNfeStore) ResponseNfe() {}

func (i *inMemoryNfeStore) AssertNfe() {}

func (i *inMemoryNfeStore) ReceiveId() {}

func (i *inMemoryNfeStore) ResponseId() {}

func (i *inMemoryNfeStore) AssertId() {}
