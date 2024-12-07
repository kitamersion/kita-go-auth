package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func CreateRole(role models.Role) (models.Role, error) {
	err := initializers.DB.Create(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func DeleteRolesByUserId(userId string) error {
	err := initializers.DB.Where("user_id = ?", userId).Delete(&models.Role{}).Error
	return err
}

func DeleteRoleByRoleId(roleId string) error {
	err := initializers.DB.Where("id = ?", roleId).Delete(&models.Role{}).Error
	return err
}

func FetchRolesByUserId(userId string) ([]models.Role, error) {
	var roles []models.Role
	err := initializers.DB.Where("user_id = ?", userId).Find(&roles).Error
	if err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}
