package insert

type InsertGateway interface {
	Processor(Request) (Response, error)
}
