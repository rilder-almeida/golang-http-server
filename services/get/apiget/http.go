package apiget

import (
	"context"
	"encoding/json"
	"net/http"

	fkerrors "github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/golang-http-server/services/get/apiget/internal/v1"
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
	router.Handle("/nfe/get", httpHandler).Methods("GET")
	return router
}

func decodeHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	const op = fkerrors.Op("apiget.decodeHTTPRequest")

	var httpRequest v1.GetHTTPRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&httpRequest.Body)
	if err != nil {
		return nil, fkerrors.E(op, err, v1.ErrCodeInvalidRequest)
	}
	return translateToEndpointRequest(httpRequest)
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	const op = fkerrors.Op("apiget.encodeHTTPResponse")

	httpResponse := translateToHTTPResponse(response.(GetEndpointResponse))
	err := shared.EncodeJSONResponse(w, httpResponse)
	if err != nil {
		return fkerrors.E(op, err, v1.ErrCodeInvalidResponse)
	}
	return nil
}

func translateToEndpointRequest(httpRequest v1.GetHTTPRequest) (GetEndpointRequest, error) {
	err := v1.ValidateGetHTTPRequest(httpRequest)
	if err != nil {
		return GetEndpointRequest{}, err
	}
	return GetEndpointRequest{
		Id: httpRequest.Body.Id,
	}, nil
}

func translateToHTTPResponse(endpointResponse GetEndpointResponse) v1.GetHTTPResponse {
	return v1.GetHTTPResponse{
		Body: endpointResponse,
	}
}
