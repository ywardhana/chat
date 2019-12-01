package http

import (
	"github.com/ywardhana/chat/model"
)

type ChatResponse struct {
	Messagge  string `json:"message"`
	CreatedAt string `json:"created_at"`
}

func serializerChat(chat *model.Chat) ChatResponse {
	return ChatResponse{
		Messagge:  chat.GetMessage(),
		CreatedAt: chat.GetTime(),
	}
}
