package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/models"
)

/*
* View: Self OR Admin
* Edit: Self OR Admin
*
* */
func CanViewEditSelf(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("User not authenticated"))
		return
	}

	u, ok := user.(models.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateResponse("User data corrupted"))
		return
	}

	// Extract the target ID from the route parameter
	targetUserID := c.Param("id")
	if targetUserID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateResponse("Target user ID is required"))
		return
	}

	userRolesTypes, err := role.GetRoleTypeForUser(u.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateResponse("Failed to fetch roles for user"))
		return
	}

	if common.UserRoleTypesContains(userRolesTypes, models.Admin) || u.ID == targetUserID {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusForbidden, common.CreateResponse("Permission denied"))
}
