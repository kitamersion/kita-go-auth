package authentication

import (
	"net/http"

	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/models"
	"github.com/kitamersion/kita-go-auth/repository"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// Check if the user exists in the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, common.CreateResponse("User not found"))
		return
	}

	u := user.(models.User)

	// Delete the refresh token for the user from the database
	err := repository.DeleteRefreshTokenByUserId(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.CreateResponse("Error deleting token for user"))
		return
	}

	// Remove cookies by setting MaxAge to -1
	// TODO: move to common domain for user deletion and logout to clear cookies
	c.SetCookie("Authorization", "", -1, "", "", common.IsProduction, true)
	c.SetCookie("RefreshToken", "", -1, "", "", common.IsProduction, true)

	// Send a success response
	c.JSON(http.StatusOK, common.CreateResponse("Logged out successfully"))
}
