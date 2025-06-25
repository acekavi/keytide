package models

import (
	"time"

	"github.com/google/uuid"
)

type ExternalIdentity struct {
	ID             uuid.UUID      `json:"id" validate:"required"`
	UserID         uuid.UUID      `json:"user_id" validate:"required"`
	Provider       string         `json:"provider" validate:"required"`
	ProviderUserID string         `json:"provider_user_id" validate:"required"`
	Email          string         `json:"email" validate:"omitempty,email"`
	Metadata       map[string]any `json:"metadata"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
