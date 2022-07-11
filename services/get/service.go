package get

import "context"

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
