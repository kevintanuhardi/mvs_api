package mysql

import (
	"context"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.warungpintar.co/sales-platform/brook/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	"gorm.io/gorm"
)

type fields struct {
	db *gorm.DB
}

func openDB() (*gorm.DB, error) {
	db := adapter.NewGormStarter()
	return db.ConnectToGorm(adapter.Config{
		Master: os.Getenv("BROOK_MASTER_DB_DSN"),
	}, &gorm.Config{})
}

func createCompany(t *testing.T, db *gorm.DB) {
	company := &entity.Company{
		Code:        "ASD",
		Name:        "Riot",
		Address:     "Riot Street 1234",
		Country:     "US",
		Province:    "California",
		City:        "LA",
		District:    "Rawa Buntu",
		Village:     "Breeze",
		PostalCode:  "123123",
		PhoneNumber: "+6211111111",
		FaxNumber:   "+6211111111",
		NPWP:        "npwnwpwnwpp",
		SPPKP:       "skipipiw",
	}

	// Adding data
	if err := db.Create(&company).Error; err != nil {
		t.Fatal(err)
	}

	// Cleaning it up
	t.Cleanup(func() {
		db.Delete(&company)
	})
}

func Test_repo_CompanyRegister(t *testing.T) {
	gormDB, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx         context.Context
		companyData *entity.Company
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully registered company",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				companyData: &entity.Company{
					Code:        "ASD2",
					Name:        "Riot",
					Address:     "Riot Street 1234",
					Country:     "US",
					Province:    "California",
					City:        "LA",
					District:    "Rawa Buntu",
					Village:     "Breeze",
					PostalCode:  "123123",
					PhoneNumber: "+6211111111",
					FaxNumber:   "+6211111111",
					NPWP:        "npwnwpwnwpp",
					SPPKP:       "skipipiw",
				},
			},
			wantErr: false,
		},
		{
			name: "Should return error when there is no company data",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:         context.Background(),
				companyData: nil,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := NewRepository(tt.fields.db)
		t.Run(tt.name, func(t *testing.T) {
			if err := r.CompanyRegister(tt.args.ctx, tt.args.companyData); (err != nil) != tt.wantErr {
				t.Errorf("repo.CompanyRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Delete data after testing
		t.Cleanup(func() {
			gormDB.Delete(tt.args.companyData)
		})
	}
}
