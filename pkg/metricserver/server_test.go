package metricserver

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
	routermock "gitlab.warungpintar.co/sales-platform/brook/pkg/router/mocks"
	webmocks "gitlab.warungpintar.co/sales-platform/brook/pkg/webservice/mocks"

	"gitlab.warungpintar.co/sales-platform/brook/pkg/webservice"
)

func TestGetDefaultConfig(t *testing.T) {
	cfg := GetDefaultConfig(3007)
	require.NotNil(t, cfg)
	require.NotNil(t, cfg.NewWebService)
	require.NotNil(t, cfg.Port)
	require.NotNil(t, cfg.Router)
}
func TestRunMetricServer(t *testing.T) {
	controller := gomock.NewController(t)
	mocksRouterRegistrator := routermock.NewMockRegistrator(controller)
	mocksWebService := webmocks.NewMockWebService(controller)
	mocksWebService.EXPECT().Run().Return(nil)
	RunMetricServer(&Config{
		Port: 3007,
		NewWebService: func(port string, routerRegistrator router.Registrator, registrators ...webservice.WebRegistrator) webservice.WebService {
			return mocksWebService
		},
		Router: mocksRouterRegistrator,
	})
}
func TestRunMetricServerError(t *testing.T) {
	controller := gomock.NewController(t)
	mocksRouterRegistrator := routermock.NewMockRegistrator(controller)
	mocksWebService := webmocks.NewMockWebService(controller)
	mocksWebService.EXPECT().Run().Return(errors.New("Something Bad Happen"))
	RunMetricServer(&Config{
		Port: 3007,
		NewWebService: func(port string, routerRegistrator router.Registrator, registrators ...webservice.WebRegistrator) webservice.WebService {
			return mocksWebService
		},
		Router: mocksRouterRegistrator,
	})
}
