package main

type inJsonNfeStore struct {
	store []XmlDocument `json:"store"`
}

func NewInJsonNfeStore() NfeStore {
	return &inJsonNfeStore{}
}

func (i *inJsonNfeStore) ReceiveNfe() {}

func (i *inJsonNfeStore) ResponseNfe() {}

func (i *inJsonNfeStore) AssertNfe() {}

func (i *inJsonNfeStore) ReceiveId() {}

func (i *inJsonNfeStore) ResponseId() {}

func (i *inJsonNfeStore) AssertId() {}
