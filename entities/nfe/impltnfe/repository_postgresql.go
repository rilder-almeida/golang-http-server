package impltnfe

import (
	"errors"
	"time"

	fkerrors "github.com/arquivei/foundationkit/errors"
	"gorm.io/gorm"

	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/xml"
)

type nfePostgresqlRepository struct {
	db *gorm.DB
}

type postgresModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement:true;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	RawXml    string
	NfeId     string
	CNPJ      string
	VNF       string
}

func (postgresModel) TableName() string {
	return "nfe"
}

func NewNFePostgresqlRepository(database *gorm.DB) nfe.Repository {
	return &nfePostgresqlRepository{
		db: database,
	}
}

func (repository *nfePostgresqlRepository) FindByID(id string) (nfe.NFeDocument, error) {
	const op = fkerrors.Op("impltnfe.postgresql.FindByID")

	var postgresModel postgresModel
	result := repository.db.Where("nfe_id = ?", id).Table("nfe").First(&postgresModel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nfe.NFeDocument{}, fkerrors.E(op, nfe.ErrDocumentNotFound, nfe.ErrCodeDocumentNotFound)
	}
	if result.Error != nil {
		return nfe.NFeDocument{}, fkerrors.E(op, result.Error, nfe.ErrCodeGetDocument)
	}

	document, err := fromPostgresModel(postgresModel)
	if err != nil {
		return nfe.NFeDocument{}, err
	}

	return document, nil
}

func (repository *nfePostgresqlRepository) Save(nfeDocument nfe.NFeDocument) error {
	const op fkerrors.Op = "impltnfe.postgresql.Save"

	postgresModel, err := toPostgresModel(nfeDocument)
	if err != nil {
		return err
	}

	result := repository.db.Create(&postgresModel)
	if result.RowsAffected == 0 {
		return fkerrors.E(op, nfe.ErrSaveDocument, nfe.ErrCodeSaveDocument)
	}

	if result.Error != nil {
		fkerrors.E(op, result.Error, nfe.ErrCodeSaveDocument)
	}

	return nil
}

func toPostgresModel(nfeDocument nfe.NFeDocument) (postgresModel, error) {
	const op = fkerrors.Op("impltnfe.postgresql.toPostgresModel")

	if nfeDocument.RawXml == "" ||
		nfeDocument.NFeXmlDocument.NFe.InfNFe.Id == "" ||
		nfeDocument.NFeXmlDocument.NFe.InfNFe.Emit.CNPJ == "" ||
		nfeDocument.NFeXmlDocument.NFe.InfNFe.Total.ICMSTot.VNF == "" {
		return postgresModel{}, fkerrors.E(op, nfe.ErrProcessDocument, nfe.ErrCodeProcessDocument)
	}

	return postgresModel{
		RawXml: nfeDocument.RawXml,
		NfeId:  nfeDocument.NFeXmlDocument.NFe.InfNFe.Id,
		CNPJ:   nfeDocument.NFeXmlDocument.NFe.InfNFe.Emit.CNPJ,
		VNF:    nfeDocument.NFeXmlDocument.NFe.InfNFe.Total.ICMSTot.VNF,
	}, nil
}

func fromPostgresModel(postgresModel postgresModel) (nfe.NFeDocument, error) {
	const op = fkerrors.Op("impltnfe.postgresql.fromPostgresModel")

	if postgresModel.RawXml == "" ||
		postgresModel.NfeId == "" ||
		postgresModel.CNPJ == "" ||
		postgresModel.VNF == "" {
		return nfe.NFeDocument{}, fkerrors.E(op, nfe.ErrProcessDocument, nfe.ErrCodeProcessDocument)
	}

	return nfe.NFeDocument{
		RawXml: postgresModel.RawXml,
		NFeXmlDocument: xml.XmlDocument{
			NFe: xml.NFeField{
				InfNFe: xml.InfNFeField{
					Id: postgresModel.NfeId,
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
