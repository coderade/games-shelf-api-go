package api

import (
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/db"
	"games-shelf-api-go/internal/logger"
	"testing"

	_ "github.com/lib/pq" // Import PostgreSQL driver for tests
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitialize(t *testing.T) {
	mockDB := new(db.MockDatabase)
	mockDB.On("PingContext", mock.Anything).Return(nil).Once()

	// Mock configuration and logger
	cfg := &config.Config{
		Port:     "4000",
		LogLevel: "info",
		Db:       config.DatabaseConfig{Dsn: "postgres://admin@localhost/games_shelf?sslmode=disable"},
		Rawg:     config.RawgConfig{ApiKey: "dummy", ApiEndpoint: "https://api.rawg.io/api"},
	}
	log := logger.NewLogger(cfg.LogLevel)

	// Initialize the server with the mock database
	server := &Server{}
	server.Initialize(cfg, log, mockDB)

	assert.NotNil(t, server.Shelf)
	assert.Equal(t, cfg, server.Config)
	assert.Equal(t, log, server.Logger)
	assert.Equal(t, mockDB, server.DBHelper)

	mockDB.AssertExpectations(t)
}

func TestClose(t *testing.T) {
	mockDB := new(db.MockDatabase)
	mockDB.On("Close").Return(nil).Once()

	// Mock configuration and logger
	cfg := &config.Config{
		Port:     "4000",
		LogLevel: "info",
		Db:       config.DatabaseConfig{Dsn: "postgres://admin@localhost/games_shelf?sslmode=disable"},
		Rawg:     config.RawgConfig{ApiKey: "dummy", ApiEndpoint: "https://api.rawg.io/api"},
	}
	log := logger.NewLogger(cfg.LogLevel)

	// Initialize the server with the mock database
	server := &Server{}
	server.Initialize(cfg, log, mockDB)

	server.Close()
	mockDB.AssertExpectations(t)
}
