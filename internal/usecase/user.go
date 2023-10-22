package usecase

import "refactoring/internal/repo"

type userUsecase struct {
	userRepoJSON repo.UserRepository
}

func NewUserUsecase() userUsecase {
	return &userUsecase{}
}
