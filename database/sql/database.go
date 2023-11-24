package sql

import (
	"tracking/database/interfaces"
	"tracking/database/sql/stores"
)

type SqlDatabase struct {
	usersStore databaseInterfaces.UsersStore
}

func New() *SqlDatabase {
	return &SqlDatabase{
		usersStore: usersStore.New(),
	}
}

func (database *SqlDatabase) Users() databaseInterfaces.UsersStore {
	return database.usersStore
}
