package services

import (
	"github.com/golang-http-server/entities"
)

const (
	POST_RESPONSE_CONTENT_TYPE = "application/json"
	JSON_FILENAME              = "store.json"
)

type NfeStore interface {
	PostRequestReceiver(JsonPostRequest) (JsonResponse, error)
	GetRequestReceiver(JsonGetRequest) (JsonResponse, error)
	RequestResponder(string, []byte, int) JsonResponse
	//utils
	AssertIdIsNew(string) bool
	StoreNfe(JsonPostRequest) ([]byte, error)
	GetNfeById(string) ([]byte, error)
	MakeJsonNfeIsNew(bool) []byte
}

type JsonResponse struct {
	contentType string
	bodyData    []byte
	httpStatus  int
}
