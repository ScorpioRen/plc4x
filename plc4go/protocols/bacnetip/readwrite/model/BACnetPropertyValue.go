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

// BACnetPropertyValue is the corresponding interface of BACnetPropertyValue
type BACnetPropertyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedDataElement
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// IsBACnetPropertyValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyValue()
	// CreateBuilder creates a BACnetPropertyValueBuilder
	CreateBACnetPropertyValueBuilder() BACnetPropertyValueBuilder
}

// _BACnetPropertyValue is the data-structure of this message
type _BACnetPropertyValue struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	PropertyArrayIndex BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedDataElement
	Priority           BACnetContextTagUnsignedInteger

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

var _ BACnetPropertyValue = (*_BACnetPropertyValue)(nil)

// NewBACnetPropertyValue factory function for _BACnetPropertyValue
func NewBACnetPropertyValue(propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedDataElement, priority BACnetContextTagUnsignedInteger, objectTypeArgument BACnetObjectType) *_BACnetPropertyValue {
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetPropertyValue must not be nil")
	}
	return &_BACnetPropertyValue{PropertyIdentifier: propertyIdentifier, PropertyArrayIndex: propertyArrayIndex, PropertyValue: propertyValue, Priority: priority, ObjectTypeArgument: objectTypeArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyValueBuilder is a builder for BACnetPropertyValue
type BACnetPropertyValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetPropertyValueBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetPropertyValueBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyValueBuilder
	// WithPropertyArrayIndex adds PropertyArrayIndex (property field)
	WithOptionalPropertyArrayIndex(BACnetContextTagUnsignedInteger) BACnetPropertyValueBuilder
	// WithOptionalPropertyArrayIndexBuilder adds PropertyArrayIndex (property field) which is build by the builder
	WithOptionalPropertyArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyValueBuilder
	// WithPropertyValue adds PropertyValue (property field)
	WithOptionalPropertyValue(BACnetConstructedDataElement) BACnetPropertyValueBuilder
	// WithOptionalPropertyValueBuilder adds PropertyValue (property field) which is build by the builder
	WithOptionalPropertyValueBuilder(func(BACnetConstructedDataElementBuilder) BACnetConstructedDataElementBuilder) BACnetPropertyValueBuilder
	// WithPriority adds Priority (property field)
	WithOptionalPriority(BACnetContextTagUnsignedInteger) BACnetPropertyValueBuilder
	// WithOptionalPriorityBuilder adds Priority (property field) which is build by the builder
	WithOptionalPriorityBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyValueBuilder
	// Build builds the BACnetPropertyValue or returns an error if something is wrong
	Build() (BACnetPropertyValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyValue
}

// NewBACnetPropertyValueBuilder() creates a BACnetPropertyValueBuilder
func NewBACnetPropertyValueBuilder() BACnetPropertyValueBuilder {
	return &_BACnetPropertyValueBuilder{_BACnetPropertyValue: new(_BACnetPropertyValue)}
}

type _BACnetPropertyValueBuilder struct {
	*_BACnetPropertyValue

	err *utils.MultiError
}

var _ (BACnetPropertyValueBuilder) = (*_BACnetPropertyValueBuilder)(nil)

func (m *_BACnetPropertyValueBuilder) WithMandatoryFields(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetPropertyValueBuilder {
	return m.WithPropertyIdentifier(propertyIdentifier)
}

func (m *_BACnetPropertyValueBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetPropertyValueBuilder {
	m.PropertyIdentifier = propertyIdentifier
	return m
}

func (m *_BACnetPropertyValueBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyValueBuilder {
	builder := builderSupplier(m.PropertyIdentifier.CreateBACnetPropertyIdentifierTaggedBuilder())
	var err error
	m.PropertyIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetPropertyIdentifierTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyValueBuilder) WithOptionalPropertyArrayIndex(propertyArrayIndex BACnetContextTagUnsignedInteger) BACnetPropertyValueBuilder {
	m.PropertyArrayIndex = propertyArrayIndex
	return m
}

func (m *_BACnetPropertyValueBuilder) WithOptionalPropertyArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyValueBuilder {
	builder := builderSupplier(m.PropertyArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.PropertyArrayIndex, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyValueBuilder) WithOptionalPropertyValue(propertyValue BACnetConstructedDataElement) BACnetPropertyValueBuilder {
	m.PropertyValue = propertyValue
	return m
}

func (m *_BACnetPropertyValueBuilder) WithOptionalPropertyValueBuilder(builderSupplier func(BACnetConstructedDataElementBuilder) BACnetConstructedDataElementBuilder) BACnetPropertyValueBuilder {
	builder := builderSupplier(m.PropertyValue.CreateBACnetConstructedDataElementBuilder())
	var err error
	m.PropertyValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetConstructedDataElementBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyValueBuilder) WithOptionalPriority(priority BACnetContextTagUnsignedInteger) BACnetPropertyValueBuilder {
	m.Priority = priority
	return m
}

func (m *_BACnetPropertyValueBuilder) WithOptionalPriorityBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyValueBuilder {
	builder := builderSupplier(m.Priority.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.Priority, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyValueBuilder) Build() (BACnetPropertyValue, error) {
	if m.PropertyIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyValue.deepCopy(), nil
}

func (m *_BACnetPropertyValueBuilder) MustBuild() BACnetPropertyValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyValueBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyValueBuilder()
}

// CreateBACnetPropertyValueBuilder creates a BACnetPropertyValueBuilder
func (m *_BACnetPropertyValue) CreateBACnetPropertyValueBuilder() BACnetPropertyValueBuilder {
	if m == nil {
		return NewBACnetPropertyValueBuilder()
	}
	return &_BACnetPropertyValueBuilder{_BACnetPropertyValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyValue) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetPropertyValue) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetPropertyValue) GetPropertyValue() BACnetConstructedDataElement {
	return m.PropertyValue
}

func (m *_BACnetPropertyValue) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyValue(structType any) BACnetPropertyValue {
	if casted, ok := structType.(BACnetPropertyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyValue) GetTypeName() string {
	return "BACnetPropertyValue"
}

func (m *_BACnetPropertyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += m.PropertyValue.GetLengthInBits(ctx)
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetPropertyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyValueParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPropertyValue, error) {
	return BACnetPropertyValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPropertyValueParseWithBufferProducer(objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyValue, error) {
		return BACnetPropertyValueParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	}
}

func BACnetPropertyValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPropertyValue, error) {
	v, err := (&_BACnetPropertyValue{ObjectTypeArgument: objectTypeArgument}).parse(ctx, readBuffer, objectTypeArgument)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPropertyValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (__bACnetPropertyValue BACnetPropertyValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var propertyArrayIndex BACnetContextTagUnsignedInteger
	_propertyArrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyArrayIndex' field"))
	}
	if _propertyArrayIndex != nil {
		propertyArrayIndex = *_propertyArrayIndex
		m.PropertyArrayIndex = propertyArrayIndex
	}

	var propertyValue BACnetConstructedDataElement
	_propertyValue, err := ReadOptionalField[BACnetConstructedDataElement](ctx, "propertyValue", ReadComplex[BACnetConstructedDataElement](BACnetConstructedDataElementParseWithBufferProducer((BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((propertyArrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValue' field"))
	}
	if _propertyValue != nil {
		propertyValue = *_propertyValue
		m.PropertyValue = propertyValue
	}

	var priority BACnetContextTagUnsignedInteger
	_priority, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	if _priority != nil {
		priority = *_priority
		m.Priority = priority
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyValue")
	}

	return m, nil
}

func (m *_BACnetPropertyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyValue")
	}

	if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", GetRef(m.GetPropertyArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyArrayIndex' field")
	}

	if err := WriteOptionalField[BACnetConstructedDataElement](ctx, "propertyValue", GetRef(m.GetPropertyValue()), WriteComplex[BACnetConstructedDataElement](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyValue' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", GetRef(m.GetPriority()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetPropertyValue) IsBACnetPropertyValue() {}

func (m *_BACnetPropertyValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyValue) deepCopy() *_BACnetPropertyValue {
	if m == nil {
		return nil
	}
	_BACnetPropertyValueCopy := &_BACnetPropertyValue{
		m.PropertyIdentifier.DeepCopy().(BACnetPropertyIdentifierTagged),
		m.PropertyArrayIndex.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.PropertyValue.DeepCopy().(BACnetConstructedDataElement),
		m.Priority.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ObjectTypeArgument,
	}
	return _BACnetPropertyValueCopy
}

func (m *_BACnetPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
