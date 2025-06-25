package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	Token      string    `json:"token"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	LastActive time.Time `json:"last_active"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
}
