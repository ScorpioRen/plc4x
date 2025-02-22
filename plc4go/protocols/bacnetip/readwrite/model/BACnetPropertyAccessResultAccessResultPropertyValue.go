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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyAccessResultAccessResultPropertyValue is the corresponding interface of BACnetPropertyAccessResultAccessResultPropertyValue
type BACnetPropertyAccessResultAccessResultPropertyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyAccessResultAccessResult
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// IsBACnetPropertyAccessResultAccessResultPropertyValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyAccessResultAccessResultPropertyValue()
	// CreateBuilder creates a BACnetPropertyAccessResultAccessResultPropertyValueBuilder
	CreateBACnetPropertyAccessResultAccessResultPropertyValueBuilder() BACnetPropertyAccessResultAccessResultPropertyValueBuilder
}

// _BACnetPropertyAccessResultAccessResultPropertyValue is the data-structure of this message
type _BACnetPropertyAccessResultAccessResultPropertyValue struct {
	BACnetPropertyAccessResultAccessResultContract
	PropertyValue BACnetConstructedData
}

var _ BACnetPropertyAccessResultAccessResultPropertyValue = (*_BACnetPropertyAccessResultAccessResultPropertyValue)(nil)
var _ BACnetPropertyAccessResultAccessResultRequirements = (*_BACnetPropertyAccessResultAccessResultPropertyValue)(nil)

