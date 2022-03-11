package main

import (
	"fmt"
	"net/http"
	"os"
)

type NfeStore interface {
	//post method
	PostRequestReceiver(JsonPostRequest) (JsonPostResponse, error)
	PostRequestResponder(string, []byte, int) JsonPostResponse
	//get method
	// GetRequestReceiver(JsonGetRequest) (JsonGetResponse, error)
	// GetRequestResponder(string, []byte, int) (JsonGetResponse, error)
	//assert utils
	AssertIdIsNew(string) bool
	//repository utils
	StoreNfe(JsonPostRequest) ([]byte, error)
	GetNfeById()
	MakeJsonNfeIsNew(bool) []byte
}

type NfeServer struct {
	storage NfeStore
}
type JsonPostRequest struct {
	XML string `json:"XML"`
}
type JsonPostResponse struct {
	contentType string
	bodyData    []byte
	httpStatus  int
}

func NewServer() *NfeServer {
	return &NfeServer{
		storage: getNfeStore(),
	}
}

func getNfeStore() NfeStore {
	envVar := os.Getenv("NFE_STORE")
	switch envVar {
	case "IN_MEMORY":
		return NewInMemoryNfeStore()
		// case "JSON": // TODO: must be implemented
		// 	return NewInJsonNfeStore()
	}
	panic(fmt.Sprintf("Bad environment variable value: %s", envVar))
}

func (n *NfeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		n.processPostRequest(w, r)
	// case http.MethodGet: // TODO: must be implemented
	// 	n.processGetRequest(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (n *NfeServer) processPostRequest(w http.ResponseWriter, r *http.Request) {
	url := urlParser(r.URL.Path, r.Method)
	if url != "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Bad URL: %s", url)
		return
	}

	body, err := requestBodyParser(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Printf("Error parsing request body: %s", err)
		return
	}

	jsonRequest := JsonPostRequest{}
	err = jsonRequestParser(body, &jsonRequest)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	response, err := n.storage.PostRequestReceiver(jsonRequest)
	if err != nil {
		fmt.Printf("Error storing NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = n.dispatchPostResponse(w, response)
	if err != nil {
		fmt.Printf("Error dispatch response NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (n *NfeServer) dispatchPostResponse(w http.ResponseWriter, response JsonPostResponse) error {
	w.Header().Set("Content-Type", response.contentType)
	w.WriteHeader(response.httpStatus)
	_, err := w.Write(response.bodyData)
	if err != nil {
		return err
	}
	return nil
}

// func (n *NfeServer) processGetRequest(w http.ResponseWriter, r *http.Request) {} // TODO: must be implemented

// func (n *NfeServer) dispatchGetResponse() {} // TODO: must be implemented
