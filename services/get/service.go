package get

import (
	"errors"

	customErrors "github.com/golang-http-server/entities/errors"
)

type Service interface {
	Get(Request) (Response, error)
}

type service struct {
	getGateway GetGateway
}

func NewService(getGateway GetGateway) Service {
	return &service{
		getGateway: getGateway,
	}
}

func (s *service) Get(request Request) (Response, error) {

	err := s.validateRequest(&request)
	if err != nil {
		return Response{}, err
	}

	response, err := s.getGateway.Processor(request)
	if err != nil {
		return Response{}, customErrors.New("FAILED_GET_NFE", "ID can not be processed by the GET gateway", err)
	}
	return response, nil
}

// assert that the http.request is valid and can be processed
func (s *service) validateRequest(request *Request) error {
	if request.Id == "" {
		return customErrors.New("ID_IS_EMPTY", "Id field cannot be empty", errors.New("ID is empty"))
	}
	return nil
}
