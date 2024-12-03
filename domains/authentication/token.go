package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
	"gorm.io/gorm"
)

func RefreshToken(c *gin.Context) {
	// Get refresh token from cookies
	refreshTokenString, err := c.Cookie("RefreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.CreateResponse("Refresh token not found"))
		return
	}

	// Find the refresh token in the database
	refreshTokenRecord, err := repository.FetchRefreshTokenByToken(refreshTokenString)
	if err != nil || refreshTokenRecord.ID == "" {
		c.JSON(http.StatusUnauthorized, common.CreateResponse("Invalid or expired refresh token"))
		return
	}

	// Check if the refresh token has expired
	if refreshTokenRecord.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, common.CreateResponse("Refresh token expired"))
		return
	}

	// Generate a new access token
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": refreshTokenRecord.UserId,
		"exp": time.Now().Add(time.Duration(common.ACCESS_TOKEN_EXPIRY) * time.Second).Unix(), // Access token expires in 1 day
	})

	// Sign and get the complete encoded access token as a string
	newAccessTokenString, err := newAccessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to generate new access token"))
		return
	}

	// Return the new access token (24 hours)
	c.SetCookie("Authorization", newAccessTokenString, common.ACCESS_TOKEN_EXPIRY, "", "", common.IsProduction, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessTokenString,
	})
}

func GenerateRefreshToken(user models.User) (string, error) {
	// Generate a new secure random token
	refreshToken, err := common.GenerateSecureToken(32) // 32 bytes
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Set expiration time
	expiresAt := time.Now().Add(time.Duration(common.REFRESH_TOKEN_EXPIRY) * time.Second)

	// Check if a refresh token already exists for the user
	existingToken, err := repository.FetchRefreshTokenByUserId(user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Unexpected error (e.g., DB connectivity issue)
		return "", fmt.Errorf("error checking for existing token: %w", err)
	}

	if err == nil {
		// Existing token found, update it
		existingToken.Token = refreshToken
		existingToken.UpdatedAt = time.Now()
		existingToken.ExpiresAt = expiresAt

		_, err := repository.UpdateRefreshTokenByUserId(user.ID, existingToken)
		if err != nil {
			return "", fmt.Errorf("failed to update refresh token: %w", err)
		}
	} else {
		// No existing token found, create a new one
		newToken := models.RefreshToken{
			ID:        uuid.New().String(),
			Token:     refreshToken,
			UserId:    user.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ExpiresAt: expiresAt,
		}

		_, err := repository.CreateRefreshToken(newToken)
		if err != nil {
			return "", fmt.Errorf("failed to create new refresh token: %w", err)
		}
	}

	return refreshToken, nil
}
