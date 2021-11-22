package usecase

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users repository.Repository
}

type ServiceManager interface {
	UserRegister(ctx context.Context, userData *entity.User) error
}

func NewService(user repository.Repository) *Service {
	return &Service{user}
}

func (s *Service) UserRegister(ctx context.Context, userData *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
			panic(err)
	}
	userData.Password = string(hashedPassword)

	err = s.users.UserRegister(ctx, userData)
	if err != nil {
		return err
	}
	return  nil
}

