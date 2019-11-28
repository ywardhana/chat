package usecase

import (
	"time"

	"github.com/ywardhana/chat/model"
)

func (uc *ChatUsecase) CreateChat(message string) *model.Chat {
	aChat := model.NewChat(message, time.Now())
	uc.repository.Insert(aChat)
	return aChat
}
