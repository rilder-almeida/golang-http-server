package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type NfeServer struct {
	storage NfeStore
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
	case "IN_JSON":
		return NewInJsonNfeStore()
	}
	panic(fmt.Sprintf("Bad environment variable value: %s", envVar))
}

func (n *NfeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		n.processPostRequest(w, r)
	case http.MethodGet:
		n.processGetRequest(w, r)
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

	body, err := requestBodyReader(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Printf("Error parsing request body: %s", err)
		return
	}

	jsonRequest := JsonPostRequest{}
	err = FromJsonRequestParser(body, &jsonRequest)
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

	err = n.dispatchResponse(w, response)
	if err != nil {
		fmt.Printf("Error dispatch response NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (n *NfeServer) dispatchResponse(w http.ResponseWriter, response JsonResponse) error {
	w.Header().Set("Content-Type", response.contentType)
	w.WriteHeader(response.httpStatus)
	_, err := w.Write(response.bodyData)
	if err != nil {
		return err
	}
	return nil
}

func (n *NfeServer) processGetRequest(w http.ResponseWriter, r *http.Request) {
	url := urlParser(r.URL.Path, r.Method)
	if url != "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Bad URL: %s", url)
		return
	}

	body, err := requestBodyReader(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Printf("Error parsing request body: %s", err)
		return
	}

	jsonRequest := JsonGetRequest{}
	err = FromJsonRequestParser(body, &jsonRequest)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	response, err := n.storage.GetRequestReceiver(jsonRequest)
	if err != nil {
		fmt.Printf("Error getting NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = n.dispatchResponse(w, response)
	if err != nil {
		fmt.Printf("Error dispatch response NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func urlParser(url string, method string) string {
	switch method {
	case "POST":
		return strings.TrimPrefix(url, "/nfe/v1")

	case "GET":
		return strings.TrimPrefix(url, "/nfe/v1")
	}
	panic(fmt.Sprintf("Bad method: %s", method))
}

func requestBodyReader(bodyRequest io.ReadCloser) ([]byte, error) {
	body64, err := ioutil.ReadAll(bodyRequest)
	if err != nil {
		return nil, err
	}
	return body64, nil
}
