package domain

import (
	"context"

	companyEntity "gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/dto"
	userEntity "gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *userEntity.User) error
	UserActivation(ctx context.Context, userData *userEntity.User) error
}

type CompanyDomainInterface interface {
	CompanyRegister(ctx context.Context, companyData *companyEntity.Company) error
}

type OtpDomainInterface interface {
	SendOTP(ctx context.Context, input *dto.SendOTPRequest) (dto.SendOTP, error)
	Verify(ctx context.Context, input *dto.VerifyOTPRequest) (dto.SendOTP, error)
	Login(ctx context.Context, input *dto.LoginRequest) (dto.Login, error)
}
type DomainService struct {
	User UserDomainInterface
	Company CompanyDomainInterface
	Otp  OtpDomainInterface
}

func NewDomain (user UserDomainInterface, company CompanyDomainInterface, otp OtpDomainInterface) DomainService {
	return DomainService{
		User: user,
		Company: company,
		Otp:  otp,
	}
}
