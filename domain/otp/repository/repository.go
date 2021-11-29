package repository

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/entity"
)

type Repository interface {
	CreateOrUpdateOtp(ctx context.Context, userData *entity.Otp) error
	FindOtp(ctx context.Context, employeeId string, otpType string) (*entity.Otp, error)
}
