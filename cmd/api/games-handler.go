package main

import (
	"encoding/json"
	"errors"
	"games-shelf-api-go/cmd/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

type response struct {
	Ok bool `json:"ok"`
}

type GamePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Publisher   string `json:"publisher"`
	Rating      string `json:"rating"`
}

func (app *application) addGame(writer http.ResponseWriter, request *http.Request) {

	var payload GamePayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	var game models.Game
	game.ID, _ = strconv.Atoi(payload.ID)
	game.Title = payload.Title
	game.Description = payload.Description
	game.Year, _ = strconv.Atoi(payload.Year)
	game.Publisher = payload.Publisher
	game.Rating, _ = strconv.Atoi(payload.Rating)
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()

	err = app.shelf.AddGame(game)

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	res := response{Ok: true}

	err = app.writeJSON(writer, http.StatusOK, res, "response")

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

}

func (app *application) editGame(writer http.ResponseWriter, request *http.Request) {

	var payload GamePayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	var game models.Game
	game.ID, _ = strconv.Atoi(payload.ID)
	game.Title = payload.Title
	game.Description = payload.Description
	game.Year, _ = strconv.Atoi(payload.Year)
	game.Publisher = payload.Publisher
	game.Rating, _ = strconv.Atoi(payload.Rating)
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()

	err = app.shelf.EditGame(game)

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	res := response{Ok: true}

	err = app.writeJSON(writer, http.StatusOK, res, "response")

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

}

func (app *application) deleteGame(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(writer, err)
		return
	}

	err = app.shelf.DeleteGame(id)

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	res := response{Ok: true}

	err = app.writeJSON(writer, http.StatusOK, res, "response")

	if err != nil {
		app.errorJSON(writer, err)
		return
	}
}

func (app *application) getGame(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(writer, err)
		return
	}

	game, err := app.shelf.GetGameById(id)

	err = app.writeJSON(writer, http.StatusOK, game, "game")

	if err != nil {
		app.errorJSON(writer, err)
		return
	}
}

func (app *application) getAllGames(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	genreId, err := strconv.Atoi(queryValues.Get("genre_id"))
	platformId, err := strconv.Atoi(queryValues.Get("platform_id"))
	games, err := app.shelf.GetAllGames(genreId, platformId)
	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	err = app.writeJSON(writer, http.StatusOK, games, "games")
	if err != nil {
		app.errorJSON(writer, err)
		return
	}
}
