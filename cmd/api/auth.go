package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"games-shelf-api-go/cmd/models"
	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func generatePasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}

func (app *application) generateJWTSecret() string {
	secret := app.config.secret
	data := "games-shelf-api"

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

var validUser = models.User{ID: 1, Email: "me@here.com", Password: generatePasswordHash("password")}

func (app *application) SignIn(writer http.ResponseWriter, request *http.Request) {
	var credentials models.Credentials

	err := json.NewDecoder(request.Body).Decode(&credentials)

	if err != nil {
		app.logger.Println(errors.New("error decoding credentials"))
		app.errorJSON(writer, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		app.logger.Println(errors.New("unauthorized"))
		app.errorJSON(writer, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.secret))
	app.writeJSON(writer, http.StatusOK, string(jwtBytes), "token")
}
