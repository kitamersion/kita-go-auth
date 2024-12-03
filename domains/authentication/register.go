package authentication

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/events"
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get email/password
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to read body"))
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to hash password"))
		return
	}

	// create user
	userId := models.UserId(uuid.New().String())
	user := models.User{
		ID:       userId,
		Email:    body.Email,
		Password: string(hash),
		ActivatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	result, err := users.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to create user"))
		return
	}

	events.EventBusGo.Publish(events.RoleAssignedEvent{
		UserId: user.ID,
		RoleId: initializers.GuestRoleId,
	})

	// response
	c.JSON(http.StatusOK, gin.H{
		"userId": result.ID,
	})
}
