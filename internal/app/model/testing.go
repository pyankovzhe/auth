package model

import "testing"

func TestAccount(t *testing.T) *Account {
	return &Account{
		Login:    "testacc",
		Password: "password123",
	}
}
