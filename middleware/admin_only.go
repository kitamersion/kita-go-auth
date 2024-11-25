package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/models"
)

func AdminOnly(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	u, ok := user.(models.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User data corrupted"})
		return
	}

	// Check if the user has the admin role
	userRolesTypes, err := role.GetRoleTypeForUser(u.ID)
	if err != nil || !common.UserRoleTypesContains(userRolesTypes, models.Admin) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	c.Next()
}
