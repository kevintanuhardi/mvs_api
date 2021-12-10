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
		Type: request.Type.String(),
	}

	err := v.Struct(parsedRequest)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if (e.Tag() == "min" && e.Field() == "PhoneNumber") {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError: true,
						ErrorType: "400",
						HumanErrorTitle: "Error Send OTP",
						HumanErrorMessage: constants.GetPhoneNumberMinError().Error(),
					},
				}, nil
			} else if (e.Tag() == "max" && e.Field() == "PhoneNumber") {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError: true,
						ErrorType: "400",
						HumanErrorTitle: "Error Send OTP",
						HumanErrorMessage: constants.GetPhoneNumberMaxError().Error(),
					},
				}, nil
			} else if (e.Tag() == "numeric" && e.Field() == "PhoneNumber") {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError: true,
						ErrorType: "400",
						HumanErrorTitle: "Error Send OTP",
						HumanErrorMessage: constants.GetNotNumericError("nomor hp").Error(),
					},
				}, nil
			} else {
				return &pb.SendOtpResponse{
					Error: &pb.ErrorPayload{
						IsError: true,
						ErrorType: "400",
						HumanErrorTitle: "Error Send OTP",
						ServerMessage: e.Error(),
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
				IsError: true,
				ErrorType: "500",
				HumanErrorTitle: "Error Send OTP",
				ServerMessage: serverMessage,
				HumanErrorMessage: humanErrorMessage,
			},
		}, nil
	}

	parsedOtp := &pb.OtpPayload{
		PhoneNumber: otp.PhoneNumber,
		Type: pb.OtpType(pb.OtpType_value[otp.Type]),
		OwnerId: otp.OwnerId,
		Verified: otp.Verified,
	}

	return &pb.SendOtpResponse{
		Otp: parsedOtp,
	}, nil
}
