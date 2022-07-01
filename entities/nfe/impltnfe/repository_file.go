package impltnfe

import (
	fkerrors "github.com/arquivei/foundationkit/errors"
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
	const op = fkerrors.Op("nfe.impltnfe.file.FindByID")

	store, err := repository.loadFileData()
	if err != nil {
		return nfe.NFeDocument{}, err
	}

	for _, nfeDocument := range store {
		if nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nil
		}
	}

	return nfe.NFeDocument{}, fkerrors.E(op, nfe.ErrCodeDocumentNotFound)
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
	const op = fkerrors.Op("nfe.impltnfe.file.loadFileData")

	data, err := shared.FromFile(repository.json_file_path)
	if err != nil {
		return nfe.NFeDocuments{}, fkerrors.E(op, err)
	}

	if string(data) == "" {
		return nfe.NFeDocuments{}, nil
	}

	store, err := shared.ToNFeDocuments(data)
	if err != nil {
		return store, fkerrors.E(op, err)
	}
	return store, nil
}

func (repository *nfeFileRepository) saveFileData(store nfe.NFeDocuments) error {
	const op = fkerrors.Op("nfe.impltnfe.file.saveFileData")

	data, err := shared.FromNFeDocuments(store)
	if err != nil {
		return fkerrors.E(op, err)
	}

	err = shared.ToFile(repository.json_file_path, data)
	if err != nil {
		return fkerrors.E(op, err)
	}
	return nil
}
