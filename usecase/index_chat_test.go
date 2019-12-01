package usecase_test

import (
	"time"

	"github.com/ywardhana/chat/errormessage"

	"github.com/ywardhana/chat/chat"

	"github.com/stretchr/testify/mock"
	"github.com/ywardhana/chat/model"
)

func (suite *ChatUsecaseTestSuite) TestIndexChat() {

	tests := []struct {
		chats         []*model.Chat
		returnError   error
		expectedError error
	}{
		{
			chats: []*model.Chat{
				model.NewChat("halo", time.Now()),
				model.NewChat("satu dua tiga", time.Now()),
			},
			returnError:   nil,
			expectedError: nil,
		},
		{
			chats:         []*model.Chat{},
			returnError:   errormessage.ErrNotFound,
			expectedError: errormessage.ErrNotFound,
		},
	}

	for _, tt := range tests {
		suite.chatRepo.On("Get", mock.Anything).Return(tt.chats, tt.returnError).Once()
		_, err := suite.usecase.IndexChat(chat.ChatIndexParam{
			Offset: 0,
			Limit:  2,
		})

		suite.Equal(tt.expectedError, err)
	}
}
