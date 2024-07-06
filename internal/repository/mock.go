// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/interface.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/interface.go -destination=./internal/repository/mock.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	entity "github.com/pagu-project/Pagu/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// AddFaucet mocks base method.
func (m *MockDatabase) AddFaucet(f *entity.PhoenixFaucet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFaucet", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFaucet indicates an expected call of AddFaucet.
func (mr *MockDatabaseMockRecorder) AddFaucet(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFaucet", reflect.TypeOf((*MockDatabase)(nil).AddFaucet), f)
}

// AddUser mocks base method.
func (m *MockDatabase) AddUser(u *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockDatabaseMockRecorder) AddUser(u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockDatabase)(nil).AddUser), u)
}

// AddVoucher mocks base method.
func (m *MockDatabase) AddVoucher(v *entity.Voucher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVoucher", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVoucher indicates an expected call of AddVoucher.
func (mr *MockDatabaseMockRecorder) AddVoucher(v any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVoucher", reflect.TypeOf((*MockDatabase)(nil).AddVoucher), v)
}

// AddZealyUser mocks base method.
func (m *MockDatabase) AddZealyUser(u *entity.ZealyUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddZealyUser", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddZealyUser indicates an expected call of AddZealyUser.
func (mr *MockDatabaseMockRecorder) AddZealyUser(u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddZealyUser", reflect.TypeOf((*MockDatabase)(nil).AddZealyUser), u)
}

// CanGetFaucet mocks base method.
func (m *MockDatabase) CanGetFaucet(user *entity.User) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanGetFaucet", user)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanGetFaucet indicates an expected call of CanGetFaucet.
func (mr *MockDatabaseMockRecorder) CanGetFaucet(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanGetFaucet", reflect.TypeOf((*MockDatabase)(nil).CanGetFaucet), user)
}

// ClaimVoucher mocks base method.
func (m *MockDatabase) ClaimVoucher(id uint, txHash string, claimer uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimVoucher", id, txHash, claimer)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimVoucher indicates an expected call of ClaimVoucher.
func (mr *MockDatabaseMockRecorder) ClaimVoucher(id, txHash, claimer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimVoucher", reflect.TypeOf((*MockDatabase)(nil).ClaimVoucher), id, txHash, claimer)
}

// GetAllZealyUser mocks base method.
func (m *MockDatabase) GetAllZealyUser() ([]*entity.ZealyUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllZealyUser")
	ret0, _ := ret[0].([]*entity.ZealyUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllZealyUser indicates an expected call of GetAllZealyUser.
func (mr *MockDatabaseMockRecorder) GetAllZealyUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllZealyUser", reflect.TypeOf((*MockDatabase)(nil).GetAllZealyUser))
}

// GetUserInApp mocks base method.
func (m *MockDatabase) GetUserInApp(appID entity.AppID, callerID string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInApp", appID, callerID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInApp indicates an expected call of GetUserInApp.
func (mr *MockDatabaseMockRecorder) GetUserInApp(appID, callerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInApp", reflect.TypeOf((*MockDatabase)(nil).GetUserInApp), appID, callerID)
}

// GetVoucherByCode mocks base method.
func (m *MockDatabase) GetVoucherByCode(code string) (entity.Voucher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVoucherByCode", code)
	ret0, _ := ret[0].(entity.Voucher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoucherByCode indicates an expected call of GetVoucherByCode.
func (mr *MockDatabaseMockRecorder) GetVoucherByCode(code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoucherByCode", reflect.TypeOf((*MockDatabase)(nil).GetVoucherByCode), code)
}

// GetZealyUser mocks base method.
func (m *MockDatabase) GetZealyUser(id string) (*entity.ZealyUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZealyUser", id)
	ret0, _ := ret[0].(*entity.ZealyUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZealyUser indicates an expected call of GetZealyUser.
func (mr *MockDatabaseMockRecorder) GetZealyUser(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZealyUser", reflect.TypeOf((*MockDatabase)(nil).GetZealyUser), id)
}

// HasUser mocks base method.
func (m *MockDatabase) HasUser(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUser", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUser indicates an expected call of HasUser.
func (mr *MockDatabaseMockRecorder) HasUser(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUser", reflect.TypeOf((*MockDatabase)(nil).HasUser), id)
}

// ListVoucher mocks base method.
func (m *MockDatabase) ListVoucher() ([]*entity.Voucher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVoucher")
	ret0, _ := ret[0].([]*entity.Voucher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVoucher indicates an expected call of ListVoucher.
func (mr *MockDatabaseMockRecorder) ListVoucher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVoucher", reflect.TypeOf((*MockDatabase)(nil).ListVoucher))
}

// UpdateZealyUser mocks base method.
func (m *MockDatabase) UpdateZealyUser(id, txHash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZealyUser", id, txHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateZealyUser indicates an expected call of UpdateZealyUser.
func (mr *MockDatabaseMockRecorder) UpdateZealyUser(id, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZealyUser", reflect.TypeOf((*MockDatabase)(nil).UpdateZealyUser), id, txHash)
}
