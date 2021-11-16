package db

import "gorm.io/gorm"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		maxPageSize := 100
		minPageSize := 10

		switch {
		case pageSize > maxPageSize:
			pageSize = maxPageSize
		case pageSize <= minPageSize:
			pageSize = minPageSize
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
