package repository

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
)

type Repository interface {
	FindOrder(ctx context.Context) ([]*entity.OrderAggregate, error)
	UserRegister(ctx context.Context, userData *entity.User) error
}
