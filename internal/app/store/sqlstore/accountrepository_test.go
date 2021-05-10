package sqlstore_test

import (
	"testing"

	"github.com/pyankovzhe/lingo/auth/internal/app/model"
	"github.com/pyankovzhe/lingo/auth/internal/app/store"
	"github.com/pyankovzhe/lingo/auth/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("accounts")

	s := sqlstore.New(db)
	a := model.TestAccount(t)
	assert.NoError(t, s.Account().Create(a))
	assert.NotNil(t, a)
}

func TestAccountRepository_FindByLogin(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("accounts")

	s := sqlstore.New(db)
	login := "login1"
	_, err := s.Account().FindByLogin(login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	a := model.TestAccount(t)
	a.Login = login
	s.Account().Create(a)
	a, err = s.Account().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestAccountRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("accounts")

	s := sqlstore.New(db)
	_, err := s.Account().Find(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	a := model.TestAccount(t)
	s.Account().Create(a)
	a, err = s.Account().Find(a.ID)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}
