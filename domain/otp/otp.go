package user

import (
	"gitlab.warungpintar.co/sales-platform/brook/config"
	otp "gitlab.warungpintar.co/sales-platform/brook/domain/otp/repository"
	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/usecase"
	user "gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
)

func NewOtp(cfg config.Config, users user.Repository, otps otp.Repository) usecase.ServiceManager {
	return usecase.NewService(users, otps)
}
