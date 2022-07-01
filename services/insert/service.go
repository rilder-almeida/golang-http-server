package insert

import (
	fkerrors "github.com/arquivei/foundationkit/errors"
)

type Service interface {
	Insert(Request) (Response, error)
}

type service struct {
	insertGateway InsertGateway
}

func NewService(insertGateway InsertGateway) Service {
	return &service{
		insertGateway: insertGateway,
	}
}

func (s *service) Insert(request Request) (Response, error) {

	err := s.validateRequest(&request)
	if err != nil {
		return Response{}, err
	}

	response, err := s.insertGateway.Processor(request)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

// TODO verificar ser a validação do request deve ser feita na camada de serviço ou endipoint
func (s *service) validateRequest(request *Request) error {
	const op = fkerrors.Op("insert.validateRequest")
	if request.XML == "" {
		return fkerrors.E(op, ErrCodeInvalidRequest)
	}
	return nil
}
