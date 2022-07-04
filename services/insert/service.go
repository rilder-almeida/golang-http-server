package insert

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

func (s *service) validateRequest(request *Request) error {
	return nil
}
