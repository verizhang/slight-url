package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"slight-url/core/exceptions"
	"slight-url/src/dtos"
	"slight-url/src/models"
	"slight-url/src/repositories"
	"slight-url/src/utils"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (UserService *UserService) FindAll(ctx *gin.Context) {
	var userQuery dtos.UserQueryDto
	err := ctx.ShouldBindQuery(&userQuery)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	var data = UserService.UserRepository.FindAll(&userQuery)

	ctx.JSON(http.StatusOK, data)
}

func (UserService *UserService) Register(ctx *gin.Context) {
	var user models.User
	var userDto dtos.UserDto
	err := ctx.ShouldBindJSON(&userDto) //validate request body
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}
	user = userDto.ToEntity()
	user.Password = utils.BcryptPassword(&user.Password)
	result := UserService.UserRepository.DB.Create(&user)
	if result.Error != nil {
		exceptions.BadRequest(ctx, result.Error)
		return
	}
	user.Token, err = utils.GetToken(&user)

	ctx.JSON(http.StatusOK, user)
}

func (UserService *UserService) Login(ctx *gin.Context) {
	var loginDto dtos.LoginDto
	err := ctx.ShouldBindJSON(&loginDto)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	var userModel models.User
	UserService.UserRepository.DB.Find(&userModel, models.User{Username: loginDto.Username})
	if reflect.DeepEqual(userModel, models.User{}) {
		exceptions.BadRequest(ctx, "User not found")
		return
	}
	if utils.ComparePassword(&userModel.Password, &loginDto.Password) {
		userModel.Token, err = utils.GetToken(&userModel)
		ctx.JSON(http.StatusOK, userModel)
	} else {
		exceptions.BadRequest(ctx, "Username or password wrong")
		return
	}
}

func (UserService *UserService) ChangePassword(ctx *gin.Context) {
	var changePasswordDto dtos.ChangePasswordDto
	err := ctx.ShouldBindJSON(&changePasswordDto)
	if err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	var user models.User
	if err = UserService.UserRepository.DB.Find(&user, models.User{Username: changePasswordDto.Username}).Error; err != nil {
		exceptions.BadRequest(ctx, err.Error())
		return
	}

	if !utils.ComparePassword(&user.Password, &changePasswordDto.OldPassword) {
		exceptions.BadRequest(ctx, "The old password was not correct")
		return
	}

	user.Password = utils.BcryptPassword(&changePasswordDto.NewPassword)
	UserService.UserRepository.DB.Save(&user)

	ctx.JSON(http.StatusOK, user)
}
