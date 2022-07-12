package apiinsert

import (
	"context"

	"github.com/golang-http-server/services/insert"

	"github.com/arquivei/foundationkit/contextmap"
	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog"
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
		data, err := service.Insert(ctx, TranslateToServiceRequest(req))

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

func (r InsertEndpointRequest) EnrichLog(
	ctx context.Context,
	zctx zerolog.Context,
) zerolog.Context {
	ctxMap := contextmap.Ctx(ctx)
	return zctx.
		Str("contextmap", ctxMap.String()).
		Str("xml", r.XML)
}

func (r InsertEndpointResponse) EnrichLog(
	ctx context.Context,
	zctx zerolog.Context,
) zerolog.Context {
	ctxMap := contextmap.Ctx(ctx)
	return zctx.
		Str("contextmap", ctxMap.String()).
		Bool("isnewnfe", r.IsNewNFe)
}
