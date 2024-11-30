package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/domains/common"
	"golang.org/x/time/rate"
)

func RateLimiter(c *gin.Context) {
	limiter := rate.NewLimiter(1, 3) // allows 3 requests per second

	if limiter.Allow() {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, common.CreateResponse("Rate limit exceed"))
	}
}
