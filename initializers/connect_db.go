package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectedDb() {
	var err error
	// Get the database connection string from the environment variable
	dsn := os.Getenv("DATABASE")
	// Assign to the global DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}

	println("Connected to database")
}
