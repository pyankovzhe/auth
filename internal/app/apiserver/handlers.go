package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pyankovzhe/auth/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

type accountRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type accountResponse struct {
	*model.Account
}

type signingResponse struct {
	Token string `json:"token"`
}

func (res *accountResponse) Render(w http.ResponseWriter, r *http.Request) error {
	res.Account.Sanitize()
	return nil
}

func (res *signingResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	req := &accountRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	a := &model.Account{
		Login:    req.Login,
		Password: req.Password,
	}

	// TODO: take out from handlers
	if err := s.store.Account().Create(a); err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
		return
	}

	// TODO: produce message to kafka
	accInBytes, err := json.Marshal(a)
	if err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
		return
	}

	s.producer.Publish([]byte(accInBytes))

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &accountResponse{Account: a})
}

func (s *server) SignIn(w http.ResponseWriter, r *http.Request) {
	req := &accountRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	a, err := s.store.Account().FindByLogin(req.Login)
	if err != nil || !(bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(req.Password)) == nil) {
		render.Render(w, r, &ErrResponse{Code: http.StatusUnauthorized, Message: errIncorrectEmailOrPassword.Error()})
		return
	}

	tokenStr, err := GenerateToken(a.Login)
	if err != nil {
		render.Render(w, r, &ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &signingResponse{Token: tokenStr})
}

func (s *server) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	a := ctx.Value(ctxKeyAccount).(*model.Account)

	render.Status(r, http.StatusOK)
	render.Render(w, r, &accountResponse{Account: a})
}
