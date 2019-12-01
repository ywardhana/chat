package chat

import (
	"github.com/ywardhana/chat/model"
)

type ChatUsecase interface {
	CreateChat(message string) *model.Chat
	IndexChat(param ChatIndexParam) ([]*model.Chat, error)
}

type ChatIndexParam struct {
	Offset int
	Limit  int
}
