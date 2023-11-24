package databaseInterfaces

import "tracking/app/entities"

type UsersStore interface {
	Create(user entities.User) error
	FindById(id string) (entities.User, error)
}

type Database interface {
	Users() UsersStore
}
