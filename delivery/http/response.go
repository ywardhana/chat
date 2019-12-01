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

func serializerListChat(chats []*model.Chat) (response []ChatResponse) {
	response = make([]ChatResponse, len(chats))
	for index, chat := range chats {
		response[index] = serializerChat(chat)
	}
	return
}
