package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func CreatePermission(permission models.Permission) (models.Permission, error) {
	err := initializers.DB.Create(&permission).Error
	if err != nil {
		return models.Permission{}, err
	}
	return permission, nil
}

func DeletePermissionsByPermissionId(permissionId string) error {
	err := initializers.DB.Where("id = ?", permissionId).Delete(&models.Permission{}).Error
	return err
}

func DeletePermissionsByUserId(userId string) error {
	err := initializers.DB.Where("user_id = ?", userId).Delete(&models.Permission{}).Error
	return err
}

func FetchPermissionsByUserId(userId string) ([]models.Permission, error) {
	var permissions []models.Permission
	err := initializers.DB.Where("user_id = ?", userId).Find(&permissions).Error
	if err != nil {
		return []models.Permission{}, err
	}

	return permissions, nil
}
