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

// role and permission join tables
type UserRole struct {
	UserId string `gorm:"primaryKey;type:uuid;index;" json:"user_id"`
	RoleId RoleId `gorm:"primaryKey;type:uuid;index;" json:"role_id"`
}

type RolePermission struct {
	RoleId       RoleId      `gorm:"primaryKey;type:uuid;index;" json:"role_id"`
	PermissionId PermssionId `gorm:"primaryKey;type:uuid;index;" json:"permission_id"`
}
