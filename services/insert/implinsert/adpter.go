package implinsert

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/shared"
)

type Adapter struct {
	repository nfe.Repository
}

func NewAdapter(repository nfe.Repository) insert.InsertGateway {
	return &Adapter{
		repository: repository,
	}
}

func (adapter *Adapter) Processor(request insert.Request) (insert.Response, error) {
	err := adapter.receiver(request)
	if err != nil {
		if err == nfe.ErrAlreadyExists {
			return adapter.responder(false), nil
		}
		return insert.Response{}, err
	}
	return adapter.responder(true), nil
}

func (adapter *Adapter) receiver(request insert.Request) error {
	xmlDocument, err := shared.ToXmlDocument([]byte(request.XML))
	if err != nil {
		return err
	}

	_, err = adapter.repository.FindByID(xmlDocument.NFe.InfNFe.Id)
	if err != nfe.ErrNotFound {
		return err
	}
	if err == nfe.ErrAlreadyExists {
		return err
	}

	nfeDocument := nfe.NFeDocument{
		RawXml:         request.XML,
		NFeXmlDocument: xmlDocument,
	}

	err = adapter.repository.Save(nfeDocument)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *Adapter) responder(isNewNFe bool) insert.Response {
	return insert.Response{
		IsNewNFe: isNewNFe,
	}
}
