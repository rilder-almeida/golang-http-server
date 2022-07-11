package insert

import "context"

type Service interface {
	Insert(context.Context, Request) (Response, error)
}

type service struct {
	insertGateway InsertGateway
}

func NewService(insertGateway InsertGateway) Service {
	return &service{
		insertGateway: insertGateway,
	}
}

func (s *service) Insert(ctx context.Context, request Request) (Response, error) {

	err := s.validateRequest(&request)
	if err != nil {
		return Response{}, err
	}

	response, err := s.insertGateway.Processor(ctx, request)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func (s *service) validateRequest(request *Request) error {
	return nil
}
