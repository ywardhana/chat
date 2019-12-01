package repository

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/ywardhana/chat/model"
)

const (
	DefaultLimit = 10
)

type ChatRepoConfig struct {
	Offset int
	Limit  int
}

func (cf *ChatRepoConfig) Validate(maxOffset int) error {
	return v.ValidateStruct(cf,
		v.Field(&cf.Offset, v.Min(0), v.Max(maxOffset)),
		v.Field(&cf.Limit, v.Min(0)),
	)
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

func (r *Chat) Get(config ChatRepoConfig) ([]*model.Chat, error) {
	if err := config.Validate(r.total); err != nil {
		return nil, err
	}

	if config.Limit == 0 {
		config.Limit = 10
	}

	start := r.total - config.Limit
	if start < 0 {
		start = 0
	}
	end := r.total - config.Offset

	return r.chats[start:end], nil
}

func (r *Chat) Count() int {
	return len(r.chats)
}
