package api

import (
	http "net/http"
	testing "testing"
	helpers "tracking/tests/api/helpers"

	assert "github.com/stretchr/testify/assert"
)

func TestWelcome(t *testing.T) {
	api := helpers.CreateAPI(t)

	response := helpers.Request(api.Router, "GET", "/health", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.JSONEq(t, "{\"ok\":true}", response.Body.String())
}
