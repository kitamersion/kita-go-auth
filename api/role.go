package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/models"
)

func AddUserRole(c *gin.Context) {
	targetUserID := c.Param("id")
	if targetUserID == "" {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to target user"))
		return
	}

	var body struct {
		Role models.RoleType
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, 
common.CreateResponse("Failed to read body"),
	)
		return
	}

	if !body.Role.IsValid() {
		c.JSON(http.StatusBadRequest, common.CreateResponse("User role is invalid"))
		return
	}

	user, userErr := users.GetUserById(targetUserID)
	if userErr != nil {
		c.JSON(http.StatusNotFound, common.CreateResponse("User not found"))
		return
	}

	userRolesTypes, err := role.GetRoleTypeForUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.CreateResponse("Failed to fetch roles for user"))
		return
	}

	if common.UserRoleTypesContains(userRolesTypes, body.Role) {
		c.JSON(http.StatusOK, common.CreateResponse("User role already exists"))
		return
	}

	userRole := models.Role{
		ID:     uuid.New().String(),
		UserID: user.ID,
		Role:   body.Role,
	}

	_, err = role.CreateRoleForUser(userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.CreateResponse("Failed to add role to user"))
		return
	}

	c.JSON(http.StatusOK, common.CreateResponse("User role added successfully"))
}

func RemoveUserRole(c *gin.Context) {
	targetUserID := c.Param("id")
	if targetUserID == "" {
		c.JSON(http.StatusBadRequest, 
common.CreateResponse("Failed to target user ID from URL pathname")
		)
		return
	}

	var body struct {
		Role models.RoleType
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, 
common.CreateResponse("Failed to read body")
		)
		return
	}

	if !body.Role.IsValid() {
		c.JSON(http.StatusNotFound, common.CreateResponse("User role is invalid"))
		return
	}

	user, userErr := users.GetUserById(targetUserID)
	if userErr != nil {
		c.JSON(http.StatusNotFound, common.CreateResponse("User not found"))
		return
	}

	userRoles, err := role.GetRolesForUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.CreateResponse("Failed to fetching roles for user"))
		return
	}

	for _, r := range userRoles {
		if r.Role == body.Role {
			err := role.DeleteRolesByRoleId(r.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.CreateResponse("Failed to remove role for user"))
				return
			}
		}
	}
	c.JSON(http.StatusOK, common.CreateResponse("User role removed sucessfully"))
}
