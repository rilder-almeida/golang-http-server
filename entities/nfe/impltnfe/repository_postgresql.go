package impltnfe

import (
	"errors"

	"gorm.io/gorm"

	customerrors "github.com/golang-http-server/entities/errors"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/xml"
)

type nfePostgresqlRepository struct {
	db *gorm.DB
}

type postgresModel struct {
	gorm.Model
	RawXml string
	NFeId  string
	CNPJ   string
	VNF    string
}

func NewNFePostgresqlRepository(database *gorm.DB) nfe.Repository {
	return &nfePostgresqlRepository{
		db: database,
	}
}

func (repository *nfePostgresqlRepository) FindByID(id string) (nfe.NFeDocument, error) {
	var postgresModel postgresModel
	result := repository.db.Where("nfe_id = ?", id).Table("nfe").First(&postgresModel)
	if result.RowsAffected == 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nfe.NFeDocument{}, nfe.ErrNotFound
	}
	if result.Error != nil {
		return nfe.NFeDocument{}, result.Error
	}

	document, err := fromPostgresModel(postgresModel)
	if err != nil {
		return nfe.NFeDocument{}, err
	}

	return document, nfe.ErrAlreadyExists
}

func (repository *nfePostgresqlRepository) Save(nfeDocument nfe.NFeDocument) error {
	postgresModel, err := toPostgresModel(nfeDocument)
	if err != nil {
		return err
	}

	result := repository.db.Create(&postgresModel)
	if result.RowsAffected == 0 {
		return customerrors.New("FAILED_TO_SAVE", "Failed to save on Postegres", nil)
	}

	return result.Error
}

func toPostgresModel(nfeDocument nfe.NFeDocument) (postgresModel, error) {
	if nfeDocument.RawXml == "" ||
		nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == "" ||
		nfeDocument.NFeXmlDocument.NFe.InfNFe.Emit.CNPJ == "" ||
		nfeDocument.NFeXmlDocument.NFe.InfNFe.Total.ICMSTot.VNF == "" {
		return postgresModel{}, customerrors.New("INVALID_NFEDOCUMENT", "Invalid NFeDocument", nil)
	}

	return postgresModel{
		RawXml: nfeDocument.RawXml,
		NFeId:  nfeDocument.NFeXmlDocument.NFe.InfNFe.Id,
		CNPJ:   nfeDocument.NFeXmlDocument.NFe.InfNFe.Emit.CNPJ,
		VNF:    nfeDocument.NFeXmlDocument.NFe.InfNFe.Total.ICMSTot.VNF,
	}, nil
}

func fromPostgresModel(postgresModel postgresModel) (nfe.NFeDocument, error) {
	if postgresModel.RawXml == "" ||
		postgresModel.NFeId == "" ||
		postgresModel.CNPJ == "" ||
		postgresModel.VNF == "" {
		return nfe.NFeDocument{}, customerrors.New("INVALID_POSTGRES_MODEL", "Postgres model is invalid", nil)
	}

	return nfe.NFeDocument{
		RawXml: postgresModel.RawXml,
		NFeXmlDocument: xml.XmlDocument{
			NFe: xml.NFeField{
				InfNFe: xml.InfNFeField{
					Id: postgresModel.NFeId,
					Emit: xml.EmitField{
						CNPJ: postgresModel.CNPJ,
					},
					Total: xml.TotalField{
						ICMSTot: xml.ICMSTotField{
							VNF: postgresModel.VNF,
						},
					},
				},
			},
		},
	}, nil
}

// TODO CRIAR O BANCO DE DADOS E A TABELA
// TODO TESTAR A CONEXAO
// TODO TESTAR A INSERCAO DE UM NOVO REGISTRO
// TODO TESTAR A INSERCAO DE UM REGISTRO QUE JA EXISTE
// TODO TESTAR A BUSCA POR ID
// TODO TESTAR A BUSCA POR ID QUE NAO EXISTE
// TODO TESTAR A BUSCA POR ID QUE JA EXISTE
