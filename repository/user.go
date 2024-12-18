package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func CreateUser(user models.User) (models.User, error) {
	err := initializers.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func FetchUserById(userId models.UserId) (models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func FetchUserByEmail(email string) (models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUserById(userId models.UserId, userUpdates models.User) (models.User, error) {
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

func DeleteUserById(userId models.UserId) error {
	err := initializers.DB.Where("id = ?", userId).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FetchAllUsers() ([]models.User, error) {
	var users []models.User
	err := initializers.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
