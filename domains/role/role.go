package role

import (
	"errors"
	"fmt"

	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
)

func AssignRoleToUser(userId models.UserId, roleId models.RoleId) (models.UserRole, error) {
	if userId == "" {
		return models.UserRole{}, errors.New("userId cannot be empty")
	}

	if roleId == "" {
		return models.UserRole{}, errors.New("roleId cannot be empty")
	}

	userRoleRecord := models.UserRole{
		UserId: userId,
		RoleId: roleId,
	}

	record, err := repository.CreateUserRole(&userRoleRecord)
	if err != nil {
		return models.UserRole{}, err
	}
	return record, nil
}

func GetRolesByRoleType(roleType models.RoleType) (models.Role, error) {
	if roleType == "" {
		return models.Role{}, errors.New("roleType cannot be empty")
	}

	role, err := repository.FetchRolesByRoleType(roleType)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func RevokeRoleForUser(userId models.UserId, roleId models.RoleId) error {
	if roleId == "" {
		return errors.New("roleId cannot be empty")
	}
	if userId == "" {
		return errors.New("userId cannot be empty")
	}

	err := repository.DeleteUserRole(userId, roleId)
	if err != nil {
		return err
	}

	return nil
}

func GetRoleTypeForUser(userId models.UserId) ([]models.RoleType, error) {
	if userId == "" {
		return nil, errors.New("userId cannot be empty")
	}

	userRoles, err := repository.FetchUserRolesByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user roles: %w", err)
	}

	roleIds := []models.RoleId{}
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	roles, err := repository.FetchRolesByRoleIds(roleIds)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch roles by IDs: %w", err)
	}

	roleTypes := []models.RoleType{}
	for _, role := range roles {
		if role.Role.IsValid() {
			roleTypes = append(roleTypes, role.Role)
		}
	}

	return roleTypes, nil
}

func RemoveUserRoleForUser(userId models.UserId, roleId models.RoleId) error {
	if userId == "" {
		return errors.New("userId cannot be empty")
	}

	if roleId == "" {
		return errors.New("roleId cannot be empty")
	}

	err := repository.DeleteUserRole(userId, roleId)
	if err != nil {
		return err
	}

	return nil
}
