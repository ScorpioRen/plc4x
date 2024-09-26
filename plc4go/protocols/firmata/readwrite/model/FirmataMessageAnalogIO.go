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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageAnalogIO is the corresponding interface of FirmataMessageAnalogIO
type FirmataMessageAnalogIO interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetData returns Data (property field)
	GetData() []int8
	// IsFirmataMessageAnalogIO is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataMessageAnalogIO()
	// CreateBuilder creates a FirmataMessageAnalogIOBuilder
	CreateFirmataMessageAnalogIOBuilder() FirmataMessageAnalogIOBuilder
}

// _FirmataMessageAnalogIO is the data-structure of this message
type _FirmataMessageAnalogIO struct {
	FirmataMessageContract
	Pin  uint8
	Data []int8
}

var _ FirmataMessageAnalogIO = (*_FirmataMessageAnalogIO)(nil)
var _ FirmataMessageRequirements = (*_FirmataMessageAnalogIO)(nil)

// NewFirmataMessageAnalogIO factory function for _FirmataMessageAnalogIO
func NewFirmataMessageAnalogIO(pin uint8, data []int8, response bool) *_FirmataMessageAnalogIO {
	_result := &_FirmataMessageAnalogIO{
		FirmataMessageContract: NewFirmataMessage(response),
		Pin:                    pin,
		Data:                   data,
	}
	_result.FirmataMessageContract.(*_FirmataMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FirmataMessageAnalogIOBuilder is a builder for FirmataMessageAnalogIO
type FirmataMessageAnalogIOBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pin uint8, data []int8) FirmataMessageAnalogIOBuilder
	// WithPin adds Pin (property field)
	WithPin(uint8) FirmataMessageAnalogIOBuilder
	// WithData adds Data (property field)
	WithData(...int8) FirmataMessageAnalogIOBuilder
	// Build builds the FirmataMessageAnalogIO or returns an error if something is wrong
	Build() (FirmataMessageAnalogIO, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FirmataMessageAnalogIO
}

// NewFirmataMessageAnalogIOBuilder() creates a FirmataMessageAnalogIOBuilder
func NewFirmataMessageAnalogIOBuilder() FirmataMessageAnalogIOBuilder {
	return &_FirmataMessageAnalogIOBuilder{_FirmataMessageAnalogIO: new(_FirmataMessageAnalogIO)}
}

type _FirmataMessageAnalogIOBuilder struct {
	*_FirmataMessageAnalogIO

	err *utils.MultiError
}

var _ (FirmataMessageAnalogIOBuilder) = (*_FirmataMessageAnalogIOBuilder)(nil)

func (m *_FirmataMessageAnalogIOBuilder) WithMandatoryFields(pin uint8, data []int8) FirmataMessageAnalogIOBuilder {
	return m.WithPin(pin).WithData(data...)
}

func (m *_FirmataMessageAnalogIOBuilder) WithPin(pin uint8) FirmataMessageAnalogIOBuilder {
	m.Pin = pin
	return m
}

func (m *_FirmataMessageAnalogIOBuilder) WithData(data ...int8) FirmataMessageAnalogIOBuilder {
	m.Data = data
	return m
}

func (m *_FirmataMessageAnalogIOBuilder) Build() (FirmataMessageAnalogIO, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._FirmataMessageAnalogIO.deepCopy(), nil
}

func (m *_FirmataMessageAnalogIOBuilder) MustBuild() FirmataMessageAnalogIO {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_FirmataMessageAnalogIOBuilder) DeepCopy() any {
	return m.CreateFirmataMessageAnalogIOBuilder()
}

// CreateFirmataMessageAnalogIOBuilder creates a FirmataMessageAnalogIOBuilder
func (m *_FirmataMessageAnalogIO) CreateFirmataMessageAnalogIOBuilder() FirmataMessageAnalogIOBuilder {
	if m == nil {
		return NewFirmataMessageAnalogIOBuilder()
	}
	return &_FirmataMessageAnalogIOBuilder{_FirmataMessageAnalogIO: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageAnalogIO) GetMessageType() uint8 {
	return 0xE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageAnalogIO) GetParent() FirmataMessageContract {
	return m.FirmataMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageAnalogIO) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataMessageAnalogIO) GetData() []int8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataMessageAnalogIO(structType any) FirmataMessageAnalogIO {
	if casted, ok := structType.(FirmataMessageAnalogIO); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageAnalogIO); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageAnalogIO) GetTypeName() string {
	return "FirmataMessageAnalogIO"
}

func (m *_FirmataMessageAnalogIO) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataMessageContract.(*_FirmataMessage).getLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 4

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_FirmataMessageAnalogIO) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataMessageAnalogIO) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataMessage, response bool) (__firmataMessageAnalogIO FirmataMessageAnalogIO, err error) {
	m.FirmataMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageAnalogIO"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageAnalogIO")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}
	m.Pin = pin

	data, err := ReadCountArrayField[int8](ctx, "data", ReadSignedByte(readBuffer, uint8(8)), uint64(int32(2)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("FirmataMessageAnalogIO"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageAnalogIO")
	}

	return m, nil
}

func (m *_FirmataMessageAnalogIO) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageAnalogIO) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageAnalogIO"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageAnalogIO")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 4), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteSignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageAnalogIO"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageAnalogIO")
		}
		return nil
	}
	return m.FirmataMessageContract.(*_FirmataMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageAnalogIO) IsFirmataMessageAnalogIO() {}

func (m *_FirmataMessageAnalogIO) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataMessageAnalogIO) deepCopy() *_FirmataMessageAnalogIO {
	if m == nil {
		return nil
	}
	_FirmataMessageAnalogIOCopy := &_FirmataMessageAnalogIO{
		m.FirmataMessageContract.(*_FirmataMessage).deepCopy(),
		m.Pin,
		utils.DeepCopySlice[int8, int8](m.Data),
	}
	m.FirmataMessageContract.(*_FirmataMessage)._SubType = m
	return _FirmataMessageAnalogIOCopy
}

func (m *_FirmataMessageAnalogIO) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
