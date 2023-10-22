package http

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error `json:"error"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	//ErrorText  string `json:"error_text,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	fmt.Println(e.Err)
	return nil
}

func ErrInvalidRequest(err error, httpStatus int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: httpStatus,
		StatusText:     "Invalid request.",
		//ErrorText:      err.Error(),
	}
}
