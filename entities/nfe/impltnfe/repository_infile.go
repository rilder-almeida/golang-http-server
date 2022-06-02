package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/shared"
)

type nfeInfileRepository struct {
	json_file_path string
}

func NewNfeInfileRepository(json_file_path string) nfe.Repository {
	return &nfeInfileRepository{
		json_file_path: json_file_path,
	}
}

func (repository *nfeInfileRepository) FindByID(id string) (nfe.NfeDocument, error) {
	store, err := repository.loadInFileData()
	if err != nil {
		return nfe.NfeDocument{}, err
	}

	for _, nfeDocument := range store {
		if nfeDocument.NfeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nfe.ErrAlreadyExists
		}
	}

	return nfe.NfeDocument{}, nfe.ErrNotFound
}

func (repository *nfeInfileRepository) Save(nfeDocument nfe.NfeDocument) error {
	store, err := repository.loadInFileData()
	if err != nil {
		return err
	}

	store = append(store, nfeDocument)

	return repository.saveInFileData(store)
}

func (repository *nfeInfileRepository) loadInFileData() (nfe.NfeDocuments, error) {
	data, err := shared.FromFile(repository.json_file_path)
	if err != nil {
		return nfe.NfeDocuments{}, err
	}

	if string(data) == "" {
		return nfe.NfeDocuments{}, nil
	}

	store, err := shared.ToNfeDocuments(data)
	if err != nil {
		return store, err
	}
	return store, nil
}

func (repository *nfeInfileRepository) saveInFileData(store nfe.NfeDocuments) error {
	data, err := shared.FromNfeDocuments(store)
	if err != nil {
		return err
	}

	err = shared.ToFile(repository.json_file_path, data)
	if err != nil {
		return err
	}
	return nil
}
