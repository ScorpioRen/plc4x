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

// BACnetAssignedLandingCallsLandingCallsList is the corresponding interface of BACnetAssignedLandingCallsLandingCallsList
type BACnetAssignedLandingCallsLandingCallsList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetLandingCalls returns LandingCalls (property field)
	GetLandingCalls() []BACnetAssignedLandingCallsLandingCallsListEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetAssignedLandingCallsLandingCallsList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAssignedLandingCallsLandingCallsList()
	// CreateBuilder creates a BACnetAssignedLandingCallsLandingCallsListBuilder
	CreateBACnetAssignedLandingCallsLandingCallsListBuilder() BACnetAssignedLandingCallsLandingCallsListBuilder
}

// _BACnetAssignedLandingCallsLandingCallsList is the data-structure of this message
type _BACnetAssignedLandingCallsLandingCallsList struct {
	OpeningTag   BACnetOpeningTag
	LandingCalls []BACnetAssignedLandingCallsLandingCallsListEntry
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetAssignedLandingCallsLandingCallsList = (*_BACnetAssignedLandingCallsLandingCallsList)(nil)

// NewBACnetAssignedLandingCallsLandingCallsList factory function for _BACnetAssignedLandingCallsLandingCallsList
func NewBACnetAssignedLandingCallsLandingCallsList(openingTag BACnetOpeningTag, landingCalls []BACnetAssignedLandingCallsLandingCallsListEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAssignedLandingCallsLandingCallsList {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetAssignedLandingCallsLandingCallsList must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetAssignedLandingCallsLandingCallsList must not be nil")
	}
	return &_BACnetAssignedLandingCallsLandingCallsList{OpeningTag: openingTag, LandingCalls: landingCalls, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAssignedLandingCallsLandingCallsListBuilder is a builder for BACnetAssignedLandingCallsLandingCallsList
type BACnetAssignedLandingCallsLandingCallsListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, landingCalls []BACnetAssignedLandingCallsLandingCallsListEntry, closingTag BACnetClosingTag) BACnetAssignedLandingCallsLandingCallsListBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetAssignedLandingCallsLandingCallsListBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAssignedLandingCallsLandingCallsListBuilder
	// WithLandingCalls adds LandingCalls (property field)
	WithLandingCalls(...BACnetAssignedLandingCallsLandingCallsListEntry) BACnetAssignedLandingCallsLandingCallsListBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetAssignedLandingCallsLandingCallsListBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAssignedLandingCallsLandingCallsListBuilder
	// Build builds the BACnetAssignedLandingCallsLandingCallsList or returns an error if something is wrong
	Build() (BACnetAssignedLandingCallsLandingCallsList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAssignedLandingCallsLandingCallsList
}

// NewBACnetAssignedLandingCallsLandingCallsListBuilder() creates a BACnetAssignedLandingCallsLandingCallsListBuilder
func NewBACnetAssignedLandingCallsLandingCallsListBuilder() BACnetAssignedLandingCallsLandingCallsListBuilder {
	return &_BACnetAssignedLandingCallsLandingCallsListBuilder{_BACnetAssignedLandingCallsLandingCallsList: new(_BACnetAssignedLandingCallsLandingCallsList)}
}

type _BACnetAssignedLandingCallsLandingCallsListBuilder struct {
	*_BACnetAssignedLandingCallsLandingCallsList

	err *utils.MultiError
}

var _ (BACnetAssignedLandingCallsLandingCallsListBuilder) = (*_BACnetAssignedLandingCallsLandingCallsListBuilder)(nil)

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, landingCalls []BACnetAssignedLandingCallsLandingCallsListEntry, closingTag BACnetClosingTag) BACnetAssignedLandingCallsLandingCallsListBuilder {
	return m.WithOpeningTag(openingTag).WithLandingCalls(landingCalls...).WithClosingTag(closingTag)
}

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetAssignedLandingCallsLandingCallsListBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAssignedLandingCallsLandingCallsListBuilder {
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

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) WithLandingCalls(landingCalls ...BACnetAssignedLandingCallsLandingCallsListEntry) BACnetAssignedLandingCallsLandingCallsListBuilder {
	m.LandingCalls = landingCalls
	return m
}

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetAssignedLandingCallsLandingCallsListBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAssignedLandingCallsLandingCallsListBuilder {
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

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) Build() (BACnetAssignedLandingCallsLandingCallsList, error) {
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
	return m._BACnetAssignedLandingCallsLandingCallsList.deepCopy(), nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) MustBuild() BACnetAssignedLandingCallsLandingCallsList {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetAssignedLandingCallsLandingCallsListBuilder) DeepCopy() any {
	return m.CreateBACnetAssignedLandingCallsLandingCallsListBuilder()
}

// CreateBACnetAssignedLandingCallsLandingCallsListBuilder creates a BACnetAssignedLandingCallsLandingCallsListBuilder
func (m *_BACnetAssignedLandingCallsLandingCallsList) CreateBACnetAssignedLandingCallsLandingCallsListBuilder() BACnetAssignedLandingCallsLandingCallsListBuilder {
	if m == nil {
		return NewBACnetAssignedLandingCallsLandingCallsListBuilder()
	}
	return &_BACnetAssignedLandingCallsLandingCallsListBuilder{_BACnetAssignedLandingCallsLandingCallsList: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetLandingCalls() []BACnetAssignedLandingCallsLandingCallsListEntry {
	return m.LandingCalls
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAssignedLandingCallsLandingCallsList(structType any) BACnetAssignedLandingCallsLandingCallsList {
	if casted, ok := structType.(BACnetAssignedLandingCallsLandingCallsList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAssignedLandingCallsLandingCallsList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetTypeName() string {
	return "BACnetAssignedLandingCallsLandingCallsList"
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.LandingCalls) > 0 {
		for _, element := range m.LandingCalls {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAssignedLandingCallsLandingCallsListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAssignedLandingCallsLandingCallsList, error) {
	return BACnetAssignedLandingCallsLandingCallsListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAssignedLandingCallsLandingCallsListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsList, error) {
		return BACnetAssignedLandingCallsLandingCallsListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetAssignedLandingCallsLandingCallsListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAssignedLandingCallsLandingCallsList, error) {
	v, err := (&_BACnetAssignedLandingCallsLandingCallsList{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetAssignedLandingCallsLandingCallsList BACnetAssignedLandingCallsLandingCallsList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedLandingCallsLandingCallsList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedLandingCallsLandingCallsList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	landingCalls, err := ReadTerminatedArrayField[BACnetAssignedLandingCallsLandingCallsListEntry](ctx, "landingCalls", ReadComplex[BACnetAssignedLandingCallsLandingCallsListEntry](BACnetAssignedLandingCallsLandingCallsListEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'landingCalls' field"))
	}
	m.LandingCalls = landingCalls

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetAssignedLandingCallsLandingCallsList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedLandingCallsLandingCallsList")
	}

	return m, nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAssignedLandingCallsLandingCallsList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedLandingCallsLandingCallsList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "landingCalls", m.GetLandingCalls(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'landingCalls' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedLandingCallsLandingCallsList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedLandingCallsLandingCallsList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAssignedLandingCallsLandingCallsList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAssignedLandingCallsLandingCallsList) IsBACnetAssignedLandingCallsLandingCallsList() {
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) deepCopy() *_BACnetAssignedLandingCallsLandingCallsList {
	if m == nil {
		return nil
	}
	_BACnetAssignedLandingCallsLandingCallsListCopy := &_BACnetAssignedLandingCallsLandingCallsList{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetAssignedLandingCallsLandingCallsListEntry, BACnetAssignedLandingCallsLandingCallsListEntry](m.LandingCalls),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetAssignedLandingCallsLandingCallsListCopy
}

func (m *_BACnetAssignedLandingCallsLandingCallsList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
