package exceptions

import "github.com/gin-gonic/gin"

func Unauthorized(message interface{}) gin.H {
	return gin.H{
		"statusCode": 401,
		"message":    message,
		"error":      "Unauthorized",
	}
}
