package handlers

import (
	"games-shelf-api-go/internal/mocks"
	"games-shelf-api-go/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllGenres(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	expectedGenres := []models.Genre{
		{ID: 1, Name: "Action"},
		{ID: 2, Name: "Adventure"},
	}

	mockShelf.On("GetAllGenres").Return(expectedGenres, nil)

	req, err := http.NewRequest(http.MethodGet, "/v1/genres", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetAllGenres(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Action")
	assert.Contains(t, rr.Body.String(), "Adventure")

	mockShelf.AssertExpectations(t)
}

func TestGetAllGenres_Error(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockShelf.On("GetAllGenres").Return(nil, assert.AnError)

	req, err := http.NewRequest(http.MethodGet, "/v1/genres", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetAllGenres(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), assert.AnError.Error())

	mockShelf.AssertExpectations(t)
}
