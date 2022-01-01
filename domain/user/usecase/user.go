package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kevintanuhardi/mvs_api/domain/user/dto"
	"github.com/kevintanuhardi/mvs_api/domain/user/entity"
	"github.com/kevintanuhardi/mvs_api/domain/user/repository"
	"github.com/kevintanuhardi/mvs_api/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users repository.Repository
}

type ServiceManager interface {
	UserRegister(ctx context.Context, userData *dto.RegisterUserRequest) (user *dto.RegisterUserResponse, err error)
	Login(ctx context.Context, body *dto.LoginRequest) (*dto.LoginResponse, error)
}

func NewService(user repository.Repository) *Service {
	return &Service{user}
}

func (s *Service) UserRegister(ctx context.Context, userData *dto.RegisterUserRequest) (user *dto.RegisterUserResponse, err error) {
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

	userEntity := &entity.User{
		ID:          nil,
		Name:        userData.Name,
		IsActive:      true,
		PhoneNumber: userData.PhoneNumber,
		Email:       userData.Email,
		Password:    userData.Password,
		Role:        userData.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userEntity, err = s.users.UserRegister(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{
		ID:          userEntity.ID,
		Name:        userEntity.Name,
		PhoneNumber: userEntity.PhoneNumber,
		Email:       userEntity.Email,
		Role:        userEntity.Role,
	}, nil
}

func (s *Service) Login(ctx context.Context, body *dto.LoginRequest) (*dto.LoginResponse, error) {
	var d *dto.LoginResponse = new(dto.LoginResponse)

	user, err := s.users.FindByEmail(ctx, body.Email)
	if err != nil {
		panic(err)
	}
	if user == nil {
		return nil, constants.GetUserNotFoundError()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return d, constants.GetWrongPassError()
	}

	d.User = &dto.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Role:        user.Role,
	}

	fmt.Println("d:", d.User)

	return d, nil
}
