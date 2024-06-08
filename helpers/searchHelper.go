package helpers

import (
	"gorm.io/gorm"
	"strings"
)

func Search(db *gorm.DB, search string, fields ...string) *gorm.DB {
	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		for i, field := range fields {
			if i == 0 {
				db = db.Where("LOWER("+field+") LIKE ?", searchPattern)
			} else {
				db = db.Or("LOWER("+field+") LIKE ?", searchPattern)
			}
		}
	}
	return db
}
