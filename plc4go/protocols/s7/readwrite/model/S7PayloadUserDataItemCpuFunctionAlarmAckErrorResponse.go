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

// S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse
type S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// IsS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder
	CreateS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder
}

// _S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse struct {
	S7PayloadUserDataItemContract
}

var _ S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse = (*_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse)(nil)

// NewS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse factory function for _S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse
func NewS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder is a builder for S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse
type S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse
}

// NewS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder() creates a S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder
func NewS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder {
	return &_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder{_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse: new(_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse)}
}

type _S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder) = (*_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder)(nil)

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder) WithMandatoryFields() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder {
	return m
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder) Build() (S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse.deepCopy(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder) DeepCopy() any {
	return m.CreateS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder()
}

// CreateS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder creates a S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder
func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) CreateS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder() S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder {
	if m == nil {
		return NewS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseBuilder{_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetCpuSubfunction() uint8 {
	return 0x0b
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse(structType any) S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) IsS7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse() {
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) deepCopy() *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseCopy := &_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
