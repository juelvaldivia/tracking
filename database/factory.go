package tracking_database

import (
	"fmt"
	"errors"

	"tracking/config"
	"tracking/database/interfaces"
	"tracking/database/memory"
	"tracking/database/sql"
)

var (
	FactoryDatabaseError = errors.New("error building database")
	UnsupportedDatabase = errors.New("unsupported database type")
  InvalidConfig       = errors.New("invalid database configuration")
)

type Factory struct {
	Config config.Config
}

func NewFactory(configuration config.Config) *Factory {
	return &Factory{
		Config: configuration,
	}
}

func (factory *Factory) BuildDatabase() (databaseInterfaces.Database, error) {
	driver := factory.Config.DatabaseDriver

	var database databaseInterfaces.Database
  var err error

	switch driver {
	case "memory":
		database = memory.New()
	case "sql":
		sqlConfig := factory.Config.SQLDatabase
		fmt.Println(sqlConfig)

		database = sql.New()
	default:
		return nil, fmt.Errorf("%w: %s", UnsupportedDatabase, driver)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %v", FactoryDatabaseError, err)
	}

	return database, nil
}
