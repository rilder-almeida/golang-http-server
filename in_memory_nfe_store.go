package main

type inMemoryNfeStore struct {
	store map[string]XmlDocument
}

func NewInMemoryNfeStore() NfeStore {
	return &inMemoryNfeStore{}
}

func (i *inMemoryNfeStore) ReceiverNfe(XmlDocument) error { return nil }

func (i *inMemoryNfeStore) ResponserNfe(XmlDocument) (JsonPostResponse, error) {
	return JsonPostResponse{}, nil
}

func (i *inMemoryNfeStore) AssertNfe() {}

func (i *inMemoryNfeStore) ReceiveId() {}

func (i *inMemoryNfeStore) ResponseId() {}

func (i *inMemoryNfeStore) AssertId() {}
