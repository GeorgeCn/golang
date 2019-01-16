package service

import "context"

// WorldService describes the service.
type WorldService interface {
	// Add your methods here
	World(ctx context.Context, s string) (rs string, err error)
}

type basicWorldService struct{}

func (b *basicWorldService) World(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of World
	return rs, err
}

// NewBasicWorldService returns a naive, stateless implementation of WorldService.
func NewBasicWorldService() WorldService {
	return &basicWorldService{}
}

// New returns a WorldService with all of the expected middleware wired in.
func New(middleware []Middleware) WorldService {
	var svc WorldService = NewBasicWorldService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
