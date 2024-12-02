package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func CreateRolePermission(rolePermission *models.RolePermission) (models.RolePermission, error) {
	err := initializers.DB.Create(&rolePermission).Error
	if err != nil {
		return models.RolePermission{}, err
	}
	return *rolePermission, nil
}

func FetchRolePermissionsByRoleId(roleId string) ([]models.RolePermission, error) {
	var rolePermissions []models.RolePermission
	err := initializers.DB.Where("role_id = ?", roleId).Find(&rolePermissions).Error
	if err != nil {
		return []models.RolePermission{}, err
	}

	return rolePermissions, nil
}

func DeleteRolePermissoinsByRoleId(roleId string) error {
	err := initializers.DB.Where("role_id = ?", roleId).Delete(&models.RolePermission{}).Error
	return err
}
