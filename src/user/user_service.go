package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"slight-url/src/core"
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

	context.JSON(http.StatusOK, ToDto(userEntity))
}

func (UserService *UserService) Login(context *gin.Context) {
	var loginDto LoginDto
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, core.BadRequestException(err.Error()))
		return
	}

	var userEntity User
	UserService.UserRepository.DB.Find(&userEntity, User{Username: loginDto.Username})
	if userEntity == (User{}) {
		context.JSON(http.StatusBadRequest, core.BadRequestException("User not found"))
		return
	}
	if ComparePassword(userEntity.Password, loginDto.Password) {
		context.JSON(http.StatusOK, ToDto(userEntity))
	} else {
		context.JSON(http.StatusBadRequest, core.BadRequestException("Username or password wrong"))
	}
}

func BcryptPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
