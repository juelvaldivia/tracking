package database

import (
	errors "errors"

	config "tracking/config"
	interfaces "tracking/database/interfaces"
	memory "tracking/database/memory"
	sql "tracking/database/sql"
)

var ErrUnsupportedDatabase = errors.New("unsupported database type")

type Factory struct {
	Config config.Config
}

func NewFactory(configuration config.Config) *Factory {
	return &Factory{
		Config: configuration,
	}
}

func (factory *Factory) BuildDatabase() (interfaces.Database, error) {
	driver := factory.Config.DatabaseDriver

	var database interfaces.Database

	switch driver {
	case "memory":
		database = memory.New()
	case "sql":
		sqlConfig := factory.Config.SQLDatabase

		database, err := sql.New(sqlConfig)

		if err != nil {
			return nil, ErrUnsupportedDatabase
		}

		return database, nil
	default:
		return nil, ErrUnsupportedDatabase
	}

	return database, nil
}
