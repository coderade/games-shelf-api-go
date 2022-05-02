package models

import (
	rawgservice "games-shelf-api-go/internal/service"
	"time"
)

// Game represents a video game entity with its details.
type Game struct {
	ID           int                     `json:"id"`                // Unique identifier for the game
	Title        string                  `json:"title"`             // Title of the game
	Description  string                  `json:"description"`       // Description of the game
	Year         int                     `json:"year"`              // Release year of the game
	Publisher    string                  `json:"publisher"`         // Publisher of the game
	Rating       int                     `json:"rating"`            // Rating of the game
	RawgId       string                  `json:"rawg_id"`           // ID of the game in the RAWG API
	RawgDetails  *rawgservice.GameResult `json:"details,omitempty"` // Details from the RAWG API
	CreatedAt    time.Time               `json:"created_at"`        // Timestamp when the game record was created
	UpdatedAt    time.Time               `json:"updated_at"`        // Timestamp when the game record was last updated
	GameGenre    []Genre                 `json:"genres"`            // List of genres associated with the game
	GamePlatform []Platform              `json:"platforms"`         // List of platforms associated with the game
}
