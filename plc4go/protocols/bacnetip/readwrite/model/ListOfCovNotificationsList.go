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

// ListOfCovNotificationsList is the corresponding interface of ListOfCovNotificationsList
type ListOfCovNotificationsList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetSpecifications returns Specifications (property field)
	GetSpecifications() []ListOfCovNotifications
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsListOfCovNotificationsList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsListOfCovNotificationsList()
	// CreateBuilder creates a ListOfCovNotificationsListBuilder
	CreateListOfCovNotificationsListBuilder() ListOfCovNotificationsListBuilder
}

// _ListOfCovNotificationsList is the data-structure of this message
type _ListOfCovNotificationsList struct {
	OpeningTag     BACnetOpeningTag
	Specifications []ListOfCovNotifications
	ClosingTag     BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ ListOfCovNotificationsList = (*_ListOfCovNotificationsList)(nil)

// NewListOfCovNotificationsList factory function for _ListOfCovNotificationsList
func NewListOfCovNotificationsList(openingTag BACnetOpeningTag, specifications []ListOfCovNotifications, closingTag BACnetClosingTag, tagNumber uint8) *_ListOfCovNotificationsList {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for ListOfCovNotificationsList must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for ListOfCovNotificationsList must not be nil")
	}
	return &_ListOfCovNotificationsList{OpeningTag: openingTag, Specifications: specifications, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ListOfCovNotificationsListBuilder is a builder for ListOfCovNotificationsList
type ListOfCovNotificationsListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, specifications []ListOfCovNotifications, closingTag BACnetClosingTag) ListOfCovNotificationsListBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) ListOfCovNotificationsListBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) ListOfCovNotificationsListBuilder
	// WithSpecifications adds Specifications (property field)
	WithSpecifications(...ListOfCovNotifications) ListOfCovNotificationsListBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) ListOfCovNotificationsListBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) ListOfCovNotificationsListBuilder
	// Build builds the ListOfCovNotificationsList or returns an error if something is wrong
	Build() (ListOfCovNotificationsList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ListOfCovNotificationsList
}

// NewListOfCovNotificationsListBuilder() creates a ListOfCovNotificationsListBuilder
func NewListOfCovNotificationsListBuilder() ListOfCovNotificationsListBuilder {
	return &_ListOfCovNotificationsListBuilder{_ListOfCovNotificationsList: new(_ListOfCovNotificationsList)}
}

type _ListOfCovNotificationsListBuilder struct {
	*_ListOfCovNotificationsList

	err *utils.MultiError
}

var _ (ListOfCovNotificationsListBuilder) = (*_ListOfCovNotificationsListBuilder)(nil)

func (m *_ListOfCovNotificationsListBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, specifications []ListOfCovNotifications, closingTag BACnetClosingTag) ListOfCovNotificationsListBuilder {
	return m.WithOpeningTag(openingTag).WithSpecifications(specifications...).WithClosingTag(closingTag)
}

func (m *_ListOfCovNotificationsListBuilder) WithOpeningTag(openingTag BACnetOpeningTag) ListOfCovNotificationsListBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_ListOfCovNotificationsListBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) ListOfCovNotificationsListBuilder {
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

func (m *_ListOfCovNotificationsListBuilder) WithSpecifications(specifications ...ListOfCovNotifications) ListOfCovNotificationsListBuilder {
	m.Specifications = specifications
	return m
}

func (m *_ListOfCovNotificationsListBuilder) WithClosingTag(closingTag BACnetClosingTag) ListOfCovNotificationsListBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_ListOfCovNotificationsListBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) ListOfCovNotificationsListBuilder {
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

func (m *_ListOfCovNotificationsListBuilder) Build() (ListOfCovNotificationsList, error) {
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
	return m._ListOfCovNotificationsList.deepCopy(), nil
}

func (m *_ListOfCovNotificationsListBuilder) MustBuild() ListOfCovNotificationsList {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ListOfCovNotificationsListBuilder) DeepCopy() any {
	return m.CreateListOfCovNotificationsListBuilder()
}

// CreateListOfCovNotificationsListBuilder creates a ListOfCovNotificationsListBuilder
func (m *_ListOfCovNotificationsList) CreateListOfCovNotificationsListBuilder() ListOfCovNotificationsListBuilder {
	if m == nil {
		return NewListOfCovNotificationsListBuilder()
	}
	return &_ListOfCovNotificationsListBuilder{_ListOfCovNotificationsList: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListOfCovNotificationsList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ListOfCovNotificationsList) GetSpecifications() []ListOfCovNotifications {
	return m.Specifications
}

func (m *_ListOfCovNotificationsList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastListOfCovNotificationsList(structType any) ListOfCovNotificationsList {
	if casted, ok := structType.(ListOfCovNotificationsList); ok {
		return casted
	}
	if casted, ok := structType.(*ListOfCovNotificationsList); ok {
		return *casted
	}
	return nil
}

func (m *_ListOfCovNotificationsList) GetTypeName() string {
	return "ListOfCovNotificationsList"
}

func (m *_ListOfCovNotificationsList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Specifications) > 0 {
		for _, element := range m.Specifications {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ListOfCovNotificationsList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ListOfCovNotificationsListParse(ctx context.Context, theBytes []byte, tagNumber uint8) (ListOfCovNotificationsList, error) {
	return ListOfCovNotificationsListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func ListOfCovNotificationsListParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotificationsList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotificationsList, error) {
		return ListOfCovNotificationsListParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func ListOfCovNotificationsListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (ListOfCovNotificationsList, error) {
	v, err := (&_ListOfCovNotificationsList{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ListOfCovNotificationsList) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__listOfCovNotificationsList ListOfCovNotificationsList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ListOfCovNotificationsList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListOfCovNotificationsList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	specifications, err := ReadTerminatedArrayField[ListOfCovNotifications](ctx, "specifications", ReadComplex[ListOfCovNotifications](ListOfCovNotificationsParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifications' field"))
	}
	m.Specifications = specifications

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("ListOfCovNotificationsList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListOfCovNotificationsList")
	}

	return m, nil
}

func (m *_ListOfCovNotificationsList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListOfCovNotificationsList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ListOfCovNotificationsList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ListOfCovNotificationsList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "specifications", m.GetSpecifications(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'specifications' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("ListOfCovNotificationsList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ListOfCovNotificationsList")
	}
	return nil
}

////
// Arguments Getter

func (m *_ListOfCovNotificationsList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_ListOfCovNotificationsList) IsListOfCovNotificationsList() {}

func (m *_ListOfCovNotificationsList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ListOfCovNotificationsList) deepCopy() *_ListOfCovNotificationsList {
	if m == nil {
		return nil
	}
	_ListOfCovNotificationsListCopy := &_ListOfCovNotificationsList{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[ListOfCovNotifications, ListOfCovNotifications](m.Specifications),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _ListOfCovNotificationsListCopy
}

func (m *_ListOfCovNotificationsList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
