package grpc

import (
	"errors"
	"net"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.warungpintar.co/sales-platform/brook/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/config"
	usecasemock "gitlab.warungpintar.co/sales-platform/brook/domain/usecase/mocks"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/metricserver"
	"gitlab.warungpintar.co/sales-platform/brook/ports/grpc/mocks"
	"gitlab.warungpintar.co/sales-platform/brook/proto/brook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type testObject struct {
	suite.Suite
	module         *server
	MockService    *usecasemock.MockServiceManager
	mockGRPCServer *mocks.MockgrpcInterface
}

func (obj *testObject) SetupTest() {
	gomockController := gomock.NewController(obj.T())
	obj.MockService = usecasemock.NewMockServiceManager(gomockController)
	obj.mockGRPCServer = mocks.NewMockgrpcInterface(gomockController)
	obj.module = &server{
		Usecase:  obj.MockService,
		server:   grpc.NewServer(),
		GrpcPort: "0.0.0.0:5077",
		networkListen: func(network, address string) (net.Listener, error) {
			return nil, nil
		},
		RegisterbrookServer: func(s *grpc.Server, srv brook.brookServer) {
		},
		RegisterReflection: func(s reflection.GRPCServer) {
		},
		transformer: func(original *grpc.Server) grpcInterface {
			return obj.mockGRPCServer
		},
	}
}

func Test_grpcServer(t *testing.T) {
	suite.Run(t, new(testObject))
}

func (obj *testObject) TestServe() {
	obj.Run("Success", func() {
		obj.mockGRPCServer.EXPECT().Serve(gomock.Any()).Return(nil)
		err := obj.module.Serve()
		obj.NoError(err)
	})
	obj.Run("ErrorOnServe", func() {
		obj.mockGRPCServer.EXPECT().Serve(gomock.Any()).Return(errors.New("something bad happen"))
		err := obj.module.Serve()
		obj.Error(err)
	})
	obj.Run("ErrorOnNetListen", func() {
		obj.module.networkListen = func(network, address string) (net.Listener, error) {
			return nil, errors.New("something bad happen")
		}
		err := obj.module.Serve()
		obj.Error(err)
	})
}

func TestNew(t *testing.T) {
	db, dbmock, err := sqlmock.New()
	require.NoError(t, err)
	dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow(1.0))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	resp := New(&localOption{
		db:  gormDB,
		Cfg: &config.Config{},
	})
	require.NotNil(t, resp)
}

func TestTransformGRPCServer(t *testing.T) {
	transformGRPCServer(&grpc.Server{})
}

func TestApplicationSuccess(t *testing.T) {
	cfg := &config.Config{}
	gormStarter := &adapter.MockGormStarter{}
	gormStarter.On("ConnectToGorm", adapter.Config{
		Master:           cfg.DB.MasterDSN,
		Replicas:         []string{cfg.DB.ReplicaDSN},
		MaxIdleConns:     cfg.DB.MaxIdle,
		MaxOpenConns:     cfg.DB.MaxOpen,
		MaxLifetimeConns: cfg.DB.MaxLifeTime,
	}, &gorm.Config{}).Return(&gorm.DB{}, nil)
	term := make(chan os.Signal, 1)
	defer close(term)
	ctrl := gomock.NewController(t)
	mockGRPCServer := mocks.NewMockServer(ctrl)
	mockGRPCServer.EXPECT().Serve().Return(nil)
	err := Application(&Option{
		GormStarter: gormStarter,
		Cfg:         cfg,
		NewGRPCServer: func(o *localOption) Server {
			return mockGRPCServer
		},
		RunMetricServer: func(cfg *metricserver.Config) {

		},
	})
	require.NoError(t, err)
}
func TestApplicationFailToServe(t *testing.T) {
	cfg := &config.Config{}
	gormStarter := &adapter.MockGormStarter{}
	gormStarter.On("ConnectToGorm", adapter.Config{
		Master:           cfg.DB.MasterDSN,
		Replicas:         []string{cfg.DB.ReplicaDSN},
		MaxIdleConns:     cfg.DB.MaxIdle,
		MaxOpenConns:     cfg.DB.MaxOpen,
		MaxLifetimeConns: cfg.DB.MaxLifeTime,
	}, &gorm.Config{}).Return(&gorm.DB{}, nil)
	ctrl := gomock.NewController(t)
	mockGRPCServer := mocks.NewMockServer(ctrl)
	mockGRPCServer.EXPECT().Serve().Return(errors.New("something bad happen"))
	err := Application(&Option{
		GormStarter: gormStarter,
		Cfg:         cfg,
		NewGRPCServer: func(o *localOption) Server {
			return mockGRPCServer
		},
		RunMetricServer: func(cfg *metricserver.Config) {},
	})
	require.Error(t, err)
}
func TestApplicationFailToConnectToDB(t *testing.T) {
	cfg := &config.Config{}
	gormStarter := &adapter.MockGormStarter{}
	gormStarter.On("ConnectToGorm", adapter.Config{
		Master:           cfg.DB.MasterDSN,
		Replicas:         []string{cfg.DB.ReplicaDSN},
		MaxIdleConns:     cfg.DB.MaxIdle,
		MaxOpenConns:     cfg.DB.MaxOpen,
		MaxLifetimeConns: cfg.DB.MaxLifeTime,
	}, &gorm.Config{}).Return(&gorm.DB{}, errors.New("Something Bad Happen"))
	ctrl := gomock.NewController(t)
	mockGRPCServer := mocks.NewMockServer(ctrl)
	err := Application(&Option{
		GormStarter: gormStarter,
		Cfg:         cfg,
		NewGRPCServer: func(o *localOption) Server {
			return mockGRPCServer
		},
	})
	require.Error(t, err)
}

func TestGetDefaultConfig(t *testing.T) {
	opt := GetDefaultOption(&config.Config{})
	require.NotNil(t, opt.Cfg)
	require.NotNil(t, opt.GormStarter)
	require.NotNil(t, opt.RunMetricServer)
	require.NotNil(t, opt.NewGRPCServer)
}
