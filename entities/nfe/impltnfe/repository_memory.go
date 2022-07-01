package impltnfe

import (
	"sync"

	fkerrors "github.com/arquivei/foundationkit/errors"
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
	const op = fkerrors.Op("impltnfe.memory.FindByID")

	for _, nfeDocument := range repository.store {
		if nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nil
		}
	}

	return nfe.NFeDocument{}, fkerrors.E(op, nfe.ErrCodeDocumentNotFound)
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
