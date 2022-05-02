package implinsert

import (
	"fmt"

	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/shared"
)

type inFileAdapter struct {
	repository nfe.Repository
}

func NewInFileAdapter(repository nfe.Repository) insert.InsertGateway {
	return &inFileAdapter{
		repository: repository,
	}
}

func (adapter *inFileAdapter) Processor(request insert.Request) (insert.Response, error) {
	err := adapter.receiver(request)
	if err != nil {
		if err == fmt.Errorf("NFe already exists") {
			return adapter.responder(false), nil
		}
		return insert.Response{}, err
	}
	return adapter.responder(true), nil
}

func (adapter *inFileAdapter) receiver(request insert.Request) error {
	xmlDocument, err := shared.ToXmlDocument([]byte(request.XML))
	if err != nil {
		return err
	}

	_, err = adapter.repository.FindByID(xmlDocument.NFe.InfNFe.Id)
	if err != fmt.Errorf("NFe not found") {
		return err
	}
	if err == nil {
		return fmt.Errorf("NFe already exists")
	}

	nfeDocument := nfe.NfeDocument{
		RawXml:         request.XML,
		NfeXmlDocument: xmlDocument,
	}

	err = adapter.repository.Save(nfeDocument)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *inFileAdapter) responder(isNewNfe bool) insert.Response {
	return insert.Response{
		IsNewNfe: isNewNfe,
	}
}
