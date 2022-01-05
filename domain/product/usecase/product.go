package usecase

import (
	"context"
	"fmt"

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
	productEntity, err := s.products.CreateProduct(
		ctx,
		&entity.Product{
			Name:       productRequest.Name,
			CategoryId: productRequest.CategoryId,
			SkuNo:      productRequest.SkuNo,
		},
	)
	fmt.Println(productEntity)
	
	if err != nil {
		return nil, err
	}

	return &dto.CreateProductResponse{}, nil
}
