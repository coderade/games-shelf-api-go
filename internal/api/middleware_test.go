package api

import (
	"games-shelf-api-go/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableCORS(t *testing.T) {
	cfg := &config.Config{}
	server := &Server{Config: cfg}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	corsHandler := server.enableCORS(handler)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	corsHandler.ServeHTTP(rr, req)

	assert.Equal(t, "*", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET,POST,PUT,DELETE,OPTIONS", rr.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type,Authorization", rr.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestValidateJWTToken(t *testing.T) {
	secret := "testsecret"
	cfg := &config.Config{Secret: secret}
	server := &Server{Config: cfg}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	validateHandler := server.validateJWTToken(handler)

	t.Run("missing authorization header", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		validateHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Contains(t, rr.Body.String(), "missing authorization header")
	})

	t.Run("invalid authorization header format", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.NoError(t, err)
		req.Header.Set("Authorization", "InvalidFormat")

		rr := httptest.NewRecorder()
		validateHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Contains(t, rr.Body.String(), "invalid authorization header format")
	})

	t.Run("invalid token", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.NoError(t, err)
		req.Header.Set("Authorization", "Bearer invalidtoken")

		rr := httptest.NewRecorder()
		validateHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Contains(t, rr.Body.String(), "unauthorized - invalid token")
	})

}
