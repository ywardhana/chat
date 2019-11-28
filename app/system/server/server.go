package server

import (
	"net/http"
	"sync"

	"github.com/ywardhana/goapi/response"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/chat/app/system/middleware"
)

type Handler interface {
	Register(r *httprouter.Router, m *middleware.Middleware)
}

type Ready struct {
	cb     chan bool
	status bool
	mutex  sync.Mutex
}

func BuildServer(ready *Ready, middleware *middleware.Middleware, handlers ...Handler) http.Handler {
	router := httprouter.New()

	for _, reg := range handlers {
		reg.Register(router, middleware)
	}

	router.NotFound = http.HandlerFunc(notFound)

	return router
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	meta := response.MetaInfo{HttpStatus: 404}

	response.OKWithMeta(w, nil, "not found", meta)
}
