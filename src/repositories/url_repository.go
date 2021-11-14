package repositories

import (
	"gorm.io/gorm"
	"slight-url/core/paginations"
	"slight-url/src/dtos"
	"slight-url/src/models"
)

type UrlRepository struct {
	DB *gorm.DB
}

func (UrlRepository *UrlRepository) FindALl(urlQuery dtos.UrlQueryDto) (data paginations.Pagination) {
	var urls []models.Url

	query := UrlRepository.DB.Preload("User")
	query = paginations.Paginate(query, paginations.PaginationOption{
		Page:  urlQuery.Page,
		Limit: urlQuery.Limit,
	})

	query.Find(&urls)
	data = paginations.Create(query, paginations.PaginationOption{
		Page:  urlQuery.Page,
		Limit: urlQuery.Limit,
		Model: []models.Url{},
		Data:  &urls,
	})
	return
}
