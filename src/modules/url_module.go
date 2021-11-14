package modules

import (
	"github.com/gin-gonic/gin"
	"slight-url/config"
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

	var UrlRoute = app.Group("/url")
	{
		UrlRoute.GET("/", UrlService.FindALl)
		UrlRoute.POST("/", UrlService.Create)
	}
}
