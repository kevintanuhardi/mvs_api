package mysql

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/db"
)

func (r *repo) FindOrder(ctx context.Context) ([]*entity.OrderAggregate, error) {
	page, pageSize := 1, 10
	rows, err := r.db.Table("sale_order").
		Scopes(db.Paginate(page, pageSize)).
		Select(`
			  sale_order.id, sale_order.customer_id, sale_order.trx_id
			, sale_order.status, sale_order.created_at, sale_order.updated_at
			, sos.id as order_store_id, sos.store_id, sos.store_trx_id
			, sosi.id as order_store_item_id, sosi.sku, sosi.name, sosi.uom
			, sosi.quantity, sosi.price_unit
		`).
		Joins("left join sale_order_store sos on sale_order.id = sos.order_id").
		Joins("left join sale_order_store_item sosi on sos.id = sosi.order_store_id").
		Rows()

	if err != nil {
		return []*entity.OrderAggregate{}, err
	}
	defer rows.Close()

	aggs := []*entity.OrderAggregate{}
	for rows.Next() {
		err := r.db.ScanRows(rows, &aggs)
		if err != nil {
			return []*entity.OrderAggregate{}, err
		}
	}

	return aggs, nil
}
