package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"slight-url/src/core/exceptions"
	"strings"
)

func Auth(c *gin.Context) {
	const bearer = "Bearer"
	var authHeader = c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.Unauthorized("Authorization header not found"))
		return
	}

	var tokenString = strings.Trim(authHeader[len(bearer):], " ")

	validatedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, isValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("error", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.Unauthorized(err.Error()))
		return
	}

	if !validatedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.Unauthorized(err.Error()))
	}
}
