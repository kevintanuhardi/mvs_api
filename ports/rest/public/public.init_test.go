package public

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/kevintanuhardi/mvs_api/domain"
// 	usecasemock "github.com/kevintanuhardi/mvs_api/domain/user/usecase/mocks"
// 	routermocks "github.com/kevintanuhardi/mvs_api/pkg/router/mocks"
// )

// func TestNewHandler(t *testing.T) {
// 	handler := NewHandler(domain.DomainService{})
// 	require.NotNil(t, handler)
// }

// type testObject struct {
// 	suite.Suite
// 	module                *Public
// 	MockService           *usecasemock.MockServiceManager
// 	mockRouterRegistrator *routermocks.MockRegistrator
// 	request               *http.Request
// 	writer                http.ResponseWriter
// }

// func (obj *testObject) SetupTest() {
// 	gomockController := gomock.NewController(obj.T())
// 	obj.mockRouterRegistrator = routermocks.NewMockRegistrator(gomockController)
// 	obj.MockService = usecasemock.NewMockServiceManager(gomockController)
// 	obj.request = &http.Request{
// 		URL: &url.URL{},
// 	}
// 	obj.writer = httptest.NewRecorder()
// 	obj.module = &Public{
// 		service: obj.MockService,
// 		prefix:  "",
// 	}
// }

// func Test_public(t *testing.T) {
// 	suite.Run(t, new(testObject))
// }

// func (obj *testObject) TestRegister() {
// 	obj.mockRouterRegistrator.EXPECT().Register(http.MethodGet, "/ping", gomock.Any()).Return()
// 	obj.mockRouterRegistrator.EXPECT().Register(http.MethodGet, "/user", gomock.Any()).Return()
// 	obj.module.Register(obj.mockRouterRegistrator)
// }

// func (obj *testObject) TestPing() {
// 	response := obj.module.PING(obj.writer, obj.request)
// 	require.NotNil(obj.T(), response)
// }
