package middleware

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/julienschmidt/httprouter"
)

type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error
type ConnWithError func(*websocket.Conn) error

type MiddlewareConfig struct {
	BasicUsername string
	BasicPassword string
}

type Middleware struct {
	config    MiddlewareConfig
	decorator *Decorator
}

func NewMiddleware(config MiddlewareConfig) *Middleware {
	return &Middleware{
		config:    config,
		decorator: NewDecorator(),
	}
}

func (m *Middleware) AuthBasic(handler HandleWithError) httprouter.Handle {
	return m.decorator.ApplyDecorator(handler, m.BasicTokenAuth())
}

func (m *Middleware) WebsocketServe(handler ConnWithError) httprouter.Handle {
	return m.decorator.ApplyWebsocketDecorator(handler)
}

func (m *Middleware) BasicTokenAuth() Auth {
	return &BasicAuth{
		BasicUsername: m.config.BasicUsername,
		BasicPassword: m.config.BasicPassword,
	}
}
