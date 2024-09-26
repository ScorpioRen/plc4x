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

// BACnetLogRecordLogDatumBooleanValue is the corresponding interface of BACnetLogRecordLogDatumBooleanValue
type BACnetLogRecordLogDatumBooleanValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogRecordLogDatum
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetContextTagBoolean
	// IsBACnetLogRecordLogDatumBooleanValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatumBooleanValue()
	// CreateBuilder creates a BACnetLogRecordLogDatumBooleanValueBuilder
	CreateBACnetLogRecordLogDatumBooleanValueBuilder() BACnetLogRecordLogDatumBooleanValueBuilder
}

// _BACnetLogRecordLogDatumBooleanValue is the data-structure of this message
type _BACnetLogRecordLogDatumBooleanValue struct {
	BACnetLogRecordLogDatumContract
	BooleanValue BACnetContextTagBoolean
}

var _ BACnetLogRecordLogDatumBooleanValue = (*_BACnetLogRecordLogDatumBooleanValue)(nil)
var _ BACnetLogRecordLogDatumRequirements = (*_BACnetLogRecordLogDatumBooleanValue)(nil)

// NewBACnetLogRecordLogDatumBooleanValue factory function for _BACnetLogRecordLogDatumBooleanValue
func NewBACnetLogRecordLogDatumBooleanValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, booleanValue BACnetContextTagBoolean, tagNumber uint8) *_BACnetLogRecordLogDatumBooleanValue {
	if booleanValue == nil {
		panic("booleanValue of type BACnetContextTagBoolean for BACnetLogRecordLogDatumBooleanValue must not be nil")
	}
	_result := &_BACnetLogRecordLogDatumBooleanValue{
		BACnetLogRecordLogDatumContract: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		BooleanValue:                    booleanValue,
	}
	_result.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogRecordLogDatumBooleanValueBuilder is a builder for BACnetLogRecordLogDatumBooleanValue
type BACnetLogRecordLogDatumBooleanValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(booleanValue BACnetContextTagBoolean) BACnetLogRecordLogDatumBooleanValueBuilder
	// WithBooleanValue adds BooleanValue (property field)
	WithBooleanValue(BACnetContextTagBoolean) BACnetLogRecordLogDatumBooleanValueBuilder
	// WithBooleanValueBuilder adds BooleanValue (property field) which is build by the builder
	WithBooleanValueBuilder(func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetLogRecordLogDatumBooleanValueBuilder
	// Build builds the BACnetLogRecordLogDatumBooleanValue or returns an error if something is wrong
	Build() (BACnetLogRecordLogDatumBooleanValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogRecordLogDatumBooleanValue
}

// NewBACnetLogRecordLogDatumBooleanValueBuilder() creates a BACnetLogRecordLogDatumBooleanValueBuilder
func NewBACnetLogRecordLogDatumBooleanValueBuilder() BACnetLogRecordLogDatumBooleanValueBuilder {
	return &_BACnetLogRecordLogDatumBooleanValueBuilder{_BACnetLogRecordLogDatumBooleanValue: new(_BACnetLogRecordLogDatumBooleanValue)}
}

type _BACnetLogRecordLogDatumBooleanValueBuilder struct {
	*_BACnetLogRecordLogDatumBooleanValue

	err *utils.MultiError
}

var _ (BACnetLogRecordLogDatumBooleanValueBuilder) = (*_BACnetLogRecordLogDatumBooleanValueBuilder)(nil)

func (m *_BACnetLogRecordLogDatumBooleanValueBuilder) WithMandatoryFields(booleanValue BACnetContextTagBoolean) BACnetLogRecordLogDatumBooleanValueBuilder {
	return m.WithBooleanValue(booleanValue)
}

func (m *_BACnetLogRecordLogDatumBooleanValueBuilder) WithBooleanValue(booleanValue BACnetContextTagBoolean) BACnetLogRecordLogDatumBooleanValueBuilder {
	m.BooleanValue = booleanValue
	return m
}

func (m *_BACnetLogRecordLogDatumBooleanValueBuilder) WithBooleanValueBuilder(builderSupplier func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetLogRecordLogDatumBooleanValueBuilder {
	builder := builderSupplier(m.BooleanValue.CreateBACnetContextTagBooleanBuilder())
	var err error
	m.BooleanValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetLogRecordLogDatumBooleanValueBuilder) Build() (BACnetLogRecordLogDatumBooleanValue, error) {
	if m.BooleanValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'booleanValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetLogRecordLogDatumBooleanValue.deepCopy(), nil
}

func (m *_BACnetLogRecordLogDatumBooleanValueBuilder) MustBuild() BACnetLogRecordLogDatumBooleanValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetLogRecordLogDatumBooleanValueBuilder) DeepCopy() any {
	return m.CreateBACnetLogRecordLogDatumBooleanValueBuilder()
}

// CreateBACnetLogRecordLogDatumBooleanValueBuilder creates a BACnetLogRecordLogDatumBooleanValueBuilder
func (m *_BACnetLogRecordLogDatumBooleanValue) CreateBACnetLogRecordLogDatumBooleanValueBuilder() BACnetLogRecordLogDatumBooleanValueBuilder {
	if m == nil {
		return NewBACnetLogRecordLogDatumBooleanValueBuilder()
	}
	return &_BACnetLogRecordLogDatumBooleanValueBuilder{_BACnetLogRecordLogDatumBooleanValue: m.deepCopy()}
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

func (m *_BACnetLogRecordLogDatumBooleanValue) GetParent() BACnetLogRecordLogDatumContract {
	return m.BACnetLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumBooleanValue) GetBooleanValue() BACnetContextTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumBooleanValue(structType any) BACnetLogRecordLogDatumBooleanValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumBooleanValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumBooleanValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumBooleanValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumBooleanValue"
}

func (m *_BACnetLogRecordLogDatumBooleanValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumBooleanValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogRecordLogDatumBooleanValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogRecordLogDatum, tagNumber uint8) (__bACnetLogRecordLogDatumBooleanValue BACnetLogRecordLogDatumBooleanValue, err error) {
	m.BACnetLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumBooleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumBooleanValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "booleanValue", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumBooleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumBooleanValue")
	}

	return m, nil
}

func (m *_BACnetLogRecordLogDatumBooleanValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumBooleanValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumBooleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumBooleanValue")
		}

		if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumBooleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumBooleanValue")
		}
		return nil
	}
	return m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumBooleanValue) IsBACnetLogRecordLogDatumBooleanValue() {}

func (m *_BACnetLogRecordLogDatumBooleanValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatumBooleanValue) deepCopy() *_BACnetLogRecordLogDatumBooleanValue {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumBooleanValueCopy := &_BACnetLogRecordLogDatumBooleanValue{
		m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum).deepCopy(),
		m.BooleanValue.DeepCopy().(BACnetContextTagBoolean),
	}
	m.BACnetLogRecordLogDatumContract.(*_BACnetLogRecordLogDatum)._SubType = m
	return _BACnetLogRecordLogDatumBooleanValueCopy
}

func (m *_BACnetLogRecordLogDatumBooleanValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
