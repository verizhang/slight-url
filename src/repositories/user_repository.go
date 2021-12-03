package repositories

import (
	"gorm.io/gorm"
	"slight-url/core/paginations"
	"slight-url/src/dtos"
	"slight-url/src/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func (UserRepository *UserRepository) FindAll(userQuery *dtos.UserQueryDto) (data paginations.Pagination) {
	var user []models.User

	query := UserRepository.DB
	query = paginations.Paginate(query, &paginations.PaginationOption{
		Page:  userQuery.Page,
		Limit: userQuery.Limit,
	})
	query.Find(&user)

	data = paginations.Create(query, &paginations.PaginationOption{
		Page:  userQuery.Page,
		Limit: userQuery.Limit,
		Model: []models.User{},
		Data:  &user,
	})

	return
}
