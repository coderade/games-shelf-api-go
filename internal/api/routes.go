package api

import (
	"context"
	"games-shelf-api-go/internal/api/handlers"
	"games-shelf-api-go/internal/models"
	rawgservice "games-shelf-api-go/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (s *Server) wrap(next http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := context.WithValue(request.Context(), "params", params)
		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}

func (s *Server) Routes() http.Handler {

	router := httprouter.New()
	secure := alice.New(s.validateJWTToken)
	rawgService := rawgservice.NewRawgService(s.Config.Rawg, s.Logger)

	router.HandlerFunc(http.MethodGet, "/status", func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(w, r, s.Config)
	})

	// public routes
	// /v1/games
	router.HandlerFunc(http.MethodGet, "/v1/games", s.handleRequest(handlers.GetAllGames))

	router.HandlerFunc(http.MethodGet, "/v1/games/:id", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGame(s.Shelf, rawgService, w, r)
	})

	router.HandlerFunc(http.MethodGet, "/v1/genres", s.handleRequest(handlers.GetAllGenres))
	router.HandlerFunc(http.MethodGet, "/v1/platforms", s.handleRequest(handlers.GetAllPlatforms))

	//auth routes
	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignIn(w, r, s.Config)
	})

	// private routes
	router.PUT("/v1/games/edit/:id", s.wrap(secure.ThenFunc(s.handleRequest(handlers.EditGame))))
	router.DELETE("/v1/games/delete/:id", s.wrap(secure.ThenFunc(s.handleRequest(handlers.DeleteGame))))
	router.POST("/v1/games/add", s.wrap(secure.ThenFunc(s.handleRequest(handlers.AddGame))))

	// graphql routes
	router.HandlerFunc(http.MethodPost, "/v1/graphql", s.handleRequest(handlers.GamesGraphQL))

	return s.enableCORS(router)

}

type RequestHandlerFunction func(shelf *models.Shelf, w http.ResponseWriter, r *http.Request)

func (s *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(s.Shelf, w, r)
	}
}
