package usecase_test

func (suite *ChatUsecaseTestSuite) TestCountChat() {
	count := 4
	suite.chatRepo.On("Count").Return(count).Once()
	result := suite.usecase.CountChat()
	suite.Assert().Equal(count, result)
}
