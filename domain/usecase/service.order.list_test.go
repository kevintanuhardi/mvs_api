package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
	mockrepo "gitlab.warungpintar.co/sales-platform/brook/domain/repository/mocks"
)

type testObject struct {
	suite.Suite
	module         *Service
	mockRepository *mockrepo.MockRepository
}

func (obj *testObject) SetupTest() {
	gomockController := gomock.NewController(obj.T())
	mockrepository := mockrepo.NewMockRepository(gomockController)
	obj.mockRepository = mockrepository
	obj.module = &Service{
		orders: mockrepository,
	}
}

func Test_service(t *testing.T) {
	suite.Run(t, new(testObject))
}

func (obj *testObject) TestOrderList() {
	obj.Run("Success", func() {
		obj.mockRepository.EXPECT().FindOrder(gomock.Any()).Return([]*entity.OrderAggregate{
			{
				ID:               1,
				CustomerID:       123,
				TrxID:            "456",
				Status:           "paid",
				OrderStoreID:     123,
				StoreID:          123,
				StoreTrxID:       "123",
				OrderStoreItemID: 123,
				SKU:              "MOL101",
				Name:             "Molto",
				Uom:              "Bungkus",
				Quantity:         10,
				PriceUnit:        10000,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			{
				ID:               1,
				CustomerID:       123,
				TrxID:            "456",
				Status:           "paid",
				OrderStoreID:     123,
				StoreID:          123,
				StoreTrxID:       "123",
				OrderStoreItemID: 123,
				SKU:              "MOL101",
				Name:             "Molto",
				Uom:              "Bungkus",
				Quantity:         10,
				PriceUnit:        10000,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
		}, nil)
		resp, err := obj.module.OrderList(context.Background())
		obj.NoError(err)
		obj.NotNil(resp)
	})
	obj.Run("Fail to get order", func() {
		obj.mockRepository.EXPECT().FindOrder(gomock.Any()).Return(nil, errors.New("Something Bad Happen"))
		resp, err := obj.module.OrderList(context.Background())
		obj.Error(err)
		obj.NotNil(resp)
	})
}
func (obj *testObject) TestNew() {
	obj.Run("Create", func() {
		resp := NewService(obj.mockRepository)
		obj.Implements((*ServiceManager)(nil), resp)
	})
}
