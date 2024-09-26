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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
type BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// IsBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger()
	// CreateBuilder creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
	CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger struct {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, integerValue BACnetApplicationTagSignedInteger, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	if integerValue == nil {
		panic("integerValue of type BACnetApplicationTagSignedInteger for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		BACnetFaultParameterFaultOutOfRangeMinNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		IntegerValue: integerValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder is a builder for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
type BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
	// WithIntegerValue adds IntegerValue (property field)
	WithIntegerValue(BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
	// WithIntegerValueBuilder adds IntegerValue (property field) which is build by the builder
	WithIntegerValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
	// Build builds the BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
}

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder() creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder {
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder{_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger: new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger)}
}

type _BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder struct {
	*_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder)(nil)

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) WithMandatoryFields(integerValue BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder {
	return m.WithIntegerValue(integerValue)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) WithIntegerValue(integerValue BACnetApplicationTagSignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder {
	m.IntegerValue = integerValue
	return m
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) WithIntegerValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder {
	builder := builderSupplier(m.IntegerValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	m.IntegerValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) Build() (BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, error) {
	if m.IntegerValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'integerValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger.deepCopy(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) MustBuild() BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder) DeepCopy() any {
	return m.CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder()
}

// CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder
func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder {
	if m == nil {
		return NewBACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder()
	}
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerBuilder{_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger: m.deepCopy()}
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

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMinNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMinNormalValueInteger BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) IsBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerCopy := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).deepCopy(),
		m.IntegerValue.DeepCopy().(BACnetApplicationTagSignedInteger),
	}
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
