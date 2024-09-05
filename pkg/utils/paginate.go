package utils

import (
	"tel/product/internal/model"

	"gorm.io/gorm"
)

func Paginate(value any, pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var total int64

	db.Model(value).Count(&total)

	pagination.TotalData = total
	pagination.TotalPages = int(total) / pagination.GetPageSize()
	div := int(total) % pagination.PageSize
	if div > 0 {
		pagination.TotalPages++
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPageSize()).Order(pagination.GetSort())
	}
}
