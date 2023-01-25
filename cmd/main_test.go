package main

import (
	"context"
	"net/http"
	"syscall"
	"testing"
	"time"

	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/db"
	"games-shelf-api-go/internal/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStartServer(t *testing.T) {
	mockDB := new(db.MockDatabase)
	mockDB.On("PingContext", mock.Anything).Return(nil).Once()
	mockDB.On("Close").Return(nil).Once()

	// Mock configuration and logger
	cfg := &config.Config{
		Port:     "4000",
		LogLevel: "info",
		Db:       config.DatabaseConfig{Dsn: "postgres://admin@localhost/games_shelf?sslmode=disable"},
		Rawg:     config.RawgConfig{ApiKey: "dummy", ApiEndpoint: "https://api.rawg.io/api"},
	}
	log = logger.NewLogger(cfg.LogLevel)

	// Start the server
	server := startServer()

	assert.NotNil(t, server)
	assert.Equal(t, ":4000", server.Addr)

	// Allow some time for the server to start
	time.Sleep(1 * time.Second)

	// Create a test request to check if the server is running
	resp, err := http.Get("http://localhost:4000/")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Simulate server shutdown
	go func() {
		time.Sleep(1 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	// Wait for server to shut down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Fatalf("Server forced to shutdown: %v", err)
	}

	// Assert server is no longer running
	resp, err = http.Get("http://localhost:4000/")
	assert.Error(t, err)

	mockDB.AssertExpectations(t)
}

func TestGracefulShutdown(t *testing.T) {
	mockDB := new(db.MockDatabase)
	mockDB.On("PingContext", mock.Anything).Return(nil).Once()
	mockDB.On("Close").Return(nil).Once()

	// Mock configuration and logger
	cfg := &config.Config{
		Port:     "4000",
		LogLevel: "info",
		Db:       config.DatabaseConfig{Dsn: "postgres://admin@localhost/games_shelf?sslmode=disable"},
		Rawg:     config.RawgConfig{ApiKey: "dummy", ApiEndpoint: "https://api.rawg.io/api"},
	}
	log = logger.NewLogger(cfg.LogLevel)

	// Start the server
	server := startServer()

	// Simulate server shutdown
	go func() {
		time.Sleep(1 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	// Wait for the shutdown signal
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Fatalf("Server forced to shutdown: %v", err)
	}

	assert.Nil(t, server.Shutdown(ctx))

	mockDB.AssertExpectations(t)
}
