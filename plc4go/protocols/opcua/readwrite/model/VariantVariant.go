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

// VariantVariant is the corresponding interface of VariantVariant
type VariantVariant interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []Variant
	// IsVariantVariant is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantVariant()
	// CreateBuilder creates a VariantVariantBuilder
	CreateVariantVariantBuilder() VariantVariantBuilder
}

// _VariantVariant is the data-structure of this message
type _VariantVariant struct {
	VariantContract
	ArrayLength *int32
	Value       []Variant
}

var _ VariantVariant = (*_VariantVariant)(nil)
var _ VariantRequirements = (*_VariantVariant)(nil)

// NewVariantVariant factory function for _VariantVariant
func NewVariantVariant(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []Variant) *_VariantVariant {
	_result := &_VariantVariant{
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

// VariantVariantBuilder is a builder for VariantVariant
type VariantVariantBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []Variant) VariantVariantBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantVariantBuilder
	// WithValue adds Value (property field)
	WithValue(...Variant) VariantVariantBuilder
	// Build builds the VariantVariant or returns an error if something is wrong
	Build() (VariantVariant, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantVariant
}

// NewVariantVariantBuilder() creates a VariantVariantBuilder
func NewVariantVariantBuilder() VariantVariantBuilder {
	return &_VariantVariantBuilder{_VariantVariant: new(_VariantVariant)}
}

type _VariantVariantBuilder struct {
	*_VariantVariant

	err *utils.MultiError
}

var _ (VariantVariantBuilder) = (*_VariantVariantBuilder)(nil)

func (m *_VariantVariantBuilder) WithMandatoryFields(value []Variant) VariantVariantBuilder {
	return m.WithValue(value...)
}

func (m *_VariantVariantBuilder) WithOptionalArrayLength(arrayLength int32) VariantVariantBuilder {
	m.ArrayLength = &arrayLength
	return m
}

func (m *_VariantVariantBuilder) WithValue(value ...Variant) VariantVariantBuilder {
	m.Value = value
	return m
}

func (m *_VariantVariantBuilder) Build() (VariantVariant, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._VariantVariant.deepCopy(), nil
}

func (m *_VariantVariantBuilder) MustBuild() VariantVariant {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_VariantVariantBuilder) DeepCopy() any {
	return m.CreateVariantVariantBuilder()
}

// CreateVariantVariantBuilder creates a VariantVariantBuilder
func (m *_VariantVariant) CreateVariantVariantBuilder() VariantVariantBuilder {
	if m == nil {
		return NewVariantVariantBuilder()
	}
	return &_VariantVariantBuilder{_VariantVariant: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantVariant) GetVariantType() uint8 {
	return uint8(24)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantVariant) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantVariant) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantVariant) GetValue() []Variant {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantVariant(structType any) VariantVariant {
	if casted, ok := structType.(VariantVariant); ok {
		return casted
	}
	if casted, ok := structType.(*VariantVariant); ok {
		return *casted
	}
	return nil
}

func (m *_VariantVariant) GetTypeName() string {
	return "VariantVariant"
}

func (m *_VariantVariant) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_VariantVariant) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantVariant) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantVariant VariantVariant, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantVariant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantVariant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[Variant](ctx, "value", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantVariant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantVariant")
	}

	return m, nil
}

func (m *_VariantVariant) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantVariant) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantVariant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantVariant")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "value", m.GetValue(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantVariant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantVariant")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantVariant) IsVariantVariant() {}

func (m *_VariantVariant) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantVariant) deepCopy() *_VariantVariant {
	if m == nil {
		return nil
	}
	_VariantVariantCopy := &_VariantVariant{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[Variant, Variant](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantVariantCopy
}

func (m *_VariantVariant) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
