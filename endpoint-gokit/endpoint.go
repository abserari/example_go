package endpoint

import "context"

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func Nop(ctx context.Context, request interface{}) (response interface{}, err error) {
	return struct{}{}, nil
}

type Middleware func(Endpoint) Endpoint

func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Endpoint) Endpoint {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type Failer interface {
	Failed() error
}
