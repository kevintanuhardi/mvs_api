package dto

type ProductUomInfo struct {
	Conversion int	`json:"conversion"`
	IsSale bool `json:"is_sale"`
	Price float32 `json:"price"`
	Name string `json:"name"`
}

type CreateProductRequest struct {
	Name string `json:"name"`
	// CategoryId int	`json:"category_id"`
	SkuNo string `json:"sku_no"`
	ProductUomList []ProductUomInfo `json:"product_uom_list"`
}

type CreateProductResponse struct {

}