package handlers

import (
	"games-shelf-api-go/internal/mocks"
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPlatforms(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	expectedPlatforms := []models.Platform{
		{ID: 1, Name: "Platform 1", Generation: "Gen 1"},
		{ID: 2, Name: "Platform 2", Generation: "Gen 2"},
	}

	mockShelf.On("GetAllPlatforms").Return(expectedPlatforms, nil)

	req, err := http.NewRequest(http.MethodGet, "/v1/platforms", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetAllPlatforms(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Platform 1")
	assert.Contains(t, rr.Body.String(), "Platform 2")

	mockShelf.AssertExpectations(t)
}

func TestGetAllPlatforms_ErrorFetching(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockShelf.On("GetAllPlatforms").Return(nil, assert.AnError)

	req, err := http.NewRequest(http.MethodGet, "/v1/platforms", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetAllPlatforms(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "Error fetching platforms")

	mockShelf.AssertExpectations(t)
}

func TestWritePlatformsResponse(t *testing.T) {
	platforms := []models.Platform{
		{ID: 1, Name: "Platform 1", Generation: "Gen 1"},
		{ID: 2, Name: "Platform 2", Generation: "Gen 2"},
	}

	rr := httptest.NewRecorder()
	err := utils.WriteJSON(rr, http.StatusOK, platforms, "platforms")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Platform 1")
	assert.Contains(t, rr.Body.String(), "Platform 2")
}
