// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/provider/kubernetes/exec (interfaces: Executor)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/k8s_exec_mock.go github.com/juju/juju/internal/provider/kubernetes/exec Executor
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	exec "github.com/juju/juju/internal/provider/kubernetes/exec"
	gomock "go.uber.org/mock/gomock"
	kubernetes "k8s.io/client-go/kubernetes"
)

// MockExecutor is a mock of Executor interface.
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor.
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance.
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Copy mocks base method.
func (m *MockExecutor) Copy(arg0 context.Context, arg1 exec.CopyParams, arg2 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockExecutorMockRecorder) Copy(arg0, arg1, arg2 any) *MockExecutorCopyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockExecutor)(nil).Copy), arg0, arg1, arg2)
	return &MockExecutorCopyCall{Call: call}
}

// MockExecutorCopyCall wrap *gomock.Call
type MockExecutorCopyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorCopyCall) Return(arg0 error) *MockExecutorCopyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorCopyCall) Do(f func(context.Context, exec.CopyParams, <-chan struct{}) error) *MockExecutorCopyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorCopyCall) DoAndReturn(f func(context.Context, exec.CopyParams, <-chan struct{}) error) *MockExecutorCopyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Exec mocks base method.
func (m *MockExecutor) Exec(arg0 context.Context, arg1 exec.ExecParams, arg2 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockExecutorMockRecorder) Exec(arg0, arg1, arg2 any) *MockExecutorExecCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockExecutor)(nil).Exec), arg0, arg1, arg2)
	return &MockExecutorExecCall{Call: call}
}

// MockExecutorExecCall wrap *gomock.Call
type MockExecutorExecCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorExecCall) Return(arg0 error) *MockExecutorExecCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorExecCall) Do(f func(context.Context, exec.ExecParams, <-chan struct{}) error) *MockExecutorExecCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorExecCall) DoAndReturn(f func(context.Context, exec.ExecParams, <-chan struct{}) error) *MockExecutorExecCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NameSpace mocks base method.
func (m *MockExecutor) NameSpace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NameSpace")
	ret0, _ := ret[0].(string)
	return ret0
}

// NameSpace indicates an expected call of NameSpace.
func (mr *MockExecutorMockRecorder) NameSpace() *MockExecutorNameSpaceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NameSpace", reflect.TypeOf((*MockExecutor)(nil).NameSpace))
	return &MockExecutorNameSpaceCall{Call: call}
}

// MockExecutorNameSpaceCall wrap *gomock.Call
type MockExecutorNameSpaceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorNameSpaceCall) Return(arg0 string) *MockExecutorNameSpaceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorNameSpaceCall) Do(f func() string) *MockExecutorNameSpaceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorNameSpaceCall) DoAndReturn(f func() string) *MockExecutorNameSpaceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RawClient mocks base method.
func (m *MockExecutor) RawClient() kubernetes.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(kubernetes.Interface)
	return ret0
}

// RawClient indicates an expected call of RawClient.
func (mr *MockExecutorMockRecorder) RawClient() *MockExecutorRawClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockExecutor)(nil).RawClient))
	return &MockExecutorRawClientCall{Call: call}
}

// MockExecutorRawClientCall wrap *gomock.Call
type MockExecutorRawClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorRawClientCall) Return(arg0 kubernetes.Interface) *MockExecutorRawClientCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorRawClientCall) Do(f func() kubernetes.Interface) *MockExecutorRawClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorRawClientCall) DoAndReturn(f func() kubernetes.Interface) *MockExecutorRawClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Status mocks base method.
func (m *MockExecutor) Status(arg0 context.Context, arg1 exec.StatusParams) (*exec.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*exec.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockExecutorMockRecorder) Status(arg0, arg1 any) *MockExecutorStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockExecutor)(nil).Status), arg0, arg1)
	return &MockExecutorStatusCall{Call: call}
}

// MockExecutorStatusCall wrap *gomock.Call
type MockExecutorStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutorStatusCall) Return(arg0 *exec.Status, arg1 error) *MockExecutorStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutorStatusCall) Do(f func(context.Context, exec.StatusParams) (*exec.Status, error)) *MockExecutorStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutorStatusCall) DoAndReturn(f func(context.Context, exec.StatusParams) (*exec.Status, error)) *MockExecutorStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
