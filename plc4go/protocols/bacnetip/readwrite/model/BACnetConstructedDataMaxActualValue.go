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

// BACnetConstructedDataMaxActualValue is the corresponding interface of BACnetConstructedDataMaxActualValue
type BACnetConstructedDataMaxActualValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxActualValue returns MaxActualValue (property field)
	GetMaxActualValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataMaxActualValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaxActualValue()
	// CreateBuilder creates a BACnetConstructedDataMaxActualValueBuilder
	CreateBACnetConstructedDataMaxActualValueBuilder() BACnetConstructedDataMaxActualValueBuilder
}

// _BACnetConstructedDataMaxActualValue is the data-structure of this message
type _BACnetConstructedDataMaxActualValue struct {
	BACnetConstructedDataContract
	MaxActualValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataMaxActualValue = (*_BACnetConstructedDataMaxActualValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaxActualValue)(nil)

// NewBACnetConstructedDataMaxActualValue factory function for _BACnetConstructedDataMaxActualValue
func NewBACnetConstructedDataMaxActualValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxActualValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxActualValue {
	if maxActualValue == nil {
		panic("maxActualValue of type BACnetApplicationTagReal for BACnetConstructedDataMaxActualValue must not be nil")
	}
	_result := &_BACnetConstructedDataMaxActualValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxActualValue:                maxActualValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaxActualValueBuilder is a builder for BACnetConstructedDataMaxActualValue
type BACnetConstructedDataMaxActualValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxActualValue BACnetApplicationTagReal) BACnetConstructedDataMaxActualValueBuilder
	// WithMaxActualValue adds MaxActualValue (property field)
	WithMaxActualValue(BACnetApplicationTagReal) BACnetConstructedDataMaxActualValueBuilder
	// WithMaxActualValueBuilder adds MaxActualValue (property field) which is build by the builder
	WithMaxActualValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataMaxActualValueBuilder
	// Build builds the BACnetConstructedDataMaxActualValue or returns an error if something is wrong
	Build() (BACnetConstructedDataMaxActualValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaxActualValue
}

// NewBACnetConstructedDataMaxActualValueBuilder() creates a BACnetConstructedDataMaxActualValueBuilder
func NewBACnetConstructedDataMaxActualValueBuilder() BACnetConstructedDataMaxActualValueBuilder {
	return &_BACnetConstructedDataMaxActualValueBuilder{_BACnetConstructedDataMaxActualValue: new(_BACnetConstructedDataMaxActualValue)}
}

type _BACnetConstructedDataMaxActualValueBuilder struct {
	*_BACnetConstructedDataMaxActualValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaxActualValueBuilder) = (*_BACnetConstructedDataMaxActualValueBuilder)(nil)

func (m *_BACnetConstructedDataMaxActualValueBuilder) WithMandatoryFields(maxActualValue BACnetApplicationTagReal) BACnetConstructedDataMaxActualValueBuilder {
	return m.WithMaxActualValue(maxActualValue)
}

func (m *_BACnetConstructedDataMaxActualValueBuilder) WithMaxActualValue(maxActualValue BACnetApplicationTagReal) BACnetConstructedDataMaxActualValueBuilder {
	m.MaxActualValue = maxActualValue
	return m
}

func (m *_BACnetConstructedDataMaxActualValueBuilder) WithMaxActualValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataMaxActualValueBuilder {
	builder := builderSupplier(m.MaxActualValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.MaxActualValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataMaxActualValueBuilder) Build() (BACnetConstructedDataMaxActualValue, error) {
	if m.MaxActualValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'maxActualValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataMaxActualValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataMaxActualValueBuilder) MustBuild() BACnetConstructedDataMaxActualValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataMaxActualValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataMaxActualValueBuilder()
}

// CreateBACnetConstructedDataMaxActualValueBuilder creates a BACnetConstructedDataMaxActualValueBuilder
func (m *_BACnetConstructedDataMaxActualValue) CreateBACnetConstructedDataMaxActualValueBuilder() BACnetConstructedDataMaxActualValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataMaxActualValueBuilder()
	}
	return &_BACnetConstructedDataMaxActualValueBuilder{_BACnetConstructedDataMaxActualValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxActualValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_ACTUAL_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetMaxActualValue() BACnetApplicationTagReal {
	return m.MaxActualValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMaxActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxActualValue(structType any) BACnetConstructedDataMaxActualValue {
	if casted, ok := structType.(BACnetConstructedDataMaxActualValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxActualValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxActualValue) GetTypeName() string {
	return "BACnetConstructedDataMaxActualValue"
}

func (m *_BACnetConstructedDataMaxActualValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxActualValue)
	lengthInBits += m.MaxActualValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxActualValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaxActualValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaxActualValue BACnetConstructedDataMaxActualValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxActualValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxActualValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "maxActualValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxActualValue' field"))
	}
	m.MaxActualValue = maxActualValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), maxActualValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxActualValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaxActualValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxActualValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxActualValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "maxActualValue", m.GetMaxActualValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxActualValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxActualValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxActualValue) IsBACnetConstructedDataMaxActualValue() {}

func (m *_BACnetConstructedDataMaxActualValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaxActualValue) deepCopy() *_BACnetConstructedDataMaxActualValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaxActualValueCopy := &_BACnetConstructedDataMaxActualValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MaxActualValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaxActualValueCopy
}

func (m *_BACnetConstructedDataMaxActualValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
