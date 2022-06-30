package internal

// TODO
import (
	"context"
	"net/http"

	"github.com/golang-http-server/services/get"

	"github.com/arquivei/foundationkit/apiutil"
	"github.com/arquivei/foundationkit/errors"
	kithttp "github.com/go-kit/kit/transport/http"
)

var errCodeInvalidRequest = errors.Code("INVALID_REQUEST")

// GetHTTPResponseError é retornado em caso de erro. O campo `code` contém um código
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
// swagger:response GetHTTPResponseError
type GetHTTPResponseError struct {
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
				getHTTPStatusCode,
				encodeHTTPError,
			),
		),
	}
}

func encodeHTTPError(ctx context.Context, err error) interface{} {
	return GetHTTPResponseError{
		Body: struct {
			Error apiutil.ErrorDescription `json:"error"`
		}{
			Error: apiutil.ParseError(err),
		},
	}.Body
}

func getHTTPStatusCode(err error) int {
	// FIXME implementar errors code do http e service
	switch errors.GetCode(err) {
	case get.ErrCodeDocumentNotFound:
		return http.StatusNotFound
	}
	return apiutil.GetDefaultErrorHTTPStatusCode(err)
}
