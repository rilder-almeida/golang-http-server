package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
)

type nfeInMemoryRepository struct {
	store nfe.NFeDocuments
}

func NewNFeInMemoryRepository() nfe.Repository {
	return &nfeInMemoryRepository{
		store: make(nfe.NFeDocuments, 0),
	}
}

func (repository *nfeInMemoryRepository) FindByID(id string) (nfe.NFeDocument, error) {
	for _, nfeDocument := range repository.store {
		if nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nfe.ErrAlreadyExists
		}
	}

	return nfe.NFeDocument{}, nfe.ErrNotFound
}

func (repository *nfeInMemoryRepository) Save(nfeDocument nfe.NFeDocument) error {
	repository.store = append(repository.store, nfeDocument)
	return nil
}
