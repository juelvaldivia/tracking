package memory

import (
	interfaces "tracking/database/interfaces"
	usersStore "tracking/database/memory/stores"
)

type MemoryDatabase struct {
	usersStore interfaces.UsersStore
}

func New() *MemoryDatabase {
	return &MemoryDatabase{
		usersStore: usersStore.New(),
	}
}

func (database *MemoryDatabase) Users() interfaces.UsersStore {
	return database.usersStore
}
