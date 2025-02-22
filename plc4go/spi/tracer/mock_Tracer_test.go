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

package tracer

import mock "github.com/stretchr/testify/mock"

// MockTracer is an autogenerated mock type for the Tracer type
type MockTracer struct {
	mock.Mock
}

type MockTracer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTracer) EXPECT() *MockTracer_Expecter {
	return &MockTracer_Expecter{mock: &_m.Mock}
}

// AddTrace provides a mock function with given fields: operation, message
func (_m *MockTracer) AddTrace(operation string, message string) {
	_m.Called(operation, message)
}

// MockTracer_AddTrace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTrace'
type MockTracer_AddTrace_Call struct {
	*mock.Call
}

// AddTrace is a helper method to define mock.On call
//   - operation string
//   - message string
func (_e *MockTracer_Expecter) AddTrace(operation interface{}, message interface{}) *MockTracer_AddTrace_Call {
	return &MockTracer_AddTrace_Call{Call: _e.mock.On("AddTrace", operation, message)}
}

func (_c *MockTracer_AddTrace_Call) Run(run func(operation string, message string)) *MockTracer_AddTrace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTracer_AddTrace_Call) Return() *MockTracer_AddTrace_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTracer_AddTrace_Call) RunAndReturn(run func(string, string)) *MockTracer_AddTrace_Call {
	_c.Call.Return(run)
	return _c
}

// AddTransactionalStartTrace provides a mock function with given fields: operation, message
func (_m *MockTracer) AddTransactionalStartTrace(operation string, message string) string {
	ret := _m.Called(operation, message)

	if len(ret) == 0 {
		panic("no return value specified for AddTransactionalStartTrace")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(operation, message)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTracer_AddTransactionalStartTrace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTransactionalStartTrace'
type MockTracer_AddTransactionalStartTrace_Call struct {
	*mock.Call
}

// AddTransactionalStartTrace is a helper method to define mock.On call
//   - operation string
//   - message string
func (_e *MockTracer_Expecter) AddTransactionalStartTrace(operation interface{}, message interface{}) *MockTracer_AddTransactionalStartTrace_Call {
	return &MockTracer_AddTransactionalStartTrace_Call{Call: _e.mock.On("AddTransactionalStartTrace", operation, message)}
}

func (_c *MockTracer_AddTransactionalStartTrace_Call) Run(run func(operation string, message string)) *MockTracer_AddTransactionalStartTrace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTracer_AddTransactionalStartTrace_Call) Return(_a0 string) *MockTracer_AddTransactionalStartTrace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTracer_AddTransactionalStartTrace_Call) RunAndReturn(run func(string, string) string) *MockTracer_AddTransactionalStartTrace_Call {
	_c.Call.Return(run)
	return _c
}

// AddTransactionalTrace provides a mock function with given fields: transactionId, operation, message
func (_m *MockTracer) AddTransactionalTrace(transactionId string, operation string, message string) {
	_m.Called(transactionId, operation, message)
}

// MockTracer_AddTransactionalTrace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTransactionalTrace'
type MockTracer_AddTransactionalTrace_Call struct {
	*mock.Call
}

// AddTransactionalTrace is a helper method to define mock.On call
//   - transactionId string
//   - operation string
//   - message string
func (_e *MockTracer_Expecter) AddTransactionalTrace(transactionId interface{}, operation interface{}, message interface{}) *MockTracer_AddTransactionalTrace_Call {
	return &MockTracer_AddTransactionalTrace_Call{Call: _e.mock.On("AddTransactionalTrace", transactionId, operation, message)}
}

func (_c *MockTracer_AddTransactionalTrace_Call) Run(run func(transactionId string, operation string, message string)) *MockTracer_AddTransactionalTrace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockTracer_AddTransactionalTrace_Call) Return() *MockTracer_AddTransactionalTrace_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTracer_AddTransactionalTrace_Call) RunAndReturn(run func(string, string, string)) *MockTracer_AddTransactionalTrace_Call {
	_c.Call.Return(run)
	return _c
}

// FilterTraces provides a mock function with given fields: traces, connectionIdFilter, transactionIdFilter, operationFilter, messageFilter
func (_m *MockTracer) FilterTraces(traces []TraceEntry, connectionIdFilter string, transactionIdFilter string, operationFilter string, messageFilter string) []TraceEntry {
	ret := _m.Called(traces, connectionIdFilter, transactionIdFilter, operationFilter, messageFilter)

	if len(ret) == 0 {
		panic("no return value specified for FilterTraces")
	}

	var r0 []TraceEntry
	if rf, ok := ret.Get(0).(func([]TraceEntry, string, string, string, string) []TraceEntry); ok {
		r0 = rf(traces, connectionIdFilter, transactionIdFilter, operationFilter, messageFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TraceEntry)
		}
	}

	return r0
}

// MockTracer_FilterTraces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterTraces'
type MockTracer_FilterTraces_Call struct {
	*mock.Call
}

