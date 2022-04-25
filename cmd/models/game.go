package models

import (
	rawgservice "games-shelf-api-go/cmd/api/service"
	"time"
)

type Game struct {
	ID           int                    `json:"id"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	Year         int                    `json:"year"`
	Publisher    string                 `json:"publisher"`
	Rating       int                    `json:"rating"`
	RawgId       string                 `json:"rawg_id"`
	RawgDetails  rawgservice.GameResult `json:"details"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	GameGenre    []Genre                `json:"genres"`
	GamePlatform []Platform             `json:"platforms"`
}
