package domain

import (
	"context"

	userEntity "gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	companyEntity "gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *userEntity.User) error
	UserActivation(ctx context.Context, userData *userEntity.User) error
}

type CompanyDomainInterface interface {
	CompanyRegister(ctx context.Context, companyData *companyEntity.Company) error
}

type DomainService struct {
	User UserDomainInterface
	Company CompanyDomainInterface
}

func NewDomain (user UserDomainInterface, company CompanyDomainInterface) DomainService {
	return DomainService{
		User: user,
		Company: company,
	}
}
