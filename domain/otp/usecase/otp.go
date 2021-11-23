package usecase

import (
	"context"
	"strconv"

	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	userRepo "gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users userRepo.Repository
}

type ServiceManager interface {
	SendOTP(ctx context.Context, input *dto.SendOTPRequest) (dto.SendOTP, error)
	Verify(ctx context.Context, input *dto.VerifyOTPRequest) (dto.SendOTP, error)
	Login(ctx context.Context, input *dto.LoginRequest) (dto.Login, error)
}

func NewService(user repository.Repository) *Service {
	return &Service{user}
}

func (s *Service) SendOTP(ctx context.Context, body *dto.SendOTPRequest) (dto.SendOTP, error) {
	var d dto.SendOTP

	user, err := s.users.FindByPhoneNumber(ctx, body.PhoneNumber)
	if err != nil {
		return d, err
	}

	d.PhoneNumber = user.PhoneNumber
	d.OwnerId = user.EmployeeId
	d.Verified = false
	d.Type = "LOGIN"

	return d, nil
}

func (s *Service) Verify(ctx context.Context, body *dto.VerifyOTPRequest) (dto.SendOTP, error) {
	var d dto.SendOTP

	if body.OTPCode != 1111 {
		return d, constants.GetErrDatabaseError()
	}

	d.PhoneNumber = body.PhoneNumber
	d.OwnerId = body.OwnerId
	d.Verified = true
	d.Type = "LOGIN"

	return d, nil
}

func (s *Service) Login(ctx context.Context, body *dto.LoginRequest) (dto.Login, error) {
	var d dto.Login
	p := strconv.Itoa(body.Pin)

	user, err := s.users.FindByEmployeeId(ctx, body.OwnerId)
	if err != nil {
		return d, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))
	if err != nil {
		return d, err
	}

	d.User = *user

	return d, nil
}
