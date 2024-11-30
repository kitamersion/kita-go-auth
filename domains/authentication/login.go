package authentication

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/users"
	"github.com/kitamersion/kita-go-auth/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// get email/password
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to read body"))
		return
	}

	// look up user by Email
	user, userErr := users.GetUserByEmail(body.Email)

	// Return early if user doesn't exist or error occurs
	if userErr != nil || user.ID == "" {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Invalid email or password"))
		return
	}

	// compare password hash against db user
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Invalid email or password"))
		return
	}

	// update users last login date
	// TODO: create a users service to update user
	user.LastLoginAt = time.Now()
	repository.UpdateUserById(user.ID, user)

	// generate access token (JWT)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Duration(common.ACCESS_TOKEN_EXPIRY) * time.Second).Unix(), // Access token expires in 1 day
	})

	// Sign and get the complete encoded access token as a string
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to generate access token"))
		return
	}

	// generate refresh token (longer expiration, e.g., 30 days)
	refreshToken, err := GenerateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.CreateResponse("Failed to generate refresh token"))
		return
	}

	// set cookies for both tokens
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", accessTokenString, common.ACCESS_TOKEN_EXPIRY, "", "", common.IsProduction, true) // 1-day expiry for access token
	c.SetCookie("RefreshToken", refreshToken, common.REFRESH_TOKEN_EXPIRY, "", "", common.IsProduction, true)      // 30-day expiry for refresh token

	// return the tokens in the response
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshToken,
	})
}
