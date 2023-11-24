package config

import (
	"fmt"
	"errors"

	"github.com/spf13/viper"
)

var (
	FailedReadConfigFile 			= errors.New("failed to read config file")
	FailedBuildConfig 				= errors.New("failed to unmarshal config")
	MissingDatabaseConfigKeys = errors.New("missing required keys in sql_database configuration")
)

type Config struct {
	DatabaseDriver string `mapstructure:"database_driver"`
	SQLDatabase    SQLConfig `mapstructure:"sql_database"`
}

type SQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func LoadConfig(configPath string) (Config, error) {
	var configuration Config

	if err := readAndUnmarshalConfig(configPath, &configuration); err != nil {
		return Config{}, err
	}

	if configuration.DatabaseDriver == "sql" {
		if err := validateSQLDatabaseConfig(configuration.SQLDatabase); err != nil {
			return Config{}, err
		}
	}

	return configuration, nil
}

func readAndUnmarshalConfig(configPath string, configuration *Config) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("%w: %v", FailedReadConfigFile, err)
	}

	if err := viper.Unmarshal(configuration); err != nil {
		return fmt.Errorf("%w: %v", FailedBuildConfig, err)
	}

	return nil
}

func validateSQLDatabaseConfig(sqlConfig SQLConfig) error {
	if sqlConfig.Host == "" ||
		sqlConfig.Port == 0 ||
		sqlConfig.Username == "" ||
		sqlConfig.Password == "" ||
		sqlConfig.Database == "" {
		return fmt.Errorf("%w expected: %v", MissingDatabaseConfigKeys, expectDatabaseConfig())
	}

	return nil
}

func expectDatabaseConfig() string {
	return `
	sql_database:
  	  host: localhost
  	  port: 5432
  	  username: your_username
  	  password: your_password
  	  database: your_database
	`
}
