package dto

import "gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"

type SendOTP struct {
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
	OwnerId     string `json:"owner_id"`
	Verified    bool   `json:"verified"`
}

type Login struct {
	RefreshToken string      `json:"refresh_token"`
	Token        string      `json:"token"`
	User         entity.User `json:"user"`
}

type SendOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
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
