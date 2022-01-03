package repository

import (
	"context"

	"github.com/kevintanuhardi/mvs_api/domain/invoice/entity"
)

type Repository interface {
	GetLastInvoice(ctx context.Context) (invoice *entity.Invoice, err error)
	CreateInvoice(ctx context.Context, invoiceData *entity.Invoice) (invoice *entity.Invoice, err error)
}
