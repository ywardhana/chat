package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/chat/app/system/middleware"
)

type Handler interface {
	Register(r *httprouter.Router, m *middleware.Middleware)
}

func BuildServer(middleware *middleware.Middleware, handlers ...Handler) http.Handler {
	router := httprouter.New()

	for _, reg := range handlers {
		reg.Register(router, middleware)
	}

	router.NotFound = http.HandlerFunc(notFound)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
