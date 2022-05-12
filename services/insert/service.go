package insert

import (
	"errors"

	customerrors "github.com/golang-http-server/entities/errors"
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
		return Response{}, customerrors.New("FAILED_INSERT_XML", "Xml can not be processed by the INSERT gateway", err)
	}
	return response, nil
}

// assert that the http.request is valid and can be processed
func (s *service) validateRequest(request *Request) error {
	if request.XML == "" {
		return customerrors.New("XML_IS_EMPTY", "Xml field cannot be empty", errors.New("XML is empty"))
	}
	return nil
}
