package models

import "time"

type Platform struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Generation string    `json:"generation"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GamePlatform struct {
	ID         int       `json:"id"`
	GameID     string    `json:"game_id"`
	PlatformID string    `json:"platform_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
