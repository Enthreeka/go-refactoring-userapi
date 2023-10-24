package usecase

import (
	"refactoring/internal/entity/dto"
	"refactoring/internal/repo"
	"testing"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	mockDB := repo.NewUserRepoJSON(`C:\Users\world\go-workspace\userapi\mock.json`)

	userUsecaseMock := NewUserUsecase(mockDB)

	request := &dto.CreateUserRequest{
		DisplayName: "d1st",
		Email:       "e1412487014@Test.ru",
	}

	userID, err := userUsecaseMock.CreateUser(request)
	if err != nil {
		t.Error(err)
	}

	if userID == "" {
		t.Errorf("user: %v , not created", request)
	}

	userStore, err := mockDB.StorageReader()
	if err != nil {
		t.Error(err)
	}

	createdUser, exist := userStore.List[userID]
	if !exist {
		t.Errorf("User with ID %s was not stored in the repository", userID)
	}

	if createdUser.DisplayName != request.DisplayName || createdUser.Email != request.Email {
		t.Errorf("Stored user does not match the input request")
	}
}
