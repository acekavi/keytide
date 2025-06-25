package models

import (
	"time"

	"github.com/google/uuid"
)

// Role represents a user role in the RBAC system.
type Role struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"` // Unique name for the role
	Description string    `json:"description"`              // Description of the role
	CreatedAt   time.Time `json:"created_at"`               // Timestamp when the role was created
	UpdatedAt   time.Time `json:"updated_at"`               // Timestamp when the role was last updated
}

// Permission represents a permission in the RBAC system.
type Permission struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"` // Unique name for the permission
	Description string    `json:"description"`              // Description of the permission
	CreatedAt   time.Time `json:"created_at"`               // Timestamp when the permission was created
	UpdatedAt   time.Time `json:"updated_at"`               // Timestamp when the permission was last updated
}

// RolePermission associates a role with a permission.
type RolePermission struct {
	RoleID       uuid.UUID `json:"role_id" validate:"required"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at"` // Timestamp when the association was created
	UpdatedAt    time.Time `json:"updated_at"` // Timestamp when the association was last updated
}

// UserRole associates a user with a role.
type UserRole struct {
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	RoleID    uuid.UUID `json:"role_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"` // Timestamp when the association was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the association was last updated
}
