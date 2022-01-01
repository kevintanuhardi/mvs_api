package domain

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/user/dto"
)

type UserDomainInterface interface {
	UserRegister(ctx context.Context, userData *dto.RegisterUserRequest) (user *dto.RegisterUserResponse, err error)
	Login(ctx context.Context, body *dto.LoginRequest) (*dto.LoginResponse, error)
}

type DomainService struct {
	User UserDomainInterface
}

func NewDomain (user UserDomainInterface) DomainService {
	return DomainService{
		User: user,
	}
}
