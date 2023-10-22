package usecase

import (
	"refactoring/internal/apperror"
	"refactoring/internal/entity"
	"refactoring/internal/entity/dto"
	"refactoring/internal/repo"
	"strconv"
	"time"
)

type userUsecase struct {
	userRepoJSON repo.UserRepository
}

func NewUserUsecase(userRepoJSON repo.UserRepository) UserUsecase {
	return &userUsecase{
		userRepoJSON,
	}
}

func (u *userUsecase) CreateUser(request *dto.CreateUserRequest) (string, error) {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return "", err
	}

	userStore.Increment++
	user := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id := strconv.Itoa(userStore.Increment)
	userStore.List[id] = user

	err = u.userRepoJSON.StorageWriter(userStore)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userUsecase) GetUser(id string) (*entity.User, error) {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return nil, err
	}

	user, ok := userStore.List[id]
	if !ok {
		return nil, apperror.ErrUserNotExist
	}

	return &user, nil
}

func (u *userUsecase) UpdateUser(request *dto.UpdateUserRequest, id string) error {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return err
	}

	if _, ok := userStore.List[id]; !ok {
		return apperror.ErrUserNotExist
	}

	user := userStore.List[id]
	user.DisplayName = request.DisplayName
	userStore.List[id] = user

	err = u.userRepoJSON.StorageWriter(userStore)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) DeleteUser(id string) error {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return err
	}

	if _, ok := userStore.List[id]; !ok {
		return apperror.ErrUserNotExist
	}

	delete(userStore.List, id)

	err = u.userRepoJSON.StorageWriter(userStore)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) SearchUsers() (entity.UserList, error) {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return nil, err
	}

	if userStore.List == nil {
		return nil, apperror.ErrStorageEmpty
	}

	return userStore.List, nil
}
