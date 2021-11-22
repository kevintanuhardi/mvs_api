package mysql

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
)


func (r *repo) UserRegister(ctx context.Context, userData *entity.User) error {
	if err := r.db.Create(userData).Error; err != nil {
		return constants.GetErrDatabaseError()
	}
	return nil
}