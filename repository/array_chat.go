package repository

import "github.com/ywardhana/chat/model"

const (
	DefaultLimit = 10
)

type ChatRepoConfig struct {
	Offset int
	Limit  int
}

type Chat struct {
	chats []*model.Chat
	total int
}

func NewChatRepository() *Chat {
	return &Chat{
		chats: make([]*model.Chat, 0),
		total: 0,
	}
}

func (r *Chat) Insert(chat *model.Chat) *Chat {
	r.chats = append(r.chats, chat)
	r.total = len(r.chats)
	return r
}

func (r *Chat) GetNewest() *model.Chat {
	return r.chats[r.total-1]
}

func (r *Chat) Get(config ChatRepoConfig) []*model.Chat {
	if config.Limit == 0 {
		config.Limit = 10
	}

	if config.Offset > config.Limit {
		config.Offset = config.Limit - 1
	}

	return r.chats[r.total-config.Limit : r.total-config.Offset-1]
}
