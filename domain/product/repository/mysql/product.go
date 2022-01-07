package mysql

import (
	"context"
	"fmt"

	"github.com/kevintanuhardi/mvs_api/domain/product/entity"
	"github.com/kevintanuhardi/mvs_api/internal/constants"
)

func (r *repo) CreateProduct(ctx context.Context, productEnt *entity.Product) (user *entity.Product, err error) {
	if err := r.db.Create(productEnt).Error; err != nil {
		fmt.Println(err)
		return nil, constants.GetErrDatabaseError()
	}
	return productEnt, nil
}
