package role

import (
	"errors"
	"log"

	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
)

func CreateRoleForUser(role models.Role) (models.Role, error) {
	response, err := repository.CreateRole(role)
	if err != nil {
		log.Println("CreateRoleForUser: error creating role for user")
		return models.Role{}, err
	}
	return response, nil
}

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

func DeleteRolesByRoleId(roleId string) error {
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
