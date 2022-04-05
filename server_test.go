package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "embed"
)

const (
	URL_API = "http://localhost:8080/nfe/v1"
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

type StubGetRequestJson struct {
	Id string `json:"id"`
}

type fakeRequest struct {
	method string
	url    string
	body   []byte
}

type tableTest struct {
	intend                 string
	fakeRequests           []fakeRequest
	expectedResponseBody   []string
	expectedResponseStatus []int
}

//go:embed test_data/valid_nfe.xml
var valid_nfe_xml []byte

//go:embed test_data/invalid_nfe.xml
var invalid_nfe_xml []byte

//go:embed test_data/existent_valid_id.json
var existent_valid_id_json []byte

//go:embed test_data/non_existent_valid_id.json
var non_existent_valid_id_json []byte

//go:embed test_data/invalid_id.json
var invalid_id_json []byte

// StubNfeStoreInMemory methods

func (s *StubNfeStoreInMemory) PostRequestReceiver(jsonRequest JsonPostRequest) (JsonResponse, error) {
	s.postRequests += 1

	bodyData, err := s.StoreNfe(jsonRequest)
	if err != nil {
		return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotAcceptable), err
	}

	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, bodyData, http.StatusOK), nil
}

func (s *StubNfeStoreInMemory) GetRequestReceiver(jsonRequest JsonGetRequest) (JsonResponse, error) {
	s.getRequests += 1

	nfeDoc, err := s.GetNfeById(jsonRequest.Id)

	if err != nil {
		return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, []byte(""), http.StatusNotFound), err
	}

	return s.RequestResponder(POST_RESPONSE_CONTENT_TYPE, nfeDoc, http.StatusOK), nil

}

func (s *StubNfeStoreInMemory) RequestResponder(contentType string, bodyData []byte, httpStatus int) JsonResponse {
	s.responses += 1

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

func newFakeRequest(request fakeRequest) *http.Request {

	switch request.method {
	case "POST":
		stubRequest := getStubPostRequestJson(request)
		jsonDataBytes, err := json.Marshal(stubRequest)
		if err != nil {
			fmt.Println(err)
		}
		req, err := http.NewRequest(request.method, request.url, bytes.NewBuffer(jsonDataBytes))
		if err != nil {
			fmt.Println(err)
		}
		return req
	case "GET":
		stubRequest := getStubGetRequestJson(request)
		jsonDataBytes, err := json.Marshal(stubRequest)
		if err != nil {
			fmt.Println(err)
		}
		req, err := http.NewRequest(request.method, request.url, bytes.NewBuffer(jsonDataBytes))
		if err != nil {
			fmt.Println(err)
		}
		return req
	default:
		return httptest.NewRequest(request.method, request.url, nil)
	}
}

func getStubPostRequestJson(request fakeRequest) StubPostRequestJson {
	return StubPostRequestJson{XML: base64.StdEncoding.EncodeToString(request.body)}
}

func getStubGetRequestJson(request fakeRequest) StubGetRequestJson {
	var stubRequest StubGetRequestJson
	err := json.Unmarshal(request.body, &stubRequest)
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

var tableTests = []tableTest{
	{
		intend: "single POST request with valid XML",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: valid_nfe_xml},
		},
		expectedResponseBody:   []string{`{"IsNewNfe":true}`},
		expectedResponseStatus: []int{http.StatusOK},
	},
	{
		intend: "multiple POST requests with valid XML",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: valid_nfe_xml},
			{method: "POST", url: URL_API, body: valid_nfe_xml},
			{method: "POST", url: URL_API, body: valid_nfe_xml},
		},
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
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
		},
		expectedResponseBody:   []string{``},
		expectedResponseStatus: []int{http.StatusInternalServerError},
	},
	{
		intend: "multiple POST requests with invalid XML",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
		},
		expectedResponseBody: []string{
			``,
			``,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusInternalServerError,
			http.StatusInternalServerError,
			http.StatusInternalServerError,
		},
	},
	{
		intend: "single POST request with invalid XML and valid XML",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
			{method: "POST", url: URL_API, body: valid_nfe_xml},
		},
		expectedResponseBody: []string{
			``,
			`{"IsNewNfe":true}`,
		},
		expectedResponseStatus: []int{
			http.StatusInternalServerError,
			http.StatusOK,
		},
	},
	{
		intend: "single POST request with valid XML and invalid XML",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: valid_nfe_xml},
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
		},
		expectedResponseBody: []string{
			`{"IsNewNfe":true}`,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusOK,
			http.StatusInternalServerError,
		},
	},
	{
		intend: "single POST request with valid XML and single GET request with valid and existent Id",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: valid_nfe_xml},
			{method: "GET", url: URL_API, body: existent_valid_id_json},
		},
		expectedResponseBody: []string{
			`{"IsNewNfe":true}`,
			`{"NFe":{"InfNFe":{"Id":"NFe35200664902000000160550010000119461000000014","Emit":{"CNPJ":"64902000000160"},"Total":{"ICMSTot":{"VNF":"5686.00"}}}}}`,
		},
		expectedResponseStatus: []int{
			http.StatusOK,
			http.StatusOK,
		},
	},
	{
		intend: "single POST request with valid XML and single GET request with valid and non-existent Id",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: valid_nfe_xml},
			{method: "GET", url: URL_API, body: non_existent_valid_id_json},
		},
		expectedResponseBody: []string{
			`{"IsNewNfe":true}`,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusOK,
			http.StatusInternalServerError,
		},
	},
	{
		intend: "single POST request with valid XML and single GET request with invalid Id",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: valid_nfe_xml},
			{method: "GET", url: URL_API, body: invalid_id_json},
		},
		expectedResponseBody: []string{
			`{"IsNewNfe":true}`,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusOK,
			http.StatusInternalServerError,
		},
	},
	{
		intend: "single POST request with invalid XML and single GET request with valid Id",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
			{method: "GET", url: URL_API, body: existent_valid_id_json},
		},
		expectedResponseBody: []string{
			``,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusInternalServerError,
			http.StatusInternalServerError,
		},
	},
	{
		intend: "single POST request with invalid XML and single GET request with invalid Id",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
			{method: "GET", url: URL_API, body: invalid_id_json},
		},
		expectedResponseBody: []string{
			``,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusInternalServerError,
			http.StatusInternalServerError,
		},
	},
	{
		intend: "single POST request with invalid XML and single GET request with valid and non-existent Id",
		fakeRequests: []fakeRequest{
			{method: "POST", url: URL_API, body: invalid_nfe_xml},
			{method: "GET", url: URL_API, body: non_existent_valid_id_json},
		},
		expectedResponseBody: []string{
			``,
			``,
		},
		expectedResponseStatus: []int{
			http.StatusInternalServerError,
			http.StatusInternalServerError,
		},
	},
}

// NfeinMemoryServer tests

func TestNfeinMemory(t *testing.T) {
	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.intend, func(t *testing.T) {
			store := &StubNfeStoreInMemory{make(map[string]NfeDocument), 0, 0, 0}
			server := &NfeServer{store}
			for j := 0; j < len(tt.fakeRequests); j++ {

				response := httptest.NewRecorder()
				request := newFakeRequest(tt.fakeRequests[j])
				server.ServeHTTP(response, request)

				assertResponseBody(t, response.Body.String(), tt.expectedResponseBody[j])
				assertStatus(t, response.Code, tt.expectedResponseStatus[j])
			}

			assertRequestTimes(t, store.postRequests+store.getRequests, len(tt.fakeRequests))
			assertRequestTimes(t, store.responses, len(tt.fakeRequests))

		})

	}
}
