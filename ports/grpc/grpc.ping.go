package grpc

import (
	"context"

	pb "github.com/kevintanuhardi/mvs_api/proto/brook"
)

func (se *server) GetPing(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "pong",
	}, nil
}
