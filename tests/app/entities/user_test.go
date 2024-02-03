package entities

import (
	testing "testing"
	entities "tracking/app/entities"

	"github.com/google/uuid"
	assert "github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	id := uuid.New()

	user := entities.User{
		Id:       id,
		Name:     "John Doe",
		Username: "john_doe",
		Password: "password123",
	}

	assert.Equal(t, id, user.Id, "User Id is not equal")
	assert.Equal(t, "John Doe", user.Name, "Name is not equal")
	assert.Equal(t, "john_doe", user.Username, "User is not equal")
	assert.Equal(t, "password123", user.Password, "Password is not equal")
}
