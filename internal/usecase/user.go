package usecase

import (
	"refactoring/internal/apperror"
	"refactoring/internal/entity"
	"refactoring/internal/entity/dto"
	"refactoring/internal/repo"
	"strconv"
	"sync"
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

	for _, users := range userStore.List {
		if users.Email == request.Email {
			return "", apperror.ErrUserExist
		}
	}

	user := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	userStore.Increment++
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

	u.mu.RLock()
	user, ok := userStore.List[id]
	if !ok {
		return nil, apperror.ErrUserNotExist
	}
	u.mu.RUnlock()

	return &user, nil
}

func (u *userUsecase) UpdateUser(request *dto.UpdateUserRequest, id string) error {
	userStore, err := u.userRepoJSON.StorageReader()
	if err != nil {
		return err
	}

	u.mu.RLock()
	if _, ok := userStore.List[id]; !ok {
		return apperror.ErrUserNotExist
	}
	u.mu.RUnlock()

	u.mu.Lock()
	defer u.mu.Unlock()
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

	u.mu.RLock()
	if _, ok := userStore.List[id]; !ok {
		return apperror.ErrUserNotExist
	}
	u.mu.RUnlock()

	delete(userStore.List, id)

	u.mu.Lock()
	defer u.mu.Unlock()
	userStore.Increment -= 1

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
