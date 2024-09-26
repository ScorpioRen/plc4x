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

// BACnetConstructedDataBitStringValuePresentValue is the corresponding interface of BACnetConstructedDataBitStringValuePresentValue
type BACnetConstructedDataBitStringValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagBitString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBitString
	// IsBACnetConstructedDataBitStringValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBitStringValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataBitStringValuePresentValueBuilder
	CreateBACnetConstructedDataBitStringValuePresentValueBuilder() BACnetConstructedDataBitStringValuePresentValueBuilder
}

// _BACnetConstructedDataBitStringValuePresentValue is the data-structure of this message
type _BACnetConstructedDataBitStringValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagBitString
}

var _ BACnetConstructedDataBitStringValuePresentValue = (*_BACnetConstructedDataBitStringValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBitStringValuePresentValue)(nil)

// NewBACnetConstructedDataBitStringValuePresentValue factory function for _BACnetConstructedDataBitStringValuePresentValue
func NewBACnetConstructedDataBitStringValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagBitString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBitStringValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagBitString for BACnetConstructedDataBitStringValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataBitStringValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBitStringValuePresentValueBuilder is a builder for BACnetConstructedDataBitStringValuePresentValue
type BACnetConstructedDataBitStringValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagBitString) BACnetConstructedDataBitStringValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagBitString) BACnetConstructedDataBitStringValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetConstructedDataBitStringValuePresentValueBuilder
	// Build builds the BACnetConstructedDataBitStringValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataBitStringValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBitStringValuePresentValue
}

// NewBACnetConstructedDataBitStringValuePresentValueBuilder() creates a BACnetConstructedDataBitStringValuePresentValueBuilder
func NewBACnetConstructedDataBitStringValuePresentValueBuilder() BACnetConstructedDataBitStringValuePresentValueBuilder {
	return &_BACnetConstructedDataBitStringValuePresentValueBuilder{_BACnetConstructedDataBitStringValuePresentValue: new(_BACnetConstructedDataBitStringValuePresentValue)}
}

type _BACnetConstructedDataBitStringValuePresentValueBuilder struct {
	*_BACnetConstructedDataBitStringValuePresentValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataBitStringValuePresentValueBuilder) = (*_BACnetConstructedDataBitStringValuePresentValueBuilder)(nil)

func (m *_BACnetConstructedDataBitStringValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagBitString) BACnetConstructedDataBitStringValuePresentValueBuilder {
	return m.WithPresentValue(presentValue)
}

func (m *_BACnetConstructedDataBitStringValuePresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagBitString) BACnetConstructedDataBitStringValuePresentValueBuilder {
	m.PresentValue = presentValue
	return m
}

func (m *_BACnetConstructedDataBitStringValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetConstructedDataBitStringValuePresentValueBuilder {
	builder := builderSupplier(m.PresentValue.CreateBACnetApplicationTagBitStringBuilder())
	var err error
	m.PresentValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBitStringBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataBitStringValuePresentValueBuilder) Build() (BACnetConstructedDataBitStringValuePresentValue, error) {
	if m.PresentValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataBitStringValuePresentValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataBitStringValuePresentValueBuilder) MustBuild() BACnetConstructedDataBitStringValuePresentValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataBitStringValuePresentValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataBitStringValuePresentValueBuilder()
}

// CreateBACnetConstructedDataBitStringValuePresentValueBuilder creates a BACnetConstructedDataBitStringValuePresentValueBuilder
func (m *_BACnetConstructedDataBitStringValuePresentValue) CreateBACnetConstructedDataBitStringValuePresentValueBuilder() BACnetConstructedDataBitStringValuePresentValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataBitStringValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataBitStringValuePresentValueBuilder{_BACnetConstructedDataBitStringValuePresentValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BITSTRING_VALUE
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetPresentValue() BACnetApplicationTagBitString {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetActualValue() BACnetApplicationTagBitString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBitString(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBitStringValuePresentValue(structType any) BACnetConstructedDataBitStringValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataBitStringValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBitStringValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataBitStringValuePresentValue"
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBitStringValuePresentValue BACnetConstructedDataBitStringValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBitStringValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBitStringValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagBitString](ctx, "presentValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagBitString](ctx, "actualValue", (*BACnetApplicationTagBitString)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBitStringValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBitStringValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBitStringValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBitStringValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagBitString](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBitStringValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBitStringValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) IsBACnetConstructedDataBitStringValuePresentValue() {
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) deepCopy() *_BACnetConstructedDataBitStringValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBitStringValuePresentValueCopy := &_BACnetConstructedDataBitStringValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetApplicationTagBitString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBitStringValuePresentValueCopy
}

func (m *_BACnetConstructedDataBitStringValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
