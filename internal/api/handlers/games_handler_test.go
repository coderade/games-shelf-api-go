package handlers

import (
	"bytes"
	"encoding/json"
	"games-shelf-api-go/internal/mocks"
	"games-shelf-api-go/internal/models"
	rawgservice "games-shelf-api-go/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddGame(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	payload := GamePayload{
		ID:          "1",
		Title:       "Test Game",
		Description: "Test Description",
		Year:        "2021",
		Publisher:   "Test Publisher",
		Rating:      "5",
	}
	game, _ := mapPayloadToGame(payload)

	mockShelf.On("AddGame", game).Return(nil)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/v1/games", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddGame(mockShelf, w, r)
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockShelf.AssertExpectations(t)
}

func TestEditGame(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	payload := GamePayload{
		ID:          "1",
		Title:       "Test Game",
		Description: "Test Description",
		Year:        "2021",
		Publisher:   "Test Publisher",
		Rating:      "5",
	}
	game, _ := mapPayloadToGame(payload)

	mockShelf.On("EditGame", game).Return(nil)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/v1/games/1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EditGame(mockShelf, w, r)
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockShelf.AssertExpectations(t)
}

func TestDeleteGame(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockShelf.On("DeleteGame", 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/v1/games/1", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteGame(mockShelf, w, r)
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockShelf.AssertExpectations(t)
}

func TestGetGame(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockRawgService := new(mocks.MockRawgService)
	game := &models.Game{
		ID:          1,
		Title:       "Test Game",
		Description: "Test Description",
		Year:        2021,
		Publisher:   "Test Publisher",
		Rating:      5,
		RawgId:      "test-rawg-id",
	}
	rawgDetails := rawgservice.GameResult{
		ID:          1,
		Slug:        "test-slug",
		Description: "test-description",
	}

	mockShelf.On("GetGameById", 1).Return(game, nil)
	mockRawgService.On("GetGameDetails", mock.Anything, "test-rawg-id").Return(rawgDetails, nil)

	req, _ := http.NewRequest("GET", "/v1/games/1", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetGame(mockShelf, mockRawgService, w, r)
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockShelf.AssertExpectations(t)
	mockRawgService.AssertExpectations(t)
}

func TestGetAllGames(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	games := []models.Game{
		{
			ID:          1,
			Title:       "Test Game",
			Description: "Test Description",
			Year:        2021,
			Publisher:   "Test Publisher",
			Rating:      5,
		},
	}

	mockShelf.On("GetAllGames", 0, 0).Return(games, nil)

	req, _ := http.NewRequest("GET", "/v1/games", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetAllGames(mockShelf, w, r)
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockShelf.AssertExpectations(t)
}
