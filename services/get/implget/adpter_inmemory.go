package implget

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/services/get"
)

type inMemoryAdapter struct {
	repository nfe.Repository
}

func NewInMemoryAdapter(repository nfe.Repository) get.GetGateway {
	return &inMemoryAdapter{
		repository: repository,
	}
}

func (adapter *inMemoryAdapter) Processor(request get.Request) (get.Response, error) {
	nfeDocument, err := adapter.receiver(request)
	if err != nil {
		return get.Response{}, err
	}
	return adapter.responder(nfeDocument), nil
}

func (adapter *inMemoryAdapter) receiver(request get.Request) (nfe.NfeDocument, error) {
	nfeDocument, err := adapter.repository.FindByID(request.Id)
	if err != nil {
		return nfe.NfeDocument{}, err
	}
	return nfeDocument, nil
}

func (adapter *inMemoryAdapter) responder(nfeDocument nfe.NfeDocument) get.Response {
	return get.Response{
		NfeDocument: nfeDocument,
	}
}
