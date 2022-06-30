package nfe

type Repository interface {
	// FIXME refatorar regra de negociopara retornar nfe.NFeDocument, bool, error
	FindByID(id string) (NFeDocument, error)
	Save(nfeDocument NFeDocument) error
}
