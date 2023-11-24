package usersStore

import (
	"tracking/app/entities"
)

type UsersStore struct {
	Users map[string]entities.User
}

// TODO: implement sql connection
func New() *UsersStore {
	return &UsersStore{
		Users: make(map[string]entities.User),
	}
}

func (store *UsersStore) Create(user entities.User) error {
	// TODO: implement sql query for create users
	return nil
}

func (store *UsersStore) FindById(id string) (entities.User, error) {
	// TODO: implement sql query for find user by id
	return entities.User{}, nil
}
