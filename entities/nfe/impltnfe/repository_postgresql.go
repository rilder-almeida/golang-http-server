package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
	"gorm.io/gorm"
)

type nfePostgresqlRepository struct {
	db *gorm.DB
}

// TODO Criar models

func NewNfePostgresqlRepository(database *gorm.DB) nfe.Repository {
	return &nfePostgresqlRepository{
		db: database,
	}
}

func (repository *nfePostgresqlRepository) FindByID(id string) (nfe.NfeDocument, error) {
	nfeDocuments, err := loadPostgresqlData(repository.db)
	if err != nil {
		return nfe.NfeDocument{}, err
	}

	for _, nfeDocument := range nfeDocuments {
		if nfeDocument.NfeXmlDocument.NFe.InfNFe.Id == id {
			return nfeDocument, nfe.ErrAlreadyExists
		}
	}

	return nfe.NfeDocument{}, nfe.ErrNotFound

}

func (repository *nfePostgresqlRepository) Save(nfeDocument nfe.NfeDocument) error {
	return savePostgresqlData(repository.db, nfeDocument)
}

func loadPostgresqlData(db *gorm.DB) (nfe.NfeDocuments, error) {
	var nfeDocuments nfe.NfeDocuments
	result := db.Find(&nfeDocuments)
	if result.Error != nil {
		return nfeDocuments, result.Error
	}

	return nfeDocuments, nil
}

func savePostgresqlData(db *gorm.DB, nfeDocument nfe.NfeDocument) error {
	result := db.Create(&nfeDocument)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// TODO SUBIR POSTGRESQL
// TODO CRIAR O BANCO DE DADOS E A TABELA
// TODO TESTAR A CONEXAO
// TODO TESTAR A INSERCAO DE UM NOVO REGISTRO
// TODO TESTAR A INSERCAO DE UM REGISTRO QUE JA EXISTE
// TODO TESTAR A BUSCA POR ID
// TODO TESTAR A BUSCA POR ID QUE NAO EXISTE
// TODO TESTAR A BUSCA POR ID QUE JA EXISTE
