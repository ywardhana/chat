package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/suite"
	"github.com/ywardhana/chat/app/system/middleware"
	"github.com/ywardhana/chat/chat/mocks"
	chatHTTP "github.com/ywardhana/chat/delivery/http"
)

type ChatHandlerTestSuite struct {
	suite.Suite
	chatUsecase *mocks.ChatUsecase
	handler     http.Handler
}

func (suite *ChatHandlerTestSuite) SetupTest() {
	suite.chatUsecase = new(mocks.ChatUsecase)
	chatHandler := chatHTTP.NewChatHandler(suite.chatUsecase)

	router := httprouter.New()
	mConfig := middleware.MiddlewareConfig{
		BasicUsername: "chat",
		BasicPassword: "chat",
	}

	middleware := middleware.NewMiddleware(mConfig)

	chatHandler.Register(router, middleware)

	suite.handler = router
}

func (suite *ChatHandlerTestSuite) Record(request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()

	suite.handler.ServeHTTP(response, request)

	return response
}

func TestChatHandler(t *testing.T) {
	suite.Run(t, new(ChatHandlerTestSuite))
}
