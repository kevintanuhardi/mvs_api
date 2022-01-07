package repository

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/product/entity"
)

type Repository interface {
	CreateProduct(ctx context.Context, productEnt *entity.Product) (*entity.Product, error)
	// CreateProductUom(ctx context.Context, productUomEnt *entity.ProductUom) (*entity.ProductUom, error)
	BulkCreateProductUom(ctx context.Context, productUomList []*entity.ProductUom) ([]*entity.ProductUom, error)
}
