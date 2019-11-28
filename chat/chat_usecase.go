package chat

import "github.com/ywardhana/chat/model"

type ChatUsecase interface {
	CreateChat(message string) *model.Chat
}
