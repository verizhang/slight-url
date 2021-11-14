package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Unauthorized(ctx *gin.Context, message interface{}) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"statusCode": 401,
		"message":    message,
		"error":      "Unauthorized",
	})
}
