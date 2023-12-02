package usersStore

import entities "tracking/app/entities"

type UsersStore struct {
	Users map[string]entities.User
}

// TODO: implement sql connection
func New() *UsersStore {
	return &UsersStore{
		Users: make(map[string]entities.User),
	}
}

func (store *UsersStore) All() ([]entities.User, error) {
	usersList := make([]entities.User, 0, len(store.Users))

	for _, user := range store.Users {
		usersList = append(usersList, user)
	}

	return usersList, nil
}

func (store *UsersStore) Create(user entities.User) error {
	// TODO: implement sql query for create users
	return nil
}

func (store *UsersStore) FindById(id string) (entities.User, error) {
	// TODO: implement sql query for find user by id
	return entities.User{}, nil
}
