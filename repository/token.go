package repository

import (
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
)

func CreateRefreshToken(token models.RefreshToken) (models.RefreshToken, error) {
	err := initializers.DB.Create(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}

func UpdateRefreshTokenByUserId(userId models.UserId, updatedToken models.RefreshToken) (models.RefreshToken, error) {
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

func DeleteRefreshTokenByUserId(userId models.UserId) error {
	err := initializers.DB.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error
	return err
}

func FetchRefreshTokenByUserId(userId models.UserId) (models.RefreshToken, error) {
	var token models.RefreshToken
	err := initializers.DB.Where("user_id = ?", userId).First(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}

func FetchRefreshTokenByToken(tokenString string) (models.RefreshToken, error) {
	var token models.RefreshToken
	err := initializers.DB.Where("token = ?", tokenString).First(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}
	return token, nil
}
