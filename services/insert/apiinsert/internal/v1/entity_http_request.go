package v1

import (
	fkerrors "github.com/arquivei/foundationkit/errors"
	"github.com/golang-http-server/entities/nfe"
)

type InsertHTTPRequest struct {
	Body struct {
		XML string
	}
}

func ValidateInsertHTTPRequest(r InsertHTTPRequest) error {
	const op = fkerrors.Op("apiinsert.v1.ValidateInsertHTTPRequest")

	if r.Body.XML == "" {
		return fkerrors.E(op, nfe.ErrEmptyXML, ErrCodeInvalidRequest)
	}
	return nil
}
