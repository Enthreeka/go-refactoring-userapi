package apperror

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotExist = NewError("not found user with the specified id", errors.New("user_not_exist"))
	ErrStorageEmpty = NewError("not found users in storage", errors.New("storage_empty"))
)

type AppError struct {
	Msg string `json:"message"`
	Err error  `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s:%v", e.Msg, e.Err)
}

func NewError(msg string, err error) *AppError {
	return &AppError{
		Err: err,
		Msg: msg,
	}
}
