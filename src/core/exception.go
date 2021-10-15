package core

import "github.com/gin-gonic/gin"

func BadRequestException(message interface{}) gin.H {
	return gin.H{
		"statusCode": 400,
		"message":    message,
		"error":      "Bad request",
	}
}
