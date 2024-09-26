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

// BACnetConstructedDataIntegerValueFaultLowLimit is the corresponding interface of BACnetConstructedDataIntegerValueFaultLowLimit
type BACnetConstructedDataIntegerValueFaultLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataIntegerValueFaultLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueFaultLowLimit()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueFaultLowLimitBuilder
	CreateBACnetConstructedDataIntegerValueFaultLowLimitBuilder() BACnetConstructedDataIntegerValueFaultLowLimitBuilder
}

// _BACnetConstructedDataIntegerValueFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataIntegerValueFaultLowLimit struct {
	BACnetConstructedDataContract
	FaultLowLimit BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataIntegerValueFaultLowLimit = (*_BACnetConstructedDataIntegerValueFaultLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueFaultLowLimit)(nil)

// NewBACnetConstructedDataIntegerValueFaultLowLimit factory function for _BACnetConstructedDataIntegerValueFaultLowLimit
func NewBACnetConstructedDataIntegerValueFaultLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultLowLimit BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueFaultLowLimit {
	if faultLowLimit == nil {
		panic("faultLowLimit of type BACnetApplicationTagSignedInteger for BACnetConstructedDataIntegerValueFaultLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueFaultLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultLowLimit:                 faultLowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueFaultLowLimitBuilder is a builder for BACnetConstructedDataIntegerValueFaultLowLimit
type BACnetConstructedDataIntegerValueFaultLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultLowLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueFaultLowLimitBuilder
	// WithFaultLowLimit adds FaultLowLimit (property field)
	WithFaultLowLimit(BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueFaultLowLimitBuilder
	// WithFaultLowLimitBuilder adds FaultLowLimit (property field) which is build by the builder
	WithFaultLowLimitBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueFaultLowLimitBuilder
	// Build builds the BACnetConstructedDataIntegerValueFaultLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueFaultLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueFaultLowLimit
}

// NewBACnetConstructedDataIntegerValueFaultLowLimitBuilder() creates a BACnetConstructedDataIntegerValueFaultLowLimitBuilder
func NewBACnetConstructedDataIntegerValueFaultLowLimitBuilder() BACnetConstructedDataIntegerValueFaultLowLimitBuilder {
	return &_BACnetConstructedDataIntegerValueFaultLowLimitBuilder{_BACnetConstructedDataIntegerValueFaultLowLimit: new(_BACnetConstructedDataIntegerValueFaultLowLimit)}
}

type _BACnetConstructedDataIntegerValueFaultLowLimitBuilder struct {
	*_BACnetConstructedDataIntegerValueFaultLowLimit

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueFaultLowLimitBuilder) = (*_BACnetConstructedDataIntegerValueFaultLowLimitBuilder)(nil)

func (m *_BACnetConstructedDataIntegerValueFaultLowLimitBuilder) WithMandatoryFields(faultLowLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueFaultLowLimitBuilder {
	return m.WithFaultLowLimit(faultLowLimit)
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimitBuilder) WithFaultLowLimit(faultLowLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueFaultLowLimitBuilder {
	m.FaultLowLimit = faultLowLimit
	return m
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimitBuilder) WithFaultLowLimitBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueFaultLowLimitBuilder {
	builder := builderSupplier(m.FaultLowLimit.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	m.FaultLowLimit, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimitBuilder) Build() (BACnetConstructedDataIntegerValueFaultLowLimit, error) {
	if m.FaultLowLimit == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'faultLowLimit' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataIntegerValueFaultLowLimit.deepCopy(), nil
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimitBuilder) MustBuild() BACnetConstructedDataIntegerValueFaultLowLimit {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimitBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataIntegerValueFaultLowLimitBuilder()
}

// CreateBACnetConstructedDataIntegerValueFaultLowLimitBuilder creates a BACnetConstructedDataIntegerValueFaultLowLimitBuilder
func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) CreateBACnetConstructedDataIntegerValueFaultLowLimitBuilder() BACnetConstructedDataIntegerValueFaultLowLimitBuilder {
	if m == nil {
		return NewBACnetConstructedDataIntegerValueFaultLowLimitBuilder()
	}
	return &_BACnetConstructedDataIntegerValueFaultLowLimitBuilder{_BACnetConstructedDataIntegerValueFaultLowLimit: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagSignedInteger {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueFaultLowLimit(structType any) BACnetConstructedDataIntegerValueFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueFaultLowLimit"
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueFaultLowLimit BACnetConstructedDataIntegerValueFaultLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultLowLimit, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "faultLowLimit", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultLowLimit' field"))
	}
	m.FaultLowLimit = faultLowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), faultLowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueFaultLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueFaultLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "faultLowLimit", m.GetFaultLowLimit(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueFaultLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) IsBACnetConstructedDataIntegerValueFaultLowLimit() {
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) deepCopy() *_BACnetConstructedDataIntegerValueFaultLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueFaultLowLimitCopy := &_BACnetConstructedDataIntegerValueFaultLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.FaultLowLimit.DeepCopy().(BACnetApplicationTagSignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueFaultLowLimitCopy
}

func (m *_BACnetConstructedDataIntegerValueFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
