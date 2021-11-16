package intools

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.warungpintar.co/sales-platform/brook/domain/usecase"
	usecasemock "gitlab.warungpintar.co/sales-platform/brook/domain/usecase/mocks"
	routermocks "gitlab.warungpintar.co/sales-platform/brook/pkg/router/mocks"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler(&usecase.Service{})
	require.NotNil(t, handler)
}

type testObject struct {
	suite.Suite
	module                *Intools
	MockService           *usecasemock.MockServiceManager
	mockRouterRegistrator *routermocks.MockRegistrator
	request               *http.Request
	writer                http.ResponseWriter
}

func (obj *testObject) SetupTest() {
	gomockController := gomock.NewController(obj.T())
	obj.mockRouterRegistrator = routermocks.NewMockRegistrator(gomockController)
	obj.MockService = usecasemock.NewMockServiceManager(gomockController)
	obj.request = &http.Request{
		URL: &url.URL{},
	}
	obj.writer = httptest.NewRecorder()
	obj.module = &Intools{
		service: obj.MockService,
		prefix:  "",
	}
}

func Test_public(t *testing.T) {
	suite.Run(t, new(testObject))
}

func (obj *testObject) TestRegister() {
	obj.mockRouterRegistrator.EXPECT().Register(http.MethodGet, "/ping", gomock.Any()).Return()
	obj.mockRouterRegistrator.EXPECT().Register(http.MethodGet, "/order", gomock.Any()).Return()
	obj.module.Register(obj.mockRouterRegistrator)
}

func (obj *testObject) TestPing() {
	response := obj.module.PING(obj.writer, obj.request)
	require.NotNil(obj.T(), response)
}
