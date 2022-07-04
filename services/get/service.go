package get

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
		return Response{}, err
	}
	return response, nil
}

func (s *service) validateRequest(request *Request) error {
	return nil
}
