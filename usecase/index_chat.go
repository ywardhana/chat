package usecase

import (
	"github.com/ywardhana/chat/chat"
	"github.com/ywardhana/chat/model"
	"github.com/ywardhana/chat/repository"
)

func (uc *ChatUsecase) IndexChat(param chat.ChatIndexParam) ([]*model.Chat, error) {
	return uc.repository.Get(repository.ChatRepoConfig{
		Limit:  param.Limit(),
		Offset: param.Offset(),
	})
}
