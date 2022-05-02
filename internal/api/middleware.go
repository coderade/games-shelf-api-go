package api

import (
	"errors"
	"games-shelf-api-go/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

// enableCORS adds Cross-Origin Resource Sharing (CORS) headers to the response.
func (s *Server) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(writer, request)
	})
}

// validateJWTToken validates the JWT token in the Authorization header.
func (s *Server) validateJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Vary", "Authorization")
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			utils.WriteErrorJson(writer, errors.New("missing authorization header"), http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.WriteErrorJson(writer, errors.New("invalid authorization header format"), http.StatusUnauthorized)
			return
		}

		token := headerParts[1]
		claims, err := jwt.HMACCheck([]byte(token), []byte(s.Config.Secret))
		if err != nil {
			utils.WriteErrorJson(writer, errors.New("unauthorized - invalid token"), http.StatusUnauthorized)
			return
		}

		if !claims.Valid(time.Now()) {
			utils.WriteErrorJson(writer, errors.New("unauthorized - token expired"), http.StatusUnauthorized)
			return
		}

		if !claims.AcceptAudience("mydomain.com") {
			utils.WriteErrorJson(writer, errors.New("unauthorized - invalid audience"), http.StatusUnauthorized)
			return
		}

		if claims.Issuer != "mydomain.com" {
			utils.WriteErrorJson(writer, errors.New("unauthorized - invalid issuer"), http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			utils.WriteErrorJson(writer, errors.New("unauthorized - invalid user ID"), http.StatusUnauthorized)
			return
		}

		log.Printf("Authenticated user ID: %d", userID)
		next.ServeHTTP(writer, request)
	})
}
