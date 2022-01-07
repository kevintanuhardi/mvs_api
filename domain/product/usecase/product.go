package usecase

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/product/dto"
	"github.com/kevintanuhardi/mvs_api/domain/product/entity"
	"github.com/kevintanuhardi/mvs_api/domain/product/repository"
)

type Service struct {
	products repository.Repository
}

type ServiceManager interface {
	CreateProduct(ctx context.Context, productRequest *dto.CreateProductRequest) (user *dto.CreateProductResponse, err error)
}

func NewService(product repository.Repository) *Service {
	return &Service{product}
}

func (s *Service) CreateProduct(ctx context.Context, productRequest *dto.CreateProductRequest) (user *dto.CreateProductResponse, err error) {
	productUomEntList := make([]*entity.ProductUom, len(productRequest.ProductUomList)) 
	productEntity, err := s.products.CreateProduct(
		ctx,
		&entity.Product{
			Name:       productRequest.Name,
			// CategoryId: productRequest.CategoryId,
			SkuNo:      productRequest.SkuNo,
		},
	)
	if err != nil {
		return nil, err
	}

	for idx, productUomReq := range productRequest.ProductUomList {
		productUomEntList[idx] = &entity.ProductUom{
			ProductId:  productEntity.ID,
			Name: 			productUomReq.Name,
			Conversion: productUomReq.Conversion,
			IsSale:     productUomReq.IsSale,
			Price:      productUomReq.Price,
		}
	}

	s.products.BulkCreateProductUom(ctx, productUomEntList)
	
	if err != nil {
		return nil, err
	}

	return &dto.CreateProductResponse{}, nil
}
