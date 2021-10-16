package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slight-url/src/core/exceptions"
)

type UserService struct {
	UserRepository UserRepository
}

func (UserService *UserService) Register(context *gin.Context) {
	var userDto UserDto
	err := context.ShouldBindJSON(&userDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userEntity = ToEntity(userDto)
	userEntity.Password = BcryptPassword(userEntity.Password)
	result := UserService.UserRepository.DB.Create(&userEntity)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	userDto = ToDto(userEntity)
	userDto.Token, err = GetToken(userEntity)
	context.JSON(http.StatusOK, userDto)
}

func (UserService *UserService) Login(context *gin.Context) {
	var loginDto LoginDto
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, exceptions.BadRequest(err.Error()))
		return
	}

	var userEntity User
	UserService.UserRepository.DB.Find(&userEntity, User{Username: loginDto.Username})
	if userEntity == (User{}) {
		context.JSON(http.StatusBadRequest, exceptions.BadRequest("User not found"))
		return
	}
	if ComparePassword(userEntity.Password, loginDto.Password) {
		var userDto UserDto = ToDto(userEntity)
		userDto.Token, err = GetToken(userEntity)
		context.JSON(http.StatusOK, userDto)
	} else {
		context.JSON(http.StatusBadRequest, exceptions.BadRequest("Username or password wrong"))
		return
	}
}
