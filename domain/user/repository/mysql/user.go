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

func (r *repo) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.First(&user, entity.User{PhoneNumber: phoneNumber}).Error; err != nil {
		return &user, constants.GetErrDatabaseError()
	}
	return &user, nil
}

func (r *repo) FindByEmployeeId(ctx context.Context, employeeId string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.First(&user, entity.User{EmployeeId: employeeId}).Error; err != nil {
		return &user, constants.GetErrDatabaseError()
	}
	return &user, nil
}
