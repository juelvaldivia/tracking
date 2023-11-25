package app

import (
	config "tracking/config"
	interfaces "tracking/database/interfaces"
)

type App struct {
	Config   config.Config
	Database interfaces.Database
}

func New(configuration config.Config, database interfaces.Database) *App {
	return &App{
		Config:   configuration,
		Database: database,
	}
}
