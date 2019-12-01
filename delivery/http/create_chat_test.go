package http_test

import (
	"net/http"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/ywardhana/chat/model"
)

func (suite *ChatHandlerTestSuite) TestCreateChat() {
	tests := []struct {
		params         interface{}
		mockFunc       func()
		expectedStatus int
	}{
		{
			params: map[string]interface{}{
				"message": "hello",
			},
			mockFunc: func() {
				suite.chatUsecase.On("CreateChat", mock.Anything).Return(model.NewChat("hello", time.Now())).Once()
			},
			expectedStatus: http.StatusOK,
		},
		{
			params: "hello",
			mockFunc: func() {
				suite.chatUsecase.On("CreateChat", mock.Anything).Return(model.NewChat("hello", time.Now())).Once()
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt.mockFunc()
		r := NewRequest().Post("/chat", tt.params).AsBasic().Build()
		w := suite.Record(r)
		suite.Assert().Equal(tt.expectedStatus, w.Code)
	}
}
