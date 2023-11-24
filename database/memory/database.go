package memory

import (
	"tracking/database/interfaces"
	"tracking/database/memory/stores"
)

type MemoryDatabase struct {
	usersStore databaseInterfaces.UsersStore
}

func New() *MemoryDatabase {
	return &MemoryDatabase{
		usersStore: usersStore.New(),
	}
}

func (database *MemoryDatabase) Users() databaseInterfaces.UsersStore {
	return database.usersStore
}
