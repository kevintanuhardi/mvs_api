package adapter

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	NewGormStarter()
}

func TestConnectToGormSuccess(t *testing.T) {
	mockOpener := &mockGormOpenInterface{}
	mockGormDB := &mockGormdb{}
	cfg, gormcfg := Config{
		Replicas: []string{""},
	}, &gorm.Config{}
	mockGormDB.On("DB").Return(&sql.DB{}, nil)
	mockGormDB.On("Use", mock.Anything).Return(nil)
	mockOpener.On("OpenConnection", cfg, gormcfg).Return(gormdb(mockGormDB), nil)
	mockOpener.On("GetOriginal").Return(&gorm.DB{})
	objTest := gormStarter{
		Opener: mockOpener,
	}
	db, err := objTest.ConnectToGorm(cfg, gormcfg)
	mockOpener.AssertExpectations(t)
	mockGormDB.AssertExpectations(t)
	require.NoError(t, err)
	require.NotNil(t, db)
}

func TestConnectToGormFailOnRegisterReplica(t *testing.T) {
	mockOpener := &mockGormOpenInterface{}
	mockGormDB := &mockGormdb{}
	cfg, gormcfg := Config{
		Replicas: []string{""},
	}, &gorm.Config{}
	mockGormDB.On("DB").Return(&sql.DB{}, nil)
	mockOpener.On("OpenConnection", cfg, gormcfg).Return(gormdb(mockGormDB), nil)
	mockGormDB.On("Use", mock.Anything).Return(errors.New("something bad happen"))
	objTest := gormStarter{
		Opener: mockOpener,
	}
	_, err := objTest.ConnectToGorm(cfg, gormcfg)
	mockOpener.AssertExpectations(t)
	mockGormDB.AssertExpectations(t)
	require.Error(t, err)
	require.Equal(t, "something bad happen", err.Error())
}

func TestConnectToGormFailOnGettingDBConnection(t *testing.T) {
	mockOpener := &mockGormOpenInterface{}
	mockGormDB := &mockGormdb{}
	cfg, gormcfg := Config{
		Replicas: []string{""},
	}, &gorm.Config{}
	mockOpener.On("OpenConnection", cfg, gormcfg).Return(gormdb(mockGormDB), nil)
	mockGormDB.On("DB").Return(&sql.DB{}, errors.New("something bad happen"))
	objTest := gormStarter{
		Opener: mockOpener,
	}
	_, err := objTest.ConnectToGorm(cfg, gormcfg)
	mockOpener.AssertExpectations(t)
	mockGormDB.AssertExpectations(t)
	require.Error(t, err)
	require.Equal(t, "something bad happen", err.Error())
}

func TestConnectToGormFailOnOpenConnection(t *testing.T) {
	mockOpener := &mockGormOpenInterface{}
	mockGormDB := &mockGormdb{}
	cfg, gormcfg := Config{
		Replicas: []string{""},
	}, &gorm.Config{}
	mockOpener.On("OpenConnection", cfg, gormcfg).Return(gormdb(mockGormDB), errors.New("something bad happen"))
	objTest := gormStarter{
		Opener: mockOpener,
	}
	_, err := objTest.ConnectToGorm(cfg, gormcfg)
	mockOpener.AssertExpectations(t)
	mockGormDB.AssertExpectations(t)
	require.Error(t, err)
	require.Equal(t, "something bad happen", err.Error())
}
