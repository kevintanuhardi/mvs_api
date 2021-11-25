package mysql

import (
	"reflect"
	"testing"

	"gitlab.warungpintar.co/sales-platform/brook/domain/company/repository"
	"gorm.io/gorm"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want repository.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
