package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubNfeStoreInMemory struct {
	store        map[string]NfeDocument
	postRequests int
	getRequests  int
	responses    int
}

type StubPostRequestJson struct {
	XML string `json:"XML"`
}

type tableTest struct {
	intend                 string
	request                *http.Request
	requestTimes           int
	expectedResponseBody   []string
	expectedResponseStatus []int
}

// StubNfeStoreInMemory methods

func (s *StubNfeStoreInMemory) PostRequestReceiver(jsonRequest JsonPostRequest) (JsonResponse, error) {
	bodyData, err := s.StoreNfe(jsonRequest)

	s.postRequests = +1

	if err != nil {
		return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotAcceptable), err
	}

	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, bodyData, http.StatusOK), nil

}

func (s *StubNfeStoreInMemory) GetRequestReceiver(jsonRequest JsonGetRequest) (JsonResponse, error) {
	nfeDoc, err := s.GetNfeById(jsonRequest.Id)

	s.getRequests = +1

	if err != nil {
		return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotFound), err
	}

	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, nfeDoc, http.StatusOK), nil

}

func (s *StubNfeStoreInMemory) RequestResponder(contentType string, bodyData []byte, httpStatus int) JsonResponse {
	s.responses = +1

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

func newFakePostRequest(method string, url string, jsonData StubPostRequestJson) *http.Request {

	jsonDataBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
	}

	data := bytes.NewBuffer(jsonDataBytes)

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		fmt.Println(err)
	}
	return req
}

func readFileToStubPostRequestJson(fileName string) (StubPostRequestJson, error) {
	var jsonData StubPostRequestJson
	jsonFileData32, err := ioutil.ReadFile(fileName)
	if err != nil {
		return jsonData, err
	}

	jsonFileData := []byte(base64.StdEncoding.EncodeToString(jsonFileData32))
	if err != nil {
		return jsonData, err
	}

	jsonData.XML = string(jsonFileData)

	return jsonData, nil
}

func getValidStubPostRequestJson() StubPostRequestJson {
	stubRequest, err := readFileToStubPostRequestJson("./test_data/valid_nfe.xml")
	if err != nil {
		fmt.Println(err)
	}
	return stubRequest
}

func getInvalidStubPostRequestJson() StubPostRequestJson {
	stubRequest, err := readFileToStubPostRequestJson("./test_data/invalid_nfe.xml")
	if err != nil {
		fmt.Println(err)
	}
	return stubRequest
}

// assertions

func assertStatus(t testing.TB, actual, expected int) {
	if actual != expected {
		t.Errorf("expected status %d, got %d", expected, actual)
	}
}

func assertResponseBody(t testing.TB, actual, expected string) {
	if actual != expected {
		t.Errorf("expected body %s, got %s", expected, actual)
	}
}

func assertRequestTimes(t testing.TB, actual, expected int) {
	if actual != expected {
		t.Errorf("expected request %d times, got %d", expected, actual)
	}
}

// Table tests

var validStubPostRequestJson = getValidStubPostRequestJson()

var invalidStubPostRequestJson = getInvalidStubPostRequestJson()

var postTableTests = []tableTest{
	{
		intend: "single POST request with valid XML",
		request: newFakePostRequest(
			"POST",
			"http://localhost:8080/nfe/v1",
			validStubPostRequestJson,
		),
		requestTimes:           1,
		expectedResponseBody:   []string{`{"IsNewNfe":true}`},
		expectedResponseStatus: []int{http.StatusOK},
	},
	{
		intend: "multiple POST requests with valid XML",
		request: newFakePostRequest(
			"POST",
			"http://localhost:8080/nfe/v1",
			validStubPostRequestJson,
		),
		requestTimes: 3,
		expectedResponseBody: []string{
			`{"IsNewNfe":true}`,
			`{"IsNewNfe":false}`,
			`{"IsNewNfe":false}`,
		},
		expectedResponseStatus: []int{
			http.StatusOK,
			http.StatusOK,
			http.StatusOK,
		},
	},
	{
		intend: "single POST request with invalid XML",
		request: newFakePostRequest(
			"POST",
			"http://localhost:8080/nfe/v1",
			invalidStubPostRequestJson,
		),
		requestTimes:           1,
		expectedResponseBody:   []string{``},
		expectedResponseStatus: []int{http.StatusInternalServerError},
	},
}

// NfeinMemoryServer tests

func TestPOSTNfeinMemory(t *testing.T) {
	for _, tt := range postTableTests {

		t.Run(tt.intend, func(t *testing.T) {
			store := &StubNfeStoreInMemory{make(map[string]NfeDocument), 0, 0, 0}
			server := &NfeServer{store}
			for j := 0; j < tt.requestTimes; j++ {
				response := httptest.NewRecorder()

				server.ServeHTTP(response, tt.request)

				assertResponseBody(t, response.Body.String(), tt.expectedResponseBody[j])

				assertStatus(t, response.Code, tt.expectedResponseStatus[j])
			}

			assertRequestTimes(t, store.postRequests, tt.requestTimes)

			assertRequestTimes(t, store.responses, tt.requestTimes)

		})

	}
}
