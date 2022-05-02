package models

import "time"

// Platform represents a gaming platform.
type Platform struct {
	ID         int       `json:"id"`         // Unique identifier for the platform
	Name       string    `json:"name"`       // Name of the platform
	Generation string    `json:"generation"` // Generation of the platform
	CreatedAt  time.Time `json:"-"`          // Timestamp when the platform record was created
	UpdatedAt  time.Time `json:"-"`          // Timestamp when the platform record was last updated
}

// GamePlatform represents the relationship between a game and a platform.
type GamePlatform struct {
	ID         int       `json:"-"`        // Unique identifier for the game-platform relationship
	GameID     int       `json:"-"`        // Unique identifier for the game
	PlatformID int       `json:"-"`        // Unique identifier for the platform
	Platform   Platform  `json:"platform"` // Embedded platform details
	CreatedAt  time.Time `json:"-"`        // Timestamp when the relationship was created
	UpdatedAt  time.Time `json:"-"`        // Timestamp when the relationship was last updated
}
