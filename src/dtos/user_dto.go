package dtos

import (
	"slight-url/core/paginations"
	"slight-url/src/models"
)

type UserDto struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (UserDto *UserDto) ToEntity() models.User {
	return models.User{
		Name:     UserDto.Name,
		Username: UserDto.Username,
		Password: UserDto.Password,
	}
}

type LoginDto struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserQueryDto struct {
	paginations.PaginationQuery
}
