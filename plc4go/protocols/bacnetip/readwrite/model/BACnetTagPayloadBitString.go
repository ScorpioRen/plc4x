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

// BACnetTagPayloadBitString is the corresponding interface of BACnetTagPayloadBitString
type BACnetTagPayloadBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetUnusedBits returns UnusedBits (property field)
	GetUnusedBits() uint8
	// GetData returns Data (property field)
	GetData() []bool
	// GetUnused returns Unused (property field)
	GetUnused() []bool
	// IsBACnetTagPayloadBitString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadBitString()
	// CreateBuilder creates a BACnetTagPayloadBitStringBuilder
	CreateBACnetTagPayloadBitStringBuilder() BACnetTagPayloadBitStringBuilder
}

// _BACnetTagPayloadBitString is the data-structure of this message
type _BACnetTagPayloadBitString struct {
	UnusedBits uint8
	Data       []bool
	Unused     []bool

	// Arguments.
	ActualLength uint32
}

var _ BACnetTagPayloadBitString = (*_BACnetTagPayloadBitString)(nil)

// NewBACnetTagPayloadBitString factory function for _BACnetTagPayloadBitString
func NewBACnetTagPayloadBitString(unusedBits uint8, data []bool, unused []bool, actualLength uint32) *_BACnetTagPayloadBitString {
	return &_BACnetTagPayloadBitString{UnusedBits: unusedBits, Data: data, Unused: unused, ActualLength: actualLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagPayloadBitStringBuilder is a builder for BACnetTagPayloadBitString
type BACnetTagPayloadBitStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unusedBits uint8, data []bool, unused []bool) BACnetTagPayloadBitStringBuilder
	// WithUnusedBits adds UnusedBits (property field)
	WithUnusedBits(uint8) BACnetTagPayloadBitStringBuilder
	// WithData adds Data (property field)
	WithData(...bool) BACnetTagPayloadBitStringBuilder
	// WithUnused adds Unused (property field)
	WithUnused(...bool) BACnetTagPayloadBitStringBuilder
	// Build builds the BACnetTagPayloadBitString or returns an error if something is wrong
	Build() (BACnetTagPayloadBitString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagPayloadBitString
}

// NewBACnetTagPayloadBitStringBuilder() creates a BACnetTagPayloadBitStringBuilder
func NewBACnetTagPayloadBitStringBuilder() BACnetTagPayloadBitStringBuilder {
	return &_BACnetTagPayloadBitStringBuilder{_BACnetTagPayloadBitString: new(_BACnetTagPayloadBitString)}
}

type _BACnetTagPayloadBitStringBuilder struct {
	*_BACnetTagPayloadBitString

	err *utils.MultiError
}

var _ (BACnetTagPayloadBitStringBuilder) = (*_BACnetTagPayloadBitStringBuilder)(nil)

func (m *_BACnetTagPayloadBitStringBuilder) WithMandatoryFields(unusedBits uint8, data []bool, unused []bool) BACnetTagPayloadBitStringBuilder {
	return m.WithUnusedBits(unusedBits).WithData(data...).WithUnused(unused...)
}

func (m *_BACnetTagPayloadBitStringBuilder) WithUnusedBits(unusedBits uint8) BACnetTagPayloadBitStringBuilder {
	m.UnusedBits = unusedBits
	return m
}

func (m *_BACnetTagPayloadBitStringBuilder) WithData(data ...bool) BACnetTagPayloadBitStringBuilder {
	m.Data = data
	return m
}

func (m *_BACnetTagPayloadBitStringBuilder) WithUnused(unused ...bool) BACnetTagPayloadBitStringBuilder {
	m.Unused = unused
	return m
}

func (m *_BACnetTagPayloadBitStringBuilder) Build() (BACnetTagPayloadBitString, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetTagPayloadBitString.deepCopy(), nil
}

func (m *_BACnetTagPayloadBitStringBuilder) MustBuild() BACnetTagPayloadBitString {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetTagPayloadBitStringBuilder) DeepCopy() any {
	return m.CreateBACnetTagPayloadBitStringBuilder()
}

// CreateBACnetTagPayloadBitStringBuilder creates a BACnetTagPayloadBitStringBuilder
func (m *_BACnetTagPayloadBitString) CreateBACnetTagPayloadBitStringBuilder() BACnetTagPayloadBitStringBuilder {
	if m == nil {
		return NewBACnetTagPayloadBitStringBuilder()
	}
	return &_BACnetTagPayloadBitStringBuilder{_BACnetTagPayloadBitString: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadBitString) GetUnusedBits() uint8 {
	return m.UnusedBits
}

func (m *_BACnetTagPayloadBitString) GetData() []bool {
	return m.Data
}

func (m *_BACnetTagPayloadBitString) GetUnused() []bool {
	return m.Unused
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadBitString(structType any) BACnetTagPayloadBitString {
	if casted, ok := structType.(BACnetTagPayloadBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadBitString) GetTypeName() string {
	return "BACnetTagPayloadBitString"
}

func (m *_BACnetTagPayloadBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (unusedBits)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 1 * uint16(len(m.Data))
	}

	// Array field
	if len(m.Unused) > 0 {
		lengthInBits += 1 * uint16(len(m.Unused))
	}

	return lengthInBits
}

func (m *_BACnetTagPayloadBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadBitStringParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadBitString, error) {
	return BACnetTagPayloadBitStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadBitStringParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadBitString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadBitString, error) {
		return BACnetTagPayloadBitStringParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetTagPayloadBitStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadBitString, error) {
	v, err := (&_BACnetTagPayloadBitString{ActualLength: actualLength}).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadBitString) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetTagPayloadBitString BACnetTagPayloadBitString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unusedBits, err := ReadSimpleField(ctx, "unusedBits", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unusedBits' field"))
	}
	m.UnusedBits = unusedBits

	data, err := ReadCountArrayField[bool](ctx, "data", ReadBoolean(readBuffer), uint64(int32((int32((int32(actualLength)-int32(int32(1))))*int32(int32(8))))-int32(unusedBits)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	unused, err := ReadCountArrayField[bool](ctx, "unused", ReadBoolean(readBuffer), uint64(unusedBits))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unused' field"))
	}
	m.Unused = unused

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadBitString")
	}

	return m, nil
}

func (m *_BACnetTagPayloadBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadBitString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadBitString")
	}

	if err := WriteSimpleField[uint8](ctx, "unusedBits", m.GetUnusedBits(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'unusedBits' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "unused", m.GetUnused(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'unused' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadBitString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadBitString")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadBitString) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadBitString) IsBACnetTagPayloadBitString() {}

func (m *_BACnetTagPayloadBitString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadBitString) deepCopy() *_BACnetTagPayloadBitString {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadBitStringCopy := &_BACnetTagPayloadBitString{
		m.UnusedBits,
		utils.DeepCopySlice[bool, bool](m.Data),
		utils.DeepCopySlice[bool, bool](m.Unused),
		m.ActualLength,
	}
	return _BACnetTagPayloadBitStringCopy
}

func (m *_BACnetTagPayloadBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
