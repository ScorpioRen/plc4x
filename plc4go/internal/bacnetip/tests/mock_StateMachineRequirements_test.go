/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package tests

import (
	bacnetip "github.com/apache/plc4x/plc4go/internal/bacnetip"
	mock "github.com/stretchr/testify/mock"
)

// MockStateMachineRequirements is an autogenerated mock type for the StateMachineRequirements type
type MockStateMachineRequirements struct {
	mock.Mock
}

type MockStateMachineRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStateMachineRequirements) EXPECT() *MockStateMachineRequirements_Expecter {
	return &MockStateMachineRequirements_Expecter{mock: &_m.Mock}
}

// AfterReceive provides a mock function with given fields: pdu
func (_m *MockStateMachineRequirements) AfterReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineRequirements_AfterReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterReceive'
type MockStateMachineRequirements_AfterReceive_Call struct {
	*mock.Call
}

// AfterReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineRequirements_Expecter) AfterReceive(pdu interface{}) *MockStateMachineRequirements_AfterReceive_Call {
	return &MockStateMachineRequirements_AfterReceive_Call{Call: _e.mock.On("AfterReceive", pdu)}
}

func (_c *MockStateMachineRequirements_AfterReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineRequirements_AfterReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineRequirements_AfterReceive_Call) Return() *MockStateMachineRequirements_AfterReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_AfterReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineRequirements_AfterReceive_Call {
	_c.Call.Return(run)
	return _c
}

// AfterSend provides a mock function with given fields: pdu
func (_m *MockStateMachineRequirements) AfterSend(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineRequirements_AfterSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AfterSend'
type MockStateMachineRequirements_AfterSend_Call struct {
	*mock.Call
}

// AfterSend is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineRequirements_Expecter) AfterSend(pdu interface{}) *MockStateMachineRequirements_AfterSend_Call {
	return &MockStateMachineRequirements_AfterSend_Call{Call: _e.mock.On("AfterSend", pdu)}
}

func (_c *MockStateMachineRequirements_AfterSend_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineRequirements_AfterSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineRequirements_AfterSend_Call) Return() *MockStateMachineRequirements_AfterSend_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_AfterSend_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineRequirements_AfterSend_Call {
	_c.Call.Return(run)
	return _c
}

// BeforeReceive provides a mock function with given fields: pdu
func (_m *MockStateMachineRequirements) BeforeReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineRequirements_BeforeReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeforeReceive'
type MockStateMachineRequirements_BeforeReceive_Call struct {
	*mock.Call
}

// BeforeReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineRequirements_Expecter) BeforeReceive(pdu interface{}) *MockStateMachineRequirements_BeforeReceive_Call {
	return &MockStateMachineRequirements_BeforeReceive_Call{Call: _e.mock.On("BeforeReceive", pdu)}
}

func (_c *MockStateMachineRequirements_BeforeReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineRequirements_BeforeReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineRequirements_BeforeReceive_Call) Return() *MockStateMachineRequirements_BeforeReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_BeforeReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineRequirements_BeforeReceive_Call {
	_c.Call.Return(run)
	return _c
}

// BeforeSend provides a mock function with given fields: pdu
func (_m *MockStateMachineRequirements) BeforeSend(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineRequirements_BeforeSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeforeSend'
type MockStateMachineRequirements_BeforeSend_Call struct {
	*mock.Call
}

// BeforeSend is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineRequirements_Expecter) BeforeSend(pdu interface{}) *MockStateMachineRequirements_BeforeSend_Call {
	return &MockStateMachineRequirements_BeforeSend_Call{Call: _e.mock.On("BeforeSend", pdu)}
}

func (_c *MockStateMachineRequirements_BeforeSend_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineRequirements_BeforeSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineRequirements_BeforeSend_Call) Return() *MockStateMachineRequirements_BeforeSend_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_BeforeSend_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineRequirements_BeforeSend_Call {
	_c.Call.Return(run)
	return _c
}

// EventSet provides a mock function with given fields: id
func (_m *MockStateMachineRequirements) EventSet(id string) {
	_m.Called(id)
}

// MockStateMachineRequirements_EventSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventSet'
type MockStateMachineRequirements_EventSet_Call struct {
	*mock.Call
}

// EventSet is a helper method to define mock.On call
//   - id string
func (_e *MockStateMachineRequirements_Expecter) EventSet(id interface{}) *MockStateMachineRequirements_EventSet_Call {
	return &MockStateMachineRequirements_EventSet_Call{Call: _e.mock.On("EventSet", id)}
}

