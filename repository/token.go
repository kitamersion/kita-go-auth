package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

// CreateRefreshToken adds a new refresh token to the database.
func CreateRefreshToken(token models.RefreshToken) (models.RefreshToken, error) {
	err := initializers.DB.Create(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}

// UpdateRefreshTokenByUserId updates a refresh token by the associated user ID.
func UpdateRefreshTokenByUserId(userId string, updatedToken models.RefreshToken) (models.RefreshToken, error) {
	var token models.RefreshToken
	err := initializers.DB.Where("user_id = ?", userId).First(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}

	token.Token = updatedToken.Token
	token.ExpiresAt = updatedToken.ExpiresAt
	err = initializers.DB.Save(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}

// DeleteRefreshTokenByUserId deletes a refresh token by the associated user ID.
func DeleteRefreshTokenByUserId(userId string) error {
	err := initializers.DB.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error
	return err
}

// FetchRefreshTokenByUserId retrieves a refresh token by the associated user ID.
func FetchRefreshTokenByUserId(userId string) (models.RefreshToken, error) {
	var token models.RefreshToken
	err := initializers.DB.Where("user_id = ?", userId).First(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}

// FetchRefreshTokenByToken retrieves a refresh token by token string when a users access token has expired.
func FetchRefreshTokenByToken(tokenString string) (models.RefreshToken, error) {
	var token models.RefreshToken
	err := initializers.DB.Where("token = ?", tokenString).First(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}
