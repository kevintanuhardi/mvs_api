package repository

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/user/entity"
)

type Repository interface {
	UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByPhoneNumberOrEmail(ctx context.Context, phoneNumber string, email string) (*entity.User, error)
}
