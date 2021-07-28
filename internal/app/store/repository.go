package store

import (
	"github.com/google/uuid"
	"github.com/pyankovzhe/auth/internal/app/model"
)

type AccountRepository interface {
	Create(*model.Account) error
	Find(uuid.UUID) (*model.Account, error)
	FindByLogin(string) (*model.Account, error)
}
