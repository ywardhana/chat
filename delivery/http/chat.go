package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/chat/app/system/middleware"
	"github.com/ywardhana/chat/chat"
)

type ChatHandler struct {
	usecase chat.ChatUsecase
}

func NewChatHandler(chatUC chat.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		usecase: chatUC,
	}
}

func (h *ChatHandler) Register(router *httprouter.Router, m *middleware.Middleware) {
	router.GET("/coba", m.AuthBasic(h.TestFunc))
	router.POST("/chat", m.AuthBasic(h.CreateChat))
}

func (h *ChatHandler) TestFunc(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	return nil
}
