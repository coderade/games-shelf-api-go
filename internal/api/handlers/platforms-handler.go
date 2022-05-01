package handlers

import (
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/utils"
	"net/http"
)

func GetAllPlatforms(shelf *models.Shelf, writer http.ResponseWriter, reader *http.Request) {
	platforms, err := shelf.GetAllPlatforms()
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}

	err = utils.WriteJson(writer, http.StatusOK, platforms, "platforms")
	if err != nil {
		utils.WriteErrorJson(writer, err, http.StatusBadRequest)
		return
	}
}
