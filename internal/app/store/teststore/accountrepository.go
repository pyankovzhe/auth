package teststore

import (
	"github.com/google/uuid"
	"github.com/pyankovzhe/auth/internal/app/model"
	"github.com/pyankovzhe/auth/internal/app/store"
)

type AccountRepository struct {
	store    *Store
	accounts map[uuid.UUID]*model.Account
}

func (r *AccountRepository) Create(a *model.Account) error {
	if err := a.Validate(); err != nil {
		return err
	}

	if err := a.EncryptPassword(); err != nil {
		return err
	}

	a.ID = uuid.New()
	r.accounts[a.ID] = a

	return nil
}

func (r *AccountRepository) Find(id uuid.UUID) (*model.Account, error) {
	a, ok := r.accounts[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return a, nil
}

func (r *AccountRepository) FindByLogin(login string) (*model.Account, error) {
	for _, acc := range r.accounts {
		if acc.Login == login {
			return acc, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
