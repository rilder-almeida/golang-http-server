package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

var postgresqlConfigDefault = postgresqlConfig{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "postgres",
	Dbname:   "postgres",
	Sslmode:  "disable",
}

type nfePostgresqlRepository struct{}

func NewNfePostgresqlRepository() nfe.Repository {
	return &nfePostgresqlRepository{}
}

func (repository *nfePostgresqlRepository) FindByID(id string) (nfe.NfeDocument, error) {
	db := getConnection(postgresqlConfigDefault)

	nfeDocuments, err := loadPostgresqlData(db)
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
	db := getConnection(postgresqlConfigDefault)
	return savePostgresqlData(db, nfeDocument)
}

func getConnection(config postgresqlConfig) *gorm.DB {
	dsn := "host=" + config.Host + " port=" + config.Port + " user=" + config.User + " password=" + config.Password + " dbname=" + config.Dbname + " sslmode=" + config.Sslmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
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

/*
https://stackoverflow.com/questions/69795462/does-gorm-automatically-close-the-connection

A gorm.DB object is intended to be reused, like a sql.DB handle.
You rarely have to explicitly close these objects. Just create it once and reuse it.

gorm.DB contains a sql.DB which uses a connection pool to manage the connections.
If it is closed, it will stop accepting new queries, wait for running queries to finish and close all connections.
*/
