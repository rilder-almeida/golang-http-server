package apiinsert

import (
	"context"

	"github.com/golang-http-server/services/insert"

	"github.com/go-kit/kit/endpoint"
)

type InsertEndpointResponse struct {
	IsNewNFe bool
}

type InsertEndpointRequest struct {
	XML string
}

func MakeAPIInsertEndpoint(service insert.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(InsertEndpointRequest)
		data, err := service.Insert(TranslateToServiceRequest(req))

		return TranslateToEndpointResponse(data), err
	}
}

func TranslateToServiceRequest(req InsertEndpointRequest) insert.Request {
	return insert.Request{
		XML: req.XML,
	}
}

func TranslateToEndpointResponse(data insert.Response) InsertEndpointResponse {
	return InsertEndpointResponse{
		IsNewNFe: data.IsNewNFe,
	}
}
