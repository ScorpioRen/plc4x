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

package spi

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockMessageCodec is an autogenerated mock type for the MessageCodec type
type MockMessageCodec struct {
	mock.Mock
}

type MockMessageCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMessageCodec) EXPECT() *MockMessageCodec_Expecter {
	return &MockMessageCodec_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields:
func (_m *MockMessageCodec) Connect() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageCodec_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockMessageCodec_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockMessageCodec_Expecter) Connect() *MockMessageCodec_Connect_Call {
	return &MockMessageCodec_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockMessageCodec_Connect_Call) Run(run func()) *MockMessageCodec_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessageCodec_Connect_Call) Return(_a0 error) *MockMessageCodec_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_Connect_Call) RunAndReturn(run func() error) *MockMessageCodec_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockMessageCodec) ConnectWithContext(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ConnectWithContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageCodec_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockMessageCodec_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMessageCodec_Expecter) ConnectWithContext(ctx interface{}) *MockMessageCodec_ConnectWithContext_Call {
	return &MockMessageCodec_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockMessageCodec_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockMessageCodec_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMessageCodec_ConnectWithContext_Call) Return(_a0 error) *MockMessageCodec_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_ConnectWithContext_Call) RunAndReturn(run func(context.Context) error) *MockMessageCodec_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// Disconnect provides a mock function with given fields:
func (_m *MockMessageCodec) Disconnect() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Disconnect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageCodec_Disconnect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disconnect'
type MockMessageCodec_Disconnect_Call struct {
	*mock.Call
}

// Disconnect is a helper method to define mock.On call
func (_e *MockMessageCodec_Expecter) Disconnect() *MockMessageCodec_Disconnect_Call {
	return &MockMessageCodec_Disconnect_Call{Call: _e.mock.On("Disconnect")}
}

func (_c *MockMessageCodec_Disconnect_Call) Run(run func()) *MockMessageCodec_Disconnect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessageCodec_Disconnect_Call) Return(_a0 error) *MockMessageCodec_Disconnect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_Disconnect_Call) RunAndReturn(run func() error) *MockMessageCodec_Disconnect_Call {
	_c.Call.Return(run)
	return _c
}

// Expect provides a mock function with given fields: ctx, acceptsMessage, handleMessage, handleError, ttl
func (_m *MockMessageCodec) Expect(ctx context.Context, acceptsMessage AcceptsMessage, handleMessage HandleMessage, handleError HandleError, ttl time.Duration) error {
	ret := _m.Called(ctx, acceptsMessage, handleMessage, handleError, ttl)

	if len(ret) == 0 {
		panic("no return value specified for Expect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AcceptsMessage, HandleMessage, HandleError, time.Duration) error); ok {
		r0 = rf(ctx, acceptsMessage, handleMessage, handleError, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageCodec_Expect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Expect'
type MockMessageCodec_Expect_Call struct {
	*mock.Call
}

// Expect is a helper method to define mock.On call
//   - ctx context.Context
//   - acceptsMessage AcceptsMessage
//   - handleMessage HandleMessage
//   - handleError HandleError
//   - ttl time.Duration
func (_e *MockMessageCodec_Expecter) Expect(ctx interface{}, acceptsMessage interface{}, handleMessage interface{}, handleError interface{}, ttl interface{}) *MockMessageCodec_Expect_Call {
	return &MockMessageCodec_Expect_Call{Call: _e.mock.On("Expect", ctx, acceptsMessage, handleMessage, handleError, ttl)}
}

func (_c *MockMessageCodec_Expect_Call) Run(run func(ctx context.Context, acceptsMessage AcceptsMessage, handleMessage HandleMessage, handleError HandleError, ttl time.Duration)) *MockMessageCodec_Expect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(AcceptsMessage), args[2].(HandleMessage), args[3].(HandleError), args[4].(time.Duration))
	})
	return _c
}

func (_c *MockMessageCodec_Expect_Call) Return(_a0 error) *MockMessageCodec_Expect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_Expect_Call) RunAndReturn(run func(context.Context, AcceptsMessage, HandleMessage, HandleError, time.Duration) error) *MockMessageCodec_Expect_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultIncomingMessageChannel provides a mock function with given fields:
func (_m *MockMessageCodec) GetDefaultIncomingMessageChannel() chan Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultIncomingMessageChannel")
	}

	var r0 chan Message
	if rf, ok := ret.Get(0).(func() chan Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan Message)
		}
	}

	return r0
}

