package teststore_test

import (
	"testing"

	"github.com/pyankovzhe/lingo/auth/internal/app/model"
	"github.com/pyankovzhe/lingo/auth/internal/app/store"
	"github.com/pyankovzhe/lingo/auth/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_Create(t *testing.T) {
	s := teststore.New()
	a := model.TestAccount(t)
	assert.NoError(t, s.Account().Create(a))
	assert.NotNil(t, a.ID)
}

func TestAccountRepository_FindByLogin(t *testing.T) {
	s := teststore.New()
	login := "user123"
	_, err := s.Account().FindByLogin(login)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	a := model.TestAccount(t)
	a.Login = login
	s.Account().Create(a)
	a, err = s.Account().FindByLogin(login)
	assert.NoError(t, err)
	assert.NotNil(t, a)
}
