package repo

import (
	"encoding/json"
	"io/fs"
	"os"
	"refactoring/internal/entity"
	"sync"
)

type userRepoJSON struct {
	store string

	mu sync.Mutex
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

// StorageWriter Использует мьютекс при записи в json файл, так как при выполнении нескольких горутин может произойти гонка данных
func (u *userRepoJSON) StorageWriter(userStore *entity.UserStore) error {
	u.mu.Lock()
	defer u.mu.Unlock()

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
