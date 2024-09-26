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

// VariantInt64 is the corresponding interface of VariantInt64
type VariantInt64 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []int64
	// IsVariantInt64 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantInt64()
	// CreateBuilder creates a VariantInt64Builder
	CreateVariantInt64Builder() VariantInt64Builder
}

// _VariantInt64 is the data-structure of this message
type _VariantInt64 struct {
	VariantContract
	ArrayLength *int32
	Value       []int64
}

var _ VariantInt64 = (*_VariantInt64)(nil)
var _ VariantRequirements = (*_VariantInt64)(nil)

// NewVariantInt64 factory function for _VariantInt64
func NewVariantInt64(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []int64) *_VariantInt64 {
	_result := &_VariantInt64{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantInt64Builder is a builder for VariantInt64
type VariantInt64Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []int64) VariantInt64Builder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantInt64Builder
	// WithValue adds Value (property field)
	WithValue(...int64) VariantInt64Builder
	// Build builds the VariantInt64 or returns an error if something is wrong
	Build() (VariantInt64, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantInt64
}

// NewVariantInt64Builder() creates a VariantInt64Builder
func NewVariantInt64Builder() VariantInt64Builder {
	return &_VariantInt64Builder{_VariantInt64: new(_VariantInt64)}
}

type _VariantInt64Builder struct {
	*_VariantInt64

	err *utils.MultiError
}

var _ (VariantInt64Builder) = (*_VariantInt64Builder)(nil)

func (m *_VariantInt64Builder) WithMandatoryFields(value []int64) VariantInt64Builder {
	return m.WithValue(value...)
}

func (m *_VariantInt64Builder) WithOptionalArrayLength(arrayLength int32) VariantInt64Builder {
	m.ArrayLength = &arrayLength
	return m
}

func (m *_VariantInt64Builder) WithValue(value ...int64) VariantInt64Builder {
	m.Value = value
	return m
}

func (m *_VariantInt64Builder) Build() (VariantInt64, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._VariantInt64.deepCopy(), nil
}

func (m *_VariantInt64Builder) MustBuild() VariantInt64 {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_VariantInt64Builder) DeepCopy() any {
	return m.CreateVariantInt64Builder()
}

// CreateVariantInt64Builder creates a VariantInt64Builder
func (m *_VariantInt64) CreateVariantInt64Builder() VariantInt64Builder {
	if m == nil {
		return NewVariantInt64Builder()
	}
	return &_VariantInt64Builder{_VariantInt64: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantInt64) GetVariantType() uint8 {
	return uint8(8)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantInt64) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantInt64) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantInt64) GetValue() []int64 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantInt64(structType any) VariantInt64 {
	if casted, ok := structType.(VariantInt64); ok {
		return casted
	}
	if casted, ok := structType.(*VariantInt64); ok {
		return *casted
	}
	return nil
}

func (m *_VariantInt64) GetTypeName() string {
	return "VariantInt64"
}

func (m *_VariantInt64) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 64 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantInt64) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantInt64) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantInt64 VariantInt64, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantInt64"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantInt64")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[int64](ctx, "value", ReadSignedLong(readBuffer, uint8(64)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantInt64"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantInt64")
	}

	return m, nil
}

func (m *_VariantInt64) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantInt64) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantInt64"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantInt64")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantInt64"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantInt64")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantInt64) IsVariantInt64() {}

func (m *_VariantInt64) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantInt64) deepCopy() *_VariantInt64 {
	if m == nil {
		return nil
	}
	_VariantInt64Copy := &_VariantInt64{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[int64, int64](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantInt64Copy
}

func (m *_VariantInt64) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
