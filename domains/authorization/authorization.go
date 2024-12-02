package authorization

import (
	"errors"
	"log"

	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
)

func CreatePermissionForUser(permission models.Permission) (models.Permission, error) {
	response, err := repository.CreatePermission(permission)
	if err != nil {
		log.Println("CreatePermissionForUser: error creating permission for user")
		return models.Permission{}, err
	}
	return response, nil
}

func GetPermissionsForUser(userId string) ([]models.Permission, error) {
	if userId == "" {
		return []models.Permission{}, errors.New("userId cannot be empty")
	}

	userPermissions, err := repository.FetchPermissionsByUserId(userId)
	if err != nil {
		return []models.Permission{}, err
	}
	return userPermissions, nil
}

func DeletePermissionsForUser(userId string) error {
	if userId == "" {
		return errors.New("userId cannot be empty")
	}

	err := repository.DeletePermissionsByUserId(userId)
	if err != nil {
		return err
	}

	return nil
}
