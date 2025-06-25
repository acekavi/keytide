package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID           uuid.UUID      `json:"id" validate:"required"`
	UserID       uuid.UUID      `json:"user_id"`
	Action       string         `json:"action" validate:"required"` // login, logout, token_issued, etc.
	ResourceType string         `json:"resource_type"`              // user, role, token, etc.
	ResourceID   string         `json:"resource_id"`
	IP           string         `json:"ip"`
	UserAgent    string         `json:"user_agent"`
	Metadata     map[string]any `json:"metadata"`
	CreatedAt    time.Time      `json:"created_at"`
}
