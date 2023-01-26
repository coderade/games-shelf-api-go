package handlers

import (
	"bytes"
	"encoding/json"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	// Setup
	cfg := &config.Config{
		Secret: "testsecret",
	}

	tests := []struct {
		name           string
		credentials    models.Credentials
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "valid credentials",
			credentials: models.Credentials{
				Email:    "me@here.com",
				Password: "pass",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "token",
		},
		{
			name: "invalid email",
			credentials: models.Credentials{
				Email:    "invalid@here.com",
				Password: "pass",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "error",
		},
		{
			name: "invalid password",
			credentials: models.Credentials{
				Email:    "me@here.com",
				Password: "wrongpass",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "error",
		},
		{
			name: "invalid JSON",
			credentials: models.Credentials{
				Email:    "",
				Password: "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			body, _ := json.Marshal(tt.credentials)
			req, err := http.NewRequest("POST", "/signin", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			// Create ResponseRecorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				SignIn(w, r, cfg)
			})
			handler.ServeHTTP(rr, req)

			// Check the status code
			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status %v, got %v", tt.expectedStatus, rr.Code)

			// Check the response body
			var responseBody map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
			if err != nil {
				t.Fatal(err)
			}

			_, exists := responseBody[tt.expectedBody]
			assert.True(t, exists, "Expected body to contain %v", tt.expectedBody)
		})
	}
}