// MockMessageCodec_GetDefaultIncomingMessageChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultIncomingMessageChannel'
type MockMessageCodec_GetDefaultIncomingMessageChannel_Call struct {
	*mock.Call
}

// GetDefaultIncomingMessageChannel is a helper method to define mock.On call
func (_e *MockMessageCodec_Expecter) GetDefaultIncomingMessageChannel() *MockMessageCodec_GetDefaultIncomingMessageChannel_Call {
	return &MockMessageCodec_GetDefaultIncomingMessageChannel_Call{Call: _e.mock.On("GetDefaultIncomingMessageChannel")}
}

func (_c *MockMessageCodec_GetDefaultIncomingMessageChannel_Call) Run(run func()) *MockMessageCodec_GetDefaultIncomingMessageChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessageCodec_GetDefaultIncomingMessageChannel_Call) Return(_a0 chan Message) *MockMessageCodec_GetDefaultIncomingMessageChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_GetDefaultIncomingMessageChannel_Call) RunAndReturn(run func() chan Message) *MockMessageCodec_GetDefaultIncomingMessageChannel_Call {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with given fields:
func (_m *MockMessageCodec) IsRunning() bool {
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

// MockMessageCodec_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type MockMessageCodec_IsRunning_Call struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *MockMessageCodec_Expecter) IsRunning() *MockMessageCodec_IsRunning_Call {
	return &MockMessageCodec_IsRunning_Call{Call: _e.mock.On("IsRunning")}
}

func (_c *MockMessageCodec_IsRunning_Call) Run(run func()) *MockMessageCodec_IsRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessageCodec_IsRunning_Call) Return(_a0 bool) *MockMessageCodec_IsRunning_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_IsRunning_Call) RunAndReturn(run func() bool) *MockMessageCodec_IsRunning_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: message
func (_m *MockMessageCodec) Send(message Message) error {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageCodec_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockMessageCodec_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - message Message
func (_e *MockMessageCodec_Expecter) Send(message interface{}) *MockMessageCodec_Send_Call {
	return &MockMessageCodec_Send_Call{Call: _e.mock.On("Send", message)}
}

func (_c *MockMessageCodec_Send_Call) Run(run func(message Message)) *MockMessageCodec_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Message))
	})
	return _c
}

func (_c *MockMessageCodec_Send_Call) Return(_a0 error) *MockMessageCodec_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_Send_Call) RunAndReturn(run func(Message) error) *MockMessageCodec_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendRequest provides a mock function with given fields: ctx, message, acceptsMessage, handleMessage, handleError, ttl
func (_m *MockMessageCodec) SendRequest(ctx context.Context, message Message, acceptsMessage AcceptsMessage, handleMessage HandleMessage, handleError HandleError, ttl time.Duration) error {
	ret := _m.Called(ctx, message, acceptsMessage, handleMessage, handleError, ttl)

	if len(ret) == 0 {
		panic("no return value specified for SendRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Message, AcceptsMessage, HandleMessage, HandleError, time.Duration) error); ok {
		r0 = rf(ctx, message, acceptsMessage, handleMessage, handleError, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageCodec_SendRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendRequest'
type MockMessageCodec_SendRequest_Call struct {
	*mock.Call
}

// SendRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - message Message
//   - acceptsMessage AcceptsMessage
//   - handleMessage HandleMessage
//   - handleError HandleError
//   - ttl time.Duration
func (_e *MockMessageCodec_Expecter) SendRequest(ctx interface{}, message interface{}, acceptsMessage interface{}, handleMessage interface{}, handleError interface{}, ttl interface{}) *MockMessageCodec_SendRequest_Call {
	return &MockMessageCodec_SendRequest_Call{Call: _e.mock.On("SendRequest", ctx, message, acceptsMessage, handleMessage, handleError, ttl)}
}

func (_c *MockMessageCodec_SendRequest_Call) Run(run func(ctx context.Context, message Message, acceptsMessage AcceptsMessage, handleMessage HandleMessage, handleError HandleError, ttl time.Duration)) *MockMessageCodec_SendRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Message), args[2].(AcceptsMessage), args[3].(HandleMessage), args[4].(HandleError), args[5].(time.Duration))
	})
	return _c
}

func (_c *MockMessageCodec_SendRequest_Call) Return(_a0 error) *MockMessageCodec_SendRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageCodec_SendRequest_Call) RunAndReturn(run func(context.Context, Message, AcceptsMessage, HandleMessage, HandleError, time.Duration) error) *MockMessageCodec_SendRequest_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMessageCodec creates a new instance of MockMessageCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageCodec {
	mock := &MockMessageCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
