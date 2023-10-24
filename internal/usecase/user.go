package usecase

import (
	"refactoring/internal/apperror"
	"refactoring/internal/entity"
	"refactoring/internal/entity/dto"
	"refactoring/internal/repo"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type userUsecase struct {
	userRepoJSON repo.UserRepository

	mu sync.RWMutex
}

func NewUserUsecase(userRepoJSON repo.UserRepository) UserUsecase {
	return &userUsecase{
		userRepoJSON: userRepoJSON,
	}
}

func (u *userUsecase) CreateUser(request *dto.CreateUserRequest) (string, error) {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return "", err
	}

	err = userStore.Range(func(key, value any) error {
		user := value.(entity.User)
		if user.Email == request.Email {
			return apperror.ErrUserExist
		}
		return nil
	})
	if err != nil {
		return "", apperror.ErrUserExist
	}

	user := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	atomic.AddInt32(&userStore.Increment, 1)
	id := strconv.FormatInt(int64(userStore.Increment), 10)
	userStore.Set(user, id)

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

	user, ok := userStore.Read(id)
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

	user, ok := userStore.Read(id)
	if !ok {
		return apperror.ErrUserNotExist
	}

	user.DisplayName = request.DisplayName

	userStore.Set(user, id)

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

	_, ok := userStore.Read(id)
	if !ok {
		return apperror.ErrUserNotExist
	}

	userStore.Delete(id)

	atomic.AddInt32(&userStore.Increment, -1)

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
