package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/shared"
)

type nfeFileRepository struct {
	json_file_path string
}

func NewNFeFileRepository(json_file_path string) nfe.Repository {
	return &nfeFileRepository{
		json_file_path: json_file_path,
	}
}

func (repository *nfeFileRepository) FindByID(id string) (nfe.NFeDocument, error) {
	store, err := repository.loadFileData()
	if err != nil {
		return nfe.NFeDocument{}, err
	}

	for _, nfeDocument := range store {
		if nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nfe.ErrAlreadyExists
		}
	}

	return nfe.NFeDocument{}, nfe.ErrNotFound
}

func (repository *nfeFileRepository) Save(nfeDocument nfe.NFeDocument) error {
	store, err := repository.loadFileData()
	if err != nil {
		return err
	}

	store = append(store, nfeDocument)

	return repository.saveFileData(store)
}

func (repository *nfeFileRepository) loadFileData() (nfe.NFeDocuments, error) {
	// FIXME refatorar return de errors
	data, err := shared.FromFile(repository.json_file_path)
	if err != nil {
		return nfe.NFeDocuments{}, err
	}

	if string(data) == "" {
		return nfe.NFeDocuments{}, nil
	}

	store, err := shared.ToNFeDocuments(data)
	if err != nil {
		return store, err
	}
	return store, nil
}

func (repository *nfeFileRepository) saveFileData(store nfe.NFeDocuments) error {
	data, err := shared.FromNFeDocuments(store)
	if err != nil {
		return err
	}

	err = shared.ToFile(repository.json_file_path, data)
	if err != nil {
		return err
	}
	return nil
}
