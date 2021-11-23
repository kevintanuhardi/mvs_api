package company

import (
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain/company/repository"
	"gitlab.warungpintar.co/sales-platform/brook/domain/company/usecase"
)

func NewCompany (cfg config.Config, company repository.Repository) usecase.ServiceManager {
	return usecase.NewService(company)
}