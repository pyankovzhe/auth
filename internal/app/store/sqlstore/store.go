package sqlstore

import (
	"database/sql"

	"github.com/pyankovzhe/auth/internal/app/store"
)

type Store struct {
	db          *sql.DB
	accountRepo *AccountRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Account() store.AccountRepository {
	if s.accountRepo != nil {
		return s.accountRepo
	}

	s.accountRepo = &AccountRepository{
		store: s,
	}

	return s.accountRepo
}
