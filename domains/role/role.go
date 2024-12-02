package role

import (
	"errors"

	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
)

func AssignRoleToUser(userId string, roleType models.RoleType) (models.Role, error) {
	if roleType == "" {
		return models.Role{}, errors.New("roleType cannot be empty")
	}

	existingRole, roleErr := GetRolesByRoleType(roleType)

	if roleErr != nil {
		return models.Role{}, roleErr
	}

	userRoleRecord := models.UserRole{
		UserId: userId,
		RoleId: existingRole.ID,
	}

	_, err := repository.CreateUserRole(&userRoleRecord)
	if err != nil {
		return models.Role{}, err
	}
	return existingRole, nil
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

func RevokeRoleForUser(userId string, roleId string) error {
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

// TODO: the below needs to be updated
func GetRolesForUser(userId string) ([]models.Role, error) {
	if userId == "" {
		return []models.Role{}, errors.New("userId cannot be empty")
	}

	userRoles, err := repository.FetchRolesByUserId(userId)
	if err != nil {
		return []models.Role{}, err
	}
	return userRoles, nil
}

func GetRoleTypeForUser(userId string) ([]models.RoleType, error) {
	if userId == "" {
		return []models.RoleType{}, errors.New("userId cannot be empty")
	}

	userRoles, err := repository.FetchRolesByUserId(userId)
	if err != nil {
		return []models.RoleType{}, err
	}
	// Initialize roleTypes as an empty slice
	roleTypes := []models.RoleType{}
	for _, role := range userRoles {
		// TODO: validator role is valid!
		roleTypes = append(roleTypes, role.Role)
	}

	return roleTypes, nil
}

func DeleteRolesByRoleId(roleId models.RoleId) error {
	if roleId == "" {
		return errors.New("roleId cannot be empty")
	}

	err := repository.DeleteRoleByRoleId(roleId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRolesByUserId(userId string) error {
	if userId == "" {
		return errors.New("userId cannot be empty")
	}

	err := repository.DeleteRolesByUserId(userId)
	if err != nil {
		return err
	}

	return nil
}
