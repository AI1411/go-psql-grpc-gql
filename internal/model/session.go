package model

import "github.com/google/uuid"

type Session struct {
	SessionID string    `json:"session_id"`
	UserID    uuid.UUID `json:"user_id"`
}
