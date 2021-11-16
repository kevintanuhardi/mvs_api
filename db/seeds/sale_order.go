package seeds

import (
	"time"
)

type SaleOrder struct {
	ID         int
	CustomerID int
	TrxID      string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (SaleOrder) TableName() string {
	return "sale_order"
}

var saleorders = []SaleOrder{
	{
		ID:         1,
		CustomerID: 666,
		TrxID:      "SO12345",
		Status:     "order.confirmed",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}

func (s Seed) SaleOrderSeed() {
	for i := range saleorders {
		c := saleorders[i]
		err := s.db.Create(&c).Error
		if err != nil {
			panic(err)
		}
	}
}
