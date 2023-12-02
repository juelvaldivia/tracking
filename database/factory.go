package database

import (
	errors "errors"
	fmt "fmt"

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
		// TODO: implement SQLConfig
		fmt.Println(sqlConfig)

		database = sql.New()
	default:
		return nil, ErrUnsupportedDatabase
	}

	return database, nil
}
