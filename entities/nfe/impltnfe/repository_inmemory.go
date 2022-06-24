package impltnfe

import (
	"sync"

	"github.com/golang-http-server/entities/nfe"
)

type nfeMemoryRepository struct {
	store nfe.NFeDocuments
}

func NewNFeMemoryRepository() nfe.Repository {
	return &nfeMemoryRepository{
		store: make(nfe.NFeDocuments, 0),
	}
}

func (repository *nfeMemoryRepository) FindByID(id string) (nfe.NFeDocument, error) {
	for _, nfeDocument := range repository.store {
		if nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nfe.ErrAlreadyExists
		}
	}

	return nfe.NFeDocument{}, nfe.ErrNotFound
}

func (repository *nfeMemoryRepository) Save(nfeDocument nfe.NFeDocument) error {
	safeStore := struct {
		sync.RWMutex
		cache nfe.NFeDocuments
	}{
		cache: repository.store,
	}
	safeStore.Lock()
	defer safeStore.Unlock()
	repository.store = append(repository.store, nfeDocument)
	return nil
}
