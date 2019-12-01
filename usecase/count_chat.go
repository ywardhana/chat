package usecase

func (uc *ChatUsecase) CountChat() int {
	return uc.repository.Count()
}
