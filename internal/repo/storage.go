package repo

import "refactoring/internal/entity"

type UserRepository interface {
	StorageReader() (*entity.UserStore, error)
	StorageWriter(userStore *entity.UserStore) error
}
