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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemClkSetRequest is the corresponding interface of S7PayloadUserDataItemClkSetRequest
type S7PayloadUserDataItemClkSetRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
	// IsS7PayloadUserDataItemClkSetRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemClkSetRequest()
	// CreateBuilder creates a S7PayloadUserDataItemClkSetRequestBuilder
	CreateS7PayloadUserDataItemClkSetRequestBuilder() S7PayloadUserDataItemClkSetRequestBuilder
}

// _S7PayloadUserDataItemClkSetRequest is the data-structure of this message
type _S7PayloadUserDataItemClkSetRequest struct {
	S7PayloadUserDataItemContract
	TimeStamp DateAndTime
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ S7PayloadUserDataItemClkSetRequest = (*_S7PayloadUserDataItemClkSetRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemClkSetRequest)(nil)

// NewS7PayloadUserDataItemClkSetRequest factory function for _S7PayloadUserDataItemClkSetRequest
func NewS7PayloadUserDataItemClkSetRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, timeStamp DateAndTime) *_S7PayloadUserDataItemClkSetRequest {
	if timeStamp == nil {
		panic("timeStamp of type DateAndTime for S7PayloadUserDataItemClkSetRequest must not be nil")
	}
	_result := &_S7PayloadUserDataItemClkSetRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		TimeStamp:                     timeStamp,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemClkSetRequestBuilder is a builder for S7PayloadUserDataItemClkSetRequest
type S7PayloadUserDataItemClkSetRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeStamp DateAndTime) S7PayloadUserDataItemClkSetRequestBuilder
	// WithTimeStamp adds TimeStamp (property field)
	WithTimeStamp(DateAndTime) S7PayloadUserDataItemClkSetRequestBuilder
	// WithTimeStampBuilder adds TimeStamp (property field) which is build by the builder
	WithTimeStampBuilder(func(DateAndTimeBuilder) DateAndTimeBuilder) S7PayloadUserDataItemClkSetRequestBuilder
	// Build builds the S7PayloadUserDataItemClkSetRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemClkSetRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemClkSetRequest
}

// NewS7PayloadUserDataItemClkSetRequestBuilder() creates a S7PayloadUserDataItemClkSetRequestBuilder
func NewS7PayloadUserDataItemClkSetRequestBuilder() S7PayloadUserDataItemClkSetRequestBuilder {
	return &_S7PayloadUserDataItemClkSetRequestBuilder{_S7PayloadUserDataItemClkSetRequest: new(_S7PayloadUserDataItemClkSetRequest)}
}

type _S7PayloadUserDataItemClkSetRequestBuilder struct {
	*_S7PayloadUserDataItemClkSetRequest

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemClkSetRequestBuilder) = (*_S7PayloadUserDataItemClkSetRequestBuilder)(nil)

func (m *_S7PayloadUserDataItemClkSetRequestBuilder) WithMandatoryFields(timeStamp DateAndTime) S7PayloadUserDataItemClkSetRequestBuilder {
	return m.WithTimeStamp(timeStamp)
}

func (m *_S7PayloadUserDataItemClkSetRequestBuilder) WithTimeStamp(timeStamp DateAndTime) S7PayloadUserDataItemClkSetRequestBuilder {
	m.TimeStamp = timeStamp
	return m
}

func (m *_S7PayloadUserDataItemClkSetRequestBuilder) WithTimeStampBuilder(builderSupplier func(DateAndTimeBuilder) DateAndTimeBuilder) S7PayloadUserDataItemClkSetRequestBuilder {
	builder := builderSupplier(m.TimeStamp.CreateDateAndTimeBuilder())
	var err error
	m.TimeStamp, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "DateAndTimeBuilder failed"))
	}
	return m
}

func (m *_S7PayloadUserDataItemClkSetRequestBuilder) Build() (S7PayloadUserDataItemClkSetRequest, error) {
	if m.TimeStamp == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'timeStamp' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._S7PayloadUserDataItemClkSetRequest.deepCopy(), nil
}

func (m *_S7PayloadUserDataItemClkSetRequestBuilder) MustBuild() S7PayloadUserDataItemClkSetRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_S7PayloadUserDataItemClkSetRequestBuilder) DeepCopy() any {
	return m.CreateS7PayloadUserDataItemClkSetRequestBuilder()
}

// CreateS7PayloadUserDataItemClkSetRequestBuilder creates a S7PayloadUserDataItemClkSetRequestBuilder
func (m *_S7PayloadUserDataItemClkSetRequest) CreateS7PayloadUserDataItemClkSetRequestBuilder() S7PayloadUserDataItemClkSetRequestBuilder {
	if m == nil {
		return NewS7PayloadUserDataItemClkSetRequestBuilder()
	}
	return &_S7PayloadUserDataItemClkSetRequestBuilder{_S7PayloadUserDataItemClkSetRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkSetRequest) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkSetRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemClkSetRequest) GetCpuSubfunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkSetRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemClkSetRequest) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkSetRequest(structType any) S7PayloadUserDataItemClkSetRequest {
	if casted, ok := structType.(S7PayloadUserDataItemClkSetRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkSetRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkSetRequest) GetTypeName() string {
	return "S7PayloadUserDataItemClkSetRequest"
}

func (m *_S7PayloadUserDataItemClkSetRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkSetRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemClkSetRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemClkSetRequest S7PayloadUserDataItemClkSetRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkSetRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkSetRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	timeStamp, err := ReadSimpleField[DateAndTime](ctx, "timeStamp", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeStamp' field"))
	}
	m.TimeStamp = timeStamp

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkSetRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkSetRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemClkSetRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkSetRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkSetRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkSetRequest")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleField[DateAndTime](ctx, "timeStamp", m.GetTimeStamp(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeStamp' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkSetRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkSetRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkSetRequest) IsS7PayloadUserDataItemClkSetRequest() {}

func (m *_S7PayloadUserDataItemClkSetRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemClkSetRequest) deepCopy() *_S7PayloadUserDataItemClkSetRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemClkSetRequestCopy := &_S7PayloadUserDataItemClkSetRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.TimeStamp.DeepCopy().(DateAndTime),
		m.reservedField0,
		m.reservedField1,
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemClkSetRequestCopy
}

func (m *_S7PayloadUserDataItemClkSetRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
