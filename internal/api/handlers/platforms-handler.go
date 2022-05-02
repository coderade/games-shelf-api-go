package handlers

import (
	"games-shelf-api-go/internal/models"
	"games-shelf-api-go/internal/repository"
	"games-shelf-api-go/internal/utils"
	"log"
	"net/http"
)

// GetAllPlatforms handles the request to get all platforms.
func GetAllPlatforms(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	platforms, err := shelf.GetAllPlatforms()
	if err != nil {
		log.Printf("Error fetching platforms: %v", err)
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
		return
	}

	writePlatformsResponse(writer, platforms)
}

// writePlatformsResponse writes the platforms response to the writer.
func writePlatformsResponse(writer http.ResponseWriter, platforms []models.Platform) {
	err := utils.WriteJSON(writer, http.StatusOK, platforms, "platforms")
	if err != nil {
		log.Printf("Error writing platforms response: %v", err)
		utils.WriteErrorJSON(writer, err, http.StatusInternalServerError)
	}
}
