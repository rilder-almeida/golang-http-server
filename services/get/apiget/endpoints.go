package apiget

import (
	"context"

	"github.com/golang-http-server/entities/xml"
	"github.com/golang-http-server/services/get"

	"github.com/go-kit/kit/endpoint"
)

type GetEndpointResponse struct {
	NFeXmlDocument xml.XmlDocument
}

type GetEndpointRequest struct {
	Id string
}

func MakeAPIGetEndpoint(service get.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetEndpointRequest)
		data, err := service.Get(TranslateToServiceRequest(req))

		return TranslateToEndpointResponse(data), err
	}
}

func TranslateToServiceRequest(req GetEndpointRequest) get.Request {
	return get.Request{
		Id: req.Id,
	}
}

func TranslateToEndpointResponse(data get.Response) GetEndpointResponse {
	return GetEndpointResponse{
		NFeXmlDocument: data.NFeXmlDocument,
	}
}
