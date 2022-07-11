package insert

import "context"

type InsertGateway interface {
	Processor(context.Context, Request) (Response, error)
}
