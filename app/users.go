package app

import (
	fmt "fmt"

	entities "tracking/app/entities"

	"github.com/google/uuid"
)

func (app *App) FindAllUsers() ([]entities.User, error) {
	return app.Database.Users().All()
}

func (app *App) CreateUser(id uuid.UUID, name string, username string, password string) (entities.User, error) {
	newUser := entities.User{
		Id:       id,
		Name:     name,
		Username: username,
		Password: password,
	}

	err := app.Database.Users().Create(newUser)
	if err != nil {
		return entities.User{}, err
	}

	return newUser, nil
}

func (app *App) FindUserById(id uuid.UUID) entities.User {
	user, user_error := app.Database.Users().FindById(id)

	if user_error != nil {
		fmt.Printf("%v", user_error)
	}

	return user
}
