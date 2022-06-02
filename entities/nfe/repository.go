package nfe

type Repository interface {
	FindByID(id string) (NfeDocument, error)
	Save(nfeDocument NfeDocument) error
}
