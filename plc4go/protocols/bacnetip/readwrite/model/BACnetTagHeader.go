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

// BACnetTagHeader is the corresponding interface of BACnetTagHeader
type BACnetTagHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTagNumber returns TagNumber (property field)
	GetTagNumber() uint8
	// GetTagClass returns TagClass (property field)
	GetTagClass() TagClass
	// GetLengthValueType returns LengthValueType (property field)
	GetLengthValueType() uint8
	// GetExtTagNumber returns ExtTagNumber (property field)
	GetExtTagNumber() *uint8
	// GetExtLength returns ExtLength (property field)
	GetExtLength() *uint8
	// GetExtExtLength returns ExtExtLength (property field)
	GetExtExtLength() *uint16
	// GetExtExtExtLength returns ExtExtExtLength (property field)
	GetExtExtExtLength() *uint32
	// GetActualTagNumber returns ActualTagNumber (virtual field)
	GetActualTagNumber() uint8
	// GetIsBoolean returns IsBoolean (virtual field)
	GetIsBoolean() bool
	// GetIsConstructed returns IsConstructed (virtual field)
	GetIsConstructed() bool
	// GetIsPrimitiveAndNotBoolean returns IsPrimitiveAndNotBoolean (virtual field)
	GetIsPrimitiveAndNotBoolean() bool
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
	// IsBACnetTagHeader is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagHeader()
	// CreateBuilder creates a BACnetTagHeaderBuilder
	CreateBACnetTagHeaderBuilder() BACnetTagHeaderBuilder
}

// _BACnetTagHeader is the data-structure of this message
type _BACnetTagHeader struct {
	TagNumber       uint8
	TagClass        TagClass
	LengthValueType uint8
	ExtTagNumber    *uint8
	ExtLength       *uint8
	ExtExtLength    *uint16
	ExtExtExtLength *uint32
}

var _ BACnetTagHeader = (*_BACnetTagHeader)(nil)

