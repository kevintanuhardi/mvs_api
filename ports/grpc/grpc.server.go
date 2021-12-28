package grpc

import (
	"fmt"
	"log"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kevintanuhardi/mvs_api/adapter"
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain"
	userDomain "github.com/kevintanuhardi/mvs_api/domain/user"
	userMysql "github.com/kevintanuhardi/mvs_api/domain/user/repository/mysql"
	"github.com/kevintanuhardi/mvs_api/pkg/metricserver"
	pb "github.com/kevintanuhardi/mvs_api/proto/brook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func transformGRPCServer(original *grpc.Server) grpcInterface {
	return original
}
func (se *server) Serve() error {
	log.Println("Serving GRPC on ", se.GrpcPort)
	lis, err := se.networkListen("tcp", se.GrpcPort)
	if err != nil {
		return err
	}
	return se.transformer(se.server).Serve(lis)
}

func Application(o *Option) error {
	db, err := o.GormStarter.ConnectToGorm(adapter.Config{
		Master:           o.Cfg.DB.MasterDSN,
		Replicas:         []string{o.Cfg.DB.ReplicaDSN},
		MaxIdleConns:     o.Cfg.DB.MaxIdle,
		MaxOpenConns:     o.Cfg.DB.MaxOpen,
		MaxLifetimeConns: o.Cfg.DB.MaxLifeTime,
	}, &gorm.Config{})
	if err != nil {
		return err
	}

	server := o.NewGRPCServer(&localOption{
		db:  db,
		Cfg: o.Cfg,
	})

	go o.RunMetricServer(metricserver.GetDefaultConfig(o.Cfg.Port.GrpcMetric))
	return server.Serve()
}
func GetDefaultOption(cfg *config.Config) *Option {
	return &Option{
		Cfg:             cfg,
		GormStarter:     adapter.NewGormStarter(),
		RunMetricServer: metricserver.RunMetricServer,
		NewGRPCServer:   New,
	}
}

func initService(db *gorm.DB) domain.DomainService {
	return domain.NewDomain(
		userDomain.NewUser(config.Config{},
			userMysql.NewRepository(db), 
		),
	)
}

func New(o *localOption) Server {
	se := &server{
		Usecase: initService(o.db),
		GrpcPort:               fmt.Sprintf("0.0.0.0:%d", o.Cfg.Port.Grpc),
		PrometheusPort:         fmt.Sprintf("0.0.0.0:%d", o.Cfg.Port.GrpcMetric),
		networkListen:          net.Listen,
		RegisterBrookServer: 		pb.RegisterBrookServer,
		RegisterReflection:     reflection.Register,
		RegisterPrometheus:     grpc_prometheus.Register,
		transformer:            transformGRPCServer,
	}
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	se.server = s

	se.RegisterBrookServer(s, se)

	// register health check
	healthpb.RegisterHealthServer(s, health.NewServer())

	se.RegisterReflection(s)
	se.RegisterPrometheus(s)

	return se
}
