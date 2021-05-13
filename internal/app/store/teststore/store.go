package teststore

import (
	"github.com/pyankovzhe/auth/internal/app/model"
	"github.com/pyankovzhe/auth/internal/app/store"
)

type Store struct {
	accountRepo *AccountRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Account() store.AccountRepository {
	if s.accountRepo != nil {
		return s.accountRepo
	}

	s.accountRepo = &AccountRepository{
		store:    s,
		accounts: make(map[int]*model.Account),
	}

	return s.accountRepo
}