// FilterTraces is a helper method to define mock.On call
//   - traces []TraceEntry
//   - connectionIdFilter string
//   - transactionIdFilter string
//   - operationFilter string
//   - messageFilter string
func (_e *MockTracer_Expecter) FilterTraces(traces interface{}, connectionIdFilter interface{}, transactionIdFilter interface{}, operationFilter interface{}, messageFilter interface{}) *MockTracer_FilterTraces_Call {
	return &MockTracer_FilterTraces_Call{Call: _e.mock.On("FilterTraces", traces, connectionIdFilter, transactionIdFilter, operationFilter, messageFilter)}
}

func (_c *MockTracer_FilterTraces_Call) Run(run func(traces []TraceEntry, connectionIdFilter string, transactionIdFilter string, operationFilter string, messageFilter string)) *MockTracer_FilterTraces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]TraceEntry), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockTracer_FilterTraces_Call) Return(_a0 []TraceEntry) *MockTracer_FilterTraces_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTracer_FilterTraces_Call) RunAndReturn(run func([]TraceEntry, string, string, string, string) []TraceEntry) *MockTracer_FilterTraces_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionId provides a mock function with given fields:
func (_m *MockTracer) GetConnectionId() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionId")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTracer_GetConnectionId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionId'
type MockTracer_GetConnectionId_Call struct {
	*mock.Call
}

// GetConnectionId is a helper method to define mock.On call
func (_e *MockTracer_Expecter) GetConnectionId() *MockTracer_GetConnectionId_Call {
	return &MockTracer_GetConnectionId_Call{Call: _e.mock.On("GetConnectionId")}
}

func (_c *MockTracer_GetConnectionId_Call) Run(run func()) *MockTracer_GetConnectionId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTracer_GetConnectionId_Call) Return(_a0 string) *MockTracer_GetConnectionId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTracer_GetConnectionId_Call) RunAndReturn(run func() string) *MockTracer_GetConnectionId_Call {
	_c.Call.Return(run)
	return _c
}

// GetTraces provides a mock function with given fields:
func (_m *MockTracer) GetTraces() []TraceEntry {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTraces")
	}

	var r0 []TraceEntry
	if rf, ok := ret.Get(0).(func() []TraceEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TraceEntry)
		}
	}

	return r0
}

// MockTracer_GetTraces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTraces'
type MockTracer_GetTraces_Call struct {
	*mock.Call
}

// GetTraces is a helper method to define mock.On call
func (_e *MockTracer_Expecter) GetTraces() *MockTracer_GetTraces_Call {
	return &MockTracer_GetTraces_Call{Call: _e.mock.On("GetTraces")}
}

func (_c *MockTracer_GetTraces_Call) Run(run func()) *MockTracer_GetTraces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTracer_GetTraces_Call) Return(_a0 []TraceEntry) *MockTracer_GetTraces_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTracer_GetTraces_Call) RunAndReturn(run func() []TraceEntry) *MockTracer_GetTraces_Call {
	_c.Call.Return(run)
	return _c
}

// ResetTraces provides a mock function with given fields:
func (_m *MockTracer) ResetTraces() {
	_m.Called()
}

// MockTracer_ResetTraces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetTraces'
type MockTracer_ResetTraces_Call struct {
	*mock.Call
}

// ResetTraces is a helper method to define mock.On call
func (_e *MockTracer_Expecter) ResetTraces() *MockTracer_ResetTraces_Call {
	return &MockTracer_ResetTraces_Call{Call: _e.mock.On("ResetTraces")}
}

func (_c *MockTracer_ResetTraces_Call) Run(run func()) *MockTracer_ResetTraces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTracer_ResetTraces_Call) Return() *MockTracer_ResetTraces_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTracer_ResetTraces_Call) RunAndReturn(run func()) *MockTracer_ResetTraces_Call {
	_c.Call.Return(run)
	return _c
}

// SetConnectionId provides a mock function with given fields: connectionId
func (_m *MockTracer) SetConnectionId(connectionId string) {
	_m.Called(connectionId)
}

// MockTracer_SetConnectionId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConnectionId'
type MockTracer_SetConnectionId_Call struct {
	*mock.Call
}

// SetConnectionId is a helper method to define mock.On call
//   - connectionId string
func (_e *MockTracer_Expecter) SetConnectionId(connectionId interface{}) *MockTracer_SetConnectionId_Call {
	return &MockTracer_SetConnectionId_Call{Call: _e.mock.On("SetConnectionId", connectionId)}
}

func (_c *MockTracer_SetConnectionId_Call) Run(run func(connectionId string)) *MockTracer_SetConnectionId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTracer_SetConnectionId_Call) Return() *MockTracer_SetConnectionId_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTracer_SetConnectionId_Call) RunAndReturn(run func(string)) *MockTracer_SetConnectionId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTracer creates a new instance of MockTracer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTracer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTracer {
	mock := &MockTracer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
