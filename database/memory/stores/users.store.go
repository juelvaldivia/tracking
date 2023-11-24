package usersStore

import (
	"errors"

	"tracking/app/entities"
)

var (
	UserNotFound = errors.New("user not found")
)

type UsersStore struct {
	Users map[string]entities.User
}

func New() *UsersStore {
	return &UsersStore{
		Users: make(map[string]entities.User),
	}
}

func (store *UsersStore) Create(user entities.User) error {
	store.Users[user.Id] = user
	return nil
}

func (store *UsersStore) FindById(id string) (entities.User, error) {
	if user, exists := store.Users[id]; exists {
		return user, nil
	}

	return entities.User{}, UserNotFound
}
