package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
)

type nfeInMemoryRepository struct {
	store nfe.NfeDocuments
}

func NewNfeInMemoryRepository() nfe.Repository {
	return &nfeInMemoryRepository{
		store: make(nfe.NfeDocuments, 0),
	}
}

func (repository *nfeInMemoryRepository) FindByID(id string) (nfe.NfeDocument, error) {
	for _, nfeDocument := range repository.store {
		if nfeDocument.NfeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nfe.ErrAlreadyExists
		}
	}

	return nfe.NfeDocument{}, nfe.ErrNotFound
}

func (repository *nfeInMemoryRepository) Save(nfeDocument nfe.NfeDocument) error {
	repository.store = append(repository.store, nfeDocument)
	return nil
}
