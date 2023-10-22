package http

import "net/http"

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (u *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *userHandler) SearchUsersHandler(w http.ResponseWriter, r *http.Request) {

}
