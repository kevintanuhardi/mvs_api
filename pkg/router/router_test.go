package router_test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router/mocks"
)

func TestRegistration(t *testing.T) {
	controller := gomock.NewController(t)
	mockRouterRegistrator := mocks.NewMockRegistrator(controller)
	mockRouterRegistrator.EXPECT().Register(http.MethodGet, "/api/ping", gomock.Any()).Return()
	mockRouterRegistrator.EXPECT().Register(http.MethodOptions, "/api/ping", gomock.Any()).Return()
	mockRouterRegistrator.EXPECT().Register(http.MethodPost, "/api/ping", gomock.Any()).Return()
	mockRouterRegistrator.EXPECT().Register(http.MethodPut, "/api/ping", gomock.Any()).Return()
	module := router.New("/api", mockRouterRegistrator)
	require.NotNil(t, module)
	module.GET("/ping", PING)
	module.OPTIONS("/ping", PING)
	module.POST("/ping", PING)
	module.PUT("/ping", PING)
}
func PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
