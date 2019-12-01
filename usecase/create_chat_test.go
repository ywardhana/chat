package usecase_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/ywardhana/chat/repository"
)

func (suite *ChatUsecaseTestSuite) TestCreateChat() {
	message := "chat"
	suite.chatRepo.On("Insert", mock.Anything).Return(repository.NewChatRepository()).Once()
	result := suite.usecase.CreateChat(message)
	suite.Assert().Equal(message, result.GetMessage())
}
