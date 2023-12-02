package sql

import (
	interfaces "tracking/database/interfaces"
	usersStore "tracking/database/sql/stores"
)

type SqlDatabase struct {
	usersStore interfaces.UsersStore
}

func New() *SqlDatabase {
	return &SqlDatabase{
		usersStore: usersStore.New(),
	}
}

func (database *SqlDatabase) Users() interfaces.UsersStore {
	return database.usersStore
}
