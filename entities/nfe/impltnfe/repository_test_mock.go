package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
	returnError bool
	err         error
}

func NewRepositoryMock(returnError bool, err error) *RepositoryMock {
	return &RepositoryMock{
		returnError: returnError,
		err:         err,
	}
}

func (r RepositoryMock) FindByID(id string) (nfe.NFeDocument, error) {
	args := r.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).(nfe.NFeDocument), args.Error(1)
	}
	return nfe.NFeDocument{}, args.Get(1).(error)
}

func (r RepositoryMock) Save(nfeDocument nfe.NFeDocument) error {
	return nil
}
