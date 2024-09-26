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

// BACnetTimeValue is the corresponding interface of BACnetTimeValue
type BACnetTimeValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// GetValue returns Value (property field)
	GetValue() BACnetConstructedDataElement
	// IsBACnetTimeValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeValue()
	// CreateBuilder creates a BACnetTimeValueBuilder
	CreateBACnetTimeValueBuilder() BACnetTimeValueBuilder
}

// _BACnetTimeValue is the data-structure of this message
type _BACnetTimeValue struct {
	TimeValue BACnetApplicationTagTime
	Value     BACnetConstructedDataElement
}

var _ BACnetTimeValue = (*_BACnetTimeValue)(nil)

// NewBACnetTimeValue factory function for _BACnetTimeValue
func NewBACnetTimeValue(timeValue BACnetApplicationTagTime, value BACnetConstructedDataElement) *_BACnetTimeValue {
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetTimeValue must not be nil")
	}
	if value == nil {
		panic("value of type BACnetConstructedDataElement for BACnetTimeValue must not be nil")
	}
	return &_BACnetTimeValue{TimeValue: timeValue, Value: value}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimeValueBuilder is a builder for BACnetTimeValue
type BACnetTimeValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeValue BACnetApplicationTagTime, value BACnetConstructedDataElement) BACnetTimeValueBuilder
	// WithTimeValue adds TimeValue (property field)
	WithTimeValue(BACnetApplicationTagTime) BACnetTimeValueBuilder
	// WithTimeValueBuilder adds TimeValue (property field) which is build by the builder
	WithTimeValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetTimeValueBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetConstructedDataElement) BACnetTimeValueBuilder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(BACnetConstructedDataElementBuilder) BACnetConstructedDataElementBuilder) BACnetTimeValueBuilder
	// Build builds the BACnetTimeValue or returns an error if something is wrong
	Build() (BACnetTimeValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimeValue
}

// NewBACnetTimeValueBuilder() creates a BACnetTimeValueBuilder
func NewBACnetTimeValueBuilder() BACnetTimeValueBuilder {
	return &_BACnetTimeValueBuilder{_BACnetTimeValue: new(_BACnetTimeValue)}
}

type _BACnetTimeValueBuilder struct {
	*_BACnetTimeValue

	err *utils.MultiError
}

var _ (BACnetTimeValueBuilder) = (*_BACnetTimeValueBuilder)(nil)

func (m *_BACnetTimeValueBuilder) WithMandatoryFields(timeValue BACnetApplicationTagTime, value BACnetConstructedDataElement) BACnetTimeValueBuilder {
	return m.WithTimeValue(timeValue).WithValue(value)
}

func (m *_BACnetTimeValueBuilder) WithTimeValue(timeValue BACnetApplicationTagTime) BACnetTimeValueBuilder {
	m.TimeValue = timeValue
	return m
}

func (m *_BACnetTimeValueBuilder) WithTimeValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetTimeValueBuilder {
	builder := builderSupplier(m.TimeValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	m.TimeValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return m
}

func (m *_BACnetTimeValueBuilder) WithValue(value BACnetConstructedDataElement) BACnetTimeValueBuilder {
	m.Value = value
	return m
}

func (m *_BACnetTimeValueBuilder) WithValueBuilder(builderSupplier func(BACnetConstructedDataElementBuilder) BACnetConstructedDataElementBuilder) BACnetTimeValueBuilder {
	builder := builderSupplier(m.Value.CreateBACnetConstructedDataElementBuilder())
	var err error
	m.Value, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetConstructedDataElementBuilder failed"))
	}
	return m
}

func (m *_BACnetTimeValueBuilder) Build() (BACnetTimeValue, error) {
	if m.TimeValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'timeValue' not set"))
	}
	if m.Value == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetTimeValue.deepCopy(), nil
}

func (m *_BACnetTimeValueBuilder) MustBuild() BACnetTimeValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetTimeValueBuilder) DeepCopy() any {
	return m.CreateBACnetTimeValueBuilder()
}

// CreateBACnetTimeValueBuilder creates a BACnetTimeValueBuilder
func (m *_BACnetTimeValue) CreateBACnetTimeValueBuilder() BACnetTimeValueBuilder {
	if m == nil {
		return NewBACnetTimeValueBuilder()
	}
	return &_BACnetTimeValueBuilder{_BACnetTimeValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeValue) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *_BACnetTimeValue) GetValue() BACnetConstructedDataElement {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimeValue(structType any) BACnetTimeValue {
	if casted, ok := structType.(BACnetTimeValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeValue) GetTypeName() string {
	return "BACnetTimeValue"
}

func (m *_BACnetTimeValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeValueParse(ctx context.Context, theBytes []byte) (BACnetTimeValue, error) {
	return BACnetTimeValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
		return BACnetTimeValueParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTimeValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeValue, error) {
	v, err := (&_BACnetTimeValue{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTimeValue) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTimeValue BACnetTimeValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	value, err := ReadSimpleField[BACnetConstructedDataElement](ctx, "value", ReadComplex[BACnetConstructedDataElement](BACnetConstructedDataElementParseWithBufferProducer((BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeValue")
	}

	return m, nil
}

func (m *_BACnetTimeValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeValue")
	}

	if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeValue' field")
	}

	if err := WriteSimpleField[BACnetConstructedDataElement](ctx, "value", m.GetValue(), WriteComplex[BACnetConstructedDataElement](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeValue")
	}
	return nil
}

func (m *_BACnetTimeValue) IsBACnetTimeValue() {}

func (m *_BACnetTimeValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimeValue) deepCopy() *_BACnetTimeValue {
	if m == nil {
		return nil
	}
	_BACnetTimeValueCopy := &_BACnetTimeValue{
		m.TimeValue.DeepCopy().(BACnetApplicationTagTime),
		m.Value.DeepCopy().(BACnetConstructedDataElement),
	}
	return _BACnetTimeValueCopy
}

func (m *_BACnetTimeValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
