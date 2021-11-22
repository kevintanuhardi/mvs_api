package usecase

import (
	"context"
	"strconv"

	"gitlab.warungpintar.co/sales-platform/brook/domain/dto"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type SendOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
	Test        string `json:"test"`
}

type VerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	OwnerId     string `json:"owner_id"`
	OTPCode     int    `json:"otp_code"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	OwnerId     string `json:"owner_id"`
	Pin         int    `json:"pin"`
}

func (s *Service) SendOTP(ctx context.Context, body *SendOTPRequest) (dto.SendOTP, error) {
	var d dto.SendOTP

	user, err := s.orders.FindByPhoneNumber(ctx, body.PhoneNumber)
	if err != nil {
		return d, err
	}

	d.PhoneNumber = user.PhoneNumber
	d.OwnerId = user.EmployeeId
	d.Verified = false
	d.Type = "LOGIN"

	return d, nil
}

func (s *Service) Verify(ctx context.Context, body *VerifyOTPRequest) (dto.SendOTP, error) {
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

func (s *Service) Login(ctx context.Context, body *LoginRequest) (dto.Login, error) {
	var d dto.Login
	p := strconv.Itoa(body.Pin)

	user, err := s.orders.FindByEmployeeId(ctx, body.OwnerId)
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
