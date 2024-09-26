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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfCovSubscriptionSpecificationEntry returns ListOfCovSubscriptionSpecificationEntry (property field)
	GetListOfCovSubscriptionSpecificationEntry() []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification()
	// CreateBuilder creates a BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
}

// _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the data-structure of this message
type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification struct {
	OpeningTag                              BACnetOpeningTag
	ListOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	ClosingTag                              BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification = (*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification)(nil)

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification factory function for _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(openingTag BACnetOpeningTag, listOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification must not be nil")
	}
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{OpeningTag: openingTag, ListOfCovSubscriptionSpecificationEntry: listOfCovSubscriptionSpecificationEntry, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder is a builder for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, closingTag BACnetClosingTag) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	// WithListOfCovSubscriptionSpecificationEntry adds ListOfCovSubscriptionSpecificationEntry (property field)
	WithListOfCovSubscriptionSpecificationEntry(...BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
	// Build builds the BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification or returns an error if something is wrong
	Build() (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
}

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder() creates a BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder{_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification: new(_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification)}
}

type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder struct {
	*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification

	err *utils.MultiError
}

var _ (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) = (*_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder)(nil)

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, closingTag BACnetClosingTag) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
	return m.WithOpeningTag(openingTag).WithListOfCovSubscriptionSpecificationEntry(listOfCovSubscriptionSpecificationEntry...).WithClosingTag(closingTag)
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
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

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) WithListOfCovSubscriptionSpecificationEntry(listOfCovSubscriptionSpecificationEntry ...BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
	m.ListOfCovSubscriptionSpecificationEntry = listOfCovSubscriptionSpecificationEntry
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
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

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) Build() (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
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
	return m._BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification.deepCopy(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) MustBuild() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) DeepCopy() any {
	return m.CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder()
}

// CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder creates a BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder
func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder {
	if m == nil {
		return NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder()
	}
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder{_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetListOfCovSubscriptionSpecificationEntry() []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	return m.ListOfCovSubscriptionSpecificationEntry
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(structType any) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfCovSubscriptionSpecificationEntry) > 0 {
		for _, element := range m.ListOfCovSubscriptionSpecificationEntry {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
		return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	v, err := (&_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfCovSubscriptionSpecificationEntry, err := ReadTerminatedArrayField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry](ctx, "listOfCovSubscriptionSpecificationEntry", ReadComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry](BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovSubscriptionSpecificationEntry' field"))
	}
	m.ListOfCovSubscriptionSpecificationEntry = listOfCovSubscriptionSpecificationEntry

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}

	return m, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfCovSubscriptionSpecificationEntry", m.GetListOfCovSubscriptionSpecificationEntry(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfCovSubscriptionSpecificationEntry' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) IsBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification() {
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) deepCopy() *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	if m == nil {
		return nil
	}
	_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationCopy := &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry](m.ListOfCovSubscriptionSpecificationEntry),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationCopy
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
