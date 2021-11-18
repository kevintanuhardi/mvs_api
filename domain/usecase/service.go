package usecase

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/repository"
)

type Service struct {
	orders repository.Repository
}
type ServiceManager interface {
	OrderList(ctx context.Context) ([]*dto.OrderDTO, error)
	UserRegister(ctx context.Context, userData *entity.User) error
}

func NewService(orders repository.Repository) *Service {
	return &Service{orders}
}
