// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/lease (interfaces: LeaseManager)
//
// Generated by this command:
//
//	mockgen -typed -package status_test -destination lease_mock_test.go github.com/juju/juju/core/lease LeaseManager
//

// Package status_test is a generated GoMock package.
package status_test

import (
	context "context"
	reflect "reflect"

	lease "github.com/juju/juju/core/lease"
	gomock "go.uber.org/mock/gomock"
)

// MockLeaseManager is a mock of LeaseManager interface.
type MockLeaseManager struct {
	ctrl     *gomock.Controller
	recorder *MockLeaseManagerMockRecorder
}

// MockLeaseManagerMockRecorder is the mock recorder for MockLeaseManager.
type MockLeaseManagerMockRecorder struct {
	mock *MockLeaseManager
}

// NewMockLeaseManager creates a new mock instance.
func NewMockLeaseManager(ctrl *gomock.Controller) *MockLeaseManager {
	mock := &MockLeaseManager{ctrl: ctrl}
	mock.recorder = &MockLeaseManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaseManager) EXPECT() *MockLeaseManagerMockRecorder {
	return m.recorder
}

// Revoke mocks base method.
func (m *MockLeaseManager) Revoke(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke.
func (mr *MockLeaseManagerMockRecorder) Revoke(arg0, arg1 any) *MockLeaseManagerRevokeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockLeaseManager)(nil).Revoke), arg0, arg1)
	return &MockLeaseManagerRevokeCall{Call: call}
}

// MockLeaseManagerRevokeCall wrap *gomock.Call
type MockLeaseManagerRevokeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseManagerRevokeCall) Return(arg0 error) *MockLeaseManagerRevokeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseManagerRevokeCall) Do(f func(string, string) error) *MockLeaseManagerRevokeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseManagerRevokeCall) DoAndReturn(f func(string, string) error) *MockLeaseManagerRevokeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Token mocks base method.
func (m *MockLeaseManager) Token(arg0, arg1 string) lease.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", arg0, arg1)
	ret0, _ := ret[0].(lease.Token)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockLeaseManagerMockRecorder) Token(arg0, arg1 any) *MockLeaseManagerTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockLeaseManager)(nil).Token), arg0, arg1)
	return &MockLeaseManagerTokenCall{Call: call}
}

// MockLeaseManagerTokenCall wrap *gomock.Call
type MockLeaseManagerTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseManagerTokenCall) Return(arg0 lease.Token) *MockLeaseManagerTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseManagerTokenCall) Do(f func(string, string) lease.Token) *MockLeaseManagerTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseManagerTokenCall) DoAndReturn(f func(string, string) lease.Token) *MockLeaseManagerTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WaitUntilExpired mocks base method.
func (m *MockLeaseManager) WaitUntilExpired(arg0 context.Context, arg1 string, arg2 chan<- struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilExpired", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilExpired indicates an expected call of WaitUntilExpired.
func (mr *MockLeaseManagerMockRecorder) WaitUntilExpired(arg0, arg1, arg2 any) *MockLeaseManagerWaitUntilExpiredCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilExpired", reflect.TypeOf((*MockLeaseManager)(nil).WaitUntilExpired), arg0, arg1, arg2)
	return &MockLeaseManagerWaitUntilExpiredCall{Call: call}
}

// MockLeaseManagerWaitUntilExpiredCall wrap *gomock.Call
type MockLeaseManagerWaitUntilExpiredCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseManagerWaitUntilExpiredCall) Return(arg0 error) *MockLeaseManagerWaitUntilExpiredCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseManagerWaitUntilExpiredCall) Do(f func(context.Context, string, chan<- struct{}) error) *MockLeaseManagerWaitUntilExpiredCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseManagerWaitUntilExpiredCall) DoAndReturn(f func(context.Context, string, chan<- struct{}) error) *MockLeaseManagerWaitUntilExpiredCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
