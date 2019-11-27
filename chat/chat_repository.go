package chat

import (
	"github.com/ywardhana/chat/model"
	"github.com/ywardhana/chat/repository"
)

type ChatRepository interface {
	Insert(chat *model.Chat) *repository.Chat
	GetNewest() *model.Chat
	Get(config ChatRepoConfig) []*model.Chat
}
