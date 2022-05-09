package httpmessage

import "fmt"

const (
	DefaultContentType = "application/json"
)

type HttpMessage struct {
	ContentType string
	BodyData    []byte
	HttpStatus  int
}

func Err(errMGs string, err error, status int) HttpMessage {
	return HttpMessage{
		ContentType: DefaultContentType,
		BodyData: []byte(fmt.Sprintf(`
			{
				"message": "%s",
				"error": "%s"
			}`, errMGs, err)),
		HttpStatus: status,
	}
}

func New(bodyData []byte, httpStatus int) HttpMessage {
	return HttpMessage{
		ContentType: DefaultContentType,
		BodyData:    bodyData,
		HttpStatus:  httpStatus,
	}
}
