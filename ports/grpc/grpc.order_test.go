package grpc

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	"gitlab.warungpintar.co/sales-platform/brook/domain/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
	"gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (obj *testObject) TestOrderList() {
	obj.Run("Success", func() {
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
		resp, err := obj.module.GetOrderList(context.Background(), &brook.GetOrderRequest{})
		obj.NoError(err)
		obj.Len(resp.Order, 2)
	})
	obj.Run("Error", func() {
		obj.MockService.EXPECT().OrderList(gomock.Any()).Return(nil, errors.New("Something Bad Happen"))
		resp, err := obj.module.GetOrderList(context.Background(), &brook.GetOrderRequest{})
		obj.Error(err)
		obj.Nil(resp)
	})
}
