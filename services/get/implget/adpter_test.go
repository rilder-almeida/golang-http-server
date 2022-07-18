package implget

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/mocks"
	"github.com/golang-http-server/services/get"
)

func TestAdapter_receiver(t *testing.T) {
	tests := []struct {
		name       string
		repository *mocks.Repository
		request    get.Request
		want       nfe.NFeDocument
		wantErr    error
	}{
		{
			name:       "should return error",
			repository: NewRepositoryMock(t, nfe.NFeDocument{}, errors.New("error")),
			request: get.Request{
				Id: "123",
			},
			wantErr: errors.New("error"),
		},
		{
			name:       "should return nil error",
			repository: NewRepositoryMock(t, nfe.NFeDocument{RawXml: "stub"}, nil),
			request: get.Request{
				Id: "123",
			},
			want:    nfe.NFeDocument{RawXml: "stub"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &Adapter{
				repository: tt.repository,
			}
			got, err := adapter.receiver(tt.request)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
			//tt.repository.AssertExpectations(t)
		})
	}
}

func NewRepositoryMock(t *testing.T, nfeDocument nfe.NFeDocument, err error) *mocks.Repository {
	repository := mocks.NewRepository(t)
	repository.On("FindByID", "123").Return(nfeDocument, err)
	return repository
}
