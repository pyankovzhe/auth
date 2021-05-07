package model

import (
	// validation "github.com/go-ozzo/ozzo-validation"
	// "github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (a *Account) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Login, validation.Required, validation.Length(5, 32)),
		validation.Field(&a.Password, validation.Required, validation.Length(6, 32)),
	)
}

func (a *Account) EncryptPassword() error {
	if len(a.Password) > 0 {
		enc, err := encryptString(a.Password)

		if err != nil {
			return err
		}

		a.EncryptedPassword = enc
	}

	return nil
}

func (a *Account) Sanitize() {
	a.Password = ""
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
