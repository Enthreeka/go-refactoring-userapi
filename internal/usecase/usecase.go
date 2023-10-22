package usecase

import (
	"refactoring/internal/entity"
	"refactoring/internal/entity/dto"
)

type UserUsecase interface {
	CreateUser(request *dto.CreateUserRequest) (string, error)
	GetUser(id string) (*entity.User, error)
	UpdateUser(request *dto.UpdateUserRequest, id string) error
	DeleteUser(id string) error
	SearchUsers() (entity.UserList, error)
}
