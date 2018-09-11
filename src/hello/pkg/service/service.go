package service

import "context"
// HelloService describes the service.
type HelloService interface {
	// Add your methods here
	Foo(ctx context.Context,s string)(rs string, err error)
}
