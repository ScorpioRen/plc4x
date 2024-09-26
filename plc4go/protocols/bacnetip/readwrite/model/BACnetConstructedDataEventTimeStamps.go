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

// BACnetConstructedDataEventTimeStamps is the corresponding interface of BACnetConstructedDataEventTimeStamps
type BACnetConstructedDataEventTimeStamps interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetEventTimeStamps returns EventTimeStamps (property field)
	GetEventTimeStamps() []BACnetTimeStamp
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormal returns ToOffnormal (virtual field)
	GetToOffnormal() BACnetTimeStamp
	// GetToFault returns ToFault (virtual field)
	GetToFault() BACnetTimeStamp
	// GetToNormal returns ToNormal (virtual field)
	GetToNormal() BACnetTimeStamp
	// IsBACnetConstructedDataEventTimeStamps is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventTimeStamps()
	// CreateBuilder creates a BACnetConstructedDataEventTimeStampsBuilder
	CreateBACnetConstructedDataEventTimeStampsBuilder() BACnetConstructedDataEventTimeStampsBuilder
}

// _BACnetConstructedDataEventTimeStamps is the data-structure of this message
type _BACnetConstructedDataEventTimeStamps struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	EventTimeStamps      []BACnetTimeStamp
}

var _ BACnetConstructedDataEventTimeStamps = (*_BACnetConstructedDataEventTimeStamps)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventTimeStamps)(nil)

