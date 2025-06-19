package models

import "time"

type User struct {
	ID         string     `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Email      string     `db:"email" json:"email"`
	Password   string     `db:"password" json:"-"` // "-" tag omits from JSON responses
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	ArchivedAt *time.Time `db:"archived_at" json:"archivedAt,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileRequest struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}

type ProfileResponse struct {
	Message string `json:"Message"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
	Email   string `json:"Email"`
}
