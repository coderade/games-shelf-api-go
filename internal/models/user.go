package models

// User represents a user in the system.
type User struct {
	ID       int    `json:"id"`       // Unique identifier for the user
	Email    string `json:"email"`    // Email address of the user
	Password string `json:"password"` // Hashed password of the user
}

// Credentials represents user credentials for authentication.
type Credentials struct {
	Email    string `json:"email"`    // Email address of the user
	Password string `json:"password"` // Plain text password provided by the user
}
