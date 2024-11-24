package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(c *gin.Context) {
	limiter := rate.NewLimiter(1, 4) // allows 4 requests per second

	if limiter.Allow() {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceed"})
	}
}
