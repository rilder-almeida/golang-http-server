package httpmessage

type HttpMessage struct {
	ContentType string
	BodyData    []byte
	HttpStatus  int
}
