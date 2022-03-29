package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type StubNfeStoreInMemory struct {
	store        map[string]NfeDocument
	postRequests [][]byte
	getRequests  [][]byte
	responses    [][]byte
}

func (s *StubNfeStoreInMemory) PostRequestReceiver(jsonRequest JsonPostRequest) (JsonResponse, error) {
	bodyData, err := s.StoreNfe(jsonRequest)

	s.postRequests = append(s.postRequests, bodyData)

	if err != nil {
		return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotAcceptable), err
	}

	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, bodyData, http.StatusOK), nil

}

func (s *StubNfeStoreInMemory) GetRequestReceiver(jsonRequest JsonGetRequest) (JsonResponse, error) {
	nfeDoc, err := s.GetNfeById(jsonRequest.Id)

	s.getRequests = append(s.getRequests, nfeDoc)

	if err != nil {
		return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotFound), err
	}

	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, nfeDoc, http.StatusOK), nil

}

func (s *StubNfeStoreInMemory) RequestResponder(contentType string, bodyData []byte, httpStatus int) JsonResponse {
	s.responses = append(s.responses, bodyData)
	return JsonResponse{
		contentType: contentType,
		bodyData:    bodyData,
		httpStatus:  httpStatus,
	}
}

func (s *StubNfeStoreInMemory) AssertIdIsNew(id string) bool {
	_, ok := s.store[id]
	return !ok
}

func (s *StubNfeStoreInMemory) StoreNfe(jsonRequest JsonPostRequest) ([]byte, error) {
	xmlParsed, err := ToXmlParser([]byte(jsonRequest.XML))
	if err != nil {
		return []byte(""), fmt.Errorf("error parsing XML: %s", err)
	}

	if s.AssertIdIsNew(xmlParsed.NFe.InfNFe.Id) {
		s.store[xmlParsed.NFe.InfNFe.Id] = NfeDocument{RawXml: jsonRequest.XML, NfeXmlDocument: xmlParsed}
		return s.MakeJsonNfeIsNew(true), nil
	}

	return s.MakeJsonNfeIsNew(false), nil
}

func (s *StubNfeStoreInMemory) GetNfeById(id string) ([]byte, error) {
	if s.AssertIdIsNew(id) {
		return []byte(""), fmt.Errorf("NFe with id %s not found", id)
	}

	return json.Marshal(s.store[id].NfeXmlDocument)
}

func (s *StubNfeStoreInMemory) MakeJsonNfeIsNew(status bool) []byte {
	data, err := json.Marshal(JsonNfeIsNew{IsNewNfe: status})
	if err != nil {
		return []byte("")
	}
	return data
}

// server requests

func newFakeRequest(method string, url string, body string) *http.Request {
	jsonData := bytes.NewBuffer([]byte(body))

	req, err := http.NewRequest(method, url, jsonData)
	if err != nil {
		panic(err)
	}
	return req
}

// TODO implement assertitions

func assertStatus(t testing.TB, got, want int) {}

func assertRequestBody(t testing.TB, got, want string) {}

func assertResponseBody(t testing.TB, got, want string) {}

func TestPOSTNfeinMemory(t *testing.T) {
	store := StubNfeStoreInMemory{
		map[string]NfeDocument{},
		nil,
		nil,
		nil,
	}
	server := &NfeServer{&store}

	// create table tests here

	// t.Run tests here
}
