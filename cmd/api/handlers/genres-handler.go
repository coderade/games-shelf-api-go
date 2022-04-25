package handlers

import (
	"games-shelf-api-go/cmd/api/utils"
	"games-shelf-api-go/cmd/models"
	"net/http"
)

func GetAllGenres(shelf *models.Shelf, writer http.ResponseWriter, request *http.Request) {

	genres, err := shelf.GetAllGenres()
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJson(writer, http.StatusOK, genres, "genres")
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}
}
