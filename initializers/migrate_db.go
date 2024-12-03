package initializers

import (
	"log"

	"github.com/kitamersion/kita-go-auth/models"
)

func MigrateDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error running migrations for User: %v", err)
	}

	err = DB.AutoMigrate(&models.RefreshToken{})
	if err != nil {
		log.Fatalf("Error running migrations for RefreshToken: %v", err)
	}

	err = DB.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatalf("Error running migrations for Role: %v", err)
	}

	err = DB.AutoMigrate(&models.Permission{})
	if err != nil {
		log.Fatalf("Error running migrations for Permission: %v", err)
	}

	err = DB.AutoMigrate(&models.UserRole{})
	if err != nil {
		log.Fatalf("Error running migrations for Permission: %v", err)
	}

	err = DB.AutoMigrate(&models.RolePermission{})
	if err != nil {
		log.Fatalf("Error running migrations for Permission: %v", err)
	}

	log.Println("Migrations complete!")
}
