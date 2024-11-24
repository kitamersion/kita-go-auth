package models

import "errors"

type RoleType string

const (
	Admin RoleType = "ADMIN"
	Basic RoleType = "BASIC"
	Guest RoleType = "GUEST"
)

type Role struct {
	ID     string   `gorm:"primaryKey;type:uuid;index;" json:"id"`
	UserID string   `gorm:"type:uuid;not null;index;" json:"user_id"`
	Role   RoleType `gorm:"type:varchar(20);not null" json:"role"`
}

func (r RoleType) IsValid() bool {
	switch r {
	case Admin, Basic, Guest:
		return true
	default:
		return false
	}
}

func (r *Role) ValidateRole() error {
	if !r.Role.IsValid() {
		return errors.New("invalid role")
	}
	return nil
}
