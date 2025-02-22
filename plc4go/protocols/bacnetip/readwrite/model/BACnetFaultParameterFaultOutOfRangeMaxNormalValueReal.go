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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal()
	// CreateBuilder creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
	CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal struct {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
	RealValue BACnetApplicationTagReal
}

var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, realValue BACnetApplicationTagReal, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal {
	if realValue == nil {
		panic("realValue of type BACnetApplicationTagReal for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal{
		BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		RealValue: realValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder is a builder for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
	// WithRealValue adds RealValue (property field)
	WithRealValue(BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
	// WithRealValueBuilder adds RealValue (property field) which is build by the builder
	WithRealValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder
	// Build builds the BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
}

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder() creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder {
	return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder{_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal: new(_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal)}
}

type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal

	parentBuilder *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder)(nil)

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) setParent(contract BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract) {
	b.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = contract
	contract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = b._BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) WithMandatoryFields(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder {
	return b.WithRealValue(realValue)
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) WithRealValue(realValue BACnetApplicationTagReal) BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder {
	b.RealValue = realValue
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) WithRealValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder {
	builder := builderSupplier(b.RealValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.RealValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) Build() (BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal, error) {
	if b.RealValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'realValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) MustBuild() BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) Done() BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder().(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) buildForBACnetFaultParameterFaultOutOfRangeMaxNormalValue() (BACnetFaultParameterFaultOutOfRangeMaxNormalValue, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder().(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder
func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder()
	}
	return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealBuilder{_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).getLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMaxNormalValueReal BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	realValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	m.RealValue = realValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "realValue", m.GetRealValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealCopy := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal{
		m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.RealValue),
	}
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealCopy.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) String() string {
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
