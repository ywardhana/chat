package usecase

import "github.com/ywardhana/chat/chat"

type ChatUsecase struct {
	repository chat.ChatRepository
}

func NewChatUsecase(repository chat.ChatRepository) *ChatUsecase {
	return &ChatUsecase{
		repository: repository,
	}
}
