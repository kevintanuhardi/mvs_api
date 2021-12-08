package grpc

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (se *server) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	parsedRequest := entity.User{
		EmployeeId: request.EmployeeId,
		CompanyId: int(request.CompanyId),
		RoleId: int(request.RoleId),
		Active: request.Active,
		PhoneNumber: request.PhoneNumber,
		Email: request.Email,
		Password: request.Password,
	}
	user, err := se.Usecase.User.UserRegister(ctx, &parsedRequest)
	if err != nil {
		return nil, err
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
	parsedRequest := entity.User{
		EmployeeId: request.EmployeeId,
		CompanyId: int(request.CompanyId),
		RoleId: int(request.RoleId),
		Active: request.Active,
		PhoneNumber: request.PhoneNumber,
		Email: request.Email,
		Password: request.Password,
	}
	err := se.Usecase.User.UserActivation(ctx, &parsedRequest)
	if err != nil {
		return nil, err
	}
	successResponse := pb.ActivateUserResponse{
		Message: "success",
	}
	return &successResponse, nil
}
