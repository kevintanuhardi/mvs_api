package usecase

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UserRegister(ctx context.Context, user *entity.User)  error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
			panic(err)
	}
	user.Password = string(hashedPassword)

	err = s.orders.UserRegister(ctx, user)
	if err != nil {
		return err
	}
	return  nil
}