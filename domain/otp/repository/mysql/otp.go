package mysql

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/entity"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
)

func (r *repo) CreateOrUpdateOtp(ctx context.Context, otpData *entity.Otp) error {
	otp := entity.Otp{}

	err := r.db.Where(map[string]interface{}{"owner_id": otpData.OwnerId, "type": "LOGIN"}).First(&otp).Error
	if err != nil {
		if err := r.db.Create(otpData).Error; err != nil {
			return constants.GetErrDatabaseError()
		}
	}

	err = r.db.Model(&otp).Where(map[string]interface{}{"owner_id": otpData.OwnerId, "type": "LOGIN"}).Update("otp", otpData.Otp).Error
	if err != nil {
		return constants.GetErrDatabaseError()
	}

	return nil
}

func (r *repo) FindOtp(ctx context.Context, employeeId string, otpType string) (*entity.Otp, error) {
	otp := entity.Otp{}

	if err := r.db.Where(map[string]interface{}{"owner_id": employeeId, "type": otpType}).First(&otp).Error; err != nil {
		return &otp, constants.GetErrDatabaseError()
	}
	return &otp, nil
}
