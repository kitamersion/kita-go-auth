package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func FetchRolesByRoleType(roleType models.RoleType) (models.Role, error) {
	var roles models.Role
	err := initializers.DB.Where("role = ?", roleType).First(&roles).Error
	if err != nil {
		return models.Role{}, err
	}

	return roles, nil
}

func FetchRolesByRoleIds(roleId []models.RoleId) ([]models.Role, error) {
	var roles []models.Role
	err := initializers.DB.Where("id in ?", roleId).Find(&roles).Error
	if err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}
