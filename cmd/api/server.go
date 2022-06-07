package main

import (
	"io"
	"io/ioutil"
	"net/http"

	customerrors "github.com/golang-http-server/entities/errors"
	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/nfe/impltnfe"
	"github.com/golang-http-server/services/get"
	"github.com/golang-http-server/services/get/apiget"
	"github.com/golang-http-server/services/get/implget"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/services/insert/apiinsert"
	"github.com/golang-http-server/services/insert/implinsert"
)

type NfeHandler struct {
	GetService    get.Service
	InsertService insert.Service
}

func NewRepository(config impltnfe.Config) nfe.Repository {
	return impltnfe.NewNfeRepository(config)
}

func NewHandler() *NfeHandler {
	repository := NewRepository(repositoryConfig)
	return &NfeHandler{
		// TODO LIMPAR O WRAPPER DE LOG E O ZEROLOG
		// TODO VERIFICAR A NECESSIDADE DE USAR O WRAPPER DE METRICS
		GetService:    get.WrapServiceWithMetrics(get.NewService(implget.WrapGetServiceWithLogging(implget.NewAdapter(repository)))),
		InsertService: insert.WrapServiceWithMetrics(insert.NewService(implinsert.WrapInsertServiceWithLogging(implinsert.NewAdapter(repository)))),
	}
}

func (n *NfeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		n.processInsertService(w, r)
	case http.MethodGet:
		n.processGetService(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}

func (n *NfeHandler) processInsertService(w http.ResponseWriter, r *http.Request) {
	body, err := requestBodyReader(r.Body)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	httpMessage := httpmessage.New(body, http.StatusOK)

	parsedRequest, err := apiinsert.HttpMessageToRequest(httpMessage)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	response, err := n.InsertService.Insert(parsedRequest)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	parsedResponse, err := apiinsert.ResponseToHttpMessage(response)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	responseDispatcher(w, httpmessage.New(parsedResponse.BodyData, http.StatusOK))
}

func (n *NfeHandler) processGetService(w http.ResponseWriter, r *http.Request) {
	body, err := requestBodyReader(r.Body)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	httpMessage := httpmessage.New(body, http.StatusOK)

	parsedRequest, err := apiget.HttpMessageToRequest(httpMessage)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	response, err := n.GetService.Get(parsedRequest)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	parsedResponse, err := apiget.ResponseToHttpMessage(response)
	if err != nil {
		responseDispatcher(w, httpmessage.New([]byte(err.Error()), StatusCode(err)))
		return
	}

	responseDispatcher(w, httpmessage.New(parsedResponse.BodyData, http.StatusOK))
}

func responseDispatcher(w http.ResponseWriter, httpMessage httpmessage.HttpMessage) {
	w.Header().Set("Content-Type", httpMessage.ContentType)
	w.WriteHeader(httpMessage.HttpStatus)
	w.Write(httpMessage.BodyData)
}

func requestBodyReader(bodyRequest io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(bodyRequest)
	defer bodyRequest.Close()

	if err != nil {
		return nil, customerrors.New("INVALID_REQUEST", "Can not read the request body", err)
	}
	return body, nil
}

func StatusCode(err error) int {
	parsedError, ok := err.(customerrors.Error)
	if !ok {
		return http.StatusInternalServerError
	}

	switch parsedError.ErrorCode {
	case "ID_IS_EMPTY":
		return http.StatusBadRequest // 400
	case "XML_IS_EMPTY":
		return http.StatusBadRequest
	case "INVALID_REQUEST":
		return http.StatusBadRequest
	case "INVALID_RESPONSE":
		return http.StatusInternalServerError // 500
	case "FAILED_INSERT_XML":
		return http.StatusInternalServerError
	case "FAILED_GET_NFE":
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
