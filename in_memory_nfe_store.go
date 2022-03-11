package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	POST_RESPONSE_CONTENT_TYPE = "application/json"
)

type NfeDocument struct {
	RawXml         string      `json:"raw_xml"`
	NfeXmlDocument XmlDocument `json:"nfe_xml_document"`
}

type JsonNfeIsNew struct {
	IsNewNfe bool `json:"IsNewNfe"`
}

type inMemoryNfeStore struct {
	store map[string]NfeDocument
}

func NewInMemoryNfeStore() NfeStore {
	return &inMemoryNfeStore{make(map[string]NfeDocument)}
}

func (i *inMemoryNfeStore) PostRequestReceiver(jsonRequest JsonPostRequest) (JsonPostResponse, error) {
	bodyData, err := i.StoreNfe(jsonRequest)

	if err != nil {
		return i.PostRequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotAcceptable), err
	}

	return i.PostRequestResponder(POST_RESPONSE_CONTENT_TYPE, bodyData, http.StatusOK), nil
}

func (i *inMemoryNfeStore) PostRequestResponder(contentType string, bodyData []byte, httpStatus int) JsonPostResponse {
	return JsonPostResponse{
		contentType: contentType,
		bodyData:    bodyData,
		httpStatus:  httpStatus,
	}
}

func (i *inMemoryNfeStore) GetRequestReceiver() {} // TODO: must be implemented

func (i *inMemoryNfeStore) GetRequestResponder() {} // TODO: must be implemented

func (i *inMemoryNfeStore) AssertIdIsNew(id string) bool {
	_, ok := i.store[id]
	return !ok
}

func (i *inMemoryNfeStore) StoreNfe(jsonRequest JsonPostRequest) ([]byte, error) {
	xmlParsed, err := xmlParser([]byte(jsonRequest.XML))
	if err != nil {
		return []byte(""), fmt.Errorf("error parsing XML: %s", err)
	}

	if i.AssertIdIsNew(xmlParsed.NFe.InfNFe.Id) {
		i.store[xmlParsed.NFe.InfNFe.Id] = NfeDocument{RawXml: jsonRequest.XML, NfeXmlDocument: xmlParsed}
		return i.MakeJsonNfeIsNew(true), nil
	}

	return i.MakeJsonNfeIsNew(false), nil
}

func (i *inMemoryNfeStore) GetNfeById() {}

func (i *inMemoryNfeStore) MakeJsonNfeIsNew(status bool) []byte {
	data, err := json.Marshal(JsonNfeIsNew{IsNewNfe: status})
	if err != nil {
		return []byte("")
	}
	return data
}
