package store

import "github.com/pyankovzhe/lingo/auth/internal/app/model"

type AccountRepository interface {
	Create(*model.Account) error
	Find(int) (*model.Account, error)
	FindByLogin(string) (*model.Account, error)
}
