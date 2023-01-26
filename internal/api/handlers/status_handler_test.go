package handlers

import (
	"games-shelf-api-go/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	cfg := &config.Config{Env: "development"}

	req, err := http.NewRequest(http.MethodGet, "/status", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		StatusHandler(w, r, cfg)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Available")
	assert.Contains(t, rr.Body.String(), "development")
	assert.Contains(t, rr.Body.String(), version)
}

func TestWriteStatusResponse(t *testing.T) {
	status := AppStatus{
		Status:      "Available",
		Environment: "development",
		Version:     version,
	}

	rr := httptest.NewRecorder()
	writeStatusResponse(rr, status)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Available")
	assert.Contains(t, rr.Body.String(), "development")
	assert.Contains(t, rr.Body.String(), version)
}
