package repository

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/dto"
)

type Repository interface {
	UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error)
	UserActivation(ctx context.Context, userData *dto.UserActivateRequest) (user *entity.User, err error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error)
	FindByPhoneNumberOrEmail(ctx context.Context, phoneNumber string, email string) (*entity.User, error)
	FindByEmployeeId(ctx context.Context, phoneNumber string) (*entity.User, error)
}
