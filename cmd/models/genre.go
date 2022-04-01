package models

import "time"

type Genre struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GameGenre struct {
	ID        int       `json:"id"`
	GameID    string    `json:"game_id"`
	GenreID   string    `json:"genre_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
