package users

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
)

func CreateUser(user models.User) (models.User, error) {
	response, err := repository.CreateUser(user)
	if err != nil {
		log.Println("CreateUser: error creating user")
		return models.User{}, err
	}
	return response, nil
}

func GetUserById(userId string) (models.User, error) {
	if userId == "" {
		return models.User{}, errors.New("user ID cannot be empty")
	}

	user, err := repository.FetchUserById(userId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		return models.User{}, errors.New("user ID cannot be empty")
	}

	user, err := repository.FetchUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func ActivateUser(userId string) error {
	if userId == "" {
		log.Println("ActivateUser: user ID cannot be empty")
		return errors.New("error activating user")
	}

	// Check if the user exists
	user, err := repository.FetchUserById(userId)
	if err != nil {
		log.Printf("ActivateUser: error fetching user with ID %s: %v\n", userId, err)
		return errors.New("error fetching user to activate")
	}

	// Update the "ActivatedAt" field to the current time
	user.ActivatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	// Save the updated user
	_, err = repository.UpdateUserById(user.ID, user)
	if err != nil {
		log.Printf("ActivateUser: error activating user with ID %s: %v\n", userId, err)
		return errors.New("error updating user record")
	}

	log.Printf("ActivateUser: user with ID %s successfully activated\n", userId)
	return nil
}

func DeactivateUser(userId string) error {
	if userId == "" {
		log.Println("DeactivateUser: user ID cannot be empty")
		return errors.New("error deactivating user")
	}

	// Check if user exists
	user, err := repository.FetchUserById(userId)
	if err != nil {
		log.Printf("DeactivateUser: failed to fetch user with ID %s: %v\n", userId, err)
		return errors.New("error fetching user to deactivate")
	}

	// Set deactivation
	user.ActivatedAt = sql.NullTime{
		Time:  time.Time{}, // Zero value for time
		Valid: false,
	}

	// Update the user
	_, err = repository.UpdateUserById(user.ID, user)
	if err != nil {
		log.Printf("DeactivateUser: failed to deactivate user with ID %s: %v\n", userId, err)
		return errors.New("error updating user record")
	}

	return nil
}

func DeleteUser(userId string) error {
	if userId == "" {
		return errors.New("user ID cannot be empty")
	}

	err := repository.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return nil
}
