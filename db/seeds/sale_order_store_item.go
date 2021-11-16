package seeds

import (
	"time"
)

type SaleOrderStoreItem struct {
	ID           int
	OrderStoreID int
	SKU          string
	Name         string
	Uom          string
	Quantity     float64
	PriceUnit    float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (SaleOrderStoreItem) TableName() string {
	return "sale_order_store_item"
}

var saleorderstoreitems = []SaleOrderStoreItem{
	{
		ID:           1,
		OrderStoreID: 1,
		SKU:          "ABC001",
		Name:         "ABC Kopi Susu",
		Uom:          "Renceng",
		Quantity:     20,
		PriceUnit:    10000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           2,
		OrderStoreID: 1,
		SKU:          "BER001",
		Name:         "Bear Brand",
		Uom:          "Dus",
		Quantity:     10,
		PriceUnit:    100000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           3,
		OrderStoreID: 1,
		SKU:          "DJI001",
		Name:         "Dji Sam Soe",
		Uom:          "Slop",
		Quantity:     5,
		PriceUnit:    250000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           4,
		OrderStoreID: 2,
		SKU:          "TPH001",
		Name:         "Teh Pucuk Harum",
		Uom:          "Dus",
		Quantity:     20,
		PriceUnit:    50000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           5,
		OrderStoreID: 2,
		SKU:          "OKY004",
		Name:         "Okky Jelly Dring Anggur",
		Uom:          "Dus",
		Quantity:     10,
		PriceUnit:    100000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		ID:           6,
		OrderStoreID: 2,
		SKU:          "SAM003",
		Name:         "Sampoerna Filter",
		Uom:          "Slop",
		Quantity:     5,
		PriceUnit:    350000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}

func (s Seed) SaleOrderStoreItemsSeed() {
	for i := range saleorderstoreitems {
		c := saleorderstoreitems[i]
		err := s.db.Create(&c).Error
		if err != nil {
			panic(err)
		}
	}
}
