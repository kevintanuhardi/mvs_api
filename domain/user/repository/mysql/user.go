package mysql

import (
	"context"
	"errors"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"

	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	"gorm.io/gorm"
)

func (r *repo) UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error) {
	if err := r.db.Create(userData).Error; err != nil {
		return nil, constants.GetErrDatabaseError()
	}
	return userData, nil
}

func (r *repo) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.First(&user, entity.User{PhoneNumber: phoneNumber}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &user, constants.GetErrDatabaseError()
	}
	return &user, nil
}

func (r *repo) FindByPhoneNumberOrEmail(ctx context.Context, phoneNumber string, email string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.Where(entity.User{PhoneNumber: phoneNumber}).Or(entity.User{Email: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
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

func (r *repo) UserActivation(ctx context.Context, userData *dto.UserActivateRequest) (*entity.User, error) {
	user := entity.User{}

	err := r.db.Where(map[string]interface{}{"employee_id": userData.EmployeeId, "active": 0}).First(&user).Error
	if err != nil {
		return nil, constants.GetEmployeeAlreadyActivatedError()
	}

	err = r.db.Model(&user).Where("employee_id = ?", userData.EmployeeId).Update("active", true).Error
	if err != nil {
		return nil, constants.GetErrDatabaseError()
	}
	return &user, nil
}