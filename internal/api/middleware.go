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

func (api *Server) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(writer, request)
	})
}

func (api *Server) validateJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Vary", "Authorization")
		authHeader := request.Header.Get("Authorization")

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			utils.WriteErrorJson(writer, errors.New("invalid Auth Header"), http.StatusUnauthorized)
			return
		}

		if headerParts[0] != "Bearer" {
			utils.WriteErrorJson(writer, errors.New("unauthorized - no Bearer"), http.StatusUnauthorized)
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(api.Config.Secret))

		if err != nil {
			utils.WriteErrorJson(writer, errors.New("unauthorized - Failed hmac check"), http.StatusUnauthorized)
			return
		}

		if !claims.Valid(time.Now()) {
			utils.WriteErrorJson(writer, errors.New("unauthorized - Token expired"), http.StatusUnauthorized)
			return
		}

		if !claims.AcceptAudience("mydomain.com") {
			utils.WriteErrorJson(writer, errors.New("unauthorized - Invalid Audience"), http.StatusUnauthorized)
			return
		}

		if claims.Issuer != "mydomain.com" {
			utils.WriteErrorJson(writer, errors.New("unauthorized - Invalid Issuer"), http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)

		if err != nil {
			utils.WriteErrorJson(writer, errors.New("unauthorized"), http.StatusUnauthorized)
			return
		}

		log.Println("Valid user: ", userID)

		next.ServeHTTP(writer, request)
	})
}
