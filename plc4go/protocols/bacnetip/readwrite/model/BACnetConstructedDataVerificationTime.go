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

// BACnetConstructedDataVerificationTime is the corresponding interface of BACnetConstructedDataVerificationTime
type BACnetConstructedDataVerificationTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetVerificationTime returns VerificationTime (property field)
	GetVerificationTime() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataVerificationTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataVerificationTime()
	// CreateBuilder creates a BACnetConstructedDataVerificationTimeBuilder
	CreateBACnetConstructedDataVerificationTimeBuilder() BACnetConstructedDataVerificationTimeBuilder
}

// _BACnetConstructedDataVerificationTime is the data-structure of this message
type _BACnetConstructedDataVerificationTime struct {
	BACnetConstructedDataContract
	VerificationTime BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataVerificationTime = (*_BACnetConstructedDataVerificationTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataVerificationTime)(nil)

// NewBACnetConstructedDataVerificationTime factory function for _BACnetConstructedDataVerificationTime
func NewBACnetConstructedDataVerificationTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, verificationTime BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVerificationTime {
	if verificationTime == nil {
		panic("verificationTime of type BACnetApplicationTagSignedInteger for BACnetConstructedDataVerificationTime must not be nil")
	}
	_result := &_BACnetConstructedDataVerificationTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		VerificationTime:              verificationTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataVerificationTimeBuilder is a builder for BACnetConstructedDataVerificationTime
type BACnetConstructedDataVerificationTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(verificationTime BACnetApplicationTagSignedInteger) BACnetConstructedDataVerificationTimeBuilder
	// WithVerificationTime adds VerificationTime (property field)
	WithVerificationTime(BACnetApplicationTagSignedInteger) BACnetConstructedDataVerificationTimeBuilder
	// WithVerificationTimeBuilder adds VerificationTime (property field) which is build by the builder
	WithVerificationTimeBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataVerificationTimeBuilder
	// Build builds the BACnetConstructedDataVerificationTime or returns an error if something is wrong
	Build() (BACnetConstructedDataVerificationTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataVerificationTime
}

// NewBACnetConstructedDataVerificationTimeBuilder() creates a BACnetConstructedDataVerificationTimeBuilder
func NewBACnetConstructedDataVerificationTimeBuilder() BACnetConstructedDataVerificationTimeBuilder {
	return &_BACnetConstructedDataVerificationTimeBuilder{_BACnetConstructedDataVerificationTime: new(_BACnetConstructedDataVerificationTime)}
}

type _BACnetConstructedDataVerificationTimeBuilder struct {
	*_BACnetConstructedDataVerificationTime

	err *utils.MultiError
}

var _ (BACnetConstructedDataVerificationTimeBuilder) = (*_BACnetConstructedDataVerificationTimeBuilder)(nil)

func (m *_BACnetConstructedDataVerificationTimeBuilder) WithMandatoryFields(verificationTime BACnetApplicationTagSignedInteger) BACnetConstructedDataVerificationTimeBuilder {
	return m.WithVerificationTime(verificationTime)
}

func (m *_BACnetConstructedDataVerificationTimeBuilder) WithVerificationTime(verificationTime BACnetApplicationTagSignedInteger) BACnetConstructedDataVerificationTimeBuilder {
	m.VerificationTime = verificationTime
	return m
}

func (m *_BACnetConstructedDataVerificationTimeBuilder) WithVerificationTimeBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataVerificationTimeBuilder {
	builder := builderSupplier(m.VerificationTime.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	m.VerificationTime, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataVerificationTimeBuilder) Build() (BACnetConstructedDataVerificationTime, error) {
	if m.VerificationTime == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'verificationTime' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataVerificationTime.deepCopy(), nil
}

func (m *_BACnetConstructedDataVerificationTimeBuilder) MustBuild() BACnetConstructedDataVerificationTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataVerificationTimeBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataVerificationTimeBuilder()
}

// CreateBACnetConstructedDataVerificationTimeBuilder creates a BACnetConstructedDataVerificationTimeBuilder
func (m *_BACnetConstructedDataVerificationTime) CreateBACnetConstructedDataVerificationTimeBuilder() BACnetConstructedDataVerificationTimeBuilder {
	if m == nil {
		return NewBACnetConstructedDataVerificationTimeBuilder()
	}
	return &_BACnetConstructedDataVerificationTimeBuilder{_BACnetConstructedDataVerificationTime: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVerificationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VERIFICATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetVerificationTime() BACnetApplicationTagSignedInteger {
	return m.VerificationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVerificationTime) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetVerificationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVerificationTime(structType any) BACnetConstructedDataVerificationTime {
	if casted, ok := structType.(BACnetConstructedDataVerificationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVerificationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVerificationTime) GetTypeName() string {
	return "BACnetConstructedDataVerificationTime"
}

func (m *_BACnetConstructedDataVerificationTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (verificationTime)
	lengthInBits += m.VerificationTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataVerificationTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataVerificationTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataVerificationTime BACnetConstructedDataVerificationTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVerificationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVerificationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	verificationTime, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "verificationTime", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'verificationTime' field"))
	}
	m.VerificationTime = verificationTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), verificationTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVerificationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVerificationTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataVerificationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVerificationTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVerificationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVerificationTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "verificationTime", m.GetVerificationTime(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'verificationTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVerificationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVerificationTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVerificationTime) IsBACnetConstructedDataVerificationTime() {}

func (m *_BACnetConstructedDataVerificationTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataVerificationTime) deepCopy() *_BACnetConstructedDataVerificationTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataVerificationTimeCopy := &_BACnetConstructedDataVerificationTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.VerificationTime.DeepCopy().(BACnetApplicationTagSignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataVerificationTimeCopy
}

func (m *_BACnetConstructedDataVerificationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
