package internal

import (
	"context"
	"net/http"

	"github.com/golang-http-server/services/insert"

	"github.com/arquivei/foundationkit/apiutil"
	fkerrors "github.com/arquivei/foundationkit/errors"
	kithttp "github.com/go-kit/kit/transport/http"
)

var errCodeInvalidRequest = fkerrors.Code("INVALID_REQUEST")

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
	// FIXME implementar errors code do http e service
	switch fkerrors.GetCode(err) {
	case insert.ErrCodeDocumentNotFound:
		return http.StatusNotFound
	}
	return apiutil.GetDefaultErrorHTTPStatusCode(err)
}
