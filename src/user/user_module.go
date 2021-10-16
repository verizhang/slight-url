package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"slight-url/src/core/middleware"
)

func UserModule(app *gin.Engine, DB *gorm.DB) {
	UserService := UserService{
		UserRepository: UserRepository{DB: DB},
	}
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	AuthRoute := app.Group("/auth")
	{
		AuthRoute.POST("/register", UserService.Register)
		AuthRoute.POST("/login", UserService.Login)
	}

	PingRoute := app.Group("/user")
	PingRoute.Use(middleware.Auth)
	{
		PingRoute.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "ping success",
			})
		})
	}
}
