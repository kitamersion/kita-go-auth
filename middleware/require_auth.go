package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"github.com/kitamersion/kita-go-auth/domains/users"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("Authorization cookie not found"))
		return
	}

	// Decode JWT and validate the signing method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key for validation
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("Invalid or malformed token"))
		return
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("Invalid token claims"))
		return
	}

	// Validate token expiration
	if exp, ok := claims["exp"].(float64); ok {
		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("Token has expired"))
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("Invalid token expiration"))
		return
	}

	// Find user by the token's subject (sub)
	sub, ok := claims["sub"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("Invalid token subject"))
		return
	}

	user, err := users.GetUserById(sub)
	if err != nil || user.ID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.CreateResponse("User not found"))
		return
	}

	// Attach the user to the context
	c.Set("user", user)

	// Continue with the next handler
	c.Next()
}
