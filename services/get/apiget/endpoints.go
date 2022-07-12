package apiget

import (
	"context"

	"github.com/arquivei/foundationkit/contextmap"
	"github.com/golang-http-server/entities/xml"
	"github.com/golang-http-server/services/get"
	"github.com/rs/zerolog"

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
		data, err := service.Get(ctx, TranslateToServiceRequest(req))

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

func (r GetEndpointRequest) EnrichLog(
	ctx context.Context,
	zctx zerolog.Context,
) zerolog.Context {
	ctxMap := contextmap.Ctx(ctx)
	return zctx.
		Str("contextmap", ctxMap.String())
}

func (r GetEndpointResponse) EnrichLog(
	ctx context.Context,
	zctx zerolog.Context,
) zerolog.Context {
	ctxMap := contextmap.Ctx(ctx)
	return zctx.
		Str("contextmap", ctxMap.String())
}
