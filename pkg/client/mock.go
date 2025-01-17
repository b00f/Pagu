// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/client/interface.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/client/interface.go -destination=./pkg/client/mock.go -package=client
//

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	gomock "go.uber.org/mock/gomock"
)

// MockIClient is a mock of IClient interface.
type MockIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIClientMockRecorder
}

// MockIClientMockRecorder is the mock recorder for MockIClient.
type MockIClientMockRecorder struct {
	mock *MockIClient
}

// NewMockIClient creates a new mock instance.
func NewMockIClient(ctrl *gomock.Controller) *MockIClient {
	mock := &MockIClient{ctrl: ctrl}
	mock.recorder = &MockIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClient) EXPECT() *MockIClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIClient)(nil).Close))
}

// GetBalance mocks base method.
func (m *MockIClient) GetBalance(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockIClientMockRecorder) GetBalance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockIClient)(nil).GetBalance), arg0, arg1)
}

// GetBlockchainHeight mocks base method.
func (m *MockIClient) GetBlockchainHeight(arg0 context.Context) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockchainHeight", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainHeight indicates an expected call of GetBlockchainHeight.
func (mr *MockIClientMockRecorder) GetBlockchainHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainHeight", reflect.TypeOf((*MockIClient)(nil).GetBlockchainHeight), arg0)
}

