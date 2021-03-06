// Code generated by MockGen. DO NOT EDIT.
// Source: domain/user/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/kevintanuhardi/mvs_api/domain/user/entity"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindByPhoneNumber mocks base method.
func (m *MockRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPhoneNumber indicates an expected call of FindByPhoneNumber.
func (mr *MockRepositoryMockRecorder) FindByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPhoneNumber", reflect.TypeOf((*MockRepository)(nil).FindByPhoneNumber), ctx, phoneNumber)
}

// FindByPhoneNumberOrEmail mocks base method.
func (m *MockRepository) FindByPhoneNumberOrEmail(ctx context.Context, phoneNumber, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhoneNumberOrEmail", ctx, phoneNumber, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPhoneNumberOrEmail indicates an expected call of FindByPhoneNumberOrEmail.
func (mr *MockRepositoryMockRecorder) FindByPhoneNumberOrEmail(ctx, phoneNumber, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPhoneNumberOrEmail", reflect.TypeOf((*MockRepository)(nil).FindByPhoneNumberOrEmail), ctx, phoneNumber, email)
}

// UserRegister mocks base method.
func (m *MockRepository) UserRegister(ctx context.Context, userData *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegister", ctx, userData)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockRepositoryMockRecorder) UserRegister(ctx, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockRepository)(nil).UserRegister), ctx, userData)
}
