package usecase

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/user/entity"
	"github.com/kevintanuhardi/mvs_api/domain/user/repository"
	"github.com/kevintanuhardi/mvs_api/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users repository.Repository
}

type ServiceManager interface {
	UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error)
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

