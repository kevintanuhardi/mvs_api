package mysql

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/product/entity"
	"github.com/kevintanuhardi/mvs_api/internal/constants"
)

func (r *repo) BulkCreateProductUom(ctx context.Context, productUomList []*entity.ProductUom) ([]*entity.ProductUom, error) {
	if err := r.db.Create(productUomList).Error; err != nil {
		return nil, constants.GetErrDatabaseError()
	}
	return productUomList, nil
}