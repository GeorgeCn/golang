package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "mqserver/mqhttp/pkg/service"
	request2 "mqserver/mqhttp/pkg/request"
	"mqserver/mqhttp/pkg/response"
)



// MakeSendToMQEndpoint returns an endpoint that invokes SendToMQ on the service.
func MakeSendToMQEndpoint(s service.MqhttpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(request2.SendRequest)
		e0 := s.SendToMQ(ctx, req)
		return e0, nil
	}
}


// ReceiveFromMQRequest collects the request parameters for the ReceiveFromMQ method.
type ReceiveFromMQRequest struct{}

// ReceiveFromMQResponse collects the response parameters for the ReceiveFromMQ method.
type ReceiveFromMQResponse struct {
	I0 interface{} `json:"i0"`
	E1 error       `json:"e1"`
}

// MakeReceiveFromMQEndpoint returns an endpoint that invokes ReceiveFromMQ on the service.
func MakeReceiveFromMQEndpoint(s service.MqhttpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s.ReceiveFromMQ(ctx)
		return nil, nil
	}
}

// Failed implements Failer.
func (r ReceiveFromMQResponse) Failed() error {
	return r.E1
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendToMQ implements Service. Primarily useful in a client.
func (e Endpoints) SendToMQ(ctx context.Context, msg request2.SendRequest) (resp response.ResultJson ){
	rep, err := e.SendToMQEndpoint(ctx, msg)
	if err != nil {
		return
	}
	return rep.(response.ResultJson)
}

// ReceiveFromMQ implements Service. Primarily useful in a client.
func (e Endpoints) ReceiveFromMQ(ctx context.Context)  (resp response.ResultJson ){
	request := ReceiveFromMQRequest{}
	rep, err := e.ReceiveFromMQEndpoint(ctx, request)
	if err != nil {
		return
	}
	return rep.(response.ResultJson)
}
