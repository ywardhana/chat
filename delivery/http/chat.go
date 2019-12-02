package http

import (
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
	router.POST("/chat", m.AuthBasic(h.CreateChat))
	router.GET("/chat", m.AuthBasic(h.IndexChat))
	router.GET("/websocket", m.WebsocketServe(h.ChatWebsocket))
}
