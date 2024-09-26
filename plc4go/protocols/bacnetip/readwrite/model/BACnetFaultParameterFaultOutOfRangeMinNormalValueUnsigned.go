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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned()
	// CreateBuilder creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
	CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned struct {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, unsignedValue BACnetApplicationTagUnsignedInteger, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetApplicationTagUnsignedInteger for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned{
		BACnetFaultParameterFaultOutOfRangeMinNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		UnsignedValue: unsignedValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder is a builder for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
	// WithUnsignedValue adds UnsignedValue (property field)
	WithUnsignedValue(BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
	// WithUnsignedValueBuilder adds UnsignedValue (property field) which is build by the builder
	WithUnsignedValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
	// Build builds the BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
}

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder() creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder {
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder{_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned: new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned)}
}

type _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder struct {
	*_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder)(nil)

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder {
	return m.WithUnsignedValue(unsignedValue)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) WithUnsignedValue(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder {
	m.UnsignedValue = unsignedValue
	return m
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) WithUnsignedValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder {
	builder := builderSupplier(m.UnsignedValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.UnsignedValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) Build() (BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned, error) {
	if m.UnsignedValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'unsignedValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned.deepCopy(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) MustBuild() BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder) DeepCopy() any {
	return m.CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder()
}

// CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder creates a BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder
func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) CreateBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder() BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder {
	if m == nil {
		return NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder()
	}
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedBuilder{_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned: m.deepCopy()}
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

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMinNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) IsBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedCopy := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned{
		m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).deepCopy(),
		m.UnsignedValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
