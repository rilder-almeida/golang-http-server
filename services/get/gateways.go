package get

import "context"

type GetGateway interface {
	Processor(context.Context, Request) (Response, error)
}
