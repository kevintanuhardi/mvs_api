package user

import (
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/usecase"
)

func NewUser (cfg config.Config, users repository.Repository) usecase.ServiceManager {
	return usecase.NewService(users)
}