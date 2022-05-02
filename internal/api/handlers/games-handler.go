package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/repository"
	rawgservice "games-shelf-api-go/internal/service"
	"games-shelf-api-go/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
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

func AddGame(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	var payload GamePayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	game, err := mapPayloadToGame(payload)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	err = shelf.AddGame(game)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	res := Response{Ok: true}
	err = utils.WriteJSON(writer, http.StatusOK, res, "response")
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}
}

func EditGame(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	var payload GamePayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	game, err := mapPayloadToGame(payload)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusBadRequest)
		return
	}

	err = shelf.EditGame(game)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	res := Response{Ok: true}
	err = utils.WriteJSON(writer, http.StatusOK, res, "response")
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}
}

func DeleteGame(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utils.WriteErrorJSON(writer, errors.New("invalid id parameter"), http.StatusBadRequest)
		return
	}

	err = shelf.DeleteGame(id)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	res := Response{Ok: true}
	err = utils.WriteJSON(writer, http.StatusOK, res, "response")
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}
}

func GetGame(shelf *repository.Shelf, rawgService *rawgservice.RawgService, writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		utils.WriteErrorJSON(writer, errors.New("invalid id parameter"), http.StatusBadRequest)
		return
	}

	game, err := shelf.GetGameById(id)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	rawgDetails, err := rawgService.GetGameDetails(ctx, game.RawgId)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}
	game.RawgDetails = &rawgDetails

	err = utils.WriteJSON(writer, http.StatusOK, game, "game")
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}
}

func GetAllGames(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	genreId, _ := strconv.Atoi(queryValues.Get("genre_id"))
	platformId, _ := strconv.Atoi(queryValues.Get("platform_id"))

	games, err := shelf.GetAllGames(genreId, platformId)
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(writer, http.StatusOK, games, "games")
	if err != nil {
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}
}

func mapPayloadToGame(payload GamePayload) (models.Game, error) {
	id, err := strconv.Atoi(payload.ID)
	if err != nil {
		return models.Game{}, errors.New("invalid id parameter")
	}

	year, err := strconv.Atoi(payload.Year)
	if err != nil {
		return models.Game{}, errors.New("invalid year parameter")
	}

	rating, err := strconv.Atoi(payload.Rating)
	if err != nil {
		return models.Game{}, errors.New("invalid rating parameter")
	}

	return models.Game{
		ID:          id,
		Title:       payload.Title,
		Description: payload.Description,
		Year:        year,
		Publisher:   payload.Publisher,
		Rating:      rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
