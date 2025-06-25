package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthorizationCode struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	ClientID    uuid.UUID `json:"client_id" validate:"required"`
	Code        string    `json:"code"`
	RedirectURI string    `json:"redirect_uri"`
	Scope       []string  `json:"scope"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}