// NewBACnetTagHeader factory function for _BACnetTagHeader
func NewBACnetTagHeader(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *_BACnetTagHeader {
	return &_BACnetTagHeader{TagNumber: tagNumber, TagClass: tagClass, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, ExtExtLength: extExtLength, ExtExtExtLength: extExtExtLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagHeaderBuilder is a builder for BACnetTagHeader
type BACnetTagHeaderBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tagNumber uint8, tagClass TagClass, lengthValueType uint8) BACnetTagHeaderBuilder
	// WithTagNumber adds TagNumber (property field)
	WithTagNumber(uint8) BACnetTagHeaderBuilder
	// WithTagClass adds TagClass (property field)
	WithTagClass(TagClass) BACnetTagHeaderBuilder
	// WithLengthValueType adds LengthValueType (property field)
	WithLengthValueType(uint8) BACnetTagHeaderBuilder
	// WithExtTagNumber adds ExtTagNumber (property field)
	WithOptionalExtTagNumber(uint8) BACnetTagHeaderBuilder
	// WithExtLength adds ExtLength (property field)
	WithOptionalExtLength(uint8) BACnetTagHeaderBuilder
	// WithExtExtLength adds ExtExtLength (property field)
	WithOptionalExtExtLength(uint16) BACnetTagHeaderBuilder
	// WithExtExtExtLength adds ExtExtExtLength (property field)
	WithOptionalExtExtExtLength(uint32) BACnetTagHeaderBuilder
	// Build builds the BACnetTagHeader or returns an error if something is wrong
	Build() (BACnetTagHeader, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagHeader
}

// NewBACnetTagHeaderBuilder() creates a BACnetTagHeaderBuilder
func NewBACnetTagHeaderBuilder() BACnetTagHeaderBuilder {
	return &_BACnetTagHeaderBuilder{_BACnetTagHeader: new(_BACnetTagHeader)}
}

type _BACnetTagHeaderBuilder struct {
	*_BACnetTagHeader

	err *utils.MultiError
}

var _ (BACnetTagHeaderBuilder) = (*_BACnetTagHeaderBuilder)(nil)

func (m *_BACnetTagHeaderBuilder) WithMandatoryFields(tagNumber uint8, tagClass TagClass, lengthValueType uint8) BACnetTagHeaderBuilder {
	return m.WithTagNumber(tagNumber).WithTagClass(tagClass).WithLengthValueType(lengthValueType)
}

func (m *_BACnetTagHeaderBuilder) WithTagNumber(tagNumber uint8) BACnetTagHeaderBuilder {
	m.TagNumber = tagNumber
	return m
}

func (m *_BACnetTagHeaderBuilder) WithTagClass(tagClass TagClass) BACnetTagHeaderBuilder {
	m.TagClass = tagClass
	return m
}

func (m *_BACnetTagHeaderBuilder) WithLengthValueType(lengthValueType uint8) BACnetTagHeaderBuilder {
	m.LengthValueType = lengthValueType
	return m
}

func (m *_BACnetTagHeaderBuilder) WithOptionalExtTagNumber(extTagNumber uint8) BACnetTagHeaderBuilder {
	m.ExtTagNumber = &extTagNumber
	return m
}

func (m *_BACnetTagHeaderBuilder) WithOptionalExtLength(extLength uint8) BACnetTagHeaderBuilder {
	m.ExtLength = &extLength
	return m
}

func (m *_BACnetTagHeaderBuilder) WithOptionalExtExtLength(extExtLength uint16) BACnetTagHeaderBuilder {
	m.ExtExtLength = &extExtLength
	return m
}

func (m *_BACnetTagHeaderBuilder) WithOptionalExtExtExtLength(extExtExtLength uint32) BACnetTagHeaderBuilder {
	m.ExtExtExtLength = &extExtExtLength
	return m
}

func (m *_BACnetTagHeaderBuilder) Build() (BACnetTagHeader, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetTagHeader.deepCopy(), nil
}

func (m *_BACnetTagHeaderBuilder) MustBuild() BACnetTagHeader {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetTagHeaderBuilder) DeepCopy() any {
	return m.CreateBACnetTagHeaderBuilder()
}

// CreateBACnetTagHeaderBuilder creates a BACnetTagHeaderBuilder
func (m *_BACnetTagHeader) CreateBACnetTagHeaderBuilder() BACnetTagHeaderBuilder {
	if m == nil {
		return NewBACnetTagHeaderBuilder()
	}
	return &_BACnetTagHeaderBuilder{_BACnetTagHeader: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagHeader) GetTagNumber() uint8 {
	return m.TagNumber
}

func (m *_BACnetTagHeader) GetTagClass() TagClass {
	return m.TagClass
}

func (m *_BACnetTagHeader) GetLengthValueType() uint8 {
	return m.LengthValueType
}

func (m *_BACnetTagHeader) GetExtTagNumber() *uint8 {
	return m.ExtTagNumber
}

func (m *_BACnetTagHeader) GetExtLength() *uint8 {
	return m.ExtLength
}

func (m *_BACnetTagHeader) GetExtExtLength() *uint16 {
	return m.ExtExtLength
}

func (m *_BACnetTagHeader) GetExtExtExtLength() *uint32 {
	return m.ExtExtExtLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagHeader) GetActualTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.GetExtTagNumber()
	_ = extTagNumber
	extLength := m.GetExtLength()
	_ = extLength
	extExtLength := m.GetExtExtLength()
	_ = extExtLength
	extExtExtLength := m.GetExtExtExtLength()
	_ = extExtExtLength
	return uint8(utils.InlineIf(bool((m.GetTagNumber()) < (15)), func() any { return uint8(m.GetTagNumber()) }, func() any { return uint8((*m.GetExtTagNumber())) }).(uint8))
}

func (m *_BACnetTagHeader) GetIsBoolean() bool {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.GetExtTagNumber()
	_ = extTagNumber
	extLength := m.GetExtLength()
	_ = extLength
	extExtLength := m.GetExtExtLength()
	_ = extExtLength
	extExtExtLength := m.GetExtExtExtLength()
	_ = extExtExtLength
	return bool(bool(bool((m.GetTagNumber()) == (1))) && bool(bool((m.GetTagClass()) == (TagClass_APPLICATION_TAGS))))
}

func (m *_BACnetTagHeader) GetIsConstructed() bool {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.GetExtTagNumber()
	_ = extTagNumber
	extLength := m.GetExtLength()
	_ = extLength
	extExtLength := m.GetExtExtLength()
	_ = extExtLength
	extExtExtLength := m.GetExtExtExtLength()
	_ = extExtExtLength
	return bool(bool(bool((m.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((m.GetLengthValueType()) == (6))))
}

func (m *_BACnetTagHeader) GetIsPrimitiveAndNotBoolean() bool {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.GetExtTagNumber()
	_ = extTagNumber
	extLength := m.GetExtLength()
	_ = extLength
	extExtLength := m.GetExtExtLength()
	_ = extExtLength
	extExtExtLength := m.GetExtExtExtLength()
	_ = extExtExtLength
	return bool(bool(!(m.GetIsConstructed())) && bool(!(m.GetIsBoolean())))
}

func (m *_BACnetTagHeader) GetActualLength() uint32 {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.GetExtTagNumber()
	_ = extTagNumber
	extLength := m.GetExtLength()
	_ = extLength
	extExtLength := m.GetExtExtLength()
	_ = extExtLength
	extExtExtLength := m.GetExtExtExtLength()
	_ = extExtExtLength
	return uint32(utils.InlineIf(bool(bool((m.GetLengthValueType()) == (5))) && bool(bool((*m.GetExtLength()) == (255))), func() any { return uint32((*m.GetExtExtExtLength())) }, func() any {
		return uint32((utils.InlineIf(bool(bool((m.GetLengthValueType()) == (5))) && bool(bool((*m.GetExtLength()) == (254))), func() any { return uint32((*m.GetExtExtLength())) }, func() any {
			return uint32((utils.InlineIf(bool((m.GetLengthValueType()) == (5)), func() any { return uint32((*m.GetExtLength())) }, func() any { return uint32(m.GetLengthValueType()) }).(uint32)))
		}).(uint32)))
	}).(uint32))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagHeader(structType any) BACnetTagHeader {
	if casted, ok := structType.(BACnetTagHeader); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagHeader); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagHeader) GetTypeName() string {
	return "BACnetTagHeader"
}

func (m *_BACnetTagHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (tagNumber)
	lengthInBits += 4

	// Simple field (tagClass)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Optional Field (extExtLength)
	if m.ExtExtLength != nil {
		lengthInBits += 16
	}

	// Optional Field (extExtExtLength)
	if m.ExtExtExtLength != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagHeaderParse(ctx context.Context, theBytes []byte) (BACnetTagHeader, error) {
	return BACnetTagHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagHeaderParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagHeader, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagHeader, error) {
		return BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTagHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagHeader, error) {
	v, err := (&_BACnetTagHeader{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagHeader) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTagHeader BACnetTagHeader, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tagNumber, err := ReadSimpleField(ctx, "tagNumber", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tagNumber' field"))
	}
	m.TagNumber = tagNumber

	tagClass, err := ReadEnumField[TagClass](ctx, "tagClass", "TagClass", ReadEnum(TagClassByValue, ReadUnsignedByte(readBuffer, uint8(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tagClass' field"))
	}
	m.TagClass = tagClass

	lengthValueType, err := ReadSimpleField(ctx, "lengthValueType", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lengthValueType' field"))
	}
	m.LengthValueType = lengthValueType

	var extTagNumber *uint8
	extTagNumber, err = ReadOptionalField[uint8](ctx, "extTagNumber", ReadUnsignedByte(readBuffer, uint8(8)), bool((tagNumber) == (15)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extTagNumber' field"))
	}
	m.ExtTagNumber = extTagNumber

	actualTagNumber, err := ReadVirtualField[uint8](ctx, "actualTagNumber", (*uint8)(nil), utils.InlineIf(bool((tagNumber) < (15)), func() any { return uint8(tagNumber) }, func() any { return uint8((*extTagNumber)) }).(uint8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualTagNumber' field"))
	}
	_ = actualTagNumber

	isBoolean, err := ReadVirtualField[bool](ctx, "isBoolean", (*bool)(nil), bool(bool((tagNumber) == (1))) && bool(bool((tagClass) == (TagClass_APPLICATION_TAGS))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isBoolean' field"))
	}
	_ = isBoolean

	isConstructed, err := ReadVirtualField[bool](ctx, "isConstructed", (*bool)(nil), bool(bool((tagClass) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((lengthValueType) == (6))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isConstructed' field"))
	}
	_ = isConstructed

	isPrimitiveAndNotBoolean, err := ReadVirtualField[bool](ctx, "isPrimitiveAndNotBoolean", (*bool)(nil), bool(!(isConstructed)) && bool(!(isBoolean)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPrimitiveAndNotBoolean' field"))
	}
	_ = isPrimitiveAndNotBoolean

	var extLength *uint8
	extLength, err = ReadOptionalField[uint8](ctx, "extLength", ReadUnsignedByte(readBuffer, uint8(8)), bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extLength' field"))
	}
	m.ExtLength = extLength

	var extExtLength *uint16
	extExtLength, err = ReadOptionalField[uint16](ctx, "extExtLength", ReadUnsignedShort(readBuffer, uint8(16)), bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (254))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extExtLength' field"))
	}
	m.ExtExtLength = extExtLength

	var extExtExtLength *uint32
	extExtExtLength, err = ReadOptionalField[uint32](ctx, "extExtExtLength", ReadUnsignedInt(readBuffer, uint8(32)), bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (255))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extExtExtLength' field"))
	}
	m.ExtExtExtLength = extExtExtLength

	actualLength, err := ReadVirtualField[uint32](ctx, "actualLength", (*uint32)(nil), utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))), func() any { return uint32((*extExtExtLength)) }, func() any {
		return uint32((utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))), func() any { return uint32((*extExtLength)) }, func() any {
			return uint32((utils.InlineIf(bool((lengthValueType) == (5)), func() any { return uint32((*extLength)) }, func() any { return uint32(lengthValueType) }).(uint32)))
		}).(uint32)))
	}).(uint32))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualLength' field"))
	}
	_ = actualLength

	if closeErr := readBuffer.CloseContext("BACnetTagHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagHeader")
	}

	return m, nil
}

func (m *_BACnetTagHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagHeader")
	}

	if err := WriteSimpleField[uint8](ctx, "tagNumber", m.GetTagNumber(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
		return errors.Wrap(err, "Error serializing 'tagNumber' field")
	}

	if err := WriteSimpleEnumField[TagClass](ctx, "tagClass", "TagClass", m.GetTagClass(), WriteEnum[TagClass, uint8](TagClass.GetValue, TagClass.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 1))); err != nil {
		return errors.Wrap(err, "Error serializing 'tagClass' field")
	}

	if err := WriteSimpleField[uint8](ctx, "lengthValueType", m.GetLengthValueType(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'lengthValueType' field")
	}

	if err := WriteOptionalField[uint8](ctx, "extTagNumber", m.GetExtTagNumber(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
		return errors.Wrap(err, "Error serializing 'extTagNumber' field")
	}
	// Virtual field
	actualTagNumber := m.GetActualTagNumber()
	_ = actualTagNumber
	if _actualTagNumberErr := writeBuffer.WriteVirtual(ctx, "actualTagNumber", m.GetActualTagNumber()); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	isBoolean := m.GetIsBoolean()
	_ = isBoolean
	if _isBooleanErr := writeBuffer.WriteVirtual(ctx, "isBoolean", m.GetIsBoolean()); _isBooleanErr != nil {
		return errors.Wrap(_isBooleanErr, "Error serializing 'isBoolean' field")
	}
	// Virtual field
	isConstructed := m.GetIsConstructed()
	_ = isConstructed
	if _isConstructedErr := writeBuffer.WriteVirtual(ctx, "isConstructed", m.GetIsConstructed()); _isConstructedErr != nil {
		return errors.Wrap(_isConstructedErr, "Error serializing 'isConstructed' field")
	}
	// Virtual field
	isPrimitiveAndNotBoolean := m.GetIsPrimitiveAndNotBoolean()
	_ = isPrimitiveAndNotBoolean
	if _isPrimitiveAndNotBooleanErr := writeBuffer.WriteVirtual(ctx, "isPrimitiveAndNotBoolean", m.GetIsPrimitiveAndNotBoolean()); _isPrimitiveAndNotBooleanErr != nil {
		return errors.Wrap(_isPrimitiveAndNotBooleanErr, "Error serializing 'isPrimitiveAndNotBoolean' field")
	}

	if err := WriteOptionalField[uint8](ctx, "extLength", m.GetExtLength(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
		return errors.Wrap(err, "Error serializing 'extLength' field")
	}

	if err := WriteOptionalField[uint16](ctx, "extExtLength", m.GetExtExtLength(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
		return errors.Wrap(err, "Error serializing 'extExtLength' field")
	}

	if err := WriteOptionalField[uint32](ctx, "extExtExtLength", m.GetExtExtExtLength(), WriteUnsignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'extExtExtLength' field")
	}
	// Virtual field
	actualLength := m.GetActualLength()
	_ = actualLength
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagHeader")
	}
	return nil
}

func (m *_BACnetTagHeader) IsBACnetTagHeader() {}

func (m *_BACnetTagHeader) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagHeader) deepCopy() *_BACnetTagHeader {
	if m == nil {
		return nil
	}
	_BACnetTagHeaderCopy := &_BACnetTagHeader{
		m.TagNumber,
		m.TagClass,
		m.LengthValueType,
		utils.CopyPtr[uint8](m.ExtTagNumber),
		utils.CopyPtr[uint8](m.ExtLength),
		utils.CopyPtr[uint16](m.ExtExtLength),
		utils.CopyPtr[uint32](m.ExtExtExtLength),
	}
	return _BACnetTagHeaderCopy
}

func (m *_BACnetTagHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
