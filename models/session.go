package models

import "time"

type Session struct {
	SessionID  string     `json:"session_id"`
	UserID     string     `json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at"`
}
