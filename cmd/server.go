package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/nfe/impltnfe"
	"github.com/golang-http-server/services/get"
	"github.com/golang-http-server/services/get/implget"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/services/insert/implinsert"
)

type NfeServer struct {
	GetService    get.Service
	InsertService insert.Service
}

func NewRepository(config impltnfe.Config) nfe.Repository {
	return impltnfe.NewNfeRepository(config)
}

func NewServer() *NfeServer {
	return &NfeServer{
		GetService:    get.NewService(implget.NewAdapter(NewRepository(Config))),
		InsertService: insert.NewService(implinsert.NewAdapter(NewRepository(Config))),
	}
}

func (n *NfeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		n.processInsertService(w, r)
	case http.MethodGet:
		n.processGetService(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (n *NfeServer) processInsertService(w http.ResponseWriter, r *http.Request) {
	body, err := requestBodyReader(r.Body)
	if err != nil {
		fmt.Printf("Error parsing request body: %s", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	httpMessage := httpmessage.HttpMessage{
		BodyData:    body,
		ContentType: r.Header.Get("Content-Type"),
	}

	parsedRequest, err := insert.HttpMessageToRequest(httpMessage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error parsing request body: %s", err)
		return
	}

	response, err := n.InsertService.Insert(parsedRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error processing request: %s", err)
		return
	}

	parsedResponse, err := insert.ResponseToHttpMessage(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error parsing response: %s", err)
		return
	}

	w.WriteHeader(parsedResponse.HttpStatus)
	w.Write(parsedResponse.BodyData)
	w.Header().Set("Content-Type", parsedResponse.ContentType)
}

func (n *NfeServer) processGetService(w http.ResponseWriter, r *http.Request) {
	body, err := requestBodyReader(r.Body)
	if err != nil {
		fmt.Printf("Error parsing request body: %s", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	httpMessage := httpmessage.HttpMessage{
		BodyData:    body,
		ContentType: r.Header.Get("Content-Type"),
	}

	parsedRequest, err := get.HttpMessageToRequest(httpMessage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error parsing request body: %s", err)
		return
	}

	response, err := n.GetService.Get(parsedRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error processing request: %s", err)
		return
	}

	parsedResponse, err := get.ResponseToHttpMessage(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error parsing response: %s", err)
		return
	}

	w.WriteHeader(parsedResponse.HttpStatus)
	w.Write(parsedResponse.BodyData)
	w.Header().Set("Content-Type", parsedResponse.ContentType)
}

func requestBodyReader(bodyRequest io.ReadCloser) ([]byte, error) {
	body64, err := ioutil.ReadAll(bodyRequest)
	defer bodyRequest.Close()

	if err != nil {
		return nil, err
	}
	return body64, nil
}
