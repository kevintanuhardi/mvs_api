package grpc

import (
	"context"

	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (se *server) GetPing(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "pong",
	}, nil
}
