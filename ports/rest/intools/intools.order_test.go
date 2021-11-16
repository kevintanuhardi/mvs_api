package intools

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.warungpintar.co/sales-platform/brook/domain/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
)

func (obj *testObject) TestGetOrderSuccess() {
	obj.MockService.EXPECT().OrderList(gomock.Any()).Return([]*dto.OrderDTO{
		{
			OrderStores: []dto.OrderStoreDTO{
				{
					OrderStoreItems: []entity.OrderStoreItem{
						{},
					},
				},
			},
		},
		{
			OrderStores: []dto.OrderStoreDTO{
				{
					OrderStoreItems: []entity.OrderStoreItem{
						{},
					},
				},
			},
		},
	}, nil)
	response := obj.module.GetOrder(obj.writer, obj.request)
	require.NotNil(obj.T(), response)
}

func (obj *testObject) TestGetOrderFailed() {
	obj.MockService.EXPECT().OrderList(gomock.Any()).Return([]*dto.OrderDTO{}, errors.New("something Bad Happen"))
	response := obj.module.GetOrder(obj.writer, obj.request)
	require.NotNil(obj.T(), response)
}
