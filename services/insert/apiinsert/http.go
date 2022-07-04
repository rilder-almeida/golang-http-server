package apiinsert

import (
	"context"
	"encoding/json"
	"net/http"

	fkerrors "github.com/arquivei/foundationkit/errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/golang-http-server/services/insert/apiinsert/internal"
	"github.com/golang-http-server/shared"
)

func MakeHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	httpHandler := kithttp.NewServer(
		endpoint,
		decodeHTTPRequest,
		encodeHTTPResponse,
		internal.GetHTTPServerOption()...,
	)

	router := mux.NewRouter()
	router.Handle("/nfe/insert", httpHandler).Methods("POST")
	return router
}

func decodeHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	const op = fkerrors.Op("apiinsert.decodeHTTPRequest")

	var httpRequest internal.InsertHTTPRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&httpRequest.Body)
	if err != nil {
		return nil, fkerrors.E(op, err, internal.ErrCodeInvalidRequest)
	}
	return translateToEndpointRequest(httpRequest)
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	const op = fkerrors.Op("apiinsert.encodeHTTPResponse")

	httpResponse := translateToHTTPResponse(response.(InsertEndpointResponse))
	err := shared.EncodeJSONResponse(w, httpResponse)
	if err != nil {
		return fkerrors.E(op, err, internal.ErrCodeInvalidResponse)
	}
	return nil
}

func translateToEndpointRequest(httpRequest internal.InsertHTTPRequest) (InsertEndpointRequest, error) {
	err := internal.ValidateInsertHTTPRequest(httpRequest)
	if err != nil {
		return InsertEndpointRequest{}, err
	}

	return InsertEndpointRequest{
		XML: httpRequest.Body.XML,
	}, nil
}

func translateToHTTPResponse(endpointResponse InsertEndpointResponse) internal.InsertHTTPResponse {
	return internal.InsertHTTPResponse{
		Body: endpointResponse,
	}
}