func (_c *MockStateMachineRequirements_EventSet_Call) Run(run func(id string)) *MockStateMachineRequirements_EventSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStateMachineRequirements_EventSet_Call) Return() *MockStateMachineRequirements_EventSet_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_EventSet_Call) RunAndReturn(run func(string)) *MockStateMachineRequirements_EventSet_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrentState provides a mock function with given fields:
func (_m *MockStateMachineRequirements) GetCurrentState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineRequirements_GetCurrentState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentState'
type MockStateMachineRequirements_GetCurrentState_Call struct {
	*mock.Call
}

// GetCurrentState is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) GetCurrentState() *MockStateMachineRequirements_GetCurrentState_Call {
	return &MockStateMachineRequirements_GetCurrentState_Call{Call: _e.mock.On("GetCurrentState")}
}

func (_c *MockStateMachineRequirements_GetCurrentState_Call) Run(run func()) *MockStateMachineRequirements_GetCurrentState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_GetCurrentState_Call) Return(_a0 State) *MockStateMachineRequirements_GetCurrentState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_GetCurrentState_Call) RunAndReturn(run func() State) *MockStateMachineRequirements_GetCurrentState_Call {
	_c.Call.Return(run)
	return _c
}

// GetStartState provides a mock function with given fields:
func (_m *MockStateMachineRequirements) GetStartState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStartState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineRequirements_GetStartState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStartState'
type MockStateMachineRequirements_GetStartState_Call struct {
	*mock.Call
}

// GetStartState is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) GetStartState() *MockStateMachineRequirements_GetStartState_Call {
	return &MockStateMachineRequirements_GetStartState_Call{Call: _e.mock.On("GetStartState")}
}

func (_c *MockStateMachineRequirements_GetStartState_Call) Run(run func()) *MockStateMachineRequirements_GetStartState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_GetStartState_Call) Return(_a0 State) *MockStateMachineRequirements_GetStartState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_GetStartState_Call) RunAndReturn(run func() State) *MockStateMachineRequirements_GetStartState_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionLog provides a mock function with given fields:
func (_m *MockStateMachineRequirements) GetTransactionLog() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionLog")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockStateMachineRequirements_GetTransactionLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionLog'
type MockStateMachineRequirements_GetTransactionLog_Call struct {
	*mock.Call
}

// GetTransactionLog is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) GetTransactionLog() *MockStateMachineRequirements_GetTransactionLog_Call {
	return &MockStateMachineRequirements_GetTransactionLog_Call{Call: _e.mock.On("GetTransactionLog")}
}

func (_c *MockStateMachineRequirements_GetTransactionLog_Call) Run(run func()) *MockStateMachineRequirements_GetTransactionLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_GetTransactionLog_Call) Return(_a0 []string) *MockStateMachineRequirements_GetTransactionLog_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_GetTransactionLog_Call) RunAndReturn(run func() []string) *MockStateMachineRequirements_GetTransactionLog_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnexpectedReceiveState provides a mock function with given fields:
func (_m *MockStateMachineRequirements) GetUnexpectedReceiveState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUnexpectedReceiveState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineRequirements_GetUnexpectedReceiveState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnexpectedReceiveState'
type MockStateMachineRequirements_GetUnexpectedReceiveState_Call struct {
	*mock.Call
}

// GetUnexpectedReceiveState is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) GetUnexpectedReceiveState() *MockStateMachineRequirements_GetUnexpectedReceiveState_Call {
	return &MockStateMachineRequirements_GetUnexpectedReceiveState_Call{Call: _e.mock.On("GetUnexpectedReceiveState")}
}

func (_c *MockStateMachineRequirements_GetUnexpectedReceiveState_Call) Run(run func()) *MockStateMachineRequirements_GetUnexpectedReceiveState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_GetUnexpectedReceiveState_Call) Return(_a0 State) *MockStateMachineRequirements_GetUnexpectedReceiveState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_GetUnexpectedReceiveState_Call) RunAndReturn(run func() State) *MockStateMachineRequirements_GetUnexpectedReceiveState_Call {
	_c.Call.Return(run)
	return _c
}

// IsFailState provides a mock function with given fields:
func (_m *MockStateMachineRequirements) IsFailState() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsFailState")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStateMachineRequirements_IsFailState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsFailState'
type MockStateMachineRequirements_IsFailState_Call struct {
	*mock.Call
}

// IsFailState is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) IsFailState() *MockStateMachineRequirements_IsFailState_Call {
	return &MockStateMachineRequirements_IsFailState_Call{Call: _e.mock.On("IsFailState")}
}

