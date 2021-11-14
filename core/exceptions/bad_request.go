package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(ctx *gin.Context, message interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"message":    message,
		"error":      "Bad request",
	})
}
