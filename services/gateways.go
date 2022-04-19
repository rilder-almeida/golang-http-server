package services

type NfeStore interface {
	Receiver(Request) (*[]byte, error)
	Responder([]byte) (Response, error)
}

type Request interface{}

type Response interface{}