func (_c *MockStateMachineRequirements_IsFailState_Call) Run(run func()) *MockStateMachineRequirements_IsFailState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_IsFailState_Call) Return(_a0 bool) *MockStateMachineRequirements_IsFailState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_IsFailState_Call) RunAndReturn(run func() bool) *MockStateMachineRequirements_IsFailState_Call {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with given fields:
func (_m *MockStateMachineRequirements) IsRunning() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRunning")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStateMachineRequirements_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type MockStateMachineRequirements_IsRunning_Call struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) IsRunning() *MockStateMachineRequirements_IsRunning_Call {
	return &MockStateMachineRequirements_IsRunning_Call{Call: _e.mock.On("IsRunning")}
}

func (_c *MockStateMachineRequirements_IsRunning_Call) Run(run func()) *MockStateMachineRequirements_IsRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_IsRunning_Call) Return(_a0 bool) *MockStateMachineRequirements_IsRunning_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_IsRunning_Call) RunAndReturn(run func() bool) *MockStateMachineRequirements_IsRunning_Call {
	_c.Call.Return(run)
	return _c
}

// IsSuccessState provides a mock function with given fields:
func (_m *MockStateMachineRequirements) IsSuccessState() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSuccessState")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockStateMachineRequirements_IsSuccessState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSuccessState'
type MockStateMachineRequirements_IsSuccessState_Call struct {
	*mock.Call
}

// IsSuccessState is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) IsSuccessState() *MockStateMachineRequirements_IsSuccessState_Call {
	return &MockStateMachineRequirements_IsSuccessState_Call{Call: _e.mock.On("IsSuccessState")}
}

func (_c *MockStateMachineRequirements_IsSuccessState_Call) Run(run func()) *MockStateMachineRequirements_IsSuccessState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_IsSuccessState_Call) Return(_a0 bool) *MockStateMachineRequirements_IsSuccessState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_IsSuccessState_Call) RunAndReturn(run func() bool) *MockStateMachineRequirements_IsSuccessState_Call {
	_c.Call.Return(run)
	return _c
}

// NewState provides a mock function with given fields: _a0
func (_m *MockStateMachineRequirements) NewState(_a0 string) State {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for NewState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func(string) State); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineRequirements_NewState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewState'
type MockStateMachineRequirements_NewState_Call struct {
	*mock.Call
}

// NewState is a helper method to define mock.On call
//   - _a0 string
func (_e *MockStateMachineRequirements_Expecter) NewState(_a0 interface{}) *MockStateMachineRequirements_NewState_Call {
	return &MockStateMachineRequirements_NewState_Call{Call: _e.mock.On("NewState", _a0)}
}

func (_c *MockStateMachineRequirements_NewState_Call) Run(run func(_a0 string)) *MockStateMachineRequirements_NewState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStateMachineRequirements_NewState_Call) Return(_a0 State) *MockStateMachineRequirements_NewState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_NewState_Call) RunAndReturn(run func(string) State) *MockStateMachineRequirements_NewState_Call {
	_c.Call.Return(run)
	return _c
}

// Receive provides a mock function with given fields: args, kwargs
func (_m *MockStateMachineRequirements) Receive(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Receive")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStateMachineRequirements_Receive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Receive'
type MockStateMachineRequirements_Receive_Call struct {
	*mock.Call
}

// Receive is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockStateMachineRequirements_Expecter) Receive(args interface{}, kwargs interface{}) *MockStateMachineRequirements_Receive_Call {
	return &MockStateMachineRequirements_Receive_Call{Call: _e.mock.On("Receive", args, kwargs)}
}

func (_c *MockStateMachineRequirements_Receive_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockStateMachineRequirements_Receive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockStateMachineRequirements_Receive_Call) Return(_a0 error) *MockStateMachineRequirements_Receive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_Receive_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockStateMachineRequirements_Receive_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields:
func (_m *MockStateMachineRequirements) Reset() {
	_m.Called()
}

// MockStateMachineRequirements_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type MockStateMachineRequirements_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) Reset() *MockStateMachineRequirements_Reset_Call {
	return &MockStateMachineRequirements_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *MockStateMachineRequirements_Reset_Call) Run(run func()) *MockStateMachineRequirements_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_Reset_Call) Return() *MockStateMachineRequirements_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_Reset_Call) RunAndReturn(run func()) *MockStateMachineRequirements_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields:
func (_m *MockStateMachineRequirements) Run() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStateMachineRequirements_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockStateMachineRequirements_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) Run() *MockStateMachineRequirements_Run_Call {
	return &MockStateMachineRequirements_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *MockStateMachineRequirements_Run_Call) Run(run func()) *MockStateMachineRequirements_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_Run_Call) Return(_a0 error) *MockStateMachineRequirements_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_Run_Call) RunAndReturn(run func() error) *MockStateMachineRequirements_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: args, kwargs
