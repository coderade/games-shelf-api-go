package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

// generatePasswordHash hashes the password using bcrypt.
func generatePasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashedPassword)
}

// generateJWTSecret creates a JWT secret using HMAC with SHA256.
func generateJWTSecret(cfg *config.Config) string {
	secret := cfg.Secret
	data := "games-shelf-api"

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

// SignIn handles user sign-in and generates a JWT token if authentication is successful.
func SignIn(writer http.ResponseWriter, request *http.Request, cfg *config.Config) {
	var credentials models.Credentials

	// Decode credentials from request body
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Error decoding credentials: %v", err)
		utils.WriteErrorJSON(writer, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// Validate user credentials
	user, err := validateUser(credentials)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		utils.WriteErrorJSON(writer, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.ID, cfg)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		utils.WriteErrorJSON(writer, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(writer, http.StatusOK, token, "token")
}

// validateUser checks if the provided credentials are valid.
func validateUser(credentials models.Credentials) (*models.User, error) {
	// For simplicity, I am using a hardcoded valid user. In production, fetch from a database.
	validUser := models.User{
		ID:       1,
		Email:    "me@here.com",
		Password: generatePasswordHash("pass"),
	}

	if credentials.Email != validUser.Email {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(credentials.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &validUser, nil
}

// generateJWT generates a JWT token for the authenticated user.
func generateJWT(userID int, cfg *config.Config) (string, error) {
	var claims jwt.Claims
	claims.Subject = fmt.Sprint(userID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(generateJWTSecret(cfg)))
	if err != nil {
		return "", err
	}

	return string(jwtBytes), nil
}
