package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/events"
	"github.com/kitamersion/kita-go-auth/models"
)

func AddUserRole(c *gin.Context) {
	urlParam := c.Param("id")

	targetUserId := models.UserId(urlParam)
	if targetUserId == "" {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to target user"))
		return
	}

	var body struct {
		RoleId models.RoleId
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest,
			common.CreateResponse("Failed to read body"),
		)
		return
	}

	user, userErr := users.GetUserById(targetUserId)
	if userErr != nil {
		c.JSON(http.StatusNotFound, common.CreateResponse("User not found"))
		return
	}

	// Publish event
	events.EventBusGo.Publish(events.RoleAssignedEvent{
		UserId: user.ID,
		RoleId: body.RoleId,
	})

	c.JSON(http.StatusOK, common.CreateResponse("User role assigned successfully"))
}

func RemoveUserRole(c *gin.Context) {
	urlParam := c.Param("id")

	targetUserId := models.UserId(urlParam)
	if targetUserId == "" {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to target user ID from URL pathname"))
		return
	}
	var body struct {
		RoleId models.RoleId
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest,
			common.CreateResponse("Failed to read body"),
		)
		return
	}

	user, userErr := users.GetUserById(targetUserId)
	if userErr != nil {
		c.JSON(http.StatusNotFound, common.CreateResponse("User not found"))
		return
	}

	// Publish event
	events.EventBusGo.Publish(events.RoleRevokedEvent{
		UserId: user.ID,
		RoleId: body.RoleId,
	})

	c.JSON(http.StatusOK, common.CreateResponse("User role revoked successfully"))
}
