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

// BACnetConfirmedServiceRequestAddListElement is the corresponding interface of BACnetConfirmedServiceRequestAddListElement
type BACnetConfirmedServiceRequestAddListElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetListOfElements returns ListOfElements (property field)
	GetListOfElements() BACnetConstructedData
	// IsBACnetConfirmedServiceRequestAddListElement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAddListElement()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAddListElementBuilder
	CreateBACnetConfirmedServiceRequestAddListElementBuilder() BACnetConfirmedServiceRequestAddListElementBuilder
}

// _BACnetConfirmedServiceRequestAddListElement is the data-structure of this message
type _BACnetConfirmedServiceRequestAddListElement struct {
	BACnetConfirmedServiceRequestContract
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	ListOfElements     BACnetConstructedData
}

var _ BACnetConfirmedServiceRequestAddListElement = (*_BACnetConfirmedServiceRequestAddListElement)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestAddListElement)(nil)

// NewBACnetConfirmedServiceRequestAddListElement factory function for _BACnetConfirmedServiceRequestAddListElement
func NewBACnetConfirmedServiceRequestAddListElement(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, listOfElements BACnetConstructedData, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAddListElement {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestAddListElement must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetConfirmedServiceRequestAddListElement must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAddListElement{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ObjectIdentifier:                      objectIdentifier,
		PropertyIdentifier:                    propertyIdentifier,
		ArrayIndex:                            arrayIndex,
		ListOfElements:                        listOfElements,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAddListElementBuilder is a builder for BACnetConfirmedServiceRequestAddListElement
type BACnetConfirmedServiceRequestAddListElementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithArrayIndex adds ArrayIndex (property field)
	WithOptionalArrayIndex(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithOptionalArrayIndexBuilder adds ArrayIndex (property field) which is build by the builder
	WithOptionalArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithListOfElements adds ListOfElements (property field)
	WithOptionalListOfElements(BACnetConstructedData) BACnetConfirmedServiceRequestAddListElementBuilder
	// WithOptionalListOfElementsBuilder adds ListOfElements (property field) which is build by the builder
	WithOptionalListOfElementsBuilder(func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetConfirmedServiceRequestAddListElementBuilder
	// Build builds the BACnetConfirmedServiceRequestAddListElement or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAddListElement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAddListElement
}

// NewBACnetConfirmedServiceRequestAddListElementBuilder() creates a BACnetConfirmedServiceRequestAddListElementBuilder
func NewBACnetConfirmedServiceRequestAddListElementBuilder() BACnetConfirmedServiceRequestAddListElementBuilder {
	return &_BACnetConfirmedServiceRequestAddListElementBuilder{_BACnetConfirmedServiceRequestAddListElement: new(_BACnetConfirmedServiceRequestAddListElement)}
}

type _BACnetConfirmedServiceRequestAddListElementBuilder struct {
	*_BACnetConfirmedServiceRequestAddListElement

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAddListElementBuilder) = (*_BACnetConfirmedServiceRequestAddListElementBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestAddListElementBuilder {
	return m.WithObjectIdentifier(objectIdentifier).WithPropertyIdentifier(propertyIdentifier)
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestAddListElementBuilder {
	m.ObjectIdentifier = objectIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestAddListElementBuilder {
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

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestAddListElementBuilder {
	m.PropertyIdentifier = propertyIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestAddListElementBuilder {
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

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithOptionalArrayIndex(arrayIndex BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestAddListElementBuilder {
	m.ArrayIndex = arrayIndex
	return m
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithOptionalArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAddListElementBuilder {
	builder := builderSupplier(m.ArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ArrayIndex, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithOptionalListOfElements(listOfElements BACnetConstructedData) BACnetConfirmedServiceRequestAddListElementBuilder {
	m.ListOfElements = listOfElements
	return m
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) WithOptionalListOfElementsBuilder(builderSupplier func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetConfirmedServiceRequestAddListElementBuilder {
	builder := builderSupplier(m.ListOfElements.CreateBACnetConstructedDataBuilder())
	var err error
	m.ListOfElements, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetConstructedDataBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) Build() (BACnetConfirmedServiceRequestAddListElement, error) {
	if m.ObjectIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if m.PropertyIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestAddListElement.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) MustBuild() BACnetConfirmedServiceRequestAddListElement {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestAddListElementBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestAddListElementBuilder()
}

// CreateBACnetConfirmedServiceRequestAddListElementBuilder creates a BACnetConfirmedServiceRequestAddListElementBuilder
func (m *_BACnetConfirmedServiceRequestAddListElement) CreateBACnetConfirmedServiceRequestAddListElementBuilder() BACnetConfirmedServiceRequestAddListElementBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestAddListElementBuilder()
	}
	return &_BACnetConfirmedServiceRequestAddListElementBuilder{_BACnetConfirmedServiceRequestAddListElement: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAddListElement) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAddListElement) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAddListElement) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestAddListElement) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestAddListElement) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetConfirmedServiceRequestAddListElement) GetListOfElements() BACnetConstructedData {
	return m.ListOfElements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAddListElement(structType any) BACnetConfirmedServiceRequestAddListElement {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAddListElement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAddListElement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAddListElement) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAddListElement"
}

func (m *_BACnetConfirmedServiceRequestAddListElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (listOfElements)
	if m.ListOfElements != nil {
		lengthInBits += m.ListOfElements.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAddListElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAddListElement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAddListElement BACnetConfirmedServiceRequestAddListElement, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAddListElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAddListElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var arrayIndex BACnetContextTagUnsignedInteger
	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
		m.ArrayIndex = arrayIndex
	}

	var listOfElements BACnetConstructedData
	_listOfElements, err := ReadOptionalField[BACnetConstructedData](ctx, "listOfElements", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(3)), (BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfElements' field"))
	}
	if _listOfElements != nil {
		listOfElements = *_listOfElements
		m.ListOfElements = listOfElements
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAddListElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAddListElement")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAddListElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAddListElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAddListElement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAddListElement")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayIndex' field")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "listOfElements", GetRef(m.GetListOfElements()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfElements' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAddListElement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAddListElement")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAddListElement) IsBACnetConfirmedServiceRequestAddListElement() {
}

func (m *_BACnetConfirmedServiceRequestAddListElement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAddListElement) deepCopy() *_BACnetConfirmedServiceRequestAddListElement {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAddListElementCopy := &_BACnetConfirmedServiceRequestAddListElement{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.ObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.PropertyIdentifier.DeepCopy().(BACnetPropertyIdentifierTagged),
		m.ArrayIndex.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ListOfElements.DeepCopy().(BACnetConstructedData),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestAddListElementCopy
}

func (m *_BACnetConfirmedServiceRequestAddListElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
