package handlers

import (
	"encoding/json"
	"errors"
	"games-shelf-api-go/cmd/api/service"
	"games-shelf-api-go/cmd/api/utils"
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

func AddGame(shelf *models.Shelf, writer http.ResponseWriter, request *http.Request) {

	var payload GamePayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
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

	err = shelf.AddGame(game)

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	res := response{Ok: true}

	err = utils.WriteJson(writer, http.StatusOK, res, "response")

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

}

func EditGame(shelf *models.Shelf, writer http.ResponseWriter, request *http.Request) {

	var payload GamePayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
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

	err = shelf.EditGame(game)

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	res := response{Ok: true}

	err = utils.WriteJson(writer, http.StatusOK, res, "response")

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

}

func DeleteGame(shelf *models.Shelf, writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		println(errors.New("invalid id parameter"))
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	err = shelf.DeleteGame(id)

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	res := response{Ok: true}

	err = utils.WriteJson(writer, http.StatusOK, res, "response")

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}
}

func GetGame(shelf *models.Shelf, writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		println(errors.New("invalid id parameter"))
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	game, err := shelf.GetGameById(id)

	game.RawgDetails = rawg_service.GetGameDetails(game.RawgId)

	err = utils.WriteJson(writer, http.StatusOK, game, "game")

	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}
}

func GetAllGames(shelf *models.Shelf, writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	genreId, err := strconv.Atoi(queryValues.Get("genre_id"))
	platformId, err := strconv.Atoi(queryValues.Get("platform_id"))
	games, err := shelf.GetAllGames(genreId, platformId)
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJson(writer, http.StatusOK, games, "games")
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}
}
