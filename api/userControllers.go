package api

import (
	"net/http"
	"time"

	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ActivatedAt *time.Time `json:"activated_at"` // Use a pointer to allow null values
	Email       string     `json:"email"`
	ID          string     `json:"id"`
}

func User(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	u := user.(models.User)

	// Handle nullable timestamp
	var activatedAt *time.Time
	if u.ActivatedAt.Valid {
		activatedAt = &u.ActivatedAt.Time
	} else {
		activatedAt = nil
	}

	response := UserResponse{
		Email:       u.Email,
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		ActivatedAt: activatedAt, // Set pointer or nil
	}

	c.JSON(http.StatusOK, response)
}

func ActivateUser(c *gin.Context) {
	var body struct {
		UserId string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user, userErr := users.GetUserById(body.UserId)
	if userErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Use the user data (e.g., user ID)
	err := users.ActivateUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while activating user"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User activated successfully"})
}

func DeactivateUser(c *gin.Context) {
	var body struct {
		UserId string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user, userErr := users.GetUserById(body.UserId)
	if userErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Use the user data (e.g., user ID)
	err := users.DeactivateUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deactivating user"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User deactivated successfully"})
}

func DeleteUser(c *gin.Context) {
	var body struct {
		UserId string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user, userErr := users.GetUserById(body.UserId)
	if userErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Use the user data (e.g., user ID)
	err := users.DeleteUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	// TODO: service for this??
	repository.DeleteRefreshTokenByUserId(user.ID)

	// TODO: move to common domain for user deletion and logout to clear cookies
	c.SetCookie("Authorization", "", -1, "", "", common.IsProduction, true)
	c.SetCookie("RefreshToken", "", -1, "", "", common.IsProduction, true)

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}
