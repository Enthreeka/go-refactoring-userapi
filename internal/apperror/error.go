package apperror

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

var (
	ErrUserNotExist = NewError("not found user with the specified id", errors.New("user_not_exist"))
	ErrStorageEmpty = NewError("not found users in storage", errors.New("storage_empty"))
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Error() string {
	return fmt.Sprintf("%s:%v", e.StatusText, e.Err)
}

func NewError(statusText string, err error) *ErrResponse {
	return &ErrResponse{
		Err:        err,
		StatusText: statusText,
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
