package internal

import (
	fkerrors "github.com/arquivei/foundationkit/errors"
	"github.com/golang-http-server/entities/nfe"
)

type GetHTTPRequest struct {
	Body struct {
		Id string
	}
}

func ValidateGetHTTPRequest(r GetHTTPRequest) error {
	const op = fkerrors.Op("apiinsert.internal.ValidateInsertHTTPRequest")

	if r.Body.Id == "" {
		return fkerrors.E(op, nfe.ErrEmptyId, ErrCodeInvalidRequest)
	}
	return nil
}
