package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
)

type UserResponse struct {
	LastLoginAt time.Time         `json:"last_login_at"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	ActivatedAt *time.Time        `json:"activated_at"` // Use a pointer to allow null values
	Email       string            `json:"email"`
	ID          string            `json:"id"`
	Roles       []models.RoleType `json:"roles"`
}

func WhoAmI(c *gin.Context) {
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

	// Fetch roles for the user
	// TODO: consider moving this to RequireAuth middleware?
	roleTypes, err := role.GetRoleTypeForUser(u.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching user roles",
		})
		return
	}

	response := UserResponse{
		Email:       u.Email,
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		LastLoginAt: u.LastLoginAt,
		ActivatedAt: activatedAt, // Set pointer or nil
		Roles:       roleTypes,   // Will be an empty array if no roles
	}

	c.JSON(http.StatusOK, response)
}

func User(c *gin.Context) {
	targetUserID := c.Param("id")
	if c.Bind(&targetUserID) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to target user ID from url pathname",
		})
		return
	}

	user, err := users.GetUserById(targetUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching user",
		})
		return
	}

	// Handle nullable timestamp
	var activatedAt *time.Time
	if user.ActivatedAt.Valid {
		activatedAt = &user.ActivatedAt.Time
	} else {
		activatedAt = nil
	}

	// Fetch roles for the user
	// TODO: consider moving this to RequireAuth middleware?
	roleTypes, err := role.GetRoleTypeForUser(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error fetching user roles",
		})
		return
	}

	response := UserResponse{
		Email:       user.Email,
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		LastLoginAt: user.LastLoginAt,
		ActivatedAt: activatedAt, // Set pointer or nil
		Roles:       roleTypes,   // Will be an empty array if no roles
	}

	c.JSON(http.StatusOK, response)
}

func ActivateUser(c *gin.Context) {
	targetUserID := c.Param("id")
	if c.Bind(&targetUserID) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to target user ID from url pathname",
		})
		return
	}

	user, userErr := users.GetUserById(targetUserID)
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
	targetUserID := c.Param("id")
	if c.Bind(&targetUserID) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to target user ID from url pathname",
		})
		return
	}

	user, userErr := users.GetUserById(targetUserID)
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

// TODO: transactional scope
func DeleteUser(c *gin.Context) {
	targetUserID := c.Param("id")
	if c.Bind(&targetUserID) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to target user ID from url pathname",
		})
		return
	}

	var user models.User
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	authUser := u.(models.User)

	if authUser.ID != targetUserID {
		userRecord, userErr := users.GetUserById(targetUserID)
		if userErr != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		user = userRecord
	} else {
		user = authUser
	}
	// Use the user data (e.g., user ID)
	err := users.DeleteUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	// TODO: service for this??
	repository.DeleteRefreshTokenByUserId(user.ID)

	role.DeleteRolesByUserId(user.ID)

	if authUser.ID == targetUserID {
		// TODO: move to common domain for user deletion and logout to clear cookies
		c.SetCookie("Authorization", "", -1, "", "", common.IsProduction, true)
		c.SetCookie("RefreshToken", "", -1, "", "", common.IsProduction, true)
	}
	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}
