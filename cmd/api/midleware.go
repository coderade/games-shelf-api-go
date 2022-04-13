package main

import (
	"errors"
	"github.com/pascaldekloe/jwt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(writer, request)
	})
}

func (app *application) validateJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Vary", "Authorization")
		authHeader := request.Header.Get("Authorization")

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			app.errorJSON(writer, errors.New("invalid Auth Header"), http.StatusUnauthorized)
			return
		}

		if headerParts[0] != "Bearer" {
			app.errorJSON(writer, errors.New("unauthorized - no Bearer"), http.StatusUnauthorized)
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.secret))

		if err != nil {
			app.errorJSON(writer, errors.New("unauthorized - Failed hmac check"), http.StatusUnauthorized)
			return
		}

		if !claims.Valid(time.Now()) {
			app.errorJSON(writer, errors.New("unauthorized - Token expired"), http.StatusUnauthorized)
			return
		}

		if !claims.AcceptAudience("mydomain.com") {
			app.errorJSON(writer, errors.New("unauthorized - Invalid Audience"), http.StatusUnauthorized)
			return
		}

		if claims.Issuer != "mydomain.com" {
			app.errorJSON(writer, errors.New("unauthorized - Invalid Issuer"), http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)

		if err != nil {
			app.errorJSON(writer, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}

		log.Println("Valid user: ", userID)

		next.ServeHTTP(writer, request)
	})
}
