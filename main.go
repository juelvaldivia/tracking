package main

import (
	fmt "fmt"

	api "tracking/api"
	app "tracking/app"
	config "tracking/config"
	database "tracking/database"
	server "tracking/server"
)

func main() {
	configuration, err := config.LoadConfig("config.yml")

	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	databaseFactory := database.NewFactory(configuration)
	database, err := databaseFactory.BuildDatabase()

	if err != nil {
		fmt.Printf("Error building database: %v\n", err)
		return
	}

	trackingApp := app.New(configuration, database)
	apiInstance := api.New(trackingApp)
	server := server.NewServerHTTP(configuration.ApiPort, apiInstance.Router)

	serverError := server.Listen()

	if serverError != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
