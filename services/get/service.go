package get

import (
	"context"
)

type Service interface {
	Get(context.Context, Request) (Response, error)
}

type service struct {
	getGateway GetGateway
}

func NewService(getGateway GetGateway) Service {
	return &service{
		getGateway: getGateway,
	}
}

func (s *service) Get(ctx context.Context, request Request) (Response, error) {

	err := s.validateRequest(&request)
	if err != nil {
		return Response{}, err
	}

	response, err := s.getGateway.Processor(ctx, request)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func (s *service) validateRequest(request *Request) error {
	return nil
}

// TODO
// log.Ctx(ctx).UpdateContext(func(zc zerolog.Context) zerolog.Context {
// 	return zc.
// 		Str("command_id", commandID).
// 		Str("namespace", namespace).
// 		Str("access_key", accesskey).
// 		Str("msg_id", item.ID).
// 		Str("msg_trace_id", traceID).
// 		Str("msg_created_at", item.CreatedAt).
// 		Str("msg_source", item.Source).
// 		Str("msg_type", item.Type)
// })
//
// TODO -- PASSAR PARA WILL
