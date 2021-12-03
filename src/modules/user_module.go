package modules

import (
	"github.com/gin-gonic/gin"
	"slight-url/config"
	"slight-url/src/middleware"
	"slight-url/src/models"
	"slight-url/src/repositories"
	"slight-url/src/services"
)

func UserModule(app *gin.Engine) {
	UserService := services.UserService{
		UserRepository: repositories.UserRepository{DB: config.DB},
	}
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	AuthRoute := app.Group("/auth")
	{
		AuthRoute.POST("/register", UserService.Register)
		AuthRoute.POST("/login", UserService.Login)
		AuthRoute.POST("/change_password", UserService.ChangePassword)
	}

	UserRoute := app.Group("/user")
	UserRoute.Use(middleware.Auth)
	{
		//UserRoute.GET("/", UserService.FindAll)
	}
}
