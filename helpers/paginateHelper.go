package helpers

import "gorm.io/gorm"

func Paginate(db *gorm.DB, page int, pageSize int) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return db.Offset(offset).Limit(pageSize)
}
