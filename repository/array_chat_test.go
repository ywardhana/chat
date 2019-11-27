package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/chat/model"
	"github.com/ywardhana/chat/repository"
)

func TestInsertAndCheckNewest(t *testing.T) {
	tests := []struct {
		chats []*model.Chat
	}{
		{
			chats: []*model.Chat{model.NewChat("halo", time.Now())},
		},
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check check", time.Now()),
			},
		},
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check 1", time.Now()),
				model.NewChat("check 2", time.Now()),
				model.NewChat("check 3", time.Now()),
			},
		},
	}

	for _, tt := range tests {
		chatRepo := repository.NewChatRepository()

		for _, chat := range tt.chats {
			chatRepo.Insert(chat)
		}

		assert.Equal(t, chatRepo.GetNewest(), tt.chats[len(tt.chats)-1])
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		chats               []*model.Chat
		chatConfig          repository.ChatRepoConfig
		expectedLastMessage *model.Chat
		expectedLen         int
		expectedErr         bool
	}{
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check 1", time.Now()),
				model.NewChat("check 2", time.Now()),
				model.NewChat("check 3", time.Now()),
			},
			chatConfig: repository.ChatRepoConfig{
				Limit:  0,
				Offset: 0,
			},
			expectedLastMessage: model.NewChat("check 3", time.Now()),
			expectedLen:         4,
			expectedErr:         false,
		},
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check 1", time.Now()),
				model.NewChat("check 2", time.Now()),
				model.NewChat("check 3", time.Now()),
			},
			chatConfig: repository.ChatRepoConfig{
				Limit:  0,
				Offset: 4,
			},
			expectedLastMessage: nil,
			expectedLen:         0,
			expectedErr:         false,
		},
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check 1", time.Now()),
				model.NewChat("check 2", time.Now()),
				model.NewChat("check 3", time.Now()),
			},
			chatConfig: repository.ChatRepoConfig{
				Limit:  2,
				Offset: 0,
			},
			expectedLastMessage: model.NewChat("check 3", time.Now()),
			expectedLen:         2,
			expectedErr:         false,
		},
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check 1", time.Now()),
				model.NewChat("check 2", time.Now()),
				model.NewChat("check 3", time.Now()),
			},
			chatConfig: repository.ChatRepoConfig{
				Limit:  2,
				Offset: 1,
			},
			expectedLastMessage: model.NewChat("check 2", time.Now()),
			expectedLen:         1,
			expectedErr:         false,
		},
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("check 1", time.Now()),
				model.NewChat("check 2", time.Now()),
				model.NewChat("check 3", time.Now()),
			},
			chatConfig: repository.ChatRepoConfig{
				Limit:  0,
				Offset: 5,
			},
			expectedLen: 0,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		chatRepo := repository.NewChatRepository()

		for _, chat := range tt.chats {
			chatRepo.Insert(chat)
		}
		chats, err := chatRepo.Get(tt.chatConfig)

		if tt.expectedLastMessage != nil {
			assert.Equal(t, chats[len(chats)-1].GetMessage(), tt.expectedLastMessage.GetMessage())
		}
		assert.Equal(t, len(chats), tt.expectedLen)
		assert.Equal(t, err != nil, tt.expectedErr)
	}
}
