package models

import (
	"database/sql"
	"time"
)

type User struct {
	Email       string       `gorm:"unique;index" json:"email"`
	Password    string       `json:"password"`
	ID          string       `gorm:"primaryKey;type:uuid;index;" json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ActivatedAt sql.NullTime `json:"activated_at"`

	// Define the relationship to the Role model
	Roles []Role `gorm:"foreignKey:UserID;references:ID;" json:"roles"`
}
