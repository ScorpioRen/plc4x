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

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder
	CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse struct {
	S7PayloadUserDataItemContract
}

var _ S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)(nil)

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder is a builder for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
}

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder() creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder {
	return &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder{_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse: new(_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)}
}

type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder) = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder)(nil)

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder) WithMandatoryFields() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder {
	return m
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder) Build() (S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse.deepCopy(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder) DeepCopy() any {
	return m.CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder()
}

// CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder creates a S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder
func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) CreateS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder() S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder {
	if m == nil {
		return NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseBuilder{_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(structType any) S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse() {
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) deepCopy() *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseCopy := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
