package domain

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *entity.User) error
}

type OtpDomainInterface interface {
	SendOTP(ctx context.Context, input *dto.SendOTPRequest) (dto.SendOTP, error)
	Verify(ctx context.Context, input *dto.VerifyOTPRequest) (dto.SendOTP, error)
	Login(ctx context.Context, input *dto.LoginRequest) (dto.Login, error)
}
type DomainService struct {
	User UserDomainInterface
	Otp  OtpDomainInterface
}

func NewDomain(user UserDomainInterface, otp OtpDomainInterface) DomainService {
	return DomainService{
		User: user,
		Otp:  otp,
	}
}
