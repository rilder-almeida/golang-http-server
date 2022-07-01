package apiget

import (
	"context"
	"encoding/json"
	"net/http"

	fkerrors "github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/golang-http-server/services/get/apiget/internal"
	"github.com/golang-http-server/shared"
)

func MakeHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	httpHandler := kithttp.NewServer(
		endpoint,
		decodeHTTPRequest,
		encodeHTTPResponse,
	)

	router := mux.NewRouter()
	router.Handle("/nfe/get", httpHandler).Methods("GET")
	return router
}

func decodeHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	const op = fkerrors.Op("apiget.decodeHTTPRequest")

	var httpRequest internal.GetHTTPRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&httpRequest.Body)
	if err != nil {
		return nil, fkerrors.E(op, internal.ErrCodeInvalidRequest, fkerrors.KV("Decode", err))
	}
	return translateToEndpointRequest(httpRequest), nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	const op = fkerrors.Op("apiget.encodeHTTPResponse")

	httpResponse := translateToHTTPResponse(response.(GetEndpointResponse))
	err := shared.EncodeJSONResponse(w, httpResponse)
	if err != nil {
		return fkerrors.E(op, internal.ErrCodeInvalidResponse, fkerrors.KV("Encode", err))
	}
	return nil
}

func translateToEndpointRequest(httpRequest internal.GetHTTPRequest) GetEndpointRequest {
	return GetEndpointRequest{
		Id: httpRequest.Body.Id,
	}
}

func translateToHTTPResponse(endpointResponse GetEndpointResponse) internal.GetHTTPResponse {
	return internal.GetHTTPResponse{
		Body: endpointResponse,
	}
}
