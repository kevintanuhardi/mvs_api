package user

import (
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/usecase"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
)

func NewOtp(cfg config.Config, users repository.Repository) usecase.ServiceManager {
	return usecase.NewService(users)
}
