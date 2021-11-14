package dtos

import (
	"slight-url/core/paginations"
	"slight-url/src/models"
)

type UrlDto struct {
	Key         string `form:"key" binding:"required"`
	Destination string `form:"destination" binding:"required"`
}

func (UrlDto *UrlDto) ToEntity() models.Url {
	return models.Url{
		Key:         UrlDto.Key,
		Destination: UrlDto.Destination,
	}
}

type UrlQueryDto struct {
	paginations.PaginationQuery
}
