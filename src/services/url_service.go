package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slight-url/core/exceptions"
	"slight-url/src/dtos"
	"slight-url/src/models"
	"slight-url/src/repositories"
)

type UrlService struct {
	UrlRepository repositories.UrlRepository
}

func (UrlService *UrlService) FindALl(ctx *gin.Context) {
	var urlQuery dtos.UrlQueryDto
	err := ctx.ShouldBindQuery(&urlQuery)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	var urls = UrlService.UrlRepository.FindALl(urlQuery)

	ctx.JSON(http.StatusOK, urls)
}

func (UrlService *UrlService) Create(ctx *gin.Context) {
	var urlDto dtos.UrlDto
	err := ctx.ShouldBindJSON(&urlDto)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	var url = urlDto.ToEntity()
	userRequest, _ := ctx.Get("user")
	url.UserID = userRequest.(models.User).ID
	UrlService.UrlRepository.DB.Create(&url)

	ctx.JSON(http.StatusOK, url)
}
