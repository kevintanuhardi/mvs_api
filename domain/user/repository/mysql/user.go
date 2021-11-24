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

func (r *repo) UserActivation(ctx context.Context, userData *entity.User) error {
	user := entity.User{}

	err := r.db.Where(map[string]interface{}{"employee_id": userData.EmployeeId, "active": 0}).First(&user).Error
	if err != nil {
		return constants.GetErrDatabaseError()
	}

	err = r.db.Model(&user).Where("employee_id = ?", userData.EmployeeId).Update("active", true).Error
	if err != nil {
		return constants.GetErrDatabaseError()
	}
	return nil
}