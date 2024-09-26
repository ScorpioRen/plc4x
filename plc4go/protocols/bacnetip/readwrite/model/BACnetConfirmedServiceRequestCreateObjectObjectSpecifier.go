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

// BACnetConfirmedServiceRequestCreateObjectObjectSpecifier is the corresponding interface of BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
type BACnetConfirmedServiceRequestCreateObjectObjectSpecifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetRawObjectType returns RawObjectType (property field)
	GetRawObjectType() BACnetContextTagEnumerated
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetIsObjectType returns IsObjectType (virtual field)
	GetIsObjectType() bool
	// GetObjectType returns ObjectType (virtual field)
	GetObjectType() BACnetObjectType
	// GetIsObjectIdentifier returns IsObjectIdentifier (virtual field)
	GetIsObjectIdentifier() bool
	// IsBACnetConfirmedServiceRequestCreateObjectObjectSpecifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestCreateObjectObjectSpecifier()
	// CreateBuilder creates a BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	CreateBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder() BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
}

// _BACnetConfirmedServiceRequestCreateObjectObjectSpecifier is the data-structure of this message
type _BACnetConfirmedServiceRequestCreateObjectObjectSpecifier struct {
	OpeningTag       BACnetOpeningTag
	RawObjectType    BACnetContextTagEnumerated
	ObjectIdentifier BACnetContextTagObjectIdentifier
	ClosingTag       BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetConfirmedServiceRequestCreateObjectObjectSpecifier = (*_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier)(nil)

// NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifier factory function for _BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
func NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(openingTag BACnetOpeningTag, rawObjectType BACnetContextTagEnumerated, objectIdentifier BACnetContextTagObjectIdentifier, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier must not be nil")
	}
	return &_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier{OpeningTag: openingTag, RawObjectType: rawObjectType, ObjectIdentifier: objectIdentifier, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder is a builder for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
type BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, closingTag BACnetClosingTag) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithRawObjectType adds RawObjectType (property field)
	WithOptionalRawObjectType(BACnetContextTagEnumerated) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithOptionalRawObjectTypeBuilder adds RawObjectType (property field) which is build by the builder
	WithOptionalRawObjectTypeBuilder(func(BACnetContextTagEnumeratedBuilder) BACnetContextTagEnumeratedBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithOptionalObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithOptionalObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithOptionalObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
	// Build builds the BACnetConfirmedServiceRequestCreateObjectObjectSpecifier or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
}

// NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder() creates a BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
func NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder() BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	return &_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder{_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier: new(_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier)}
}

type _BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder struct {
	*_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) = (*_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, closingTag BACnetClosingTag) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	return m.WithOpeningTag(openingTag).WithClosingTag(closingTag)
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
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

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithOptionalRawObjectType(rawObjectType BACnetContextTagEnumerated) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	m.RawObjectType = rawObjectType
	return m
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithOptionalRawObjectTypeBuilder(builderSupplier func(BACnetContextTagEnumeratedBuilder) BACnetContextTagEnumeratedBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	builder := builderSupplier(m.RawObjectType.CreateBACnetContextTagEnumeratedBuilder())
	var err error
	m.RawObjectType, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagEnumeratedBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithOptionalObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	m.ObjectIdentifier = objectIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithOptionalObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	builder := builderSupplier(m.ObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
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

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) Build() (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error) {
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
	return m._BACnetConfirmedServiceRequestCreateObjectObjectSpecifier.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) MustBuild() BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder()
}

// CreateBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder creates a BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder
func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) CreateBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder() BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder()
	}
	return &_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder{_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetRawObjectType() BACnetContextTagEnumerated {
	return m.RawObjectType
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetIsObjectType() bool {
	ctx := context.Background()
	_ = ctx
	rawObjectType := m.GetRawObjectType()
	_ = rawObjectType
	objectIdentifier := m.GetObjectIdentifier()
	_ = objectIdentifier
	return bool(bool((m.GetRawObjectType()) != (nil)))
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetObjectType() BACnetObjectType {
	ctx := context.Background()
	_ = ctx
	rawObjectType := m.GetRawObjectType()
	_ = rawObjectType
	objectIdentifier := m.GetObjectIdentifier()
	_ = objectIdentifier
	return CastBACnetObjectType(MapBACnetObjectType(ctx, (m.GetRawObjectType())))
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetIsObjectIdentifier() bool {
	ctx := context.Background()
	_ = ctx
	rawObjectType := m.GetRawObjectType()
	_ = rawObjectType
	objectIdentifier := m.GetObjectIdentifier()
	_ = objectIdentifier
	return bool(bool((m.GetObjectIdentifier()) != (nil)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(structType any) BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	if casted, ok := structType.(BACnetConfirmedServiceRequestCreateObjectObjectSpecifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestCreateObjectObjectSpecifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetTypeName() string {
	return "BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Optional Field (rawObjectType)
	if m.RawObjectType != nil {
		lengthInBits += m.RawObjectType.GetLengthInBits(ctx)
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)
	}

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error) {
	return BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error) {
		return BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error) {
	v, err := (&_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetConfirmedServiceRequestCreateObjectObjectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	var rawObjectType BACnetContextTagEnumerated
	_rawObjectType, err := ReadOptionalField[BACnetContextTagEnumerated](ctx, "rawObjectType", ReadComplex[BACnetContextTagEnumerated](BACnetContextTagParseWithBufferProducer[BACnetContextTagEnumerated]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_ENUMERATED)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawObjectType' field"))
	}
	if _rawObjectType != nil {
		rawObjectType = *_rawObjectType
		m.RawObjectType = rawObjectType
	}

	isObjectType, err := ReadVirtualField[bool](ctx, "isObjectType", (*bool)(nil), bool((rawObjectType) != (nil)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isObjectType' field"))
	}
	_ = isObjectType

	objectType, err := ReadVirtualField[BACnetObjectType](ctx, "objectType", (*BACnetObjectType)(nil), MapBACnetObjectType(ctx, (rawObjectType)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectType' field"))
	}
	_ = objectType

	var objectIdentifier BACnetContextTagObjectIdentifier
	_objectIdentifier, err := ReadOptionalField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	if _objectIdentifier != nil {
		objectIdentifier = *_objectIdentifier
		m.ObjectIdentifier = objectIdentifier
	}

	isObjectIdentifier, err := ReadVirtualField[bool](ctx, "isObjectIdentifier", (*bool)(nil), bool((objectIdentifier) != (nil)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isObjectIdentifier' field"))
	}
	_ = isObjectIdentifier

	// Validation
	if !(bool(isObjectType) || bool(isObjectIdentifier)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "either we need a objectType or a objectIdentifier"})
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteOptionalField[BACnetContextTagEnumerated](ctx, "rawObjectType", GetRef(m.GetRawObjectType()), WriteComplex[BACnetContextTagEnumerated](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'rawObjectType' field")
	}
	// Virtual field
	isObjectType := m.GetIsObjectType()
	_ = isObjectType
	if _isObjectTypeErr := writeBuffer.WriteVirtual(ctx, "isObjectType", m.GetIsObjectType()); _isObjectTypeErr != nil {
		return errors.Wrap(_isObjectTypeErr, "Error serializing 'isObjectType' field")
	}
	// Virtual field
	objectType := m.GetObjectType()
	_ = objectType
	if _objectTypeErr := writeBuffer.WriteVirtual(ctx, "objectType", m.GetObjectType()); _objectTypeErr != nil {
		return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
	}

	if err := WriteOptionalField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", GetRef(m.GetObjectIdentifier()), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}
	// Virtual field
	isObjectIdentifier := m.GetIsObjectIdentifier()
	_ = isObjectIdentifier
	if _isObjectIdentifierErr := writeBuffer.WriteVirtual(ctx, "isObjectIdentifier", m.GetIsObjectIdentifier()); _isObjectIdentifierErr != nil {
		return errors.Wrap(_isObjectIdentifierErr, "Error serializing 'isObjectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) IsBACnetConfirmedServiceRequestCreateObjectObjectSpecifier() {
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) deepCopy() *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestCreateObjectObjectSpecifierCopy := &_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.RawObjectType.DeepCopy().(BACnetContextTagEnumerated),
		m.ObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetConfirmedServiceRequestCreateObjectObjectSpecifierCopy
}

func (m *_BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
