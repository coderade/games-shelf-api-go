package handlers

import (
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/repository"
	"games-shelf-api-go/internal/utils"
	"log"
	"net/http"
)

// GetAllGenres handles the request to get all genres.
func GetAllGenres(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	genres, err := shelf.GetAllGenres()
	if err != nil {
		log.Printf("Error fetching genres: %v", err)
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	writeGenresResponse(writer, genres)
}

// writeGenresResponse writes the genres response to the writer.
func writeGenresResponse(writer http.ResponseWriter, genres []models.Genre) {
	err := utils.WriteJSON(writer, http.StatusOK, genres, "genres")
	if err != nil {
		log.Printf("Error writing genres response: %v", err)
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
	}
}
