// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/client (interfaces: Backend)
//
// Generated by this command:
//
//	mockgen -typed -package client_test -destination package_mock_test.go github.com/juju/juju/apiserver/facades/client/client Backend
//

// Package client_test is a generated GoMock package.
package client_test

import (
	reflect "reflect"
	time "time"

	state "github.com/juju/juju/state"
	names "github.com/juju/names/v6"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// AllIPAddresses mocks base method.
func (m *MockBackend) AllIPAddresses() ([]*state.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllIPAddresses")
	ret0, _ := ret[0].([]*state.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllIPAddresses indicates an expected call of AllIPAddresses.
func (mr *MockBackendMockRecorder) AllIPAddresses() *MockBackendAllIPAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllIPAddresses", reflect.TypeOf((*MockBackend)(nil).AllIPAddresses))
	return &MockBackendAllIPAddressesCall{Call: call}
}

// MockBackendAllIPAddressesCall wrap *gomock.Call
type MockBackendAllIPAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendAllIPAddressesCall) Return(arg0 []*state.Address, arg1 error) *MockBackendAllIPAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendAllIPAddressesCall) Do(f func() ([]*state.Address, error)) *MockBackendAllIPAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendAllIPAddressesCall) DoAndReturn(f func() ([]*state.Address, error)) *MockBackendAllIPAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllLinkLayerDevices mocks base method.
func (m *MockBackend) AllLinkLayerDevices() ([]*state.LinkLayerDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllLinkLayerDevices")
	ret0, _ := ret[0].([]*state.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLinkLayerDevices indicates an expected call of AllLinkLayerDevices.
func (mr *MockBackendMockRecorder) AllLinkLayerDevices() *MockBackendAllLinkLayerDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLinkLayerDevices", reflect.TypeOf((*MockBackend)(nil).AllLinkLayerDevices))
	return &MockBackendAllLinkLayerDevicesCall{Call: call}
}

// MockBackendAllLinkLayerDevicesCall wrap *gomock.Call
type MockBackendAllLinkLayerDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendAllLinkLayerDevicesCall) Return(arg0 []*state.LinkLayerDevice, arg1 error) *MockBackendAllLinkLayerDevicesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendAllLinkLayerDevicesCall) Do(f func() ([]*state.LinkLayerDevice, error)) *MockBackendAllLinkLayerDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendAllLinkLayerDevicesCall) DoAndReturn(f func() ([]*state.LinkLayerDevice, error)) *MockBackendAllLinkLayerDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllMachines mocks base method.
func (m *MockBackend) AllMachines() ([]*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachines")
	ret0, _ := ret[0].([]*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachines indicates an expected call of AllMachines.
func (mr *MockBackendMockRecorder) AllMachines() *MockBackendAllMachinesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachines", reflect.TypeOf((*MockBackend)(nil).AllMachines))
	return &MockBackendAllMachinesCall{Call: call}
}

// MockBackendAllMachinesCall wrap *gomock.Call
type MockBackendAllMachinesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendAllMachinesCall) Return(arg0 []*state.Machine, arg1 error) *MockBackendAllMachinesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendAllMachinesCall) Do(f func() ([]*state.Machine, error)) *MockBackendAllMachinesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendAllMachinesCall) DoAndReturn(f func() ([]*state.Machine, error)) *MockBackendAllMachinesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllStatus mocks base method.
func (m *MockBackend) AllStatus() (*state.AllStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllStatus")
	ret0, _ := ret[0].(*state.AllStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllStatus indicates an expected call of AllStatus.
func (mr *MockBackendMockRecorder) AllStatus() *MockBackendAllStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllStatus", reflect.TypeOf((*MockBackend)(nil).AllStatus))
	return &MockBackendAllStatusCall{Call: call}
}

// MockBackendAllStatusCall wrap *gomock.Call
type MockBackendAllStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendAllStatusCall) Return(arg0 *state.AllStatus, arg1 error) *MockBackendAllStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendAllStatusCall) Do(f func() (*state.AllStatus, error)) *MockBackendAllStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendAllStatusCall) DoAndReturn(f func() (*state.AllStatus, error)) *MockBackendAllStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerNodes mocks base method.
func (m *MockBackend) ControllerNodes() ([]state.ControllerNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNodes")
	ret0, _ := ret[0].([]state.ControllerNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerNodes indicates an expected call of ControllerNodes.
func (mr *MockBackendMockRecorder) ControllerNodes() *MockBackendControllerNodesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNodes", reflect.TypeOf((*MockBackend)(nil).ControllerNodes))
	return &MockBackendControllerNodesCall{Call: call}
}

// MockBackendControllerNodesCall wrap *gomock.Call
type MockBackendControllerNodesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendControllerNodesCall) Return(arg0 []state.ControllerNode, arg1 error) *MockBackendControllerNodesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendControllerNodesCall) Do(f func() ([]state.ControllerNode, error)) *MockBackendControllerNodesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendControllerNodesCall) DoAndReturn(f func() ([]state.ControllerNode, error)) *MockBackendControllerNodesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerTimestamp mocks base method.
func (m *MockBackend) ControllerTimestamp() (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTimestamp")
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerTimestamp indicates an expected call of ControllerTimestamp.
func (mr *MockBackendMockRecorder) ControllerTimestamp() *MockBackendControllerTimestampCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTimestamp", reflect.TypeOf((*MockBackend)(nil).ControllerTimestamp))
	return &MockBackendControllerTimestampCall{Call: call}
}

