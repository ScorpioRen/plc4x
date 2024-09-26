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

// BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged is the corresponding interface of BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
type BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice
	// IsBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
	CreateBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
}

// _BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged is the data-structure of this message
type _BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged struct {
	Header BACnetTagHeader
	Value  BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged = (*_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged)(nil)

// NewBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged factory function for _BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
func NewBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged(header BACnetTagHeader, value BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice, tagNumber uint8, tagClass TagClass) *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged must not be nil")
	}
	return &_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder is a builder for BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
type BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
	// Build builds the BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
}

// NewBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder() creates a BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
func NewBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder {
	return &_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder{_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged: new(_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged)}
}

type _BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder struct {
	*_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) = (*_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder {
	return m.WithHeader(header).WithValue(value)
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder {
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

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) WithValue(value BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder {
	m.Value = value
	return m
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) Build() (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) MustBuild() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder()
}

// CreateBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder creates a BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder
func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) CreateBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder()
	}
	return &_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedBuilder{_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetValue() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged(structType any) BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged"
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, error) {
	return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, error) {
		return BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, error) {
	v, err := (&_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged")
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

	value, err := ReadManualField[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice](ctx, "value", readBuffer, EnsureType[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice_COLDSTART)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDevice](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) IsBACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged() {
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) deepCopy() *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedCopy := &_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedCopy
}

func (m *_BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
