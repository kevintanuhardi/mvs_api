package usecase

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/company/repository"
)

type Service struct {
	company repository.Repository
}

type ServiceManager interface {
	CompanyRegister(ctx context.Context, companyData *entity.Company) error
}

func NewService(company repository.Repository) *Service {
	return &Service{company}
}

func (s *Service) CompanyRegister(ctx context.Context, company *entity.Company)  error {
	err := s.company.CompanyRegister(ctx, company)
	if err != nil {
		return err
	}
	return  nil
}

