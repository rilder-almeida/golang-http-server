package implinsert

import (
	fkerrors "github.com/arquivei/foundationkit/errors"

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
	err, ok := adapter.receiver(request)
	if err != nil {
		return insert.Response{}, err
	}

	return adapter.responder(ok), nil
}

func (adapter *Adapter) receiver(request insert.Request) (error, bool) {
	const op = fkerrors.Op("implinsert.receiver")

	xmlDocument, err := shared.ToXmlDocument([]byte(request.XML))
	if err != nil {
		return fkerrors.E(op, fkerrors.KV("ToXmlDocument", err), insert.ErrCodeInvalidRequest), false
	}

	document, err := adapter.repository.FindByID(xmlDocument.NFe.InfNFe.Id)
	if err.(fkerrors.Error).Code == nfe.ErrCodeDocumentNotFound && document.IsEmpty() {
		nfeDocument := nfe.NFeDocument{
			RawXml:         request.XML,
			NFeXmlDocument: xmlDocument,
		}

		err = adapter.repository.Save(nfeDocument)
		if err != nil {
			return err, false
		}

		return nil, true
	}

	return err, false
}

func (adapter *Adapter) responder(isNewNFe bool) insert.Response {
	return insert.Response{
		IsNewNFe: isNewNFe,
	}
}
