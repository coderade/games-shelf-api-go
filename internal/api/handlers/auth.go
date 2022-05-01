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
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

func generatePasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}

func generateJWTSecret(cfg *config.Config) string {
	secret := cfg.Secret
	data := "games-shelf-api"

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

var validUser = models.User{
	ID:       1,
	Email:    "me@here.com",
	Password: generatePasswordHash("pass"),
}

func SignIn(writer http.ResponseWriter, request *http.Request, cfg *config.Config) {
	var credentials models.Credentials

	err := json.NewDecoder(request.Body).Decode(&credentials)

	if err != nil {
		println(errors.New("error decoding credentials"))
		utils.WriteErrorJson(writer, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		println(errors.New("unauthorized"))
		utils.WriteErrorJson(writer, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, _ := claims.HMACSign(jwt.HS256, []byte(generateJWTSecret(cfg)))

	token := string(jwtBytes)
	utils.WriteJson(writer, http.StatusOK, token, "token")
}
