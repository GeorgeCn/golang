// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "mqserver/mqhttp/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SendToMQEndpoint      endpoint.Endpoint
	ReceiveFromMQEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.MqhttpService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		ReceiveFromMQEndpoint: MakeReceiveFromMQEndpoint(s),
		SendToMQEndpoint:      MakeSendToMQEndpoint(s),
	}
	for _, m := range mdw["SendToMQ"] {
		eps.SendToMQEndpoint = m(eps.SendToMQEndpoint)
	}
	for _, m := range mdw["ReceiveFromMQ"] {
		eps.ReceiveFromMQEndpoint = m(eps.ReceiveFromMQEndpoint)
	}
	return eps
}
