package usersController

import (
	http "net/http"

	responses "tracking/api/responses"
)

func (controller *UsersController) GetUsers(response http.ResponseWriter, request *http.Request) {
	result, err := controller.App.FindAllUsers()

	if err != nil {
		responses.Json(response, http.StatusInternalServerError, err.Error())
		return
	}

	userMap := make(map[string]interface{})
	userMap["users"] = result

	responses.Json(response, http.StatusOK, userMap)
}
