package dto

import "gitlab.warungpintar.co/sales-platform/brook/domain/entity"

type OrderDTO struct {
	Order       entity.Order    `json:"order"`
	OrderStores []OrderStoreDTO `json:"order_stores"`
}

type OrderStoreDTO struct {
	entity.OrderStore
	OrderStoreItems []entity.OrderStoreItem `json:"order_store_items"`
}
