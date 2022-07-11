package apiinsert

import (
	"context"
	"encoding/json"
	"net/http"

	fkerrors "github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/golang-http-server/services/insert/apiinsert/internal/v1"
	"github.com/golang-http-server/shared"
)

func MakeHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	httpHandler := kithttp.NewServer(
		endpoint,
		decodeHTTPRequest,
		encodeHTTPResponse,
		v1.GetHTTPServerOption()...,
	)

	router := mux.NewRouter()
	router.Handle("/nfe/insert", httpHandler).Methods("POST")
	return router
}

func decodeHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	const op = fkerrors.Op("apiinsert.decodeHTTPRequest")

	var httpRequest v1.InsertHTTPRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&httpRequest.Body)
	if err != nil {
		return nil, fkerrors.E(op, err, v1.ErrCodeInvalidRequest)
	}
	return translateToEndpointRequest(httpRequest)
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	const op = fkerrors.Op("apiinsert.encodeHTTPResponse")

	httpResponse := translateToHTTPResponse(response.(InsertEndpointResponse))
	err := shared.EncodeJSONResponse(w, httpResponse)
	if err != nil {
		return fkerrors.E(op, err, v1.ErrCodeInvalidResponse)
	}
	return nil
}

func translateToEndpointRequest(httpRequest v1.InsertHTTPRequest) (InsertEndpointRequest, error) {
	err := v1.ValidateInsertHTTPRequest(httpRequest)
	if err != nil {
		return InsertEndpointRequest{}, err
	}

	return InsertEndpointRequest{
		XML: httpRequest.Body.XML,
	}, nil
}

func translateToHTTPResponse(endpointResponse InsertEndpointResponse) v1.InsertHTTPResponse {
	return v1.InsertHTTPResponse{
		Body: endpointResponse,
	}
}
