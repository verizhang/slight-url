package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"os/user"
	"reflect"
	"slight-url/config"
	"slight-url/core/exceptions"
	"slight-url/src/models"
	"strings"
)

func Auth(ctx *gin.Context) {
	const bearer = "Bearer"
	var authHeader = ctx.GetHeader("Authorization")
	if authHeader == "" {
		exceptions.Unauthorized(ctx, "Authorization header not found")
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
		exceptions.Unauthorized(ctx, err.Error())
	}

	payload := validatedToken.Claims.(jwt.MapClaims)

	var findUser models.User
	config.DB.Find(&findUser, payload["id"])
	if reflect.DeepEqual(findUser, user.User{}) {
		exceptions.Unauthorized(ctx, "User not found")
	}
	ctx.Set("user", findUser)
}
