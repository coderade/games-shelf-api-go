package handlers

import (
	"encoding/json"
	"games-shelf-api-go/internal/config"
	"log"
	"net/http"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

const version = "1.0.0"

// StatusHandler handles the request to get the application's status.
func StatusHandler(writer http.ResponseWriter, request *http.Request, cfg *config.Config) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: cfg.Env,
		Version:     version,
	}

	writeStatusResponse(writer, currentStatus)
}

// writeStatusResponse writes the status response to the writer.
func writeStatusResponse(writer http.ResponseWriter, status AppStatus) {
	j, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		log.Printf("Error marshalling status response: %v", err)
		http.Error(writer, "Failed to generate status response", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(j)
	if err != nil {
		log.Printf("Error writing status response: %v", err)
		http.Error(writer, "Failed to write status response", http.StatusInternalServerError)
	}
}
