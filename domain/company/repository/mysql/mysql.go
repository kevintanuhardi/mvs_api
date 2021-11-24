package mysql

import (
	"gitlab.warungpintar.co/sales-platform/brook/domain/company/repository"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB // we can change all *gorm.DB to *sqlx.DB, it's optional.
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &repo{db}
}
