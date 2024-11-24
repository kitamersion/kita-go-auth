package initializers

import (
	"log"

	"github.com/kitamersion/kita-go-auth/models"
)

func MigrateDatabase() {
	// Migrate User model
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error running migrations for User: %v", err)
	}

	// Migrate RefreshToken model
	err = DB.AutoMigrate(&models.RefreshToken{})
	if err != nil {
		log.Fatalf("Error running migrations for RefreshToken: %v", err)
	}

	// Migrate Role model
	err = DB.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatalf("Error running migrations for Role: %v", err)
	}

	log.Println("Migrations complete!")
}
