package nfe

type Repository interface {
	FindByID(id string) (NFeDocument, error)
	Save(nfeDocument NFeDocument) error
}
