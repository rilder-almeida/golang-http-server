package apiget

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/golang-http-server/services/get/apiget/internal"
	"github.com/golang-http-server/shared"
	"github.com/gorilla/mux"
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
	var httpRequest internal.GetHTTPRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&httpRequest.Body)
	if err != nil {
		// TODO: http custom error
		return nil, err
	}
	return translateToEndpointRequest(httpRequest), nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	httpResponse := translateToHTTPResponse(response.(GetEndpointResponse))
	err := shared.EncodeJSONResponse(w, httpResponse)
	if err != nil {
		// TODO: http custom error
		return err
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
