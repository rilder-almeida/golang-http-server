package insert

import "errors"

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

	err := s.prepareRequest(&request)
	if err != nil {
		return Response{}, err
	}

	return s.insertGateway.Processor(request)
}

// assert that the http.request is valid and can be processed
func (s *service) prepareRequest(request *Request) error {
	if request.XML == "" {
		return errors.New("XML is empty")
	}
	return nil
}
