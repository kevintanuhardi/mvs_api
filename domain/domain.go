package domain

import (
	"context"

	userEntity "github.com/kevintanuhardi/mvs_api/domain/user/entity"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *userEntity.User) (user *userEntity.User, err error)
}

type DomainService struct {
	User UserDomainInterface
}

func NewDomain (user UserDomainInterface) DomainService {
	return DomainService{
		User: user,
	}
}
