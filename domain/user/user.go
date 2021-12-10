package user

import (
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	companyRepository "gitlab.warungpintar.co/sales-platform/brook/domain/company/repository"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/usecase"
)

func NewUser (cfg config.Config, users repository.Repository, company companyRepository.Repository) usecase.ServiceManager {
	return usecase.NewService(users, company)
}