package get

type GetGateway interface {
	Processor(Request) (Response, error)
}
