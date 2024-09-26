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

// BACnetDateRangeEnclosed is the corresponding interface of BACnetDateRangeEnclosed
type BACnetDateRangeEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetDateRange returns DateRange (property field)
	GetDateRange() BACnetDateRange
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetDateRangeEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDateRangeEnclosed()
	// CreateBuilder creates a BACnetDateRangeEnclosedBuilder
	CreateBACnetDateRangeEnclosedBuilder() BACnetDateRangeEnclosedBuilder
}

// _BACnetDateRangeEnclosed is the data-structure of this message
type _BACnetDateRangeEnclosed struct {
	OpeningTag BACnetOpeningTag
	DateRange  BACnetDateRange
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetDateRangeEnclosed = (*_BACnetDateRangeEnclosed)(nil)

// NewBACnetDateRangeEnclosed factory function for _BACnetDateRangeEnclosed
func NewBACnetDateRangeEnclosed(openingTag BACnetOpeningTag, dateRange BACnetDateRange, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDateRangeEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetDateRangeEnclosed must not be nil")
	}
	if dateRange == nil {
		panic("dateRange of type BACnetDateRange for BACnetDateRangeEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetDateRangeEnclosed must not be nil")
	}
	return &_BACnetDateRangeEnclosed{OpeningTag: openingTag, DateRange: dateRange, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetDateRangeEnclosedBuilder is a builder for BACnetDateRangeEnclosed
type BACnetDateRangeEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, dateRange BACnetDateRange, closingTag BACnetClosingTag) BACnetDateRangeEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetDateRangeEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetDateRangeEnclosedBuilder
	// WithDateRange adds DateRange (property field)
	WithDateRange(BACnetDateRange) BACnetDateRangeEnclosedBuilder
	// WithDateRangeBuilder adds DateRange (property field) which is build by the builder
	WithDateRangeBuilder(func(BACnetDateRangeBuilder) BACnetDateRangeBuilder) BACnetDateRangeEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetDateRangeEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetDateRangeEnclosedBuilder
	// Build builds the BACnetDateRangeEnclosed or returns an error if something is wrong
	Build() (BACnetDateRangeEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetDateRangeEnclosed
}

// NewBACnetDateRangeEnclosedBuilder() creates a BACnetDateRangeEnclosedBuilder
func NewBACnetDateRangeEnclosedBuilder() BACnetDateRangeEnclosedBuilder {
	return &_BACnetDateRangeEnclosedBuilder{_BACnetDateRangeEnclosed: new(_BACnetDateRangeEnclosed)}
}

type _BACnetDateRangeEnclosedBuilder struct {
	*_BACnetDateRangeEnclosed

	err *utils.MultiError
}

var _ (BACnetDateRangeEnclosedBuilder) = (*_BACnetDateRangeEnclosedBuilder)(nil)

func (m *_BACnetDateRangeEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, dateRange BACnetDateRange, closingTag BACnetClosingTag) BACnetDateRangeEnclosedBuilder {
	return m.WithOpeningTag(openingTag).WithDateRange(dateRange).WithClosingTag(closingTag)
}

func (m *_BACnetDateRangeEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetDateRangeEnclosedBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetDateRangeEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetDateRangeEnclosedBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetDateRangeEnclosedBuilder) WithDateRange(dateRange BACnetDateRange) BACnetDateRangeEnclosedBuilder {
	m.DateRange = dateRange
	return m
}

func (m *_BACnetDateRangeEnclosedBuilder) WithDateRangeBuilder(builderSupplier func(BACnetDateRangeBuilder) BACnetDateRangeBuilder) BACnetDateRangeEnclosedBuilder {
	builder := builderSupplier(m.DateRange.CreateBACnetDateRangeBuilder())
	var err error
	m.DateRange, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDateRangeBuilder failed"))
	}
	return m
}

func (m *_BACnetDateRangeEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetDateRangeEnclosedBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetDateRangeEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetDateRangeEnclosedBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetDateRangeEnclosedBuilder) Build() (BACnetDateRangeEnclosed, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.DateRange == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'dateRange' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetDateRangeEnclosed.deepCopy(), nil
}

func (m *_BACnetDateRangeEnclosedBuilder) MustBuild() BACnetDateRangeEnclosed {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetDateRangeEnclosedBuilder) DeepCopy() any {
	return m.CreateBACnetDateRangeEnclosedBuilder()
}

// CreateBACnetDateRangeEnclosedBuilder creates a BACnetDateRangeEnclosedBuilder
func (m *_BACnetDateRangeEnclosed) CreateBACnetDateRangeEnclosedBuilder() BACnetDateRangeEnclosedBuilder {
	if m == nil {
		return NewBACnetDateRangeEnclosedBuilder()
	}
	return &_BACnetDateRangeEnclosedBuilder{_BACnetDateRangeEnclosed: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateRangeEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDateRangeEnclosed) GetDateRange() BACnetDateRange {
	return m.DateRange
}

func (m *_BACnetDateRangeEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetDateRangeEnclosed(structType any) BACnetDateRangeEnclosed {
	if casted, ok := structType.(BACnetDateRangeEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateRangeEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateRangeEnclosed) GetTypeName() string {
	return "BACnetDateRangeEnclosed"
}

func (m *_BACnetDateRangeEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateRangeEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateRangeEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetDateRangeEnclosed, error) {
	return BACnetDateRangeEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetDateRangeEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRangeEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRangeEnclosed, error) {
		return BACnetDateRangeEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetDateRangeEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDateRangeEnclosed, error) {
	v, err := (&_BACnetDateRangeEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetDateRangeEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetDateRangeEnclosed BACnetDateRangeEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRangeEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRangeEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	dateRange, err := ReadSimpleField[BACnetDateRange](ctx, "dateRange", ReadComplex[BACnetDateRange](BACnetDateRangeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateRange' field"))
	}
	m.DateRange = dateRange

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetDateRangeEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRangeEnclosed")
	}

	return m, nil
}

func (m *_BACnetDateRangeEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateRangeEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDateRangeEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRangeEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetDateRange](ctx, "dateRange", m.GetDateRange(), WriteComplex[BACnetDateRange](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'dateRange' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRangeEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRangeEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDateRangeEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDateRangeEnclosed) IsBACnetDateRangeEnclosed() {}

func (m *_BACnetDateRangeEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetDateRangeEnclosed) deepCopy() *_BACnetDateRangeEnclosed {
	if m == nil {
		return nil
	}
	_BACnetDateRangeEnclosedCopy := &_BACnetDateRangeEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.DateRange.DeepCopy().(BACnetDateRange),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetDateRangeEnclosedCopy
}

func (m *_BACnetDateRangeEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
