package usersController

import (
	http "net/http"

	chi "github.com/go-chi/chi/v5"

	responses "tracking/api/responses"
)

func (controller *UsersController) FindUserById(response http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "id")

	user := controller.App.FindUserById(userId)

	responses.Json(response, http.StatusOK, user)
}
