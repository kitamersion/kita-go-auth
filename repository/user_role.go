package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func CreateUserRole(userRole *models.UserRole) (models.UserRole, error) {
	err := initializers.DB.Create(&userRole).Error
	if err != nil {
		return models.UserRole{}, err
	}
	return *userRole, nil
}

func FetchUserRolesByUserId(userId string) ([]models.UserRole, error) {
	var userRoles []models.UserRole
	err := initializers.DB.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return []models.UserRole{}, err
	}

	return userRoles, nil
}

func DeleteUserRole(userId string, roleId string) error {
	err := initializers.DB.
		Where("user_id = ?", userId).
		Where("role_id = ?", roleId).
		Delete(&models.UserRole{}).Error
	return err
}
