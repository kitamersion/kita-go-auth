package models

import "errors"

type (
	RoleType string
	RoleId   string
)

const (
	Admin RoleType = "ADMIN"
	Basic RoleType = "BASIC"
	Guest RoleType = "GUEST"
)

type Role struct {
	ID   RoleId   `gorm:"primaryKey;type:uuid;index;" json:"id"`
	Role RoleType `json:"role"`
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
