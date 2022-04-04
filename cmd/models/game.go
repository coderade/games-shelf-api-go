package models

import "time"

type Game struct {
	ID           int            `json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Year         int            `json:"year"`
	Publisher    string         `json:"publisher"`
	Rating       string         `json:"rating"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	GameGenre    []GameGenre    `json:"genre"`
	GamePlatform []GamePlatform `json:"platform"`
}
