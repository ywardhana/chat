package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ywardhana/chat/chat"
	"github.com/ywardhana/chat/chat/mocks"
	"github.com/ywardhana/chat/usecase"
)

type ChatUsecaseTestSuite struct {
	suite.Suite
	chatRepo *mocks.ChatRepository
	usecase  chat.ChatUsecase
}

func (suite *ChatUsecaseTestSuite) SetupTest() {
	suite.chatRepo = new(mocks.ChatRepository)

	suite.usecase = usecase.NewChatUsecase(
		suite.chatRepo,
	)
}

func TestChatUsecase(t *testing.T) {
	suite.Run(t, new(ChatUsecaseTestSuite))
}
