package entity

import "time"

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CategoryId int	`json:"category_id"`
	SkuNo string `json:"sku_no"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}