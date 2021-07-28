package sqlstore

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/pyankovzhe/auth/internal/app/model"
	"github.com/pyankovzhe/auth/internal/app/store"
)

type AccountRepository struct {
	store *Store
}

func (r *AccountRepository) Create(a *model.Account) error {
	if err := a.EncryptPassword(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO accounts (login, encrypted_password) VALUES ($1, $2) RETURNING id",
		a.Login,
		a.EncryptedPassword,
	).Scan(&a.ID)
}

func (r *AccountRepository) FindByLogin(login string) (*model.Account, error) {
	a := &model.Account{}

	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM accounts WHERE login = $1",
		login,
	).Scan(&a.ID, &a.Login, &a.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return a, nil
}

func (r *AccountRepository) Find(id uuid.UUID) (*model.Account, error) {
	a := &model.Account{}

	if err := r.store.db.QueryRow(
		"SELECT id, login, encrypted_password FROM accounts WHERE id = $1",
		id,
	).Scan(&a.ID, &a.Login, &a.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return a, nil
}

// func (r *Repo) GetUsers(ctx context.Context) ([]repository.User, error) {
// 	rows, err := r.db.QueryContext(ctx, `
// 		SELECT id, name, email FROM users
// 	`)
// 	defer rows.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	var users []repository.User

// 	for rows.Next() {
// 		var u repository.User
// 		if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
// 			return nil, err
// 		}

// 		users = append(users, u)
// 	}

// 	return users, rows.Err()
// }
