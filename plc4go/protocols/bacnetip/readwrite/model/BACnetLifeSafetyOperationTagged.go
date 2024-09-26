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

// BACnetLifeSafetyOperationTagged is the corresponding interface of BACnetLifeSafetyOperationTagged
type BACnetLifeSafetyOperationTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLifeSafetyOperation
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetLifeSafetyOperationTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLifeSafetyOperationTagged()
	// CreateBuilder creates a BACnetLifeSafetyOperationTaggedBuilder
	CreateBACnetLifeSafetyOperationTaggedBuilder() BACnetLifeSafetyOperationTaggedBuilder
}

// _BACnetLifeSafetyOperationTagged is the data-structure of this message
type _BACnetLifeSafetyOperationTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLifeSafetyOperation
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLifeSafetyOperationTagged = (*_BACnetLifeSafetyOperationTagged)(nil)

// NewBACnetLifeSafetyOperationTagged factory function for _BACnetLifeSafetyOperationTagged
func NewBACnetLifeSafetyOperationTagged(header BACnetTagHeader, value BACnetLifeSafetyOperation, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLifeSafetyOperationTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLifeSafetyOperationTagged must not be nil")
	}
	return &_BACnetLifeSafetyOperationTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLifeSafetyOperationTaggedBuilder is a builder for BACnetLifeSafetyOperationTagged
type BACnetLifeSafetyOperationTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetLifeSafetyOperation, proprietaryValue uint32) BACnetLifeSafetyOperationTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetLifeSafetyOperationTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLifeSafetyOperationTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetLifeSafetyOperation) BACnetLifeSafetyOperationTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetLifeSafetyOperationTaggedBuilder
	// Build builds the BACnetLifeSafetyOperationTagged or returns an error if something is wrong
	Build() (BACnetLifeSafetyOperationTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLifeSafetyOperationTagged
}

// NewBACnetLifeSafetyOperationTaggedBuilder() creates a BACnetLifeSafetyOperationTaggedBuilder
func NewBACnetLifeSafetyOperationTaggedBuilder() BACnetLifeSafetyOperationTaggedBuilder {
	return &_BACnetLifeSafetyOperationTaggedBuilder{_BACnetLifeSafetyOperationTagged: new(_BACnetLifeSafetyOperationTagged)}
}

type _BACnetLifeSafetyOperationTaggedBuilder struct {
	*_BACnetLifeSafetyOperationTagged

	err *utils.MultiError
}

var _ (BACnetLifeSafetyOperationTaggedBuilder) = (*_BACnetLifeSafetyOperationTaggedBuilder)(nil)

func (m *_BACnetLifeSafetyOperationTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetLifeSafetyOperation, proprietaryValue uint32) BACnetLifeSafetyOperationTaggedBuilder {
	return m.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetLifeSafetyOperationTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLifeSafetyOperationTaggedBuilder {
	builder := builderSupplier(m.Header.CreateBACnetTagHeaderBuilder())
	var err error
	m.Header, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) WithValue(value BACnetLifeSafetyOperation) BACnetLifeSafetyOperationTaggedBuilder {
	m.Value = value
	return m
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetLifeSafetyOperationTaggedBuilder {
	m.ProprietaryValue = proprietaryValue
	return m
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) Build() (BACnetLifeSafetyOperationTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetLifeSafetyOperationTagged.deepCopy(), nil
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) MustBuild() BACnetLifeSafetyOperationTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetLifeSafetyOperationTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetLifeSafetyOperationTaggedBuilder()
}

// CreateBACnetLifeSafetyOperationTaggedBuilder creates a BACnetLifeSafetyOperationTaggedBuilder
func (m *_BACnetLifeSafetyOperationTagged) CreateBACnetLifeSafetyOperationTaggedBuilder() BACnetLifeSafetyOperationTaggedBuilder {
	if m == nil {
		return NewBACnetLifeSafetyOperationTaggedBuilder()
	}
	return &_BACnetLifeSafetyOperationTaggedBuilder{_BACnetLifeSafetyOperationTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLifeSafetyOperationTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLifeSafetyOperationTagged) GetValue() BACnetLifeSafetyOperation {
	return m.Value
}

func (m *_BACnetLifeSafetyOperationTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLifeSafetyOperationTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLifeSafetyOperationTagged(structType any) BACnetLifeSafetyOperationTagged {
	if casted, ok := structType.(BACnetLifeSafetyOperationTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLifeSafetyOperationTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLifeSafetyOperationTagged) GetTypeName() string {
	return "BACnetLifeSafetyOperationTagged"
}

func (m *_BACnetLifeSafetyOperationTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetLifeSafetyOperationTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLifeSafetyOperationTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLifeSafetyOperationTagged, error) {
	return BACnetLifeSafetyOperationTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLifeSafetyOperationTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLifeSafetyOperationTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLifeSafetyOperationTagged, error) {
		return BACnetLifeSafetyOperationTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLifeSafetyOperationTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLifeSafetyOperationTagged, error) {
	v, err := (&_BACnetLifeSafetyOperationTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLifeSafetyOperationTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLifeSafetyOperationTagged BACnetLifeSafetyOperationTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLifeSafetyOperationTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLifeSafetyOperationTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetLifeSafetyOperation](ctx, "value", readBuffer, EnsureType[BACnetLifeSafetyOperation](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetLifeSafetyOperationTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLifeSafetyOperationTagged")
	}

	return m, nil
}

func (m *_BACnetLifeSafetyOperationTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLifeSafetyOperationTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLifeSafetyOperationTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLifeSafetyOperationTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLifeSafetyOperation](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLifeSafetyOperationTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLifeSafetyOperationTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLifeSafetyOperationTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLifeSafetyOperationTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLifeSafetyOperationTagged) IsBACnetLifeSafetyOperationTagged() {}

func (m *_BACnetLifeSafetyOperationTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLifeSafetyOperationTagged) deepCopy() *_BACnetLifeSafetyOperationTagged {
	if m == nil {
		return nil
	}
	_BACnetLifeSafetyOperationTaggedCopy := &_BACnetLifeSafetyOperationTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetLifeSafetyOperationTaggedCopy
}

func (m *_BACnetLifeSafetyOperationTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
