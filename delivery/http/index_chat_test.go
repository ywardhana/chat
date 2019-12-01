package http_test

import (
	"net/http"
	"time"

	"github.com/ywardhana/chat/errormessage"

	"github.com/stretchr/testify/mock"
	"github.com/ywardhana/chat/model"
)

func (suite *ChatHandlerTestSuite) TestIndexChat() {
	tests := []struct {
		url            string
		mockFunc       func()
		expectedStatus int
	}{
		{
			url: "/chat",
			mockFunc: func() {
				chats := []*model.Chat{model.NewChat("hello", time.Now())}
				suite.chatUsecase.On("IndexChat", mock.Anything).Return(chats, nil).Once()
				suite.chatUsecase.On("CountChat").Return(1).Once()
			},
			expectedStatus: http.StatusOK,
		},
		{
			url: "/chat?limit=xxx",
			mockFunc: func() {
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			url: "/chat?offset=xxx",
			mockFunc: func() {
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			url: "/chat",
			mockFunc: func() {
				chats := []*model.Chat{model.NewChat("hello", time.Now())}
				suite.chatUsecase.On("IndexChat", mock.Anything).Return(chats, errormessage.ErrorUnexpected).Once()
				suite.chatUsecase.On("CountChat").Return(1).Once()
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt.mockFunc()
		r := NewRequest().Get(tt.url).AsBasic().Build()
		w := suite.Record(r)
		suite.Assert().Equal(tt.expectedStatus, w.Code)
	}
}
