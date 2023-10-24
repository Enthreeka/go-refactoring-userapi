package entity

import (
	"sync"
	"time"
)

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	UserList  map[string]User
	UserStore struct {
		Increment int32    `json:"increment"`
		List      UserList `json:"list"`

		mu sync.RWMutex
	}
)

// Set потокобезопасная запись структуры пользователя в map
func (u *UserStore) Set(user User, id string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.List[id] = user
}

// Read потокобезопасные поиск пользователя в map
func (u *UserStore) Read(id string) (User, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	user, ok := u.List[id]
	if !ok {
		return User{}, false
	}

	return user, true
}

// Delete потокобезопасное удаление пользователя
func (u *UserStore) Delete(id string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.List, id)
}

// Range потокобезопасные перебор данных с выполением какой-то функции для каждого элемента
func (u *UserStore) Range(f func(key, value any) error) error {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for key, user := range u.List {
		if err := f(key, user); err != nil {
			return err
		}
	}

	return nil
}
