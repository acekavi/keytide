package models

import (
	"time"

	"github.com/google/uuid"
)

type UserConsent struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	ClientID  uuid.UUID `json:"client_id" validate:"required"`
	Scope     []string  `json:"scope"`
	GrantedAt time.Time `json:"granted_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