// NewBACnetConstructedDataEventTimeStamps factory function for _BACnetConstructedDataEventTimeStamps
func NewBACnetConstructedDataEventTimeStamps(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, eventTimeStamps []BACnetTimeStamp, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventTimeStamps {
	_result := &_BACnetConstructedDataEventTimeStamps{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		EventTimeStamps:               eventTimeStamps,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventTimeStampsBuilder is a builder for BACnetConstructedDataEventTimeStamps
type BACnetConstructedDataEventTimeStampsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventTimeStamps []BACnetTimeStamp) BACnetConstructedDataEventTimeStampsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataEventTimeStampsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataEventTimeStampsBuilder
	// WithEventTimeStamps adds EventTimeStamps (property field)
	WithEventTimeStamps(...BACnetTimeStamp) BACnetConstructedDataEventTimeStampsBuilder
	// Build builds the BACnetConstructedDataEventTimeStamps or returns an error if something is wrong
	Build() (BACnetConstructedDataEventTimeStamps, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventTimeStamps
}

// NewBACnetConstructedDataEventTimeStampsBuilder() creates a BACnetConstructedDataEventTimeStampsBuilder
func NewBACnetConstructedDataEventTimeStampsBuilder() BACnetConstructedDataEventTimeStampsBuilder {
	return &_BACnetConstructedDataEventTimeStampsBuilder{_BACnetConstructedDataEventTimeStamps: new(_BACnetConstructedDataEventTimeStamps)}
}

type _BACnetConstructedDataEventTimeStampsBuilder struct {
	*_BACnetConstructedDataEventTimeStamps

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventTimeStampsBuilder) = (*_BACnetConstructedDataEventTimeStampsBuilder)(nil)

func (m *_BACnetConstructedDataEventTimeStampsBuilder) WithMandatoryFields(eventTimeStamps []BACnetTimeStamp) BACnetConstructedDataEventTimeStampsBuilder {
	return m.WithEventTimeStamps(eventTimeStamps...)
}

func (m *_BACnetConstructedDataEventTimeStampsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataEventTimeStampsBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataEventTimeStampsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataEventTimeStampsBuilder {
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

func (m *_BACnetConstructedDataEventTimeStampsBuilder) WithEventTimeStamps(eventTimeStamps ...BACnetTimeStamp) BACnetConstructedDataEventTimeStampsBuilder {
	m.EventTimeStamps = eventTimeStamps
	return m
}

func (m *_BACnetConstructedDataEventTimeStampsBuilder) Build() (BACnetConstructedDataEventTimeStamps, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataEventTimeStamps.deepCopy(), nil
}

func (m *_BACnetConstructedDataEventTimeStampsBuilder) MustBuild() BACnetConstructedDataEventTimeStamps {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataEventTimeStampsBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataEventTimeStampsBuilder()
}

// CreateBACnetConstructedDataEventTimeStampsBuilder creates a BACnetConstructedDataEventTimeStampsBuilder
func (m *_BACnetConstructedDataEventTimeStamps) CreateBACnetConstructedDataEventTimeStampsBuilder() BACnetConstructedDataEventTimeStampsBuilder {
	if m == nil {
		return NewBACnetConstructedDataEventTimeStampsBuilder()
	}
	return &_BACnetConstructedDataEventTimeStampsBuilder{_BACnetConstructedDataEventTimeStamps: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventTimeStamps) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_TIME_STAMPS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataEventTimeStamps) GetEventTimeStamps() []BACnetTimeStamp {
	return m.EventTimeStamps
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventTimeStamps) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *_BACnetConstructedDataEventTimeStamps) GetToOffnormal() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return CastBACnetTimeStamp(CastBACnetTimeStamp(utils.InlineIf(bool((len(m.GetEventTimeStamps())) == (3)), func() any { return CastBACnetTimeStamp(m.GetEventTimeStamps()[0]) }, func() any { return CastBACnetTimeStamp(nil) })))
}

func (m *_BACnetConstructedDataEventTimeStamps) GetToFault() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return CastBACnetTimeStamp(CastBACnetTimeStamp(utils.InlineIf(bool((len(m.GetEventTimeStamps())) == (3)), func() any { return CastBACnetTimeStamp(m.GetEventTimeStamps()[1]) }, func() any { return CastBACnetTimeStamp(nil) })))
}

func (m *_BACnetConstructedDataEventTimeStamps) GetToNormal() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return CastBACnetTimeStamp(CastBACnetTimeStamp(utils.InlineIf(bool((len(m.GetEventTimeStamps())) == (3)), func() any { return CastBACnetTimeStamp(m.GetEventTimeStamps()[2]) }, func() any { return CastBACnetTimeStamp(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventTimeStamps(structType any) BACnetConstructedDataEventTimeStamps {
	if casted, ok := structType.(BACnetConstructedDataEventTimeStamps); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventTimeStamps); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventTimeStamps) GetTypeName() string {
	return "BACnetConstructedDataEventTimeStamps"
}

func (m *_BACnetConstructedDataEventTimeStamps) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.EventTimeStamps) > 0 {
		for _, element := range m.EventTimeStamps {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventTimeStamps) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventTimeStamps) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventTimeStamps BACnetConstructedDataEventTimeStamps, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventTimeStamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventTimeStamps")
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

	eventTimeStamps, err := ReadTerminatedArrayField[BACnetTimeStamp](ctx, "eventTimeStamps", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventTimeStamps' field"))
	}
	m.EventTimeStamps = eventTimeStamps

	toOffnormal, err := ReadVirtualField[BACnetTimeStamp](ctx, "toOffnormal", (*BACnetTimeStamp)(nil), CastBACnetTimeStamp(utils.InlineIf(bool((len(eventTimeStamps)) == (3)), func() any { return CastBACnetTimeStamp(eventTimeStamps[0]) }, func() any { return CastBACnetTimeStamp(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormal' field"))
	}
	_ = toOffnormal

	toFault, err := ReadVirtualField[BACnetTimeStamp](ctx, "toFault", (*BACnetTimeStamp)(nil), CastBACnetTimeStamp(utils.InlineIf(bool((len(eventTimeStamps)) == (3)), func() any { return CastBACnetTimeStamp(eventTimeStamps[1]) }, func() any { return CastBACnetTimeStamp(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFault' field"))
	}
	_ = toFault

	toNormal, err := ReadVirtualField[BACnetTimeStamp](ctx, "toNormal", (*BACnetTimeStamp)(nil), CastBACnetTimeStamp(utils.InlineIf(bool((len(eventTimeStamps)) == (3)), func() any { return CastBACnetTimeStamp(eventTimeStamps[2]) }, func() any { return CastBACnetTimeStamp(nil) })))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormal' field"))
	}
	_ = toNormal

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventTimeStamps)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "eventTimeStamps should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventTimeStamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventTimeStamps")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventTimeStamps) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventTimeStamps) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventTimeStamps"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventTimeStamps")
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

		if err := WriteComplexTypeArrayField(ctx, "eventTimeStamps", m.GetEventTimeStamps(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'eventTimeStamps' field")
		}
		// Virtual field
		toOffnormal := m.GetToOffnormal()
		_ = toOffnormal
		if _toOffnormalErr := writeBuffer.WriteVirtual(ctx, "toOffnormal", m.GetToOffnormal()); _toOffnormalErr != nil {
			return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
		}
		// Virtual field
		toFault := m.GetToFault()
		_ = toFault
		if _toFaultErr := writeBuffer.WriteVirtual(ctx, "toFault", m.GetToFault()); _toFaultErr != nil {
			return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
		}
		// Virtual field
		toNormal := m.GetToNormal()
		_ = toNormal
		if _toNormalErr := writeBuffer.WriteVirtual(ctx, "toNormal", m.GetToNormal()); _toNormalErr != nil {
			return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventTimeStamps"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventTimeStamps")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventTimeStamps) IsBACnetConstructedDataEventTimeStamps() {}

func (m *_BACnetConstructedDataEventTimeStamps) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventTimeStamps) deepCopy() *_BACnetConstructedDataEventTimeStamps {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventTimeStampsCopy := &_BACnetConstructedDataEventTimeStamps{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetTimeStamp, BACnetTimeStamp](m.EventTimeStamps),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventTimeStampsCopy
}

func (m *_BACnetConstructedDataEventTimeStamps) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
