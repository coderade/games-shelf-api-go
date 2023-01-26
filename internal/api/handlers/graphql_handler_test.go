package handlers

import (
	"bytes"
	"errors"
	"games-shelf-api-go/internal/mocks"
	"games-shelf-api-go/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGamesGraphQL(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	expectedGames := []models.Game{
		{ID: 1, Title: "Test Game 1"},
		{ID: 2, Title: "Test Game 2"},
	}

	mockShelf.On("GetAllGames", 0, 0).Return(expectedGames, nil)
	mockShelf.On("GetGameById", mock.AnythingOfType("int")).Return(&models.Game{}, nil)
	mockShelf.On("GetAllGenres").Return([]models.Genre{}, nil)
	mockShelf.On("GetAllPlatforms").Return([]models.Platform{}, nil)

	query := `{ list { id title } }`
	req, err := http.NewRequest(http.MethodPost, "/v1/graphql", bytes.NewBufferString(query))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GamesGraphQL(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test Game 1")
	assert.Contains(t, rr.Body.String(), "Test Game 2")

	mockShelf.AssertExpectations(t)
}

func TestGamesGraphQL_SchemaError(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockShelf.On("GetAllGames", 0, 0).Return(nil, errors.New("schema error"))

	query := `{ list { id title } }`
	req, err := http.NewRequest(http.MethodPost, "/v1/graphql", bytes.NewBufferString(query))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GamesGraphQL(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "schema error")

	mockShelf.AssertExpectations(t)
}

func TestGamesGraphQL_BodyError(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockShelf.On("GetAllGames", 0, 0).Return([]models.Game{}, nil)
	mockShelf.On("GetGameById", mock.AnythingOfType("int")).Return(&models.Game{}, nil)
	mockShelf.On("GetAllGenres").Return([]models.Genre{}, nil)
	mockShelf.On("GetAllPlatforms").Return([]models.Platform{}, nil)

	req, err := http.NewRequest(http.MethodPost, "/v1/graphql", nil) // Sending a nil body to simulate error
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GamesGraphQL(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid request body")

	mockShelf.AssertExpectations(t)
}

func TestGamesGraphQL_GraphQLError(t *testing.T) {
	mockShelf := new(mocks.MockShelf)
	mockShelf.On("GetAllGames", 0, 0).Return([]models.Game{}, nil)
	mockShelf.On("GetGameById", mock.AnythingOfType("int")).Return(&models.Game{}, nil)
	mockShelf.On("GetAllGenres").Return([]models.Genre{}, nil)
	mockShelf.On("GetAllPlatforms").Return([]models.Platform{}, nil)

	query := `{ invalidQuery }` // An invalid query to simulate GraphQL error
	req, err := http.NewRequest(http.MethodPost, "/v1/graphql", bytes.NewBufferString(query))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GamesGraphQL(mockShelf, w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "GraphQL errors")

	mockShelf.AssertExpectations(t)
}
