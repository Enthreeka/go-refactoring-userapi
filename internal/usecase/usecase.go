package usecase

type UserUsecase interface {
	CreateUser()
	GetUser()
	UpdateUser()
	DeleteUser()
	SearchUsers()
}