// MockBackendControllerTimestampCall wrap *gomock.Call
type MockBackendControllerTimestampCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendControllerTimestampCall) Return(arg0 *time.Time, arg1 error) *MockBackendControllerTimestampCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendControllerTimestampCall) Do(f func() (*time.Time, error)) *MockBackendControllerTimestampCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendControllerTimestampCall) DoAndReturn(f func() (*time.Time, error)) *MockBackendControllerTimestampCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HAPrimaryMachine mocks base method.
func (m *MockBackend) HAPrimaryMachine() (names.MachineTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HAPrimaryMachine")
	ret0, _ := ret[0].(names.MachineTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HAPrimaryMachine indicates an expected call of HAPrimaryMachine.
func (mr *MockBackendMockRecorder) HAPrimaryMachine() *MockBackendHAPrimaryMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HAPrimaryMachine", reflect.TypeOf((*MockBackend)(nil).HAPrimaryMachine))
	return &MockBackendHAPrimaryMachineCall{Call: call}
}

// MockBackendHAPrimaryMachineCall wrap *gomock.Call
type MockBackendHAPrimaryMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendHAPrimaryMachineCall) Return(arg0 names.MachineTag, arg1 error) *MockBackendHAPrimaryMachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendHAPrimaryMachineCall) Do(f func() (names.MachineTag, error)) *MockBackendHAPrimaryMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendHAPrimaryMachineCall) DoAndReturn(f func() (names.MachineTag, error)) *MockBackendHAPrimaryMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MachineConstraints mocks base method.
func (m *MockBackend) MachineConstraints() (*state.MachineConstraints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineConstraints")
	ret0, _ := ret[0].(*state.MachineConstraints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineConstraints indicates an expected call of MachineConstraints.
func (mr *MockBackendMockRecorder) MachineConstraints() *MockBackendMachineConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineConstraints", reflect.TypeOf((*MockBackend)(nil).MachineConstraints))
	return &MockBackendMachineConstraintsCall{Call: call}
}

// MockBackendMachineConstraintsCall wrap *gomock.Call
type MockBackendMachineConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendMachineConstraintsCall) Return(arg0 *state.MachineConstraints, arg1 error) *MockBackendMachineConstraintsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendMachineConstraintsCall) Do(f func() (*state.MachineConstraints, error)) *MockBackendMachineConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendMachineConstraintsCall) DoAndReturn(f func() (*state.MachineConstraints, error)) *MockBackendMachineConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
