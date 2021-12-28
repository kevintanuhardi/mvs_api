package mysql

import (
	"context"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kevintanuhardi/mvs_api/adapter"
	"github.com/kevintanuhardi/mvs_api/domain/user/dto"
	"github.com/kevintanuhardi/mvs_api/domain/user/entity"
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

func createUser(t *testing.T, db *gorm.DB) {
	user := &entity.User{
		EmployeeId:  "JINX-666",
		Name:        "Budi",
		CompanyId:   1,
		Active:      false,
		PhoneNumber: "08123123123123",
		Email:       "jane@doe.com",
		Password:    "jinxpowpow",
	}

	// Adding data
	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	// Cleaning it up
	t.Cleanup(func() {
		db.Delete(&user)
	})
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
					EmployeeId:  "testEmployeeId",
					CompanyId:   1,
					Active:      true,
					PhoneNumber: "08123123123123",
					Email:       "test@mail.com",
					Password:    "test",
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
			if _, err := r.UserRegister(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("repo.UserRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Cleanup(func() {
				gormDB.Delete(tt.args.userData)
			})
		})
	}
}

func Test_repo_FindByPhoneNumber(t *testing.T) {
	type args struct {
		ctx         context.Context
		phoneNumber string
	}
	tests := []struct {
		name    string
		r       *repo
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FindByPhoneNumber(tt.args.ctx, tt.args.phoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.FindByPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.FindByPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_FindByEmployeeId(t *testing.T) {
	type args struct {
		ctx        context.Context
		employeeId string
	}
	tests := []struct {
		name    string
		r       *repo
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FindByEmployeeId(tt.args.ctx, tt.args.employeeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.FindByEmployeeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.FindByEmployeeId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_UserActivation(t *testing.T) {
	gormDB, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx      context.Context
		userData *dto.UserActivateRequest
	}

	createUser(t, gormDB)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully update field active to 1 in user table",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				userData: &dto.UserActivateRequest{
					EmployeeId:  "JINX-666",
					CompanyCode: "ASD",
				},
			},
			wantErr: false,
		},
		{
			name: "Should return Company Code tidak ditemukan",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				userData: &dto.UserActivateRequest{
					EmployeeId:  "JINX-666",
					CompanyCode: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Should return BFF ID tidak ditemukan",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				userData: &dto.UserActivateRequest{
					EmployeeId:  "",
					CompanyCode: "ASD",
				},
			},
			wantErr: true,
		},
		{
			name: "Failed to activate user with active = 1",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				userData: &dto.UserActivateRequest{
					EmployeeId:  "JINX-666",
					CompanyCode: "ASD",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		r := NewRepository(tt.fields.db)
		t.Run(tt.name, func(t *testing.T) {
			if _, err := r.UserActivation(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("repo.UserActivation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
