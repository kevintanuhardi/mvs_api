package grpc

import (
	"context"

	validator "github.com/go-playground/validator/v10"
	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/dto"
	"gitlab.warungpintar.co/sales-platform/brook/internal/constants"
	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (se *server) SendOtp(ctx context.Context, request *pb.SendOtpRequest) (*pb.SendOtpResponse, error) {
	v := validator.New()

	parsedRequest := dto.SendOTPRequest{
		PhoneNumber: request.PhoneNumber,
		Type:        request.Type.String(),
	}

	err := v.Struct(parsedRequest)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Tag() == "min" && e.Field() == "PhoneNumber" {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error Send OTP",
						HumanErrorMessage: constants.GetPhoneNumberMinError().Error(),
					},
				}, nil
			} else if e.Tag() == "max" && e.Field() == "PhoneNumber" {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error Send OTP",
						HumanErrorMessage: constants.GetPhoneNumberMaxError().Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "PhoneNumber" {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error Send OTP",
						HumanErrorMessage: constants.GetNotNumericError("nomor hp").Error(),
					},
				}, nil
			} else {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:         true,
						ErrorType:       "400",
						HumanErrorTitle: "Error Send OTP",
						ServerMessage:   e.Error(),
					},
				}, nil
			}
		}
	}

	otp, err := se.Usecase.Otp.SendOTP(ctx, &parsedRequest)
	if err != nil {
		var humanErrorMessage string
		var serverMessage string
		switch t := err.(type) {
		case *constants.CustomError:
			humanErrorMessage = t.Error()
		default:
			serverMessage = t.Error()
		}
		return &pb.SendOtpResponse{
			Error: &pb.ErrorPayload{
				IsError:           true,
				ErrorType:         "500",
				HumanErrorTitle:   "Error Send OTP",
				ServerMessage:     serverMessage,
				HumanErrorMessage: humanErrorMessage,
			},
		}, nil
	}

	parsedOtp := &pb.OtpPayload{
		PhoneNumber: otp.PhoneNumber,
		Type:        pb.OtpType(pb.OtpType_value[otp.Type]),
		OwnerId:     otp.OwnerId,
		Verified:    otp.Verified,
	}

	return &pb.SendOtpResponse{
		Otp: parsedOtp,
	}, nil
}

func (se *server) VerifyOtp(ctx context.Context, request *pb.VerifyOtpRequest) (*pb.VerifyOtpResponse, error) {
	v := validator.New()

	parsedRequest := dto.VerifyOTPRequest{
		PhoneNumber: request.PhoneNumber,
		OwnerId:     request.OwnerId,
		OTPCode:     request.OtpCode,
		Type:        request.Type.String(),
	}

	err := v.Struct(parsedRequest)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Tag() == "min" && e.Field() == "PhoneNumber" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetPhoneNumberMinError().Error(),
					},
				}, nil
			} else if e.Tag() == "max" && e.Field() == "PhoneNumber" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetPhoneNumberMaxError().Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "PhoneNumber" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetNotNumericError("nomor hp").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "PhoneNumber" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("nomor hp").Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "OwnerId" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetNotNumericError("owner id").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "OwnerId" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("owner id").Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "OTPCode" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetNotNumericError("kode otp").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "OTPCode" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("kode otp").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "Type" {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("type").Error(),
					},
				}, nil
			} else {
				return &pb.VerifyOtpResponse{
					Error: &pb.ErrorPayload{
						IsError:         true,
						ErrorType:       "400",
						HumanErrorTitle: "Error",
						ServerMessage:   e.Error(),
					},
				}, nil
			}
		}
	}

	verify, err := se.Usecase.Otp.Verify(ctx, &parsedRequest)
	if err != nil {
		var humanErrorMessage string
		var serverMessage string
		switch t := err.(type) {
		case *constants.CustomError:
			humanErrorMessage = t.Error()
		default:
			serverMessage = t.Error()
		}
		return &pb.VerifyOtpResponse{
			Error: &pb.ErrorPayload{
				IsError:           true,
				ErrorType:         "500",
				HumanErrorTitle:   "Error",
				ServerMessage:     serverMessage,
				HumanErrorMessage: humanErrorMessage,
			},
		}, nil
	}

	parsedOtp := &pb.OtpPayload{
		PhoneNumber: verify.PhoneNumber,
		Type:        pb.OtpType(pb.OtpType_value[verify.Type]),
		OwnerId:     verify.OwnerId,
		Verified:    verify.Verified,
	}

	return &pb.VerifyOtpResponse{
		Otp: parsedOtp,
	}, nil
}

func (se *server) AuthLogin(ctx context.Context, request *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	v := validator.New()

	parsedRequest := dto.LoginRequest{
		PhoneNumber: request.PhoneNumber,
		OwnerId:     request.OwnerId,
		Password:    request.Password,
	}

	err := v.Struct(parsedRequest)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Tag() == "min" && e.Field() == "PhoneNumber" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error Send OTP",
						HumanErrorMessage: constants.GetPhoneNumberMinError().Error(),
					},
				}, nil
			} else if e.Tag() == "max" && e.Field() == "PhoneNumber" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetPhoneNumberMaxError().Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "PhoneNumber" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetNotNumericError("nomor hp").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "PhoneNumber" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("nomor hp").Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "OwnerId" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetNotNumericError("owner id").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "OwnerId" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("owner id").Error(),
					},
				}, nil
			} else if e.Tag() == "numeric" && e.Field() == "Password" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetNotNumericError("password").Error(),
					},
				}, nil
			} else if e.Tag() == "required" && e.Field() == "Password" {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:           true,
						ErrorType:         "400",
						HumanErrorTitle:   "Error",
						HumanErrorMessage: constants.GetFieldRequiredError("password").Error(),
					},
				}, nil
			} else {
				return &pb.AuthLoginResponse{
					Error: &pb.ErrorPayload{
						IsError:         true,
						ErrorType:       "400",
						HumanErrorTitle: "Error Send OTP",
						ServerMessage:   e.Error(),
					},
				}, nil
			}
		}
	}

	login, err := se.Usecase.Otp.Login(ctx, &parsedRequest)
	if err != nil {
		var humanErrorMessage string
		var serverMessage string
		switch t := err.(type) {
		case *constants.CustomError:
			humanErrorMessage = t.Error()
		default:
			serverMessage = t.Error()
		}
		return &pb.AuthLoginResponse{
			Error: &pb.ErrorPayload{
				IsError:           true,
				ErrorType:         "500",
				HumanErrorTitle:   "Error",
				ServerMessage:     serverMessage,
				HumanErrorMessage: humanErrorMessage,
			},
		}, nil
	}

	parsedUser := &pb.UserInfo{
		CompanyId:   int32(login.User.CompanyId),
		RoleId:      int32(login.User.RoleId),
		EmployeeId:  login.User.EmployeeId,
		Active:      login.User.Active,
		PhoneNumber: login.User.PhoneNumber,
		Email:       login.User.Email,
		Name:        login.User.Name,
		// Branch: login.User.Branch,
	}

	return &pb.AuthLoginResponse{
		RefreshToken: login.RefreshToken,
		Token:        login.Token,
		User:         parsedUser,
	}, nil
}
