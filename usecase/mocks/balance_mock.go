// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nawa/cryptoexchange-dashboard/usecase (interfaces: BalanceUsecases)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nawa/cryptoexchange-dashboard/model"
)

// MockBalanceUsecases is a mock of BalanceUsecases interface
type MockBalanceUsecases struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceUsecasesMockRecorder
}

// MockBalanceUsecasesMockRecorder is the mock recorder for MockBalanceUsecases
type MockBalanceUsecasesMockRecorder struct {
	mock *MockBalanceUsecases
}

// NewMockBalanceUsecases creates a new mock instance
func NewMockBalanceUsecases(ctrl *gomock.Controller) *MockBalanceUsecases {
	mock := &MockBalanceUsecases{ctrl: ctrl}
	mock.recorder = &MockBalanceUsecasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBalanceUsecases) EXPECT() *MockBalanceUsecasesMockRecorder {
	return m.recorder
}

// FetchAll mocks base method
func (m *MockBalanceUsecases) FetchAll(arg0 string) ([]model.CurrencyBalance, error) {
	ret := m.ctrl.Call(m, "FetchAll", arg0)
	ret0, _ := ret[0].([]model.CurrencyBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAll indicates an expected call of FetchAll
func (mr *MockBalanceUsecasesMockRecorder) FetchAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAll", reflect.TypeOf((*MockBalanceUsecases)(nil).FetchAll), arg0)
}

// FetchHourly mocks base method
func (m *MockBalanceUsecases) FetchHourly(arg0 string, arg1 int) ([]model.CurrencyBalance, error) {
	ret := m.ctrl.Call(m, "FetchHourly", arg0, arg1)
	ret0, _ := ret[0].([]model.CurrencyBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHourly indicates an expected call of FetchHourly
func (mr *MockBalanceUsecasesMockRecorder) FetchHourly(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHourly", reflect.TypeOf((*MockBalanceUsecases)(nil).FetchHourly), arg0, arg1)
}

// FetchMonthly mocks base method
func (m *MockBalanceUsecases) FetchMonthly(arg0 string) ([]model.CurrencyBalance, error) {
	ret := m.ctrl.Call(m, "FetchMonthly", arg0)
	ret0, _ := ret[0].([]model.CurrencyBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMonthly indicates an expected call of FetchMonthly
func (mr *MockBalanceUsecasesMockRecorder) FetchMonthly(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMonthly", reflect.TypeOf((*MockBalanceUsecases)(nil).FetchMonthly), arg0)
}

// FetchWeekly mocks base method
func (m *MockBalanceUsecases) FetchWeekly(arg0 string) ([]model.CurrencyBalance, error) {
	ret := m.ctrl.Call(m, "FetchWeekly", arg0)
	ret0, _ := ret[0].([]model.CurrencyBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchWeekly indicates an expected call of FetchWeekly
func (mr *MockBalanceUsecasesMockRecorder) FetchWeekly(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWeekly", reflect.TypeOf((*MockBalanceUsecases)(nil).FetchWeekly), arg0)
}

// GetActiveCurrencies mocks base method
func (m *MockBalanceUsecases) GetActiveCurrencies() ([]model.CurrencyBalance, error) {
	ret := m.ctrl.Call(m, "GetActiveCurrencies")
	ret0, _ := ret[0].([]model.CurrencyBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveCurrencies indicates an expected call of GetActiveCurrencies
func (mr *MockBalanceUsecasesMockRecorder) GetActiveCurrencies() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveCurrencies", reflect.TypeOf((*MockBalanceUsecases)(nil).GetActiveCurrencies))
}

// StartSyncFromExchangePeriodically mocks base method
func (m *MockBalanceUsecases) StartSyncFromExchangePeriodically(arg0 time.Duration) (func(), error) {
	ret := m.ctrl.Call(m, "StartSyncFromExchangePeriodically", arg0)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSyncFromExchangePeriodically indicates an expected call of StartSyncFromExchangePeriodically
func (mr *MockBalanceUsecasesMockRecorder) StartSyncFromExchangePeriodically(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSyncFromExchangePeriodically", reflect.TypeOf((*MockBalanceUsecases)(nil).StartSyncFromExchangePeriodically), arg0)
}

// SyncFromExchange mocks base method
func (m *MockBalanceUsecases) SyncFromExchange() error {
	ret := m.ctrl.Call(m, "SyncFromExchange")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncFromExchange indicates an expected call of SyncFromExchange
func (mr *MockBalanceUsecasesMockRecorder) SyncFromExchange() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncFromExchange", reflect.TypeOf((*MockBalanceUsecases)(nil).SyncFromExchange))
}