func (_m *MockStateMachineRequirements) Send(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStateMachineRequirements_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockStateMachineRequirements_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockStateMachineRequirements_Expecter) Send(args interface{}, kwargs interface{}) *MockStateMachineRequirements_Send_Call {
	return &MockStateMachineRequirements_Send_Call{Call: _e.mock.On("Send", args, kwargs)}
}

func (_c *MockStateMachineRequirements_Send_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockStateMachineRequirements_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockStateMachineRequirements_Send_Call) Return(_a0 error) *MockStateMachineRequirements_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_Send_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockStateMachineRequirements_Send_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockStateMachineRequirements) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockStateMachineRequirements_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockStateMachineRequirements_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) String() *MockStateMachineRequirements_String_Call {
	return &MockStateMachineRequirements_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockStateMachineRequirements_String_Call) Run(run func()) *MockStateMachineRequirements_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_String_Call) Return(_a0 string) *MockStateMachineRequirements_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_String_Call) RunAndReturn(run func() string) *MockStateMachineRequirements_String_Call {
	_c.Call.Return(run)
	return _c
}

// UnexpectedReceive provides a mock function with given fields: pdu
func (_m *MockStateMachineRequirements) UnexpectedReceive(pdu bacnetip.PDU) {
	_m.Called(pdu)
}

// MockStateMachineRequirements_UnexpectedReceive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnexpectedReceive'
type MockStateMachineRequirements_UnexpectedReceive_Call struct {
	*mock.Call
}

// UnexpectedReceive is a helper method to define mock.On call
//   - pdu bacnetip.PDU
func (_e *MockStateMachineRequirements_Expecter) UnexpectedReceive(pdu interface{}) *MockStateMachineRequirements_UnexpectedReceive_Call {
	return &MockStateMachineRequirements_UnexpectedReceive_Call{Call: _e.mock.On("UnexpectedReceive", pdu)}
}

func (_c *MockStateMachineRequirements_UnexpectedReceive_Call) Run(run func(pdu bacnetip.PDU)) *MockStateMachineRequirements_UnexpectedReceive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.PDU))
	})
	return _c
}

func (_c *MockStateMachineRequirements_UnexpectedReceive_Call) Return() *MockStateMachineRequirements_UnexpectedReceive_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_UnexpectedReceive_Call) RunAndReturn(run func(bacnetip.PDU)) *MockStateMachineRequirements_UnexpectedReceive_Call {
	_c.Call.Return(run)
	return _c
}

// getCurrentState provides a mock function with given fields:
func (_m *MockStateMachineRequirements) getCurrentState() State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getCurrentState")
	}

	var r0 State
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(State)
		}
	}

	return r0
}

// MockStateMachineRequirements_getCurrentState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getCurrentState'
type MockStateMachineRequirements_getCurrentState_Call struct {
	*mock.Call
}

// getCurrentState is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) getCurrentState() *MockStateMachineRequirements_getCurrentState_Call {
	return &MockStateMachineRequirements_getCurrentState_Call{Call: _e.mock.On("getCurrentState")}
}

func (_c *MockStateMachineRequirements_getCurrentState_Call) Run(run func()) *MockStateMachineRequirements_getCurrentState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_getCurrentState_Call) Return(_a0 State) *MockStateMachineRequirements_getCurrentState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_getCurrentState_Call) RunAndReturn(run func() State) *MockStateMachineRequirements_getCurrentState_Call {
	_c.Call.Return(run)
	return _c
}

// getMachineGroup provides a mock function with given fields:
func (_m *MockStateMachineRequirements) getMachineGroup() *StateMachineGroup {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getMachineGroup")
	}

	var r0 *StateMachineGroup
	if rf, ok := ret.Get(0).(func() *StateMachineGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StateMachineGroup)
		}
	}

	return r0
}

// MockStateMachineRequirements_getMachineGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getMachineGroup'
type MockStateMachineRequirements_getMachineGroup_Call struct {
	*mock.Call
}

// getMachineGroup is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) getMachineGroup() *MockStateMachineRequirements_getMachineGroup_Call {
	return &MockStateMachineRequirements_getMachineGroup_Call{Call: _e.mock.On("getMachineGroup")}
}

func (_c *MockStateMachineRequirements_getMachineGroup_Call) Run(run func()) *MockStateMachineRequirements_getMachineGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_getMachineGroup_Call) Return(_a0 *StateMachineGroup) *MockStateMachineRequirements_getMachineGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_getMachineGroup_Call) RunAndReturn(run func() *StateMachineGroup) *MockStateMachineRequirements_getMachineGroup_Call {
	_c.Call.Return(run)
	return _c
}

