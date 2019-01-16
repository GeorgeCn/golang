// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "mqserver/mqhttp/pkg/endpoint"
	http1 "mqserver/mqhttp/pkg/http"
	service "mqserver/mqhttp/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"ReceiveFromMQ": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ReceiveFromMQ", logger))},
		"SendToMQ":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SendToMQ", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["SendToMQ"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SendToMQ")), endpoint.InstrumentingMiddleware(duration.With("method", "SendToMQ"))}
	mw["ReceiveFromMQ"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ReceiveFromMQ")), endpoint.InstrumentingMiddleware(duration.With("method", "ReceiveFromMQ"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"SendToMQ", "ReceiveFromMQ"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
