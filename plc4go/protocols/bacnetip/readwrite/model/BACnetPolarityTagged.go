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

// BACnetPolarityTagged is the corresponding interface of BACnetPolarityTagged
type BACnetPolarityTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetPolarity
	// IsBACnetPolarityTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPolarityTagged()
	// CreateBuilder creates a BACnetPolarityTaggedBuilder
	CreateBACnetPolarityTaggedBuilder() BACnetPolarityTaggedBuilder
}

// _BACnetPolarityTagged is the data-structure of this message
type _BACnetPolarityTagged struct {
	Header BACnetTagHeader
	Value  BACnetPolarity

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetPolarityTagged = (*_BACnetPolarityTagged)(nil)

// NewBACnetPolarityTagged factory function for _BACnetPolarityTagged
func NewBACnetPolarityTagged(header BACnetTagHeader, value BACnetPolarity, tagNumber uint8, tagClass TagClass) *_BACnetPolarityTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetPolarityTagged must not be nil")
	}
	return &_BACnetPolarityTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPolarityTaggedBuilder is a builder for BACnetPolarityTagged
type BACnetPolarityTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetPolarity) BACnetPolarityTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetPolarityTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetPolarityTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetPolarity) BACnetPolarityTaggedBuilder
	// Build builds the BACnetPolarityTagged or returns an error if something is wrong
	Build() (BACnetPolarityTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPolarityTagged
}

// NewBACnetPolarityTaggedBuilder() creates a BACnetPolarityTaggedBuilder
func NewBACnetPolarityTaggedBuilder() BACnetPolarityTaggedBuilder {
	return &_BACnetPolarityTaggedBuilder{_BACnetPolarityTagged: new(_BACnetPolarityTagged)}
}

type _BACnetPolarityTaggedBuilder struct {
	*_BACnetPolarityTagged

	err *utils.MultiError
}

var _ (BACnetPolarityTaggedBuilder) = (*_BACnetPolarityTaggedBuilder)(nil)

func (m *_BACnetPolarityTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetPolarity) BACnetPolarityTaggedBuilder {
	return m.WithHeader(header).WithValue(value)
}

func (m *_BACnetPolarityTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetPolarityTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetPolarityTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetPolarityTaggedBuilder {
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

func (m *_BACnetPolarityTaggedBuilder) WithValue(value BACnetPolarity) BACnetPolarityTaggedBuilder {
	m.Value = value
	return m
}

func (m *_BACnetPolarityTaggedBuilder) Build() (BACnetPolarityTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPolarityTagged.deepCopy(), nil
}

func (m *_BACnetPolarityTaggedBuilder) MustBuild() BACnetPolarityTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPolarityTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetPolarityTaggedBuilder()
}

// CreateBACnetPolarityTaggedBuilder creates a BACnetPolarityTaggedBuilder
func (m *_BACnetPolarityTagged) CreateBACnetPolarityTaggedBuilder() BACnetPolarityTaggedBuilder {
	if m == nil {
		return NewBACnetPolarityTaggedBuilder()
	}
	return &_BACnetPolarityTaggedBuilder{_BACnetPolarityTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPolarityTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetPolarityTagged) GetValue() BACnetPolarity {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPolarityTagged(structType any) BACnetPolarityTagged {
	if casted, ok := structType.(BACnetPolarityTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPolarityTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPolarityTagged) GetTypeName() string {
	return "BACnetPolarityTagged"
}

func (m *_BACnetPolarityTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetPolarityTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPolarityTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetPolarityTagged, error) {
	return BACnetPolarityTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetPolarityTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPolarityTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPolarityTagged, error) {
		return BACnetPolarityTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetPolarityTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetPolarityTagged, error) {
	v, err := (&_BACnetPolarityTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPolarityTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetPolarityTagged BACnetPolarityTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPolarityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPolarityTagged")
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

	value, err := ReadManualField[BACnetPolarity](ctx, "value", readBuffer, EnsureType[BACnetPolarity](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetPolarity_NORMAL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetPolarityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPolarityTagged")
	}

	return m, nil
}

func (m *_BACnetPolarityTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPolarityTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPolarityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPolarityTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetPolarity](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPolarityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPolarityTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPolarityTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetPolarityTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetPolarityTagged) IsBACnetPolarityTagged() {}

func (m *_BACnetPolarityTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPolarityTagged) deepCopy() *_BACnetPolarityTagged {
	if m == nil {
		return nil
	}
	_BACnetPolarityTaggedCopy := &_BACnetPolarityTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetPolarityTaggedCopy
}

func (m *_BACnetPolarityTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
