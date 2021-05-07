package model_test

import (
	"testing"

	"github.com/pyankovzhe/lingo/auth/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestAccount_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		acc     func() *model.Account
		isValid bool
	}{
		{
			name: "valid",
			acc: func() *model.Account {
				return model.TestAccount(t)
			},
			isValid: true,
		},
		{
			name: "with empty login",
			acc: func() *model.Account {
				a := model.TestAccount(t)
				a.Login = ""
				return a
			},
			isValid: false,
		},
		{
			name: "with short login",
			acc: func() *model.Account {
				a := model.TestAccount(t)
				a.Login = "lo"
				return a
			},
			isValid: false,
		},
		{
			name: "with empty password",
			acc: func() *model.Account {
				a := model.TestAccount(t)
				a.Password = ""
				return a
			},
			isValid: false,
		},
		{
			name: "with short password",
			acc: func() *model.Account {
				a := model.TestAccount(t)
				a.Password = "pa"
				return a
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.acc().Validate())
			} else {
				assert.Error(t, tc.acc().Validate())
			}
		})
	}
}

func TestAccount_EncryptPassword(t *testing.T) {
	acc := model.TestAccount(t)
	assert.NoError(t, acc.EncryptPassword())
	assert.NotEmpty(t, acc.EncryptedPassword)
}
