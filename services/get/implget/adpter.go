package implget

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/services/get"
)

type Adapter struct {
	repository nfe.Repository
}

func NewAdapter(repository nfe.Repository) get.GetGateway {
	return &Adapter{
		repository: repository,
	}
}

func (adapter *Adapter) Processor(request get.Request) (get.Response, error) {
	nfeDocument, err := adapter.receiver(request)
	if err != nil {
		return get.Response{}, err
	}
	return adapter.responder(nfeDocument), nil
}

func (adapter *Adapter) receiver(request get.Request) (nfe.NFeDocument, error) {
	nfeDocument, err := adapter.repository.FindByID(request.Id)
	if err != nil {
		return nfeDocument, err
	}
	return nfeDocument, nil
}

func (adapter *Adapter) responder(nfeDocument nfe.NFeDocument) get.Response {
	return get.Response{
		NFeXmlDocument: nfeDocument.NFeXmlDocument,
	}
}
