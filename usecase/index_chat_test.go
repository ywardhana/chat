package usecase_test

import (
	"time"

	"github.com/ywardhana/chat/delivery/http"
	"github.com/ywardhana/chat/errormessage"

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
		conditions := map[string][]string{
			"offset": []string{"0"},
			"limit":  []string{"2"},
		}
		param, err := http.NewChatIndexParam(conditions)
		if err != nil {
			suite.Nil(err)
		}
		_, err = suite.usecase.IndexChat(param)

		suite.Equal(tt.expectedError, err)
	}
}
