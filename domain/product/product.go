package product

import (
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain/product/repository"
	"github.com/kevintanuhardi/mvs_api/domain/product/usecase"
)

func NewProduct (cfg config.Config, products repository.Repository) usecase.ServiceManager {
	return usecase.NewService(products)
}