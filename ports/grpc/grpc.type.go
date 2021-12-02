package grpc

import (
	"net"

	"gitlab.warungpintar.co/sales-platform/brook/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/metricserver"
	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type server struct {
	Usecase                domain.DomainService
	server                 *grpc.Server
	GrpcPort               string
	PrometheusPort         string
	transformer            transformer
	networkListen          networkListen
	RegisterBrookServer 		RegisterBrookServer
	RegisterReflection     RegisterReflection
	RegisterPrometheus     RegisterPrometheus
}
type Option struct {
	Cfg             *config.Config
	GormStarter     adapter.GormStarter
	RunMetricServer metricserver.RunMetricServerFunc
	NewGRPCServer   NewGRPCServer
}
type Server interface {
	Serve() error
}
type RegisterPrometheus func(server *grpc.Server)
type networkListen func(network, address string) (net.Listener, error)
type RegisterBrookServer func(s *grpc.Server, srv pb.BrookServer)
type RegisterReflection func(s reflection.GRPCServer)
type transformer func(original *grpc.Server) grpcInterface
type NewGRPCServer func(o *localOption) Server
type grpcInterface interface {
	Serve(lis net.Listener) error
}

type localOption struct {
	db  *gorm.DB
	Cfg *config.Config
}
