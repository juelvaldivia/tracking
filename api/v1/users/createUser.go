package usersController

import (
	json "encoding/json"
	http "net/http"

	responses "tracking/api/responses"
	entities "tracking/app/entities"
)

func (controller *UsersController) CreateUser(response http.ResponseWriter, request *http.Request) {
	var newUser entities.User

	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := controller.App.CreateUser(newUser.Id, newUser.Name, newUser.User, newUser.Password)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	responses.Json(response, http.StatusCreated, user)
}
