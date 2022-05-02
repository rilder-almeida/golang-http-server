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

	err := s.prepareRequest(&request)
	if err != nil {
		return Response{}, err
	}

	return s.getGateway.Processor(request)
}

// assert that the http.request is valid and can be processed
func (s *service) prepareRequest(request *Request) error {
	return nil
}
