package authentication

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// get email/password
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create user
	userId := uuid.New().String()
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"user_id": result.ID,
	})
}
