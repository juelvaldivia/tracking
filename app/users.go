package app

import (
	fmt "fmt"

	entities "tracking/app/entities"
)

func (app *App) FindAllUsers() ([]entities.User, error) {
	return app.Database.Users().All()
}

func (app *App) CreateUser(id string, name string, username string, password string) (entities.User, error) {
	newUser := entities.User{
		Id:       id,
		Name:     name,
		User:     username,
		Password: password,
	}

	err := app.Database.Users().Create(newUser)
	if err != nil {
		return entities.User{}, err
	}

	return newUser, nil
}

func (app *App) FindUserById(id string) entities.User {
	user, user_error := app.Database.Users().FindById(id)

	if user_error != nil {
		fmt.Printf("%v", user_error)
	}

	return user
}
