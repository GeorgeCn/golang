package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "service/world/pkg/service"
)

// WorldRequest collects the request parameters for the World method.
type WorldRequest struct {
	S string `json:"s"`
}

// WorldResponse collects the response parameters for the World method.
type WorldResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeWorldEndpoint returns an endpoint that invokes World on the service.
func MakeWorldEndpoint(s service.WorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WorldRequest)
		rs, err := s.World(ctx, req.S)
		return WorldResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r WorldResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// World implements Service. Primarily useful in a client.
func (e Endpoints) World(ctx context.Context, s string) (rs string, err error) {
	request := WorldRequest{S: s}
	response, err := e.WorldEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WorldResponse).Rs, response.(WorldResponse).Err
}
