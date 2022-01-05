package entity

import (
	"time"
)

type ProductUom struct {
	ID int `json:"id"`
	ProductId string `json:"product_id"`
	// Conversion to smallest UOM
	Conversion int	`json:"conversion"`
	IsSale bool `json:"is_sale"`
	Price float32 `json:"price"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ProductUom) TableName() string {
	return "product_uom"
}