// GetBlockchainInfo mocks base method.
func (m *MockIClient) GetBlockchainInfo(arg0 context.Context) (*pactus.GetBlockchainInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockchainInfo", arg0)
	ret0, _ := ret[0].(*pactus.GetBlockchainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainInfo indicates an expected call of GetBlockchainInfo.
func (mr *MockIClientMockRecorder) GetBlockchainInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainInfo", reflect.TypeOf((*MockIClient)(nil).GetBlockchainInfo), arg0)
}

// GetFee mocks base method.
func (m *MockIClient) GetFee(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFee", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFee indicates an expected call of GetFee.
func (mr *MockIClientMockRecorder) GetFee(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFee", reflect.TypeOf((*MockIClient)(nil).GetFee), arg0, arg1)
}

// GetNetworkInfo mocks base method.
func (m *MockIClient) GetNetworkInfo(arg0 context.Context) (*pactus.GetNetworkInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkInfo", arg0)
	ret0, _ := ret[0].(*pactus.GetNetworkInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkInfo indicates an expected call of GetNetworkInfo.
func (mr *MockIClientMockRecorder) GetNetworkInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkInfo", reflect.TypeOf((*MockIClient)(nil).GetNetworkInfo), arg0)
}

// GetTransactionData mocks base method.
func (m *MockIClient) GetTransactionData(arg0 context.Context, arg1 string) (*pactus.GetTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionData", arg0, arg1)
	ret0, _ := ret[0].(*pactus.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionData indicates an expected call of GetTransactionData.
func (mr *MockIClientMockRecorder) GetTransactionData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionData", reflect.TypeOf((*MockIClient)(nil).GetTransactionData), arg0, arg1)
}

// GetValidatorInfo mocks base method.
func (m *MockIClient) GetValidatorInfo(arg0 context.Context, arg1 string) (*pactus.GetValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorInfo", arg0, arg1)
	ret0, _ := ret[0].(*pactus.GetValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorInfo indicates an expected call of GetValidatorInfo.
func (mr *MockIClientMockRecorder) GetValidatorInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorInfo", reflect.TypeOf((*MockIClient)(nil).GetValidatorInfo), arg0, arg1)
}

// GetValidatorInfoByNumber mocks base method.
func (m *MockIClient) GetValidatorInfoByNumber(arg0 context.Context, arg1 int32) (*pactus.GetValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorInfoByNumber", arg0, arg1)
	ret0, _ := ret[0].(*pactus.GetValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorInfoByNumber indicates an expected call of GetValidatorInfoByNumber.
func (mr *MockIClientMockRecorder) GetValidatorInfoByNumber(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorInfoByNumber", reflect.TypeOf((*MockIClient)(nil).GetValidatorInfoByNumber), arg0, arg1)
}

// LastBlockTime mocks base method.
func (m *MockIClient) LastBlockTime(arg0 context.Context) (uint32, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBlockTime", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastBlockTime indicates an expected call of LastBlockTime.
func (mr *MockIClientMockRecorder) LastBlockTime(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockTime", reflect.TypeOf((*MockIClient)(nil).LastBlockTime), arg0)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddClient mocks base method.
func (m *MockManager) AddClient(c IClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddClient", c)
}

// AddClient indicates an expected call of AddClient.
func (mr *MockManagerMockRecorder) AddClient(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockManager)(nil).AddClient), c)
}

// GetBalance mocks base method.
func (m *MockManager) GetBalance(addr string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", addr)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockManagerMockRecorder) GetBalance(addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockManager)(nil).GetBalance), addr)
}

// GetBlockchainHeight mocks base method.
func (m *MockManager) GetBlockchainHeight() (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockchainHeight")
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainHeight indicates an expected call of GetBlockchainHeight.
func (mr *MockManagerMockRecorder) GetBlockchainHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainHeight", reflect.TypeOf((*MockManager)(nil).GetBlockchainHeight))
}

// GetBlockchainInfo mocks base method.
func (m *MockManager) GetBlockchainInfo() (*pactus.GetBlockchainInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockchainInfo")
	ret0, _ := ret[0].(*pactus.GetBlockchainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainInfo indicates an expected call of GetBlockchainInfo.
func (mr *MockManagerMockRecorder) GetBlockchainInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainInfo", reflect.TypeOf((*MockManager)(nil).GetBlockchainInfo))
}

// GetCirculatingSupply mocks base method.
func (m *MockManager) GetCirculatingSupply() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCirculatingSupply")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCirculatingSupply indicates an expected call of GetCirculatingSupply.
func (mr *MockManagerMockRecorder) GetCirculatingSupply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCirculatingSupply", reflect.TypeOf((*MockManager)(nil).GetCirculatingSupply))
}

// GetFee mocks base method.
func (m *MockManager) GetFee(amt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFee", amt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFee indicates an expected call of GetFee.
func (mr *MockManagerMockRecorder) GetFee(amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFee", reflect.TypeOf((*MockManager)(nil).GetFee), amt)
}

// GetLastBlockTime mocks base method.
func (m *MockManager) GetLastBlockTime() (uint32, uint32) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastBlockTime")
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(uint32)
	return ret0, ret1
}

// GetLastBlockTime indicates an expected call of GetLastBlockTime.
func (mr *MockManagerMockRecorder) GetLastBlockTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastBlockTime", reflect.TypeOf((*MockManager)(nil).GetLastBlockTime))
}

// GetNetworkInfo mocks base method.
func (m *MockManager) GetNetworkInfo() (*pactus.GetNetworkInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkInfo")
	ret0, _ := ret[0].(*pactus.GetNetworkInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkInfo indicates an expected call of GetNetworkInfo.
func (mr *MockManagerMockRecorder) GetNetworkInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkInfo", reflect.TypeOf((*MockManager)(nil).GetNetworkInfo))
}

// GetPeerInfo mocks base method.
func (m *MockManager) GetPeerInfo(address string) (*pactus.PeerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerInfo", address)
	ret0, _ := ret[0].(*pactus.PeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerInfo indicates an expected call of GetPeerInfo.
func (mr *MockManagerMockRecorder) GetPeerInfo(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerInfo", reflect.TypeOf((*MockManager)(nil).GetPeerInfo), address)
}

// GetRandomClient mocks base method.
func (m *MockManager) GetRandomClient() IClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomClient")
	ret0, _ := ret[0].(IClient)
	return ret0
}

// GetRandomClient indicates an expected call of GetRandomClient.
func (mr *MockManagerMockRecorder) GetRandomClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomClient", reflect.TypeOf((*MockManager)(nil).GetRandomClient))
}

// GetTransactionData mocks base method.
func (m *MockManager) GetTransactionData(txID string) (*pactus.GetTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionData", txID)
	ret0, _ := ret[0].(*pactus.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionData indicates an expected call of GetTransactionData.
func (mr *MockManagerMockRecorder) GetTransactionData(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionData", reflect.TypeOf((*MockManager)(nil).GetTransactionData), txID)
}

// GetValidatorInfo mocks base method.
func (m *MockManager) GetValidatorInfo(address string) (*pactus.GetValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorInfo", address)
	ret0, _ := ret[0].(*pactus.GetValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorInfo indicates an expected call of GetValidatorInfo.
func (mr *MockManagerMockRecorder) GetValidatorInfo(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorInfo", reflect.TypeOf((*MockManager)(nil).GetValidatorInfo), address)
}

// GetValidatorInfoByNumber mocks base method.
func (m *MockManager) GetValidatorInfoByNumber(num int32) (*pactus.GetValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorInfoByNumber", num)
	ret0, _ := ret[0].(*pactus.GetValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorInfoByNumber indicates an expected call of GetValidatorInfoByNumber.
func (mr *MockManagerMockRecorder) GetValidatorInfoByNumber(num any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorInfoByNumber", reflect.TypeOf((*MockManager)(nil).GetValidatorInfoByNumber), num)
}

// Start mocks base method.
func (m *MockManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop))
}

// getLocalClient mocks base method.
func (m *MockManager) getLocalClient() IClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getLocalClient")
	ret0, _ := ret[0].(IClient)
	return ret0
}

// getLocalClient indicates an expected call of getLocalClient.
func (mr *MockManagerMockRecorder) getLocalClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getLocalClient", reflect.TypeOf((*MockManager)(nil).getLocalClient))
}

// updateValMap mocks base method.
func (m *MockManager) updateValMap() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateValMap")
}

// updateValMap indicates an expected call of updateValMap.
func (mr *MockManagerMockRecorder) updateValMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateValMap", reflect.TypeOf((*MockManager)(nil).updateValMap))
}
