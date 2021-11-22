package domain

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *entity.User) error
}

type DomainService struct {
	User UserDomainInterface
}

func NewDomain (user UserDomainInterface) DomainService {
	return DomainService{
		User: user,
	}
}
