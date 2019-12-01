package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/chat/app/system/middleware"
)

type ChatHandler struct {
	// usecase chat.
}

func (h *ChatHandler) Register(router *httprouter.Router, m *middleware.Middleware) {
	router.GET("/coba", m.AuthBasic(h.TestFunc))
	// router.POST("/chat", m.AuthBasic(h.CreateChatHandler))
}

func (h *ChatHandler) TestFunc(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	return nil
}
