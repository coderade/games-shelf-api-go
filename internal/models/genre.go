package models

import (
	"time"
)

// Genre represents a game genre.
type Genre struct {
	ID        int       `json:"id"`   // Unique identifier for the genre
	Name      string    `json:"name"` // Name of the genre
	CreatedAt time.Time `json:"-"`    // Timestamp when the genre record was created
	UpdatedAt time.Time `json:"-"`    // Timestamp when the genre record was last updated
}

// GameGenre represents the relationship between a game and a genre.
type GameGenre struct {
	ID        int       `json:"-"`     // Unique identifier for the game-genre relationship
	GameID    int       `json:"-"`     // Unique identifier for the game
	GenreID   int       `json:"-"`     // Unique identifier for the genre
	Genre     Genre     `json:"genre"` // Embedded genre details
	CreatedAt time.Time `json:"-"`     // Timestamp when the relationship was created
	UpdatedAt time.Time `json:"-"`     // Timestamp when the relationship was last updated
}
