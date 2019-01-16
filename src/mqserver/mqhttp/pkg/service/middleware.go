package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	"mqserver/mqhttp/pkg/request"
	"mqserver/mqhttp/pkg/response"
)

// Middleware describes a service middleware.
type Middleware func(MqhttpService) MqhttpService

type loggingMiddleware struct {
	logger log.Logger
	next   MqhttpService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a MqhttpService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MqhttpService) MqhttpService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendToMQ(ctx context.Context, msg request.SendRequest) (resp response.ResultJson ) {
	defer func() {
		l.logger.Log("method", "SendToMQ", "request", msg.String(), "response", resp.String())
	}()
	return l.next.SendToMQ(ctx, msg)
}
func (l loggingMiddleware) ReceiveFromMQ(ctx context.Context) {
	defer func() {
		l.logger.Log("method", "ReceiveFromMQ", "request", nil, "response", nil)
	}()
}
