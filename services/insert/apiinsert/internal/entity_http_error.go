package internal

import (
	"context"
	"net/http"

	"github.com/arquivei/foundationkit/apiutil"
	fkerrors "github.com/arquivei/foundationkit/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/services/insert"
)

var ErrCodeInvalidRequest = fkerrors.Code("INVALID_REQUEST")
var ErrCodeInvalidResponse = fkerrors.Code("INVALID_RESPONSE")

// InsertHTTPResponseError é retornado em caso de erro. O campo `code` contém um código
// para ser usado no tratamento dos erros enquanto que o campo `message` contém um texto descritivo
// sobre o que aconteceu. O campo `message` não deve ser usado para ocmparação pois não há garantias
// sobre o formato da mensagem ou conteúdo.
//
// ```
// Error Code             | HTTP | Descrição
// =======================|======|==========
// INVALID_REQUEST        | 400  | Requisição inválida ou mal formatada
// DOCUMENT_NOT_FOUND     | 404  | NFe não encontrada no banco
// FAILED_GET_DOCUMENT    | 500  | Algum problema em buscar a NFSe no banco
// INTERNAL_ERROR         | 500  | Alguma falha ocorreu e não foi tratada corretamente
// ```
//
// swagger:response InsertHTTPResponseError
type InsertHTTPResponseError struct {
	// in:body
	Body struct {
		Error apiutil.ErrorDescription `json:"error"`
	}
}

// GetHTTPServerOption returns a http server option with an error encoder
func GetHTTPServerOption() []kithttp.ServerOption {
	return []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(
			apiutil.NewHTTPErrorJSONEncoder(
				insertHTTPStatusCode,
				encodeHTTPError,
			),
		),
	}
}

func encodeHTTPError(ctx context.Context, err error) interface{} {
	return InsertHTTPResponseError{
		Body: struct {
			Error apiutil.ErrorDescription `json:"error"`
		}{
			Error: apiutil.ParseError(err),
		},
	}.Body
}

func insertHTTPStatusCode(err error) int {
	switch fkerrors.GetCode(err) {
	case nfe.ErrCodeDocumentNotFound:
		return http.StatusNotFound
	case nfe.ErrCodeProcessDocument, nfe.ErrCodeSaveDocument, nfe.ErrCodeGetDocument, ErrCodeInvalidResponse:
		return http.StatusInternalServerError
	case insert.ErrCodeInvalidRequest, ErrCodeInvalidRequest, ErrCodeInvalidResponse:
		return http.StatusBadRequest
	}
	return apiutil.GetDefaultErrorHTTPStatusCode(err)
}
