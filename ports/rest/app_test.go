package rest

// import (
// 	"errors"
// 	"io"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/opentracing/opentracing-go"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/kevintanuhardi/mvs_api/adapter"
// 	"github.com/kevintanuhardi/mvs_api/config"
// 	"github.com/kevintanuhardi/mvs_api/pkg/metricserver"
// 	"github.com/kevintanuhardi/mvs_api/pkg/router"
// 	routermock "github.com/kevintanuhardi/mvs_api/pkg/router/mocks"
// 	"github.com/kevintanuhardi/mvs_api/pkg/webservice"
// 	webmocks "github.com/kevintanuhardi/mvs_api/pkg/webservice/mocks"
// 	"gorm.io/gorm"
// )

// type testSuite struct {
// 	suite.Suite
// }

// func (obj *testSuite) TestApplication() {
// 	cfg := &config.Config{}

// 	obj.Run("Error jaeger service empty", func() {
// 		gormStarter := obj.connectToGorm(cfg)
// 		err := Application(&Config{
// 			Cfg:         cfg,
// 			GormStarter: gormStarter,
// 			InitTracer: func(serviceName string) (opentracing.Tracer, io.Closer, error) {
// 				return nil, nil, errors.New("Something is wrong")
// 			},
// 		})

// 		require.Error(obj.T(), err)
// 	})

// 	obj.Run("Error database error", func() {
// 		gormStarter := obj.notConnectToGorm(cfg)
// 		err := Application(&Config{
// 			Cfg:         cfg,
// 			GormStarter: gormStarter,
// 		})

// 		require.Error(obj.T(), err)
// 	})
// }
// func TestApplicationSuccessConnect(t *testing.T) {
// 	cfg := &config.Config{}
// 	controller := gomock.NewController(t)
// 	mockWebServer := webmocks.NewMockWebService(controller)
// 	// mockWebRegistrator := webmocks.NewMockWebRegistrator(controller)
// 	mockRouterRegistrator := routermock.NewMockRegistrator(controller)
// 	gormStarter := &adapter.MockGormStarter{}
// 	gormStarter.On("ConnectToGorm", adapter.Config{
// 		Master:           cfg.DB.MasterDSN,
// 		Replicas:         []string{cfg.DB.ReplicaDSN},
// 		MaxIdleConns:     cfg.DB.MaxIdle,
// 		MaxOpenConns:     cfg.DB.MaxOpen,
// 		MaxLifetimeConns: cfg.DB.MaxLifeTime,
// 	}, &gorm.Config{}).Return(&gorm.DB{}, nil)
// 	mockWebServer.EXPECT().Run().Return(nil)
// 	err := Application(&Config{
// 		Cfg:         cfg,
// 		GormStarter: gormStarter,
// 		NewWebService: func(port string, routerRegistrator router.Registrator, registrators ...webservice.WebRegistrator) webservice.WebService {
// 			return mockWebServer
// 		},
// 		// GenerateWebRegistrator: func(service usecase.ServiceManager) []webservice.WebRegistrator {
// 		// 	return []webservice.WebRegistrator{mockWebRegistrator}
// 		// },
// 		GenerateRouter: func(tracer opentracing.Tracer) router.Registrator {
// 			return mockRouterRegistrator
// 		},
// 		InitTracer: func(serviceName string) (opentracing.Tracer, io.Closer, error) {
// 			return opentracing.GlobalTracer(), nil, nil
// 		},
// 		RunMetricServer: func(cfg *metricserver.Config) {

// 		},
// 	})

// 	require.NoError(t, err)
// }
// func (obj *testSuite) connectToGorm(cfg *config.Config) *adapter.MockGormStarter {
// 	gormStarter := &adapter.MockGormStarter{}
// 	gormStarter.On("ConnectToGorm", adapter.Config{
// 		Master:           cfg.DB.MasterDSN,
// 		Replicas:         []string{cfg.DB.ReplicaDSN},
// 		MaxIdleConns:     cfg.DB.MaxIdle,
// 		MaxOpenConns:     cfg.DB.MaxOpen,
// 		MaxLifetimeConns: cfg.DB.MaxLifeTime,
// 	}, &gorm.Config{}).Return(&gorm.DB{}, nil)
// 	return gormStarter
// }

// func (obj *testSuite) notConnectToGorm(cfg *config.Config) *adapter.MockGormStarter {
// 	gormStarter := &adapter.MockGormStarter{}
// 	gormStarter.On("ConnectToGorm", adapter.Config{
// 		Master:           cfg.DB.MasterDSN,
// 		Replicas:         []string{cfg.DB.ReplicaDSN},
// 		MaxIdleConns:     cfg.DB.MaxIdle,
// 		MaxOpenConns:     cfg.DB.MaxOpen,
// 		MaxLifetimeConns: cfg.DB.MaxLifeTime,
// 	}, &gorm.Config{}).Return(&gorm.DB{}, errors.New("db error"))
// 	return gormStarter
// }

// func Test_Application(t *testing.T) {
// 	suite.Run(t, new(testSuite))
// }
// func TestGetRouter(t *testing.T) {
// 	module := getRouter(opentracing.GlobalTracer())
// 	require.NotNil(t, module)
// }
// func TestGetWebRegistrator(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	// mockUsecase := usecasemocks.NewMockServiceManager(controller)
// 	registrators := getWebRegistrator(mockUsecase)
// 	require.Len(t, registrators, 2)
// }
// func TestGetDefaultConfig(t *testing.T) {
// 	defaultConfig := GetDefaultConfig(&config.Config{})
// 	require.NotNil(t, defaultConfig)
// 	require.NotNil(t, defaultConfig.GormStarter)
// 	require.NotNil(t, defaultConfig.NewWebService)
// 	require.NotNil(t, defaultConfig.GenerateRouter)
// 	require.NotNil(t, defaultConfig.GenerateWebRegistrator)
// 	require.NotNil(t, defaultConfig.InitTracer)
// }
