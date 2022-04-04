package models

import "time"

type Platform struct {
	ID         int       `json:"-"`
	Name       string    `json:"name"`
	Generation string    `json:"generation"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type GamePlatform struct {
	ID         int       `json:"-"`
	GameID     int       `json:"-"`
	PlatformID int       `json:"-"`
	Platform   Platform  `json:"platform"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
