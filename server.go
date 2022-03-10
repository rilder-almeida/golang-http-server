package main

import (

	// "encoding/xml"

	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type NfeStore interface {
	//post method
	ReceiverNfe(XmlDocument) error
	ResponserNfe(XmlDocument) (JsonPostResponse, error)
	AssertNfe()
	//get method
	ReceiveId()
	ResponseId()
	AssertId()
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
	case "JSON":
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

	body, err := requestBodyParser(r)
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

	// TODO: xmlParser should be call inside *_nfe_store files
	xmlParsed, err := xmlParser([]byte(jsonRequest.XML))
	if err != nil {
		fmt.Printf("Error parsing XML: %s", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// TODO: ReceiverNfe should receive a JsonPostRequest and return a JsonPostResponse and a error from ResponserNfe (response, err)
	err = n.storage.ReceiverNfe(xmlParsed)
	if err != nil {
		fmt.Printf("Error storing NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: ResponserNfe should be call by ReceriverNfe return and returns a JsonPostResponse and a error
	response, err := n.storage.ResponserNfe(xmlParsed)
	if err != nil {
		fmt.Printf("Error receive response NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// fmt.Println("Id: ", xmlParsed.NFe.InfNFe.Id)
	// fmt.Println("CNPJ: ", xmlParsed.NFe.InfNFe.Emit.CNPJ)
	// fmt.Println("Total: ", xmlParsed.NFe.InfNFe.Total.ICMSTot.VNF)

	err = n.dispatchPostResponse(response)
	if err != nil {
		fmt.Printf("Error dispatch response NFe: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (n *NfeServer) dispatchPostResponse(response JsonPostResponse) error { return nil }

func (n *NfeServer) processGetRequest(w http.ResponseWriter, r *http.Request) {
	n.dispatchGetResponse()
}

func (n *NfeServer) dispatchGetResponse() {}

// TODO: move to a separate file

func urlParser(url string, method string) string {
	switch method {
	case "POST":
		return strings.TrimPrefix(url, "/nfe/v1/")

	case "GET":
		return strings.TrimPrefix(url, "/nfe/v1/")
	}
	panic(fmt.Sprintf("Bad method: %s", method))
}

func requestBodyParser(r *http.Request) ([]byte, error) {
	body64, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body64, nil
}

func jsonRequestParser(body []byte, v interface{}) error {
	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return nil
}

func xmlParser(body64 []byte) (XmlDocument, error) {
	var xmlDoc XmlDocument

	body, err := base64.StdEncoding.DecodeString(string(body64))
	if err != nil {
		return xmlDoc, err
	}

	err = xml.Unmarshal(body, &xmlDoc)
	if err != nil {
		return xmlDoc, err
	}
	return xmlDoc, nil
}
