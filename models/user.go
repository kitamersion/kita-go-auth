package models

import (
	"database/sql"
	"time"
)

type User struct {
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	LastLoginAt time.Time    `json:"last_login_at"`
	ActivatedAt sql.NullTime `json:"activated_at"`
	Email       string       `gorm:"unique;index" json:"email"`
	Password    string       `json:"password"`
	ID          string       `gorm:"primaryKey;type:uuid;index;" json:"id"`
}
