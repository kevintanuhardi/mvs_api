package mysql

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
)


func (r *repo) CompanyRegister(ctx context.Context, companyData *entity.Company) error {
	if err := r.db.Create(companyData).Error; err != nil {
		return constants.GetErrDatabaseError()
	}
	return nil
}