package usecase

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users repository.Repository
}

type ServiceManager interface {
	UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error)
	UserActivation(ctx context.Context, userData *entity.User) error
}

func NewService(user repository.Repository) *Service {
	return &Service{user}
}

func (s *Service) UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error) {
	// check if phone number or email already registered
	existingUser, err := s.users.FindByPhoneNumberOrEmail(ctx, userData.PhoneNumber, userData.Email)
	if err != nil {
		panic(err)
	}
	if existingUser != nil {
		return nil, constants.GetDuplicateUserError()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	userData.Password = string(hashedPassword)

	user, err = s.users.UserRegister(ctx, userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (s *Service) UserActivation(ctx context.Context, userData *entity.User) error {
	err := s.users.UserActivation(ctx, userData)
	if err != nil {
		return err
	}
	return  nil
}

