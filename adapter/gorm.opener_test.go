package adapter

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gorm.io/gorm"
)

func TestGormOpenerNew(t *testing.T) {
	newGormOpener()
}

func TestGormOpenerOpenConnection(t *testing.T) {
	objTest := gormOpener{
		gormOpen: func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
			return &gorm.DB{}, nil
		},
	}

	gormdb, err := objTest.OpenConnection(Config{}, &gorm.Config{})
	require.NoError(t, err)
	require.NotNil(t, gormdb)
}

func TestGormOpenerGetOriginal(t *testing.T) {
	objTest := gormOpener{
		gormOpen: func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
			return &gorm.DB{}, nil
		},
	}

	_, err := objTest.OpenConnection(Config{}, &gorm.Config{})
	require.NoError(t, err)
	resp := objTest.GetOriginal()
	require.NotNil(t, resp)
}
