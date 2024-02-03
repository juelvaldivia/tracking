package usersStore

import (
	"fmt"
	entities "tracking/app/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UsersStore struct {
	*sqlx.DB
}

// TODO: implement sql connection
func New(connection *sqlx.DB) *UsersStore {
	return &UsersStore{
		DB: connection,
	}
}

func (store *UsersStore) All() ([]entities.User, error) {
	var users []entities.User

	if err := store.Select(&users, `SELECT * FROM users`); err != nil {
		return []entities.User{}, fmt.Errorf("error selecting users: %w", err)
	}

	return users, nil
}

func (store *UsersStore) Create(user entities.User) error {
	if err := store.Get(
		user, `INSERT INTO users (name, username, password) VALUES ($1, $2, $3, $4) RETURNING *`,
		user.Name,
		user.Username,
		user.Password,
	); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (store *UsersStore) FindById(id uuid.UUID) (entities.User, error) {
	var user entities.User

	if err := store.Get(&user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return entities.User{}, fmt.Errorf("error fetching user: %w", err)
	}

	return entities.User{}, nil
}
