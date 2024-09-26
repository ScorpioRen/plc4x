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

// BACnetConstructedDataCOVIncrement is the corresponding interface of BACnetConstructedDataCOVIncrement
type BACnetConstructedDataCOVIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataCOVIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCOVIncrement()
	// CreateBuilder creates a BACnetConstructedDataCOVIncrementBuilder
	CreateBACnetConstructedDataCOVIncrementBuilder() BACnetConstructedDataCOVIncrementBuilder
}

// _BACnetConstructedDataCOVIncrement is the data-structure of this message
type _BACnetConstructedDataCOVIncrement struct {
	BACnetConstructedDataContract
	CovIncrement BACnetApplicationTagReal
}

var _ BACnetConstructedDataCOVIncrement = (*_BACnetConstructedDataCOVIncrement)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCOVIncrement)(nil)

// NewBACnetConstructedDataCOVIncrement factory function for _BACnetConstructedDataCOVIncrement
func NewBACnetConstructedDataCOVIncrement(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, covIncrement BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVIncrement {
	if covIncrement == nil {
		panic("covIncrement of type BACnetApplicationTagReal for BACnetConstructedDataCOVIncrement must not be nil")
	}
	_result := &_BACnetConstructedDataCOVIncrement{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovIncrement:                  covIncrement,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCOVIncrementBuilder is a builder for BACnetConstructedDataCOVIncrement
type BACnetConstructedDataCOVIncrementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(covIncrement BACnetApplicationTagReal) BACnetConstructedDataCOVIncrementBuilder
	// WithCovIncrement adds CovIncrement (property field)
	WithCovIncrement(BACnetApplicationTagReal) BACnetConstructedDataCOVIncrementBuilder
	// WithCovIncrementBuilder adds CovIncrement (property field) which is build by the builder
	WithCovIncrementBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataCOVIncrementBuilder
	// Build builds the BACnetConstructedDataCOVIncrement or returns an error if something is wrong
	Build() (BACnetConstructedDataCOVIncrement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCOVIncrement
}

// NewBACnetConstructedDataCOVIncrementBuilder() creates a BACnetConstructedDataCOVIncrementBuilder
func NewBACnetConstructedDataCOVIncrementBuilder() BACnetConstructedDataCOVIncrementBuilder {
	return &_BACnetConstructedDataCOVIncrementBuilder{_BACnetConstructedDataCOVIncrement: new(_BACnetConstructedDataCOVIncrement)}
}

type _BACnetConstructedDataCOVIncrementBuilder struct {
	*_BACnetConstructedDataCOVIncrement

	err *utils.MultiError
}

var _ (BACnetConstructedDataCOVIncrementBuilder) = (*_BACnetConstructedDataCOVIncrementBuilder)(nil)

func (m *_BACnetConstructedDataCOVIncrementBuilder) WithMandatoryFields(covIncrement BACnetApplicationTagReal) BACnetConstructedDataCOVIncrementBuilder {
	return m.WithCovIncrement(covIncrement)
}

func (m *_BACnetConstructedDataCOVIncrementBuilder) WithCovIncrement(covIncrement BACnetApplicationTagReal) BACnetConstructedDataCOVIncrementBuilder {
	m.CovIncrement = covIncrement
	return m
}

func (m *_BACnetConstructedDataCOVIncrementBuilder) WithCovIncrementBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataCOVIncrementBuilder {
	builder := builderSupplier(m.CovIncrement.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.CovIncrement, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataCOVIncrementBuilder) Build() (BACnetConstructedDataCOVIncrement, error) {
	if m.CovIncrement == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'covIncrement' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataCOVIncrement.deepCopy(), nil
}

func (m *_BACnetConstructedDataCOVIncrementBuilder) MustBuild() BACnetConstructedDataCOVIncrement {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataCOVIncrementBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataCOVIncrementBuilder()
}

// CreateBACnetConstructedDataCOVIncrementBuilder creates a BACnetConstructedDataCOVIncrementBuilder
func (m *_BACnetConstructedDataCOVIncrement) CreateBACnetConstructedDataCOVIncrementBuilder() BACnetConstructedDataCOVIncrementBuilder {
	if m == nil {
		return NewBACnetConstructedDataCOVIncrementBuilder()
	}
	return &_BACnetConstructedDataCOVIncrementBuilder{_BACnetConstructedDataCOVIncrement: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVIncrement) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVIncrement) GetCovIncrement() BACnetApplicationTagReal {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCOVIncrement) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVIncrement(structType any) BACnetConstructedDataCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataCOVIncrement"
}

func (m *_BACnetConstructedDataCOVIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCOVIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCOVIncrement BACnetConstructedDataCOVIncrement, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covIncrement, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "covIncrement", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	m.CovIncrement = covIncrement

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), covIncrement)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVIncrement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCOVIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVIncrement")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "covIncrement", m.GetCovIncrement(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVIncrement")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVIncrement) IsBACnetConstructedDataCOVIncrement() {}

func (m *_BACnetConstructedDataCOVIncrement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCOVIncrement) deepCopy() *_BACnetConstructedDataCOVIncrement {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCOVIncrementCopy := &_BACnetConstructedDataCOVIncrement{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.CovIncrement.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCOVIncrementCopy
}

func (m *_BACnetConstructedDataCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
