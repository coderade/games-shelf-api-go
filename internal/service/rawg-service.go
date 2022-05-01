package rawgservice

import (
	"context"
	"encoding/json"
	"fmt"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/logger"
	"io"
	"net/http"
	"time"
)

type GameResult struct {
	ID              int     `json:"id"`
	Slug            string  `json:"slug"`
	Description     string  `json:"description"`
	Metacritic      int     `json:"metacritic"`
	MetacriticUrl   string  `json:"metacritic_url"`
	BackgroundImage string  `json:"background_image"`
	Publisher       string  `json:"publisher"`
	Rating          float32 `json:"rating"`
}

type RawgService struct {
	Config config.RawgConfig
	Logger *logger.Logger
	Client *http.Client
}

func NewRawgService(cfg config.RawgConfig, logger *logger.Logger) *RawgService {
	return &RawgService{
		Config: cfg,
		Logger: logger,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *RawgService) GetGameDetails(ctx context.Context, rawgId string) (GameResult, error) {
	path := fmt.Sprintf("games/%s", rawgId)
	return s.doRequest(ctx, path)
}

func (s *RawgService) doRequest(ctx context.Context, path string) (GameResult, error) {
	url := fmt.Sprintf("%s/%s?key=%s", s.Config.ApiEndpoint, path, s.Config.ApiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		s.Logger.Errorf("Failed to create request: %v", err)
		return GameResult{}, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		s.Logger.Errorf("Failed to perform request: %v", err)
		return GameResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Rawg request error - Path: %s : %s ", path, resp.Status)
		s.Logger.Error(errorMessage)
		return GameResult{}, fmt.Errorf(errorMessage)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.Logger.Errorf("Failed to read response body: %v", err)
		return GameResult{}, err
	}

	var gameResult GameResult
	err = json.Unmarshal(body, &gameResult)
	if err != nil {
		s.Logger.Errorf("Failed to unmarshal response body: %v", err)
		return GameResult{}, err
	}

	return gameResult, nil
}
