package repository

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/product/entity"
)

type Repository interface {
	CreateProduct(ctx context.Context, userData *entity.Product) (user *entity.Product, err error)
	CreateProductUom(ctx context.Context, userData *entity.ProductUom) (user *entity.ProductUom, err error)
}
