package models

import (
	"time"

	"github.com/google/uuid"
)

type OAuthClient struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	Name         string    `json:"name"`
	ClientID     uuid.UUID `json:"client_id" validate:"required"`
	ClientSecret string    `json:"client_secret"`
	RedirectURIs []string  `json:"redirect_uris"`
	GrantTypes   []string  `json:"grant_types"`
	Scopes       []string  `json:"scopes"`
	OwnerUserID  uuid.UUID `json:"owner_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
