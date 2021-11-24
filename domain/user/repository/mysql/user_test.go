package mysql

import (
	"context"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.warungpintar.co/sales-platform/brook/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gorm.io/gorm"
)

func openDB() (*gorm.DB, error) {
	db := adapter.NewGormStarter()
	return db.ConnectToGorm(adapter.Config{
		Master: os.Getenv("BROOK_MASTER_DB_DSN"),
	}, &gorm.Config{})
}

func Test_repo_UserRegister(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipped if not in short mode")
	}

	gormDB, err := openDB()
	if err != nil {
		t.Fatal(err)
	}
	
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		userData *entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully registered user",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				userData: &entity.User{
					EmployeeId: "testEmployeeId",
					CompanyId: 1,
					Active: true,
					PhoneNumber: "08123123123123",
					Email: "test@mail.com",
					Password: "test",
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// r := &repo{
			// 	db: tt.fields.db,
			// }
			r := NewRepository(tt.fields.db)
			if err := r.UserRegister(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("repo.UserRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
