package metricserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	routermock "gitlab.warungpintar.co/sales-platform/brook/pkg/router/mocks"
)

func TestRegister(t *testing.T) {
	controller := gomock.NewController(t)
	mocksRouterRegistrator := routermock.NewMockRegistrator(controller)

	mocksRouterRegistrator.EXPECT().Register(http.MethodGet, "/metrics", gomock.Any()).Return()
	module := &Metric{}

	module.Register(mocksRouterRegistrator)
}
func TestNewHandler(t *testing.T) {
	module := NewHandler()
	require.Equal(t, module.prefix, "")
}
func TestHandleMetrics(t *testing.T) {
	module := NewHandler()
	module.HandlerMetrics(httptest.NewRecorder(), &http.Request{})
}
