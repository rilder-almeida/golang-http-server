package get

import (
	"github.com/golang-http-server/entities/nfe"
)

type Response struct {
	NfeDocument nfe.NfeDocument `json:"nfe_document"`
}
