package usecase

import (
	"context"
	"fmt"

	"gitlab.warungpintar.co/sales-platform/brook/domain/dto"
	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
)

func (s *Service) OrderList(ctx context.Context) ([]*dto.OrderDTO, error) {
	orders := make([]*dto.OrderDTO, 0)

	aggs, err := s.orders.FindOrder(ctx)
	if err != nil {
		return orders, err
	}

	orders = s.corders(ctx, aggs)
	return orders, nil
}

func (s *Service) corders(ctx context.Context, aggs []*entity.OrderAggregate) []*dto.OrderDTO {
	maporderstore := s.orderstore(ctx, aggs)
	orders := make([]*dto.OrderDTO, 0)
	maporder := make(map[int]bool)
	for _, aa := range aggs {
		if _, found := maporder[aa.ID]; found {
			continue
		}
		maporder[aa.ID] = true

		orderStore := make([]dto.OrderStoreDTO, 0)
		if oo, found := maporderstore[aa.ID]; found {
			orderStore = oo
		}

		orders = append(orders, &dto.OrderDTO{
			Order: entity.Order{
				ID:         aa.ID,
				CustomerID: aa.CustomerID,
				TrxID:      aa.TrxID,
				Status:     aa.Status,
				CreatedAt:  aa.CreatedAt,
				UpdatedAt:  aa.UpdatedAt,
			},
			OrderStores: orderStore,
		})
	}

	return orders
}

func (s *Service) orderstore(ctx context.Context, aggs []*entity.OrderAggregate) map[int][]dto.OrderStoreDTO {
	maporderstoreitem := s.orderstoreitems(ctx, aggs)
	ordersStore := make(map[int][]dto.OrderStoreDTO)
	maporderstore := make(map[string]bool)
	for _, aa := range aggs {
		keyosi := fmt.Sprintf("%d%d", aa.ID, aa.OrderStoreID)
		if _, found := maporderstore[keyosi]; found {
			continue
		}
		maporderstore[keyosi] = true

		orderStoreItem := make([]entity.OrderStoreItem, 0)
		if oo, found := maporderstoreitem[keyosi]; found {
			orderStoreItem = oo
		}

		key := aa.ID
		ordersStore[key] = append(ordersStore[key], dto.OrderStoreDTO{
			OrderStore: entity.OrderStore{
				ID:         aa.OrderStoreID,
				OrderID:    aa.ID,
				StoreID:    aa.StoreID,
				StoreTrxID: aa.StoreTrxID,
				CreatedAt:  aa.CreatedAt,
				UpdatedAt:  aa.UpdatedAt,
			},
			OrderStoreItems: orderStoreItem,
		})
	}

	return ordersStore
}

func (s *Service) orderstoreitems(_ context.Context, aggs []*entity.OrderAggregate) map[string][]entity.OrderStoreItem {
	ordersStoreItem := make(map[string][]entity.OrderStoreItem)
	for _, aa := range aggs {
		key := fmt.Sprintf("%d%d", aa.ID, aa.OrderStoreID)
		ordersStoreItem[key] = append(ordersStoreItem[key], entity.OrderStoreItem{
			ID:           aa.OrderStoreItemID,
			OrderStoreID: aa.OrderStoreID,
			SKU:          aa.SKU,
			Name:         aa.Name,
			Uom:          aa.Uom,
			Quantity:     aa.Quantity,
			PriceUnit:    aa.PriceUnit,
			CreatedAt:    aa.CreatedAt,
			UpdatedAt:    aa.UpdatedAt,
		})
	}

	return ordersStoreItem
}
