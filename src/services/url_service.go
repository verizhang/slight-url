package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slight-url/core/exceptions"
	"slight-url/src/dtos"
	"slight-url/src/models"
	"slight-url/src/repositories"
	"slight-url/src/utils"
	"strconv"
	"time"
)

type UrlService struct {
	UrlRepository repositories.UrlRepository
}

func (UrlService *UrlService) Create(ctx *gin.Context) {
	var urlDto dtos.UrlDto
	err := ctx.ShouldBindJSON(&urlDto)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	err = utils.UrlValidator(urlDto.Destination)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	var url = urlDto.ToEntity()
	userRequest, exists := ctx.Get("user")
	if exists == true {
		var userID = userRequest.(models.User).ID
		url.UserID = &userID
	}
	if err := UrlService.UrlRepository.DB.Create(&url).Error; err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	var day = time.Hour * 24
	var duration = time.Duration(3)
	if exists == true {
		duration = time.Duration(7)
	}
	time.AfterFunc(day*duration, func() {
		fmt.Println("delete url" + url.Key)
		UrlService.UrlRepository.DB.Delete(&url)
	})

	ctx.JSON(http.StatusOK, url)
}

func (UrlService *UrlService) FindALl(ctx *gin.Context) {
	userRequest, _ := ctx.Get("user")
	var user = userRequest.(models.User)

	var urlQuery dtos.UrlQueryDto
	err := ctx.ShouldBindQuery(&urlQuery)
	urlQuery.User = user
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	var urls = UrlService.UrlRepository.FindALl(&urlQuery)

	ctx.JSON(http.StatusOK, urls)
}

func (UrlService *UrlService) FindOne(ctx *gin.Context) {
	var url models.Url
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	if err = UrlService.UrlRepository.DB.First(&url, id).Error; err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, url)
}

func (UrlService *UrlService) Remove(ctx *gin.Context) {
	var url models.Url
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	if err = UrlService.UrlRepository.DB.First(&url, id).Error; err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	if err = UrlService.UrlRepository.DB.Delete(&url).Error; err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "url removed"})
}

func (UrlService *UrlService) HandleRedirect(ctx *gin.Context) {
	var url models.Url
	var key = ctx.Param("key")

	if err := UrlService.UrlRepository.DB.First(&url, models.Url{Key: key}).Error; err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, url.Destination)
}
