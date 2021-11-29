package mysql

import (
	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/repository"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &repo{db}
}
