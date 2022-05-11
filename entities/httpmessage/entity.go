package httpmessage

const (
	DefaultContentType = "application/json"
)

type HttpMessage struct {
	ContentType string
	BodyData    []byte
	HttpStatus  int
}

func Err(err error, status int) HttpMessage {
	return HttpMessage{
		ContentType: DefaultContentType,
		BodyData:    []byte(err.Error()),
		HttpStatus:  status,
	}
}

func New(bodyData []byte, httpStatus int) HttpMessage {
	return HttpMessage{
		ContentType: DefaultContentType,
		BodyData:    bodyData,
		HttpStatus:  httpStatus,
	}
}
