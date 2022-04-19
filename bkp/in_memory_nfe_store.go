package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type inMemoryNfeStore struct {
	store map[string]NfeDocument
}

func NewInMemoryNfeStore() NfeStore {
	return &inMemoryNfeStore{make(map[string]NfeDocument)}
}

func (i *inMemoryNfeStore) PostRequestReceiver(jsonRequest JsonPostRequest) (JsonResponse, error) {
	bodyData, err := i.StoreNfe(jsonRequest)

	if err != nil {
		return i.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotAcceptable), err
	}

	return i.RequestResponder(POST_RESPONSE_CONTENT_TYPE, bodyData, http.StatusOK), nil
}

func (i *inMemoryNfeStore) GetRequestReceiver(jsonRequest JsonGetRequest) (JsonResponse, error) {
	nfeDoc, err := i.GetNfeById(jsonRequest.Id)

	if err != nil {
		return i.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotFound), err
	}

	return i.RequestResponder(POST_RESPONSE_CONTENT_TYPE, nfeDoc, http.StatusOK), nil
}

func (i *inMemoryNfeStore) RequestResponder(contentType string, bodyData []byte, httpStatus int) JsonResponse {
	return JsonResponse{
		contentType: contentType,
		bodyData:    bodyData,
		httpStatus:  httpStatus,
	}
}

func (i *inMemoryNfeStore) AssertIdIsNew(id string) bool {
	_, ok := i.store[id]
	return !ok
}

func (i *inMemoryNfeStore) StoreNfe(jsonRequest JsonPostRequest) ([]byte, error) {
	xmlParsed, err := ToXmlParser([]byte(jsonRequest.XML))
	if err != nil {
		return []byte(""), fmt.Errorf("error parsing XML: %s", err)
	}

	if i.AssertIdIsNew(xmlParsed.NFe.InfNFe.Id) {
		i.store[xmlParsed.NFe.InfNFe.Id] = NfeDocument{RawXml: jsonRequest.XML, NfeXmlDocument: xmlParsed}
		return i.MakeJsonNfeIsNew(true), nil
	}

	return i.MakeJsonNfeIsNew(false), nil
}

func (i *inMemoryNfeStore) GetNfeById(id string) ([]byte, error) {
	if i.AssertIdIsNew(id) {
		return []byte(""), fmt.Errorf("NFe with id %s not found", id)
	}

	return json.Marshal(i.store[id].NfeXmlDocument)
}

func (i *inMemoryNfeStore) MakeJsonNfeIsNew(status bool) []byte {
	data, err := json.Marshal(JsonNfeIsNew{IsNewNfe: status})
	if err != nil {
		return []byte("")
	}
	return data
}
