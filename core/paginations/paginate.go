package paginations

import (
	"gorm.io/gorm"
	"math"
)

func GetTotalItems(db *gorm.DB, model interface{}) int64 {
	var totalItems int64
	db.Model(model).Count(&totalItems)
	return totalItems
}

func Paginate(query *gorm.DB, options *PaginationOption) (tx *gorm.DB) {
	var page = options.Page
	if page == 0 {
		page = 1
	}

	var limit = options.Limit
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	offset := (page - 1) * limit
	tx = query.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	})

	return
}

func Create(query *gorm.DB, options *PaginationOption) (pagination Pagination) {
	totalRows := GetTotalItems(query, options.Model)
	pagination.Meta.Page = options.Page
	pagination.Meta.Limit = options.Limit
	pagination.Meta.ItemCount = totalRows
	pagination.Meta.PageCount = int(math.Ceil(float64(totalRows) / float64(options.Limit)))
	pagination.Data = options.Data
	return
}
