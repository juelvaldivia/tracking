package entities

import (
	testing "testing"
	entities "tracking/app/entities"

	assert "github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := entities.User{
		Id:       "1",
		Name:     "John Doe",
		User:     "john_doe",
		Password: "password123",
	}

	assert.Equal(t, "1", user.Id, "User Id is not equal")
	assert.Equal(t, "John Doe", user.Name, "Name is not equal")
	assert.Equal(t, "john_doe", user.User, "User is not equal")
	assert.Equal(t, "password123", user.Password, "Password is not equal")
}
