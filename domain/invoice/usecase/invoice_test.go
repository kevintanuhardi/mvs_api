package usecase

import (
	"context"
	"testing"
)

func TestService_GetInvoice(t *testing.T) {
	type args struct {
		ctx context.Context
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
			s:    &Service{},
			args: args{
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if err := s.GetInvoice(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Service.GetInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
