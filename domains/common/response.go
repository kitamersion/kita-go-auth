package common

import "github.com/gin-gonic/gin"

func CreateResponse(message string) gin.H {
	return gin.H{
		"message": message,
	}
}
