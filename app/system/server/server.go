package server

import (
	"github.com/bukalapak/mitra/app/api/middleware"
	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Register(r *httprouter.Router, m *middleware.Middleware)
}
