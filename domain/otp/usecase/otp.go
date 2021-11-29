package usecase

import (
	"context"
	"math/rand"
	"strconv"

	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/entity"
	otpRepo "gitlab.warungpintar.co/sales-platform/brook/domain/otp/repository"
	userRepo "gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/time"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users userRepo.Repository
	otps  otpRepo.Repository
}

type ServiceManager interface {
	SendOTP(ctx context.Context, input *dto.SendOTPRequest) (dto.SendOTP, error)
	Verify(ctx context.Context, input *dto.VerifyOTPRequest) (dto.SendOTP, error)
	Login(ctx context.Context, input *dto.LoginRequest) (dto.Login, error)
}

func NewService(users userRepo.Repository, otps otpRepo.Repository) *Service {
	return &Service{users, otps}
}

func (s *Service) SendOTP(ctx context.Context, body *dto.SendOTPRequest) (dto.SendOTP, error) {
	var d dto.SendOTP

	user, err := s.users.FindByPhoneNumber(ctx, body.PhoneNumber)
	if err != nil {
		return d, constants.GetCustomError("User not found.")
	}

	timestamp := time.GetCurrentTimeAdd15Min()

	v := rand.Intn(9000-1000) + 1000

	otp := entity.Otp{
		OwnerId: user.EmployeeId,
		Otp:     strconv.Itoa(v),
		Type:    body.Type,
		ExpTime: timestamp,
	}

	d = dto.SendOTP{
		PhoneNumber: user.PhoneNumber,
		OwnerId:     user.EmployeeId,
		Verified:    false,
		Type:        body.Type,
	}

	err = s.otps.CreateOrUpdateOtp(ctx, &otp)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (s *Service) Verify(ctx context.Context, body *dto.VerifyOTPRequest) (dto.SendOTP, error) {
	var d dto.SendOTP

	currentTimestamp := time.GetCurrentTime()

	otp, err := s.otps.FindOtp(ctx, body.OwnerId, body.Type)
	if err != nil {
		return d, constants.GetCustomError("OTP Not Found.")
	}

	if otp.Otp != body.OTPCode {
		return d, constants.GetCustomError("OTP Code Invalid.")
	}

	if currentTimestamp.After(otp.ExpTime) {
		return d, constants.GetCustomError("OTP code expired. Please resend OTP code.")
	}

	d.PhoneNumber = body.PhoneNumber
	d.OwnerId = body.OwnerId
	d.Verified = true
	d.Type = body.Type

	return d, nil
}

func (s *Service) Login(ctx context.Context, body *dto.LoginRequest) (dto.Login, error) {
	var d dto.Login

	user, err := s.users.FindByEmployeeId(ctx, body.OwnerId)
	if err != nil {
		return d, constants.GetCustomError("User not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return d, constants.GetCustomError("Wrong Password!")
	}

	d.User = *user

	return d, nil
}
