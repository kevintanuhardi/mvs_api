package grpc

import (
	"context"

	validator "github.com/go-playground/validator/v10"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/dto"
	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (se *server) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

	v := validator.New()
	
	parsedRequest := entity.User{
		EmployeeId: request.EmployeeId,
		CompanyId: int(request.CompanyId),
		RoleId: int(request.RoleId),
		Active: request.Active,
		PhoneNumber: request.PhoneNumber,
		Email: request.Email,
		Password: request.Password,
		Name: request.Name,
	}

	err := v.Struct(parsedRequest)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Tag() == "email" {
				return &pb.RegisterUserResponse{
					Error: &pb.ErrorPayload{
						IsError: true,
						ErrorType: "400",
						HumanErrorTitle: "Error Registering User",
						HumanErrorMessage: "email is not valid",
					},
				}, nil
			} else {
				return &pb.RegisterUserResponse{
					Error: &pb.ErrorPayload{
						IsError: true,
						ErrorType: "400",
						HumanErrorTitle: "Error Registering User",
						ServerMessage: e.Error(),
					},
				}, nil
			}
		}
	}
	
	user, err := se.Usecase.User.UserRegister(ctx, &parsedRequest)
	if err != nil {
		return &pb.RegisterUserResponse{
			Error: &pb.ErrorPayload{
				IsError: true,
				ErrorType: "500",
				HumanErrorTitle: "Error Registering User",
				ServerMessage: err.Error(),
			},
		}, nil
	}
	parsedUser := &pb.UserInfo{
		CompanyId: int32(user.CompanyId),
		RoleId: int32(user.RoleId),
		EmployeeId: user.EmployeeId,
		Active: user.Active,
		PhoneNumber: user.PhoneNumber,
		Email: user.Email,
		Name: user.Name,
		// Branch: user.Branch,
	}
	successResponse := pb.RegisterUserResponse{
		User: parsedUser,
	}
	return &successResponse, nil
}

func (se *server) ActivateUser(ctx context.Context, request *pb.ActivateUserRequest) (*pb.ActivateUserResponse, error) {

	v := validator.New()

	parsedRequest := dto.UserActivateRequest{
		EmployeeId: request.EmployeeId,
		CompanyCode: request.CompanyCode,
	}

	err := v.Struct(parsedRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return &pb.ActivateUserResponse{
				Error: &pb.ErrorPayload{
					IsError: true,
					ErrorType: "400",
					HumanErrorTitle: "Error Aktivasi Sales",
					ServerMessage: e.Error(),
				},
			}, nil
		}
	}

	user, company, err := se.Usecase.User.UserActivation(ctx, &parsedRequest)
	if err != nil {
		return &pb.ActivateUserResponse{
			Error: &pb.ErrorPayload{
				IsError: true,
				ErrorType: "400",
				HumanErrorTitle: "Error Aktivasi Sales",
				ServerMessage: err.Error(),
			},
		}, nil
	}

	parsedCompany := &pb.Company{
		Code: company.Code,
		Name: company.Name,
		Address: company.Address,
		Country: company.Country,
		Province: company.Province,
		City: company.City,
		District: company.District,
		Village: company.Village,
		PostalCode: company.PostalCode,
		PhoneNumber: company.PhoneNumber,
		FaxNumber: company.FaxNumber,
		Npwp: company.NPWP,
		Sppkp: company.SPPKP,
	}

	parsedUser := &pb.UserInfo{
		CompanyId: int32(user.CompanyId),
		RoleId: int32(user.RoleId),
		EmployeeId: user.EmployeeId,
		Active: user.Active,
		PhoneNumber: user.PhoneNumber,
		Email: user.Email,
		Name: user.Name,
		Company: parsedCompany,
	}

	successResponse := pb.ActivateUserResponse{
		User: parsedUser,
	}
	return &successResponse, nil
}
