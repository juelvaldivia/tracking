package interfaces

import entities "tracking/app/entities"

type UsersStore interface {
	All() ([]entities.User, error)
	Create(user entities.User) error
	FindById(id string) (entities.User, error)
}

type Database interface {
	Users() UsersStore
}
