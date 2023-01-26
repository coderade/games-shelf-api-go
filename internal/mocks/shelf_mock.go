package mocks

import (
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/repository"

	"github.com/stretchr/testify/mock"
)

// MockShelf is a mock implementation of the ShelfRepository
type MockShelf struct {
	mock.Mock
}

// Ensure MockShelf implements the ShelfRepository interface
var _ repository.ShelfRepository = (*MockShelf)(nil)

func (m *MockShelf) GetAllGenres() ([]models.Genre, error) {
	args := m.Called()
	return args.Get(0).([]models.Genre), args.Error(1)
}

func (m *MockShelf) GetAllPlatforms() ([]models.Platform, error) {
	args := m.Called()
	return args.Get(0).([]models.Platform), args.Error(1)
}

func (m *MockShelf) GetGameById(id int) (*models.Game, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Game), args.Error(1)
}

func (m *MockShelf) GetAllGames(genreID int, platformID int) ([]models.Game, error) {
	args := m.Called(genreID, platformID)
	return args.Get(0).([]models.Game), args.Error(1)
}

func (m *MockShelf) AddGame(game models.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockShelf) EditGame(game models.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockShelf) DeleteGame(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockShelf) GetGenresAndPlatformsByGameId(id int) ([]models.Genre, []models.Platform, error) {
	args := m.Called(id)
	return args.Get(0).([]models.Genre), args.Get(1).([]models.Platform), args.Error(2)
}
