package usersController

import (
	http "net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	responses "tracking/api/responses"
)

func (controller *UsersController) FindUserById(response http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "id")

	id, err := uuid.Parse(userId)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	user := controller.App.FindUserById(id)

	responses.Json(response, http.StatusOK, user)
}
