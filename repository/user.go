package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

// CreateUser creates a new user in the database.
func CreateUser(user models.User) (models.User, error) {
	err := initializers.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// FetchUserById retrieves a user by their ID.
func FetchUserById(userId string) (models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// FetchUserByEmail retrieves a user by their email.
func FetchUserByEmail(email string) (models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateUserById updates a user record using the provided User struct.
func UpdateUserById(userId string, userUpdates models.User) (models.User, error) {
	var user models.User

	// Ensure we only update the fields provided in userUpdates
	err := initializers.DB.Model(&models.User{}).
		Where("id = ?", userId).
		Select("Email", "Password", "UpdatedAt", "ActivatedAt", "LastLoginAt").
		Updates(models.User{
			Email:       userUpdates.Email,
			Password:    userUpdates.Password,
			UpdatedAt:   userUpdates.UpdatedAt,
			ActivatedAt: userUpdates.ActivatedAt,
			LastLoginAt: userUpdates.LastLoginAt,
		}).Error
	if err != nil {
		return models.User{}, err
	}

	// Re-fetch the updated user to ensure consistency
	err = initializers.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// DeleteUserById removes a user by their ID.
func DeleteUserById(userId string) error {
	err := initializers.DB.Where("id = ?", userId).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

// FetchAllUsers retrieves all users from the database.
func FetchAllUsers() ([]models.User, error) {
	var users []models.User
	err := initializers.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
