package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}
