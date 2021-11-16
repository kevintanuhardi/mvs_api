package grpc

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/dto"
	pb "gitlab.warungpintar.co/sales-platform/brook/proto/brook"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (se *server) GetOrderList(context.Context, *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	data, err := se.Usecase.OrderList(context.Background())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return translateOrderList(data), nil
}
func translateOrderList(req []*dto.OrderDTO) *pb.GetOrderResponse {
	resp := &pb.GetOrderResponse{}
	for _, order := range req {
		respOrder := &pb.Order{
			Id:         int64(order.Order.ID),
			CustomerID: int64(order.Order.CustomerID),
			TrxID:      order.Order.TrxID,
			Status:     order.Order.Status,
			CreatedAt:  order.Order.CreatedAt.Unix(),
			UpdatedAt:  order.Order.UpdatedAt.Unix(),
			Timezone:   order.Order.CreatedAt.Location().String(),
		}

		respOrderStores := make([]*pb.OrderStores, 0)
		for i := range order.OrderStores {
			respOrderStores := &pb.OrderStores{
				ID:              int64(order.OrderStores[i].ID),
				OrderID:         int64(order.OrderStores[i].OrderID),
				StoreID:         int64(order.OrderStores[i].StoreID),
				StoreTrxID:      order.OrderStores[i].StoreTrxID,
				CreatedAt:       order.OrderStores[i].CreatedAt.Unix(),
				UpdatedAt:       order.OrderStores[i].UpdatedAt.Unix(),
				Timezone:        order.OrderStores[i].CreatedAt.Location().String(),
				OrderStoreItems: make([]*pb.OrderStoreItems, 0),
			}
			for j := range order.OrderStores[i].OrderStoreItems {
				respOrderStores.OrderStoreItems = append(respOrderStores.OrderStoreItems,
					&pb.OrderStoreItems{
						ID:           int64(order.OrderStores[i].OrderStoreItems[j].ID),
						OrderStoreID: int64(order.OrderStores[i].OrderStoreItems[j].OrderStoreID),
						SKU:          order.OrderStores[i].OrderStoreItems[j].SKU,
						Name:         order.OrderStores[i].OrderStoreItems[j].Name,
						Uom:          order.OrderStores[i].OrderStoreItems[j].Uom,
						Quantity:     order.OrderStores[i].OrderStoreItems[j].Quantity,
						PriceUnit:    order.OrderStores[i].OrderStoreItems[j].PriceUnit,
						CreatedAt:    order.OrderStores[i].OrderStoreItems[j].CreatedAt.Unix(),
						UpdatedAt:    order.OrderStores[i].OrderStoreItems[j].UpdatedAt.Unix(),
						Timezone:     order.OrderStores[i].OrderStoreItems[j].CreatedAt.Location().String(),
					})
			}
		}
		respOrder.OrderStores = respOrderStores
		resp.Order = append(resp.Order, respOrder)
	}
	return resp
}
