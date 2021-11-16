package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.warungpintar.co/sales-platform/brook/domain/entity"
	"gitlab.warungpintar.co/sales-platform/brook/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OrderSuite struct {
	suite.Suite
	order  repository.Repository
	db     *gorm.DB
	dbmock sqlmock.Sqlmock
}

func TestRepoMain(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}

func (o *OrderSuite) SetupTest() {
	db, dbmock, err := sqlmock.New()
	require.NoError(o.T(), err)
	o.dbmock = dbmock
	dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow(1.0))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(o.T(), err)
	o.db = gormDB
	o.order = NewRepository(o.db)
}

func (o *OrderSuite) TestFindOrder() {
	orderAggregate := entity.OrderAggregate{}
	o.dbmock.ExpectQuery(("SELECT sale_order.id, " +
		"sale_order.customer_id, " +
		"sale_order.trx_id , " +
		"sale_order.status, " +
		"sale_order.created_at, " +
		"sale_order.updated_at , " +
		"sos.id as order_store_id, " +
		"sos.store_id, " +
		"sos.store_trx_id , " +
		"sosi.id as order_store_item_id, " +
		"sosi.sku, " +
		"sosi.name, " +
		"sosi.uom , " +
		"sosi.quantity, " +
		"sosi.price_unit FROM `sale_order` " +
		"left join sale_order_store sos on sale_order.id = sos.order_id " +
		"left join sale_order_store_item sosi on sos.id = sosi.order_store_id " +
		"LIMIT 10")).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"customer_id",
		"trx_id",
		"status",
		"created_at",
		"updated_at",
		"order_store_id",
		"store_id",
		"store_trx_id",
		"sku",
		"name",
		"uom",
		"quantity",
	}).AddRow(
		orderAggregate.ID,
		orderAggregate.CustomerID,
		orderAggregate.TrxID,
		orderAggregate.Status,
		orderAggregate.CreatedAt,
		orderAggregate.UpdatedAt,
		orderAggregate.OrderStoreID,
		orderAggregate.StoreID,
		orderAggregate.StoreTrxID,
		orderAggregate.SKU,
		orderAggregate.Name,
		orderAggregate.Uom,
		orderAggregate.Quantity,
	))
	aggs, err := o.order.FindOrder(context.Background())
	require.NoError(o.T(), err)
	require.Greater(o.T(), len(aggs), 0)
}
