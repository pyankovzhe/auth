package apiserver

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	errInvalidAccessToken       = errors.New("Invalid access token")
	errNotAuthenticated         = errors.New("Not authenticated")
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type ErrResponse struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}
