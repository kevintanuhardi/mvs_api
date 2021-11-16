package seeds

import (
	"time"
)

type SaleOrderStore struct {
	ID         int
	OrderID    int
	StoreID    int
	StoreTrxID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (SaleOrderStore) TableName() string {
	return "sale_order_store"
}

var saleorderstores = []SaleOrderStore{
	{
		ID:         1,
		OrderID:    1,
		StoreID:    10,
		StoreTrxID: "SO-10-9891239",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
	{
		ID:         2,
		OrderID:    1,
		StoreID:    20,
		StoreTrxID: "SO-10-9876543",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}

func (s Seed) SaleOrderStoresSeed() {
	for i := range saleorderstores {
		c := saleorderstores[i]
		err := s.db.Create(&c).Error
		if err != nil {
			panic(err)
		}
	}
}
