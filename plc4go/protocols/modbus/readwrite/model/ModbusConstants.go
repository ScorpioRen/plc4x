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

// Constant values.
const ModbusConstants_MODBUSTCPDEFAULTPORT uint16 = uint16(502)

// ModbusConstants is the corresponding interface of ModbusConstants
type ModbusConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsModbusConstants is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusConstants()
	// CreateBuilder creates a ModbusConstantsBuilder
	CreateModbusConstantsBuilder() ModbusConstantsBuilder
}

// _ModbusConstants is the data-structure of this message
type _ModbusConstants struct {
}

var _ ModbusConstants = (*_ModbusConstants)(nil)

// NewModbusConstants factory function for _ModbusConstants
func NewModbusConstants() *_ModbusConstants {
	return &_ModbusConstants{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusConstantsBuilder is a builder for ModbusConstants
type ModbusConstantsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ModbusConstantsBuilder
	// Build builds the ModbusConstants or returns an error if something is wrong
	Build() (ModbusConstants, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusConstants
}

// NewModbusConstantsBuilder() creates a ModbusConstantsBuilder
func NewModbusConstantsBuilder() ModbusConstantsBuilder {
	return &_ModbusConstantsBuilder{_ModbusConstants: new(_ModbusConstants)}
}

type _ModbusConstantsBuilder struct {
	*_ModbusConstants

	err *utils.MultiError
}

var _ (ModbusConstantsBuilder) = (*_ModbusConstantsBuilder)(nil)

func (m *_ModbusConstantsBuilder) WithMandatoryFields() ModbusConstantsBuilder {
	return m
}

func (m *_ModbusConstantsBuilder) Build() (ModbusConstants, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ModbusConstants.deepCopy(), nil
}

func (m *_ModbusConstantsBuilder) MustBuild() ModbusConstants {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ModbusConstantsBuilder) DeepCopy() any {
	return m.CreateModbusConstantsBuilder()
}

// CreateModbusConstantsBuilder creates a ModbusConstantsBuilder
func (m *_ModbusConstants) CreateModbusConstantsBuilder() ModbusConstantsBuilder {
	if m == nil {
		return NewModbusConstantsBuilder()
	}
	return &_ModbusConstantsBuilder{_ModbusConstants: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusConstants) GetModbusTcpDefaultPort() uint16 {
	return ModbusConstants_MODBUSTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusConstants(structType any) ModbusConstants {
	if casted, ok := structType.(ModbusConstants); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusConstants); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusConstants) GetTypeName() string {
	return "ModbusConstants"
}

func (m *_ModbusConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (modbusTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusConstantsParse(ctx context.Context, theBytes []byte) (ModbusConstants, error) {
	return ModbusConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusConstantsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusConstants, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusConstants, error) {
		return ModbusConstantsParseWithBuffer(ctx, readBuffer)
	}
}

func ModbusConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusConstants, error) {
	v, err := (&_ModbusConstants{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ModbusConstants) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__modbusConstants ModbusConstants, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	modbusTcpDefaultPort, err := ReadConstField[uint16](ctx, "modbusTcpDefaultPort", ReadUnsignedShort(readBuffer, uint8(16)), ModbusConstants_MODBUSTCPDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modbusTcpDefaultPort' field"))
	}
	_ = modbusTcpDefaultPort

	if closeErr := readBuffer.CloseContext("ModbusConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusConstants")
	}

	return m, nil
}

func (m *_ModbusConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ModbusConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusConstants")
	}

	if err := WriteConstField(ctx, "modbusTcpDefaultPort", ModbusConstants_MODBUSTCPDEFAULTPORT, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'modbusTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("ModbusConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusConstants")
	}
	return nil
}

func (m *_ModbusConstants) IsModbusConstants() {}

func (m *_ModbusConstants) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusConstants) deepCopy() *_ModbusConstants {
	if m == nil {
		return nil
	}
	_ModbusConstantsCopy := &_ModbusConstants{}
	return _ModbusConstantsCopy
}

func (m *_ModbusConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
