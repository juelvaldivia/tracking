package main


import (
	"fmt"

	"tracking/app"
	"tracking/config"
	"tracking/database"
)

func main() {
	configuration, err := config.LoadConfig("config.yml")

	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	databaseFactory := tracking_database.NewFactory(configuration)
	database, err := databaseFactory.BuildDatabase()

	if err != nil {
		fmt.Printf("Error building database: %v\n", err)
		return
	}

	trackingApp := app.New(configuration, database)
	trackingApp.Start()
}
