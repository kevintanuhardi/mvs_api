package grpc

import (
	"context"

	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (se *server) RegisterCompany(context.Context, *pb.RegisterCompanyRequest) (*pb.RegisterCompanyResponse, error) {
	return &pb.RegisterCompanyResponse{
		Message: "pong",
	}, nil
}
