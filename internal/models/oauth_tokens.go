package models

import (
	"time"

	"github.com/google/uuid"
)

type OAuthToken struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	ClientID  uuid.UUID `json:"client_id" validate:"required"`
	TokenType string    `json:"token_type"` // e.g., access, refresh, id
	Token     string    `json:"token"`      // Opaque or JWT fingerprint
	Scope     []string  `json:"scope"`      // Scopes associated with the token
	ExpiresAt time.Time `json:"expires_at"` // Expiration time of the token
	IssuedAt  time.Time `json:"issued_at"`  // Issued time of the token
	RevokedAt time.Time `json:"revoked_at"` // Revocation time of the token
}
