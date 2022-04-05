package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type inJsonNfeStore struct {
	Store    map[string]NfeDocument `json:"Store"`
	Filename string                 `json:"-"`
}

func NewInJsonNfeStore() NfeStore {
	return &inJsonNfeStore{
		make(map[string]NfeDocument),
		JSON_FILENAME,
	}
}

func (j *inJsonNfeStore) PostRequestReceiver(jsonRequest JsonPostRequest) (JsonResponse, error) {
	bodyData, err := j.StoreNfe(jsonRequest)

	if err != nil {
		return j.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotAcceptable), err
	}

	return j.RequestResponder(POST_RESPONSE_CONTENT_TYPE, bodyData, http.StatusOK), nil
}

func (j *inJsonNfeStore) RequestResponder(contentType string, bodyData []byte, httpStatus int) JsonResponse {
	return JsonResponse{
		contentType: contentType,
		bodyData:    bodyData,
		httpStatus:  httpStatus,
	}
}

func (j *inJsonNfeStore) GetRequestReceiver(jsonRequest JsonGetRequest) (JsonResponse, error) {
	nfeDoc, err := j.GetNfeById(jsonRequest.Id)

	if err != nil {
		return j.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotFound), err
	}

	return j.RequestResponder(POST_RESPONSE_CONTENT_TYPE, nfeDoc, http.StatusOK), nil
}

func (j *inJsonNfeStore) AssertIdIsNew(id string) bool {
	_, ok := j.Store[id]
	return !ok
}

func (j *inJsonNfeStore) StoreNfe(jsonRequest JsonPostRequest) ([]byte, error) {
	err := FromJsonFile(j.Filename, j)
	if err != nil {
		panic(err)
	}

	xmlParsed, err := ToXmlParser([]byte(jsonRequest.XML))
	if err != nil {
		return []byte(""), fmt.Errorf("error parsing XML: %s", err)
	}

	if j.AssertIdIsNew(xmlParsed.NFe.InfNFe.Id) {
		j.Store[xmlParsed.NFe.InfNFe.Id] = NfeDocument{RawXml: jsonRequest.XML, NfeXmlDocument: xmlParsed}
		return j.MakeJsonNfeIsNew(true), nil
	}

	err = ToJsonFile(j.Filename, j)
	if err != nil {
		panic(err)
	}

	return j.MakeJsonNfeIsNew(false), nil
}

func (j *inJsonNfeStore) GetNfeById(id string) ([]byte, error) {
	err := FromJsonFile(j.Filename, j)
	if err != nil {
		panic(err)
	}

	if j.AssertIdIsNew(id) {
		return []byte(""), fmt.Errorf("NFe with id %s not found", id)
	}

	return json.Marshal(j.Store[id].NfeXmlDocument)
}

func (j *inJsonNfeStore) MakeJsonNfeIsNew(status bool) []byte {
	data, err := json.Marshal(JsonNfeIsNew{IsNewNfe: status})
	if err != nil {
		return []byte("")
	}
	return data
}
