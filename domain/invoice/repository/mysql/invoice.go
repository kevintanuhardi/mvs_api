package mysql

import (
	"context"
	"errors"

	"github.com/kevintanuhardi/mvs_api/domain/invoice/entity"
	"github.com/kevintanuhardi/mvs_api/internal/constants"
	"gorm.io/gorm"
)

func (r *repo) GetLastInvoice(ctx context.Context) (*entity.Invoice, error) {
	lastInvoice := &entity.Invoice{}


	if err := r.db.Last(lastInvoice).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return lastInvoice, constants.GetErrDatabaseError()
	}
	return lastInvoice, nil
}

func (r *repo) FindInvoiceByInvoiceId(ctx context.Context, invoiceId int) (*entity.Invoice, error) {
	invoice := &entity.Invoice{}


	if err := r.db.Where(&entity.Invoice{ID: &invoiceId}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return invoice, constants.GetErrDatabaseError()
	}
	return invoice, nil
}

func (r *repo) CreateInvoice(ctx context.Context, invoiceData *entity.Invoice) (invoice *entity.Invoice, err error) {
	if err := r.db.Create(invoiceData).Error; err != nil {
		return nil, constants.GetErrDatabaseError()
	}
	return invoiceData, nil
}
