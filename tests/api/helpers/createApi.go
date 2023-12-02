package helpers

import (
	testing "testing"
	api "tracking/api"
	app "tracking/app"
	entities "tracking/app/entities"
	config "tracking/config"
	database "tracking/database"

	assert "github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type MockApp struct {
	mock.Mock
}

func (m *MockApp) FindAllUsers() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func CreateAPI(t *testing.T) *api.API {
	testConfig := config.Config{DatabaseDriver: "memory"}
	databaseFactory := database.NewFactory(testConfig)

	database, err := databaseFactory.BuildDatabase()
	assert.NoError(t, err, "Error building database")

	app := app.New(testConfig, database)

	return api.New(app)
}

func CreateAPIWithApp(t *testing.T, app *app.App) *api.API {
	return api.New(app)
}
