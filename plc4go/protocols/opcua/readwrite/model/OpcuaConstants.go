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
const OpcuaConstants_PROTOCOLVERSION uint8 = uint8(0)

// OpcuaConstants is the corresponding interface of OpcuaConstants
type OpcuaConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsOpcuaConstants is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaConstants()
	// CreateBuilder creates a OpcuaConstantsBuilder
	CreateOpcuaConstantsBuilder() OpcuaConstantsBuilder
}

// _OpcuaConstants is the data-structure of this message
type _OpcuaConstants struct {
}

var _ OpcuaConstants = (*_OpcuaConstants)(nil)

// NewOpcuaConstants factory function for _OpcuaConstants
func NewOpcuaConstants() *_OpcuaConstants {
	return &_OpcuaConstants{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaConstantsBuilder is a builder for OpcuaConstants
type OpcuaConstantsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() OpcuaConstantsBuilder
	// Build builds the OpcuaConstants or returns an error if something is wrong
	Build() (OpcuaConstants, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaConstants
}

// NewOpcuaConstantsBuilder() creates a OpcuaConstantsBuilder
func NewOpcuaConstantsBuilder() OpcuaConstantsBuilder {
	return &_OpcuaConstantsBuilder{_OpcuaConstants: new(_OpcuaConstants)}
}

type _OpcuaConstantsBuilder struct {
	*_OpcuaConstants

	err *utils.MultiError
}

var _ (OpcuaConstantsBuilder) = (*_OpcuaConstantsBuilder)(nil)

func (m *_OpcuaConstantsBuilder) WithMandatoryFields() OpcuaConstantsBuilder {
	return m
}

func (m *_OpcuaConstantsBuilder) Build() (OpcuaConstants, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._OpcuaConstants.deepCopy(), nil
}

func (m *_OpcuaConstantsBuilder) MustBuild() OpcuaConstants {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_OpcuaConstantsBuilder) DeepCopy() any {
	return m.CreateOpcuaConstantsBuilder()
}

// CreateOpcuaConstantsBuilder creates a OpcuaConstantsBuilder
func (m *_OpcuaConstants) CreateOpcuaConstantsBuilder() OpcuaConstantsBuilder {
	if m == nil {
		return NewOpcuaConstantsBuilder()
	}
	return &_OpcuaConstantsBuilder{_OpcuaConstants: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_OpcuaConstants) GetProtocolVersion() uint8 {
	return OpcuaConstants_PROTOCOLVERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaConstants(structType any) OpcuaConstants {
	if casted, ok := structType.(OpcuaConstants); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaConstants); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaConstants) GetTypeName() string {
	return "OpcuaConstants"
}

func (m *_OpcuaConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolVersion)
	lengthInBits += 8

	return lengthInBits
}

func (m *_OpcuaConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaConstantsParse(ctx context.Context, theBytes []byte) (OpcuaConstants, error) {
	return OpcuaConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaConstantsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaConstants, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaConstants, error) {
		return OpcuaConstantsParseWithBuffer(ctx, readBuffer)
	}
}

func OpcuaConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaConstants, error) {
	v, err := (&_OpcuaConstants{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_OpcuaConstants) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__opcuaConstants OpcuaConstants, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolVersion, err := ReadConstField[uint8](ctx, "protocolVersion", ReadUnsignedByte(readBuffer, uint8(8)), OpcuaConstants_PROTOCOLVERSION)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersion' field"))
	}
	_ = protocolVersion

	if closeErr := readBuffer.CloseContext("OpcuaConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaConstants")
	}

	return m, nil
}

func (m *_OpcuaConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("OpcuaConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for OpcuaConstants")
	}

	if err := WriteConstField(ctx, "protocolVersion", OpcuaConstants_PROTOCOLVERSION, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'protocolVersion' field")
	}

	if popErr := writeBuffer.PopContext("OpcuaConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for OpcuaConstants")
	}
	return nil
}

func (m *_OpcuaConstants) IsOpcuaConstants() {}

func (m *_OpcuaConstants) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaConstants) deepCopy() *_OpcuaConstants {
	if m == nil {
		return nil
	}
	_OpcuaConstantsCopy := &_OpcuaConstants{}
	return _OpcuaConstantsCopy
}

func (m *_OpcuaConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
