package models

type Role struct {
	ID     string `gorm:"primaryKey;type:uuid;default" json:"id"`
	UserID string `gorm:"type:uuid;not null;index;" json:"user_id"`
	Role   string `gorm:"not null" json:"role"`

	// Define the relationship to the User model
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"user"`
}
