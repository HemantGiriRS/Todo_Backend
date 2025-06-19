package models

import "time"

type todo struct {
	ID          string     `db:"id" json:"id"`
	User_id     string     `db:"user_id" json:"user_id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	IsCompleted bool       `db:"is_completed" json:"is_completed"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	ArchivedAt  *time.Time `db:"archived_at" json:"archived_at"`
}

type AddTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTask struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateTask struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
type UpdateStatus struct {
	Id          string `json:"id"`
	IsCompleted bool   `json:"is_completed"`
}
