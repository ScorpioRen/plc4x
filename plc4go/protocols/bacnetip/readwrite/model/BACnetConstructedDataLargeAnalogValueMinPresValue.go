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

// BACnetConstructedDataLargeAnalogValueMinPresValue is the corresponding interface of BACnetConstructedDataLargeAnalogValueMinPresValue
type BACnetConstructedDataLargeAnalogValueMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueMinPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueMinPresValue()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
	CreateBACnetConstructedDataLargeAnalogValueMinPresValueBuilder() BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
}

// _BACnetConstructedDataLargeAnalogValueMinPresValue is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueMinPresValue struct {
	BACnetConstructedDataContract
	MinPresValue BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueMinPresValue = (*_BACnetConstructedDataLargeAnalogValueMinPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueMinPresValue)(nil)

// NewBACnetConstructedDataLargeAnalogValueMinPresValue factory function for _BACnetConstructedDataLargeAnalogValueMinPresValue
func NewBACnetConstructedDataLargeAnalogValueMinPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minPresValue BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueMinPresValue {
	if minPresValue == nil {
		panic("minPresValue of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueMinPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueMinPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinPresValue:                  minPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueMinPresValueBuilder is a builder for BACnetConstructedDataLargeAnalogValueMinPresValue
type BACnetConstructedDataLargeAnalogValueMinPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minPresValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
	// WithMinPresValue adds MinPresValue (property field)
	WithMinPresValue(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
	// WithMinPresValueBuilder adds MinPresValue (property field) which is build by the builder
	WithMinPresValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueMinPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueMinPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueMinPresValue
}

// NewBACnetConstructedDataLargeAnalogValueMinPresValueBuilder() creates a BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
func NewBACnetConstructedDataLargeAnalogValueMinPresValueBuilder() BACnetConstructedDataLargeAnalogValueMinPresValueBuilder {
	return &_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder{_BACnetConstructedDataLargeAnalogValueMinPresValue: new(_BACnetConstructedDataLargeAnalogValueMinPresValue)}
}

type _BACnetConstructedDataLargeAnalogValueMinPresValueBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueMinPresValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) = (*_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder)(nil)

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) WithMandatoryFields(minPresValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMinPresValueBuilder {
	return m.WithMinPresValue(minPresValue)
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) WithMinPresValue(minPresValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMinPresValueBuilder {
	m.MinPresValue = minPresValue
	return m
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) WithMinPresValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueMinPresValueBuilder {
	builder := builderSupplier(m.MinPresValue.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	m.MinPresValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) Build() (BACnetConstructedDataLargeAnalogValueMinPresValue, error) {
	if m.MinPresValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'minPresValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataLargeAnalogValueMinPresValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueMinPresValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataLargeAnalogValueMinPresValueBuilder()
}

// CreateBACnetConstructedDataLargeAnalogValueMinPresValueBuilder creates a BACnetConstructedDataLargeAnalogValueMinPresValueBuilder
func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) CreateBACnetConstructedDataLargeAnalogValueMinPresValueBuilder() BACnetConstructedDataLargeAnalogValueMinPresValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataLargeAnalogValueMinPresValueBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueMinPresValueBuilder{_BACnetConstructedDataLargeAnalogValueMinPresValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetMinPresValue() BACnetApplicationTagDouble {
	return m.MinPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueMinPresValue(structType any) BACnetConstructedDataLargeAnalogValueMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueMinPresValue"
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueMinPresValue BACnetConstructedDataLargeAnalogValueMinPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minPresValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "minPresValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPresValue' field"))
	}
	m.MinPresValue = minPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), minPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueMinPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueMinPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "minPresValue", m.GetMinPresValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueMinPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) IsBACnetConstructedDataLargeAnalogValueMinPresValue() {
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) deepCopy() *_BACnetConstructedDataLargeAnalogValueMinPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueMinPresValueCopy := &_BACnetConstructedDataLargeAnalogValueMinPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MinPresValue.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueMinPresValueCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueMinPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
