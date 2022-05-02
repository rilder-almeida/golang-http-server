package impltnfe

import (
	"fmt"

	"github.com/golang-http-server/entities/nfe"
)

type nfeInMemoryRepository struct {
	store nfe.NfeDocumentList
}

func NewNfeInMemoryRepository() nfe.Repository {
	return &nfeInMemoryRepository{
		store: make([]nfe.NfeDocument, 0),
	}
}

func (repository *nfeInMemoryRepository) FindByID(id string) (nfe.NfeDocument, error) {
	for _, nfeDocument := range repository.store {
		if nfeDocument.NfeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nil
		}
	}

	return nfe.NfeDocument{}, fmt.Errorf("NFe not found")
}

func (repository *nfeInMemoryRepository) Save(nfeDocument nfe.NfeDocument) error {
	repository.store = append(repository.store, nfeDocument)
	return nil
}
