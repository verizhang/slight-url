package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		AuthRoute.POST("/register", func(context *gin.Context) {
			UserService.Register(context)
		})

		AuthRoute.POST("/login", func(context *gin.Context) {
			UserService.Login(context)
		})
	}
}
