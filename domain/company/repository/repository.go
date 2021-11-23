package repository

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
)

type Repository interface {
	CompanyRegister(ctx context.Context, companyData *entity.Company) error
}