// NewBACnetPropertyAccessResultAccessResultPropertyValue factory function for _BACnetPropertyAccessResultAccessResultPropertyValue
func NewBACnetPropertyAccessResultAccessResultPropertyValue(peekedTagHeader BACnetTagHeader, propertyValue BACnetConstructedData, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPropertyAccessResultAccessResultPropertyValue {
	if propertyValue == nil {
		panic("propertyValue of type BACnetConstructedData for BACnetPropertyAccessResultAccessResultPropertyValue must not be nil")
	}
	_result := &_BACnetPropertyAccessResultAccessResultPropertyValue{
		BACnetPropertyAccessResultAccessResultContract: NewBACnetPropertyAccessResultAccessResult(peekedTagHeader, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument),
		PropertyValue: propertyValue,
	}
	_result.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyAccessResultAccessResultPropertyValueBuilder is a builder for BACnetPropertyAccessResultAccessResultPropertyValue
type BACnetPropertyAccessResultAccessResultPropertyValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(propertyValue BACnetConstructedData) BACnetPropertyAccessResultAccessResultPropertyValueBuilder
	// WithPropertyValue adds PropertyValue (property field)
	WithPropertyValue(BACnetConstructedData) BACnetPropertyAccessResultAccessResultPropertyValueBuilder
	// WithPropertyValueBuilder adds PropertyValue (property field) which is build by the builder
	WithPropertyValueBuilder(func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetPropertyAccessResultAccessResultPropertyValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyAccessResultAccessResultBuilder
	// Build builds the BACnetPropertyAccessResultAccessResultPropertyValue or returns an error if something is wrong
	Build() (BACnetPropertyAccessResultAccessResultPropertyValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyAccessResultAccessResultPropertyValue
}

// NewBACnetPropertyAccessResultAccessResultPropertyValueBuilder() creates a BACnetPropertyAccessResultAccessResultPropertyValueBuilder
func NewBACnetPropertyAccessResultAccessResultPropertyValueBuilder() BACnetPropertyAccessResultAccessResultPropertyValueBuilder {
	return &_BACnetPropertyAccessResultAccessResultPropertyValueBuilder{_BACnetPropertyAccessResultAccessResultPropertyValue: new(_BACnetPropertyAccessResultAccessResultPropertyValue)}
}

type _BACnetPropertyAccessResultAccessResultPropertyValueBuilder struct {
	*_BACnetPropertyAccessResultAccessResultPropertyValue

	parentBuilder *_BACnetPropertyAccessResultAccessResultBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyAccessResultAccessResultPropertyValueBuilder) = (*_BACnetPropertyAccessResultAccessResultPropertyValueBuilder)(nil)

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) setParent(contract BACnetPropertyAccessResultAccessResultContract) {
	b.BACnetPropertyAccessResultAccessResultContract = contract
	contract.(*_BACnetPropertyAccessResultAccessResult)._SubType = b._BACnetPropertyAccessResultAccessResultPropertyValue
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) WithMandatoryFields(propertyValue BACnetConstructedData) BACnetPropertyAccessResultAccessResultPropertyValueBuilder {
	return b.WithPropertyValue(propertyValue)
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) WithPropertyValue(propertyValue BACnetConstructedData) BACnetPropertyAccessResultAccessResultPropertyValueBuilder {
	b.PropertyValue = propertyValue
	return b
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) WithPropertyValueBuilder(builderSupplier func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetPropertyAccessResultAccessResultPropertyValueBuilder {
	builder := builderSupplier(b.PropertyValue.CreateBACnetConstructedDataBuilder())
	var err error
	b.PropertyValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConstructedDataBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) Build() (BACnetPropertyAccessResultAccessResultPropertyValue, error) {
	if b.PropertyValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyAccessResultAccessResultPropertyValue.deepCopy(), nil
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) MustBuild() BACnetPropertyAccessResultAccessResultPropertyValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) Done() BACnetPropertyAccessResultAccessResultBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyAccessResultAccessResultBuilder().(*_BACnetPropertyAccessResultAccessResultBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) buildForBACnetPropertyAccessResultAccessResult() (BACnetPropertyAccessResultAccessResult, error) {
	return b.Build()
}

func (b *_BACnetPropertyAccessResultAccessResultPropertyValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyAccessResultAccessResultPropertyValueBuilder().(*_BACnetPropertyAccessResultAccessResultPropertyValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyAccessResultAccessResultPropertyValueBuilder creates a BACnetPropertyAccessResultAccessResultPropertyValueBuilder
func (b *_BACnetPropertyAccessResultAccessResultPropertyValue) CreateBACnetPropertyAccessResultAccessResultPropertyValueBuilder() BACnetPropertyAccessResultAccessResultPropertyValueBuilder {
	if b == nil {
		return NewBACnetPropertyAccessResultAccessResultPropertyValueBuilder()
	}
	return &_BACnetPropertyAccessResultAccessResultPropertyValueBuilder{_BACnetPropertyAccessResultAccessResultPropertyValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetParent() BACnetPropertyAccessResultAccessResultContract {
	return m.BACnetPropertyAccessResultAccessResultContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResultAccessResultPropertyValue(structType any) BACnetPropertyAccessResultAccessResultPropertyValue {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResultPropertyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResultPropertyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResultPropertyValue"
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult).getLengthInBits(ctx))

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyAccessResultAccessResult, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetPropertyAccessResultAccessResultPropertyValue BACnetPropertyAccessResultAccessResultPropertyValue, err error) {
	m.BACnetPropertyAccessResultAccessResultContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResultPropertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResultPropertyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	propertyValue, err := ReadSimpleField[BACnetConstructedData](ctx, "propertyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(4)), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(propertyIdentifierArgument), (BACnetTagPayloadUnsignedInteger)(propertyArrayIndexArgument)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValue' field"))
	}
	m.PropertyValue = propertyValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResultPropertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResultPropertyValue")
	}

	return m, nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResultPropertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResultPropertyValue")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "propertyValue", m.GetPropertyValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResultPropertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResultPropertyValue")
		}
		return nil
	}
	return m.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) IsBACnetPropertyAccessResultAccessResultPropertyValue() {
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) deepCopy() *_BACnetPropertyAccessResultAccessResultPropertyValue {
	if m == nil {
		return nil
	}
	_BACnetPropertyAccessResultAccessResultPropertyValueCopy := &_BACnetPropertyAccessResultAccessResultPropertyValue{
		m.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult).deepCopy(),
		utils.DeepCopy[BACnetConstructedData](m.PropertyValue),
	}
	_BACnetPropertyAccessResultAccessResultPropertyValueCopy.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult)._SubType = m
	return _BACnetPropertyAccessResultAccessResultPropertyValueCopy
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
