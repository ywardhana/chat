package chat

import (
	"github.com/ywardhana/chat/model"
)

type ChatUsecase interface {
	CreateChat(message string) *model.Chat
	IndexChat(param ChatIndexParam) ([]*model.Chat, error)
	CountChat() int
}

type ChatIndexParam interface {
	Limit() int
	Offset() int
}
