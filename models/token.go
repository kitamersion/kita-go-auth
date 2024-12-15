package models

import (
	"time"
)

type RefreshToken struct {
	ID        string    `gorm:"primaryKey;index;type:uuid;" json:"id"`
	UserId    UserId    `gorm:"index" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Token     string    `gorm:"unique" json:"token"`
}