// getStateTimeoutTask provides a mock function with given fields:
func (_m *MockStateMachineRequirements) getStateTimeoutTask() *TimeoutTask {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getStateTimeoutTask")
	}

	var r0 *TimeoutTask
	if rf, ok := ret.Get(0).(func() *TimeoutTask); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TimeoutTask)
		}
	}

	return r0
}

// MockStateMachineRequirements_getStateTimeoutTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getStateTimeoutTask'
type MockStateMachineRequirements_getStateTimeoutTask_Call struct {
	*mock.Call
}

// getStateTimeoutTask is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) getStateTimeoutTask() *MockStateMachineRequirements_getStateTimeoutTask_Call {
	return &MockStateMachineRequirements_getStateTimeoutTask_Call{Call: _e.mock.On("getStateTimeoutTask")}
}

func (_c *MockStateMachineRequirements_getStateTimeoutTask_Call) Run(run func()) *MockStateMachineRequirements_getStateTimeoutTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_getStateTimeoutTask_Call) Return(_a0 *TimeoutTask) *MockStateMachineRequirements_getStateTimeoutTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_getStateTimeoutTask_Call) RunAndReturn(run func() *TimeoutTask) *MockStateMachineRequirements_getStateTimeoutTask_Call {
	_c.Call.Return(run)
	return _c
}

// getStates provides a mock function with given fields:
func (_m *MockStateMachineRequirements) getStates() []State {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getStates")
	}

	var r0 []State
	if rf, ok := ret.Get(0).(func() []State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]State)
		}
	}

	return r0
}

// MockStateMachineRequirements_getStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getStates'
type MockStateMachineRequirements_getStates_Call struct {
	*mock.Call
}

// getStates is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) getStates() *MockStateMachineRequirements_getStates_Call {
	return &MockStateMachineRequirements_getStates_Call{Call: _e.mock.On("getStates")}
}

func (_c *MockStateMachineRequirements_getStates_Call) Run(run func()) *MockStateMachineRequirements_getStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_getStates_Call) Return(_a0 []State) *MockStateMachineRequirements_getStates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_getStates_Call) RunAndReturn(run func() []State) *MockStateMachineRequirements_getStates_Call {
	_c.Call.Return(run)
	return _c
}

// halt provides a mock function with given fields:
func (_m *MockStateMachineRequirements) halt() {
	_m.Called()
}

// MockStateMachineRequirements_halt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'halt'
type MockStateMachineRequirements_halt_Call struct {
	*mock.Call
}

// halt is a helper method to define mock.On call
func (_e *MockStateMachineRequirements_Expecter) halt() *MockStateMachineRequirements_halt_Call {
	return &MockStateMachineRequirements_halt_Call{Call: _e.mock.On("halt")}
}

func (_c *MockStateMachineRequirements_halt_Call) Run(run func()) *MockStateMachineRequirements_halt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateMachineRequirements_halt_Call) Return() *MockStateMachineRequirements_halt_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_halt_Call) RunAndReturn(run func()) *MockStateMachineRequirements_halt_Call {
	_c.Call.Return(run)
	return _c
}

// setMachineGroup provides a mock function with given fields: machineGroup
func (_m *MockStateMachineRequirements) setMachineGroup(machineGroup *StateMachineGroup) {
	_m.Called(machineGroup)
}

// MockStateMachineRequirements_setMachineGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setMachineGroup'
type MockStateMachineRequirements_setMachineGroup_Call struct {
	*mock.Call
}

// setMachineGroup is a helper method to define mock.On call
//   - machineGroup *StateMachineGroup
func (_e *MockStateMachineRequirements_Expecter) setMachineGroup(machineGroup interface{}) *MockStateMachineRequirements_setMachineGroup_Call {
	return &MockStateMachineRequirements_setMachineGroup_Call{Call: _e.mock.On("setMachineGroup", machineGroup)}
}

func (_c *MockStateMachineRequirements_setMachineGroup_Call) Run(run func(machineGroup *StateMachineGroup)) *MockStateMachineRequirements_setMachineGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*StateMachineGroup))
	})
	return _c
}

func (_c *MockStateMachineRequirements_setMachineGroup_Call) Return() *MockStateMachineRequirements_setMachineGroup_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStateMachineRequirements_setMachineGroup_Call) RunAndReturn(run func(*StateMachineGroup)) *MockStateMachineRequirements_setMachineGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStateMachineRequirements creates a new instance of MockStateMachineRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStateMachineRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStateMachineRequirements {
	mock := &MockStateMachineRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
