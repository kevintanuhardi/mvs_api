package usecase

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	companyEntity "gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/repository"
	companyRepository "gitlab.warungpintar.co/sales-platform/brook/domain/company/repository"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users repository.Repository
	company companyRepository.Repository
}

type ServiceManager interface {
	UserRegister(ctx context.Context, userData *entity.User) (user *entity.User, err error)
	UserActivation(ctx context.Context, userData *dto.UserActivateRequest) (user *entity.User, company *companyEntity.Company, err error)
}

func NewService(user repository.Repository, company companyRepository.Repository) *Service {
	return &Service{user, company}
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


func (s *Service) UserActivation(ctx context.Context, userData *dto.UserActivateRequest) (user *entity.User, company *companyEntity.Company, err error) {

	company, err = s.company.FindByCompanyCode(ctx, userData.CompanyCode)
	if err != nil {
		return nil, nil, constants.GetCompanyCodeNotFoundError()
	}

	user, err = s.users.FindByEmployeeId(ctx, userData.EmployeeId)
	if err != nil {
		return nil, nil, constants.GetEmployeeIdNotFoundError()
	}

	_, err = s.users.UserActivation(ctx, userData)
	if err != nil {
		return nil, nil, err
	}
	return user, company, nil
}

