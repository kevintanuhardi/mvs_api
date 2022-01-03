package usecase

import (
	"context"
	"testing"

	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/dto"
	"github.com/kevintanuhardi/mvs_api/domain/invoice/repository/mysql"
	"github.com/kevintanuhardi/mvs_api/ports/rest"
)

func TestService_GetInvoice(t *testing.T) {
	cfg := config.Load()
	db, _ := rest.AppWithGorm(rest.GetDefaultConfig(cfg))
	type args struct {
		ctx context.Context
		request dto.GetInvoiceRequest
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			// s:    new{
			// 	mysql.NewRepository(db),
			// },
			s: NewService(mysql.NewRepository(db)),
			args: args{
				ctx: nil,
				request: dto.GetInvoiceRequest{
					InvoiceDate:          "03-01-2021",
					PaymentPeriodInDays:  4,
					ShippingContactName:  "drh Test",
					ShippingContactPhone: "081828282",
					ShippingAddress:      "Jl. Test",
					ShippingCity:         "Test",
					ShippingPostalCode:   "11515",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s
			if _, err := s.GetInvoice(tt.args.ctx, &tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Service.GetInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
