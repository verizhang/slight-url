package modules

import (
	"github.com/gin-gonic/gin"
	"slight-url/config"
	"slight-url/src/middleware"
	"slight-url/src/models"
	"slight-url/src/repositories"
	"slight-url/src/services"
)

func UrlModule(app *gin.Engine) {
	var UrlService = services.UrlService{
		UrlRepository: repositories.UrlRepository{
			DB: config.DB,
		},
	}
	err := config.DB.AutoMigrate(&models.Url{})
	if err != nil {
		panic(err)
	}

	var PublicUrlRoute = app.Group("/url")
	PublicUrlRoute.Use(middleware.OptionalAuth)
	{
		PublicUrlRoute.POST("/", UrlService.Create)
	}
	app.GET("/:key", UrlService.HandleRedirect)

	var UrlRoute = app.Group("/url/user")
	UrlRoute.Use(middleware.Auth)
	{
		UrlRoute.GET("/", UrlService.FindALl)
		UrlRoute.GET("/:id", UrlService.FindOne)
		UrlRoute.DELETE("/:id", UrlService.Remove)
	}

}
