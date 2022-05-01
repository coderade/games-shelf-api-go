package handlers

import (
	"encoding/json"
	"games-shelf-api-go/internal/config"
	"net/http"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

const version = "1.0.0"

func StatusHandler(writer http.ResponseWriter, reader *http.Request, cfg *config.Config) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: cfg.Env,
		Version:     version,
	}

	j, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(j)
	if err != nil {
		println(err)
	}
}
