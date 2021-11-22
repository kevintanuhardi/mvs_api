package repository

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
)

type Repository interface {
	UserRegister(ctx context.Context, userData *entity.User) error
	UserActivation(ctx context.Context, userData *entity.User) error
}
