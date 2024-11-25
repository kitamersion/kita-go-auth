package common

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/kitamersion/kita-go-auth/models"
)

// GenerateSecureToken creates a cryptographically secure random token
func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate secure token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func UserRoleContains(slice []models.Role, value models.RoleType) bool {
	for _, v := range slice {
		if v.Role == value {
			return true
		}
	}
	return false
}

func UserRoleTypesContains(slice []models.RoleType, value models.RoleType) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
