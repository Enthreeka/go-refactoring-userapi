package http

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"refactoring/internal/apperror"
	"refactoring/internal/entity/dto"
	"refactoring/internal/usecase"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *userHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (u *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := u.userUsecase.GetUser(id)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotExist) {
			_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func (u *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	request := &dto.CreateUserRequest{}

	if err := render.Bind(r, request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	id, err := u.userUsecase.CreateUser(request)
	if err != nil {
		if errors.Is(err, apperror.ErrUserExist) {
			render.Render(w, r, ErrInvalidRequest(err, http.StatusBadRequest))
			return
		}
		render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"status":  "created",
		"user_id": id,
	})
}

func (u *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := u.userUsecase.DeleteUser(id)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotExist) {
			_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (u *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	request := &dto.UpdateUserRequest{}

	if err := render.Bind(r, request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	id := chi.URLParam(r, "id")

	err := u.userUsecase.UpdateUser(request, id)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotExist) {
			_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (u *userHandler) SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := u.userUsecase.SearchUsers()
	if err != nil {
		if errors.Is(err, apperror.ErrStorageEmpty) {
			_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusNotFound))
			return
		}
		_ = render.Render(w, r, ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, users)
}
