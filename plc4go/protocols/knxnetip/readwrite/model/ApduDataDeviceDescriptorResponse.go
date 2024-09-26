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

// ApduDataDeviceDescriptorResponse is the corresponding interface of ApduDataDeviceDescriptorResponse
type ApduDataDeviceDescriptorResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// GetDescriptorType returns DescriptorType (property field)
	GetDescriptorType() uint8
	// GetData returns Data (property field)
	GetData() []byte
	// IsApduDataDeviceDescriptorResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataDeviceDescriptorResponse()
	// CreateBuilder creates a ApduDataDeviceDescriptorResponseBuilder
	CreateApduDataDeviceDescriptorResponseBuilder() ApduDataDeviceDescriptorResponseBuilder
}

// _ApduDataDeviceDescriptorResponse is the data-structure of this message
type _ApduDataDeviceDescriptorResponse struct {
	ApduDataContract
	DescriptorType uint8
	Data           []byte
}

var _ ApduDataDeviceDescriptorResponse = (*_ApduDataDeviceDescriptorResponse)(nil)
var _ ApduDataRequirements = (*_ApduDataDeviceDescriptorResponse)(nil)

// NewApduDataDeviceDescriptorResponse factory function for _ApduDataDeviceDescriptorResponse
func NewApduDataDeviceDescriptorResponse(descriptorType uint8, data []byte, dataLength uint8) *_ApduDataDeviceDescriptorResponse {
	_result := &_ApduDataDeviceDescriptorResponse{
		ApduDataContract: NewApduData(dataLength),
		DescriptorType:   descriptorType,
		Data:             data,
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataDeviceDescriptorResponseBuilder is a builder for ApduDataDeviceDescriptorResponse
type ApduDataDeviceDescriptorResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(descriptorType uint8, data []byte) ApduDataDeviceDescriptorResponseBuilder
	// WithDescriptorType adds DescriptorType (property field)
	WithDescriptorType(uint8) ApduDataDeviceDescriptorResponseBuilder
	// WithData adds Data (property field)
	WithData(...byte) ApduDataDeviceDescriptorResponseBuilder
	// Build builds the ApduDataDeviceDescriptorResponse or returns an error if something is wrong
	Build() (ApduDataDeviceDescriptorResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataDeviceDescriptorResponse
}

// NewApduDataDeviceDescriptorResponseBuilder() creates a ApduDataDeviceDescriptorResponseBuilder
func NewApduDataDeviceDescriptorResponseBuilder() ApduDataDeviceDescriptorResponseBuilder {
	return &_ApduDataDeviceDescriptorResponseBuilder{_ApduDataDeviceDescriptorResponse: new(_ApduDataDeviceDescriptorResponse)}
}

type _ApduDataDeviceDescriptorResponseBuilder struct {
	*_ApduDataDeviceDescriptorResponse

	err *utils.MultiError
}

var _ (ApduDataDeviceDescriptorResponseBuilder) = (*_ApduDataDeviceDescriptorResponseBuilder)(nil)

func (m *_ApduDataDeviceDescriptorResponseBuilder) WithMandatoryFields(descriptorType uint8, data []byte) ApduDataDeviceDescriptorResponseBuilder {
	return m.WithDescriptorType(descriptorType).WithData(data...)
}

func (m *_ApduDataDeviceDescriptorResponseBuilder) WithDescriptorType(descriptorType uint8) ApduDataDeviceDescriptorResponseBuilder {
	m.DescriptorType = descriptorType
	return m
}

func (m *_ApduDataDeviceDescriptorResponseBuilder) WithData(data ...byte) ApduDataDeviceDescriptorResponseBuilder {
	m.Data = data
	return m
}

func (m *_ApduDataDeviceDescriptorResponseBuilder) Build() (ApduDataDeviceDescriptorResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduDataDeviceDescriptorResponse.deepCopy(), nil
}

func (m *_ApduDataDeviceDescriptorResponseBuilder) MustBuild() ApduDataDeviceDescriptorResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduDataDeviceDescriptorResponseBuilder) DeepCopy() any {
	return m.CreateApduDataDeviceDescriptorResponseBuilder()
}

// CreateApduDataDeviceDescriptorResponseBuilder creates a ApduDataDeviceDescriptorResponseBuilder
func (m *_ApduDataDeviceDescriptorResponse) CreateApduDataDeviceDescriptorResponseBuilder() ApduDataDeviceDescriptorResponseBuilder {
	if m == nil {
		return NewApduDataDeviceDescriptorResponseBuilder()
	}
	return &_ApduDataDeviceDescriptorResponseBuilder{_ApduDataDeviceDescriptorResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataDeviceDescriptorResponse) GetApciType() uint8 {
	return 0xD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataDeviceDescriptorResponse) GetParent() ApduDataContract {
	return m.ApduDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataDeviceDescriptorResponse) GetDescriptorType() uint8 {
	return m.DescriptorType
}

func (m *_ApduDataDeviceDescriptorResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataDeviceDescriptorResponse(structType any) ApduDataDeviceDescriptorResponse {
	if casted, ok := structType.(ApduDataDeviceDescriptorResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataDeviceDescriptorResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataDeviceDescriptorResponse) GetTypeName() string {
	return "ApduDataDeviceDescriptorResponse"
}

func (m *_ApduDataDeviceDescriptorResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	// Simple field (descriptorType)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataDeviceDescriptorResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataDeviceDescriptorResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataDeviceDescriptorResponse ApduDataDeviceDescriptorResponse, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataDeviceDescriptorResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	descriptorType, err := ReadSimpleField(ctx, "descriptorType", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'descriptorType' field"))
	}
	m.DescriptorType = descriptorType

	data, err := readBuffer.ReadByteArray("data", int(utils.InlineIf((bool((dataLength) < (1))), func() any { return int32(int32(0)) }, func() any { return int32(int32(dataLength) - int32(int32(1))) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataDeviceDescriptorResponse")
	}

	return m, nil
}

func (m *_ApduDataDeviceDescriptorResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataDeviceDescriptorResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataDeviceDescriptorResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "descriptorType", m.GetDescriptorType(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
			return errors.Wrap(err, "Error serializing 'descriptorType' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataDeviceDescriptorResponse")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataDeviceDescriptorResponse) IsApduDataDeviceDescriptorResponse() {}

func (m *_ApduDataDeviceDescriptorResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataDeviceDescriptorResponse) deepCopy() *_ApduDataDeviceDescriptorResponse {
	if m == nil {
		return nil
	}
	_ApduDataDeviceDescriptorResponseCopy := &_ApduDataDeviceDescriptorResponse{
		m.ApduDataContract.(*_ApduData).deepCopy(),
		m.DescriptorType,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataDeviceDescriptorResponseCopy
}

func (m *_ApduDataDeviceDescriptorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
