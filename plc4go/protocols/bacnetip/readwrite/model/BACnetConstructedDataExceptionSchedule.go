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

// BACnetConstructedDataExceptionSchedule is the corresponding interface of BACnetConstructedDataExceptionSchedule
type BACnetConstructedDataExceptionSchedule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetExceptionSchedule returns ExceptionSchedule (property field)
	GetExceptionSchedule() []BACnetSpecialEvent
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataExceptionSchedule is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataExceptionSchedule()
	// CreateBuilder creates a BACnetConstructedDataExceptionScheduleBuilder
	CreateBACnetConstructedDataExceptionScheduleBuilder() BACnetConstructedDataExceptionScheduleBuilder
}

// _BACnetConstructedDataExceptionSchedule is the data-structure of this message
type _BACnetConstructedDataExceptionSchedule struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ExceptionSchedule    []BACnetSpecialEvent
}

var _ BACnetConstructedDataExceptionSchedule = (*_BACnetConstructedDataExceptionSchedule)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataExceptionSchedule)(nil)

// NewBACnetConstructedDataExceptionSchedule factory function for _BACnetConstructedDataExceptionSchedule
func NewBACnetConstructedDataExceptionSchedule(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, exceptionSchedule []BACnetSpecialEvent, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExceptionSchedule {
	_result := &_BACnetConstructedDataExceptionSchedule{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ExceptionSchedule:             exceptionSchedule,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataExceptionScheduleBuilder is a builder for BACnetConstructedDataExceptionSchedule
type BACnetConstructedDataExceptionScheduleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(exceptionSchedule []BACnetSpecialEvent) BACnetConstructedDataExceptionScheduleBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExceptionScheduleBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataExceptionScheduleBuilder
	// WithExceptionSchedule adds ExceptionSchedule (property field)
	WithExceptionSchedule(...BACnetSpecialEvent) BACnetConstructedDataExceptionScheduleBuilder
	// Build builds the BACnetConstructedDataExceptionSchedule or returns an error if something is wrong
	Build() (BACnetConstructedDataExceptionSchedule, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataExceptionSchedule
}

// NewBACnetConstructedDataExceptionScheduleBuilder() creates a BACnetConstructedDataExceptionScheduleBuilder
func NewBACnetConstructedDataExceptionScheduleBuilder() BACnetConstructedDataExceptionScheduleBuilder {
	return &_BACnetConstructedDataExceptionScheduleBuilder{_BACnetConstructedDataExceptionSchedule: new(_BACnetConstructedDataExceptionSchedule)}
}

type _BACnetConstructedDataExceptionScheduleBuilder struct {
	*_BACnetConstructedDataExceptionSchedule

	err *utils.MultiError
}

var _ (BACnetConstructedDataExceptionScheduleBuilder) = (*_BACnetConstructedDataExceptionScheduleBuilder)(nil)

func (m *_BACnetConstructedDataExceptionScheduleBuilder) WithMandatoryFields(exceptionSchedule []BACnetSpecialEvent) BACnetConstructedDataExceptionScheduleBuilder {
	return m.WithExceptionSchedule(exceptionSchedule...)
}

func (m *_BACnetConstructedDataExceptionScheduleBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataExceptionScheduleBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataExceptionScheduleBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataExceptionScheduleBuilder {
	builder := builderSupplier(m.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataExceptionScheduleBuilder) WithExceptionSchedule(exceptionSchedule ...BACnetSpecialEvent) BACnetConstructedDataExceptionScheduleBuilder {
	m.ExceptionSchedule = exceptionSchedule
	return m
}

func (m *_BACnetConstructedDataExceptionScheduleBuilder) Build() (BACnetConstructedDataExceptionSchedule, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataExceptionSchedule.deepCopy(), nil
}

func (m *_BACnetConstructedDataExceptionScheduleBuilder) MustBuild() BACnetConstructedDataExceptionSchedule {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataExceptionScheduleBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataExceptionScheduleBuilder()
}

// CreateBACnetConstructedDataExceptionScheduleBuilder creates a BACnetConstructedDataExceptionScheduleBuilder
func (m *_BACnetConstructedDataExceptionSchedule) CreateBACnetConstructedDataExceptionScheduleBuilder() BACnetConstructedDataExceptionScheduleBuilder {
	if m == nil {
		return NewBACnetConstructedDataExceptionScheduleBuilder()
	}
	return &_BACnetConstructedDataExceptionScheduleBuilder{_BACnetConstructedDataExceptionSchedule: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExceptionSchedule) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXCEPTION_SCHEDULE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataExceptionSchedule) GetExceptionSchedule() []BACnetSpecialEvent {
	return m.ExceptionSchedule
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExceptionSchedule(structType any) BACnetConstructedDataExceptionSchedule {
	if casted, ok := structType.(BACnetConstructedDataExceptionSchedule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExceptionSchedule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExceptionSchedule) GetTypeName() string {
	return "BACnetConstructedDataExceptionSchedule"
}

func (m *_BACnetConstructedDataExceptionSchedule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ExceptionSchedule) > 0 {
		for _, element := range m.ExceptionSchedule {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataExceptionSchedule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataExceptionSchedule) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataExceptionSchedule BACnetConstructedDataExceptionSchedule, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExceptionSchedule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExceptionSchedule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	exceptionSchedule, err := ReadTerminatedArrayField[BACnetSpecialEvent](ctx, "exceptionSchedule", ReadComplex[BACnetSpecialEvent](BACnetSpecialEventParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceptionSchedule' field"))
	}
	m.ExceptionSchedule = exceptionSchedule

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExceptionSchedule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExceptionSchedule")
	}

	return m, nil
}

func (m *_BACnetConstructedDataExceptionSchedule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataExceptionSchedule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExceptionSchedule"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExceptionSchedule")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "exceptionSchedule", m.GetExceptionSchedule(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'exceptionSchedule' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExceptionSchedule"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExceptionSchedule")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExceptionSchedule) IsBACnetConstructedDataExceptionSchedule() {}

func (m *_BACnetConstructedDataExceptionSchedule) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataExceptionSchedule) deepCopy() *_BACnetConstructedDataExceptionSchedule {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataExceptionScheduleCopy := &_BACnetConstructedDataExceptionSchedule{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetSpecialEvent, BACnetSpecialEvent](m.ExceptionSchedule),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataExceptionScheduleCopy
}

func (m *_BACnetConstructedDataExceptionSchedule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
