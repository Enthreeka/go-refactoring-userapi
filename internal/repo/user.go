package repo

import (
	"encoding/json"
	"io/fs"
	"os"
	"refactoring/internal/entity"
)

type userRepoJSON struct {
	store string
}

func NewUserRepoJSON(store string) UserRepository {
	return &userRepoJSON{
		store: store,
	}
}

func (u *userRepoJSON) StorageReader() (*entity.UserStore, error) {
	fileByte, err := os.ReadFile(u.store)
	if err != nil {
		return nil, err
	}

	userStore := &entity.UserStore{}

	err = json.Unmarshal(fileByte, &userStore)
	if err != nil {
		return nil, err
	}

	return userStore, nil
}

func (u *userRepoJSON) StorageWriter(userStore *entity.UserStore) error {
	b, err := json.MarshalIndent(&userStore, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(u.store, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
