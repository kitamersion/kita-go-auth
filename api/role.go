package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/models"
)

func AddUserRole(c *gin.Context) {
	var body struct {
		UserId string // TODO: handle userId for only admins
		Role   models.RoleType
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	if !body.Role.IsValid() {
		c.JSON(http.StatusNotFound, gin.H{"error": "User role is invalid"})
		return
	}

	user, exists := c.Get("user")
	u := user.(models.User)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userRolesTypes, err := role.GetRoleTypeForUser(u.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetching roles for user"})
		return
	}

	if common.UserRoleTypesContains(userRolesTypes, body.Role) {
		c.JSON(http.StatusOK, gin.H{"message": "User role added sucessfully", "role": body.Role})
		return
	}

	userRole := models.Role{
		ID:     uuid.New().String(),
		UserID: u.ID,
		Role:   body.Role,
	}

	roleRes, err := role.CreateRoleForUser(userRole)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to add role to user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User role added sucessfully", "role": roleRes.Role})
}

func RemoveUserRole(c *gin.Context) {
	var body struct {
		UserId string // TODO: handle user id for only admins
		Role   models.RoleType
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	if !body.Role.IsValid() {
		c.JSON(http.StatusNotFound, gin.H{"error": "User role is invalid"})
		return
	}

	user, exists := c.Get("user")
	u := user.(models.User)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userRoles, err := role.GetRolesForUser(u.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetching roles for user"})
		return
	}

	for _, r := range userRoles {
		if r.Role == body.Role {
			err := role.DeleteRolesByRoleId(r.ID)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"message": "Failed to remove role for user"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "User role removed sucessfully", "role": body.Role})
}
