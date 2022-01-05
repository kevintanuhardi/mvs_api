package mysql

import (
	"context"
	"errors"

	"github.com/kevintanuhardi/mvs_api/domain/user/entity"

	"github.com/kevintanuhardi/mvs_api/internal/constants"
	"gorm.io/gorm"
)

func (r *repo) UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error) {
	if err := r.db.Create(userData).Error; err != nil {
		return nil, constants.GetErrDatabaseError()
	}
	return userData, nil
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.Where(entity.User{Email: email}).First(&user).Error; err != nil {
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
