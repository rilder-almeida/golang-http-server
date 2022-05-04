package impltnfe

import (
	"fmt"

	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/shared"
)

type nfeInfileRepository struct {
	store          nfe.NfeDocumentList
	json_file_path string
}

func NewNfeInfileRepository(json_file_path string) nfe.Repository {
	return &nfeInfileRepository{
		store:          make([]nfe.NfeDocument, 0),
		json_file_path: json_file_path,
	}
}

func (repository *nfeInfileRepository) FindByID(id string) (nfe.NfeDocument, error) {
	err := repository.loadInFileData()
	if err != nil {
		return nfe.NfeDocument{}, err
	}

	for _, nfeDocument := range repository.store {
		if nfeDocument.NfeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nil
		}
	}

	return nfe.NfeDocument{}, fmt.Errorf("NFe not found")
}

func (repository *nfeInfileRepository) Save(nfeDocument nfe.NfeDocument) error {
	err := repository.loadInFileData()
	if err != nil {
		return err
	}

	repository.store = append(repository.store, nfeDocument)

	return repository.saveInFileData()
}

func (repository *nfeInfileRepository) loadInFileData() error {
	data, err := shared.FromFile(repository.json_file_path)
	if err != nil {
		return err
	}

	repository.store, err = shared.ToNfeDocumentList(data)
	if err != nil {
		return err
	}
	return nil
}

func (repository *nfeInfileRepository) saveInFileData() error {
	data, err := shared.FromNfeDocumentList(repository.store)
	if err != nil {
		return err
	}

	err = shared.ToFile(repository.json_file_path, data)
	if err != nil {
		return err
	}
	return nil
}