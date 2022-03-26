package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubNfeStoreInMemory struct {
	store     map[string]NfeDocument
	postCalls []string
	getCalls  []string
}

type StubNfeDocument struct {
	RawXml         string      `json:"raw_xml"`
	NfeXmlDocument XmlDocument `json:"nfe_xml_document"`
}

// teste 2

func (s *StubNfeStoreInMemory) PostRequestReceiver(j JsonPostRequest) (JsonResponse, error) {
	s.postCalls = append(s.postCalls, j.XML)
	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(j.XML), http.StatusOK)
}

func (s *StubNfeStoreInMemory) GetRequestReceiver(j JsonGetRequest) (JsonResponse, error) {
	s.getCalls = append(s.getCalls, j.Id)
	return s.RequestResponder(GET_RESPONSE_CONTENT_TYPE, []byte(j.Id), http.StatusOK)
}

func (s *StubNfeStoreInMemory) RequestResponder(contentType string, bodyData []byte, httpStatus int) JsonResponse {
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

func TestPOSTNfeinMemory(t *testing.T) {
	store := StubNfeStoreInMemory{
		map[string]NfeDocument{},
		nil,
		nil,
	}
	server := &NfeServer{&store}

	t.Run("returns http status code 201", func(t *testing.T) {
		request := newPostNfeRequest("<xml>")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)
	})

	t.Run("returns the content type as application/json", func(t *testing.T) {
		request := newPostNfeRequest("<xml>")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertContentType(t, response.Header().Get("Content-Type"), "application/json")
	})

	t.Run("creates a new nfe document", func(t *testing.T) {
		request := newPostNfeRequest("<xml>")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)
		assertResponseBody(t, response.Body.String(), `{"raw_xml":"<xml>","nfe_xml_document":{"XMLName":{"Space":"","Local":"nfeProc"},"NFe":{"XMLName":{"Space":"","Local":"NFe"},"InfNFe":{"XMLName":{"Space":"","Local":"infNFe"},"Id":"","Emit":{"XMLName":{"Space":"","Local":"emit"},"CNPJ":""},"Total":{"XMLName":{"Space":"","Local":"total"},"ICMSTot":{"XMLName":{"Space":"","Local":"ICMSTot"},"VNF":""}}}}}}`)
	})

	t.Run("returns a 500 internal server error if the xml is invalid", func(t *testing.T) {
		request := newPostNfeRequest("<xml>")
		response := httptest.NewRecorder()

	})
}
