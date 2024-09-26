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

// BACnetConstructedDataCarDoorStatus is the corresponding interface of BACnetConstructedDataCarDoorStatus
type BACnetConstructedDataCarDoorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetCarDoorStatus returns CarDoorStatus (property field)
	GetCarDoorStatus() []BACnetDoorStatusTagged
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataCarDoorStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCarDoorStatus()
	// CreateBuilder creates a BACnetConstructedDataCarDoorStatusBuilder
	CreateBACnetConstructedDataCarDoorStatusBuilder() BACnetConstructedDataCarDoorStatusBuilder
}

// _BACnetConstructedDataCarDoorStatus is the data-structure of this message
type _BACnetConstructedDataCarDoorStatus struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	CarDoorStatus        []BACnetDoorStatusTagged
}

var _ BACnetConstructedDataCarDoorStatus = (*_BACnetConstructedDataCarDoorStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCarDoorStatus)(nil)

// NewBACnetConstructedDataCarDoorStatus factory function for _BACnetConstructedDataCarDoorStatus
func NewBACnetConstructedDataCarDoorStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, carDoorStatus []BACnetDoorStatusTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarDoorStatus {
	_result := &_BACnetConstructedDataCarDoorStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		CarDoorStatus:                 carDoorStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCarDoorStatusBuilder is a builder for BACnetConstructedDataCarDoorStatus
type BACnetConstructedDataCarDoorStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(carDoorStatus []BACnetDoorStatusTagged) BACnetConstructedDataCarDoorStatusBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCarDoorStatusBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCarDoorStatusBuilder
	// WithCarDoorStatus adds CarDoorStatus (property field)
	WithCarDoorStatus(...BACnetDoorStatusTagged) BACnetConstructedDataCarDoorStatusBuilder
	// Build builds the BACnetConstructedDataCarDoorStatus or returns an error if something is wrong
	Build() (BACnetConstructedDataCarDoorStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCarDoorStatus
}

// NewBACnetConstructedDataCarDoorStatusBuilder() creates a BACnetConstructedDataCarDoorStatusBuilder
func NewBACnetConstructedDataCarDoorStatusBuilder() BACnetConstructedDataCarDoorStatusBuilder {
	return &_BACnetConstructedDataCarDoorStatusBuilder{_BACnetConstructedDataCarDoorStatus: new(_BACnetConstructedDataCarDoorStatus)}
}

type _BACnetConstructedDataCarDoorStatusBuilder struct {
	*_BACnetConstructedDataCarDoorStatus

	err *utils.MultiError
}

var _ (BACnetConstructedDataCarDoorStatusBuilder) = (*_BACnetConstructedDataCarDoorStatusBuilder)(nil)

func (m *_BACnetConstructedDataCarDoorStatusBuilder) WithMandatoryFields(carDoorStatus []BACnetDoorStatusTagged) BACnetConstructedDataCarDoorStatusBuilder {
	return m.WithCarDoorStatus(carDoorStatus...)
}

func (m *_BACnetConstructedDataCarDoorStatusBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCarDoorStatusBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataCarDoorStatusBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCarDoorStatusBuilder {
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

func (m *_BACnetConstructedDataCarDoorStatusBuilder) WithCarDoorStatus(carDoorStatus ...BACnetDoorStatusTagged) BACnetConstructedDataCarDoorStatusBuilder {
	m.CarDoorStatus = carDoorStatus
	return m
}

func (m *_BACnetConstructedDataCarDoorStatusBuilder) Build() (BACnetConstructedDataCarDoorStatus, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataCarDoorStatus.deepCopy(), nil
}

func (m *_BACnetConstructedDataCarDoorStatusBuilder) MustBuild() BACnetConstructedDataCarDoorStatus {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataCarDoorStatusBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataCarDoorStatusBuilder()
}

// CreateBACnetConstructedDataCarDoorStatusBuilder creates a BACnetConstructedDataCarDoorStatusBuilder
func (m *_BACnetConstructedDataCarDoorStatus) CreateBACnetConstructedDataCarDoorStatusBuilder() BACnetConstructedDataCarDoorStatusBuilder {
	if m == nil {
		return NewBACnetConstructedDataCarDoorStatusBuilder()
	}
	return &_BACnetConstructedDataCarDoorStatusBuilder{_BACnetConstructedDataCarDoorStatus: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarDoorStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarDoorStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_DOOR_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarDoorStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarDoorStatus) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataCarDoorStatus) GetCarDoorStatus() []BACnetDoorStatusTagged {
	return m.CarDoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarDoorStatus) GetZero() uint64 {
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
func CastBACnetConstructedDataCarDoorStatus(structType any) BACnetConstructedDataCarDoorStatus {
	if casted, ok := structType.(BACnetConstructedDataCarDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarDoorStatus) GetTypeName() string {
	return "BACnetConstructedDataCarDoorStatus"
}

func (m *_BACnetConstructedDataCarDoorStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.CarDoorStatus) > 0 {
		for _, element := range m.CarDoorStatus {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCarDoorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCarDoorStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCarDoorStatus BACnetConstructedDataCarDoorStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarDoorStatus")
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

	carDoorStatus, err := ReadTerminatedArrayField[BACnetDoorStatusTagged](ctx, "carDoorStatus", ReadComplex[BACnetDoorStatusTagged](BACnetDoorStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'carDoorStatus' field"))
	}
	m.CarDoorStatus = carDoorStatus

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarDoorStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCarDoorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarDoorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarDoorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarDoorStatus")
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

		if err := WriteComplexTypeArrayField(ctx, "carDoorStatus", m.GetCarDoorStatus(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'carDoorStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarDoorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarDoorStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarDoorStatus) IsBACnetConstructedDataCarDoorStatus() {}

func (m *_BACnetConstructedDataCarDoorStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCarDoorStatus) deepCopy() *_BACnetConstructedDataCarDoorStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCarDoorStatusCopy := &_BACnetConstructedDataCarDoorStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetDoorStatusTagged, BACnetDoorStatusTagged](m.CarDoorStatus),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCarDoorStatusCopy
}

func (m *_BACnetConstructedDataCarDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
