package insert

import (
	"errors"

	customErrors "github.com/golang-http-server/entities/errors"
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
		return Response{}, customErrors.Error{
			ErrorCode:        "FAILED_INSERT_XML",
			Message:          "Xml can not be processed by the INSERT gateway",
			ApplicationError: err,
		}
	}
	return response, nil
}

// assert that the http.request is valid and can be processed
func (s *service) validateRequest(request *Request) error {
	if request.XML == "" {
		return customErrors.Error{
			ErrorCode:        "XML_IS_EMPTY",
			Message:          "Xml field cannot be empty",
			ApplicationError: errors.New("XML is empty"),
		}
	}
	return nil
}
