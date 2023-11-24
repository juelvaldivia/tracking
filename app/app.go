package app

import (
	"fmt"

	"tracking/app/entities"
	"tracking/config"
	"tracking/database/interfaces"
)

type App struct {
	Config config.Config
	Database databaseInterfaces.Database
}

func New(configuration config.Config, database databaseInterfaces.Database) *App {
	return &App{
		Config: configuration,
		Database: database,
	}
}

func (app *App) Start() {
	// how to use database into App
	app.Database.Users().Create(entities.User{"1","Joel", "juel", "123"})
	user, user_error := app.Database.Users().FindById("1")

	if user_error != nil {
		fmt.Printf("%v", user_error)
	}

	// TODO: create API

	fmt.Println(user)
	fmt.Println("App started!")
}
