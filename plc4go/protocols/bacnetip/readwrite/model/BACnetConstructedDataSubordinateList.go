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

// BACnetConstructedDataSubordinateList is the corresponding interface of BACnetConstructedDataSubordinateList
type BACnetConstructedDataSubordinateList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSubordinateList returns SubordinateList (property field)
	GetSubordinateList() []BACnetDeviceObjectReference
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataSubordinateList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSubordinateList()
	// CreateBuilder creates a BACnetConstructedDataSubordinateListBuilder
	CreateBACnetConstructedDataSubordinateListBuilder() BACnetConstructedDataSubordinateListBuilder
}

// _BACnetConstructedDataSubordinateList is the data-structure of this message
type _BACnetConstructedDataSubordinateList struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	SubordinateList      []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataSubordinateList = (*_BACnetConstructedDataSubordinateList)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSubordinateList)(nil)

// NewBACnetConstructedDataSubordinateList factory function for _BACnetConstructedDataSubordinateList
func NewBACnetConstructedDataSubordinateList(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, subordinateList []BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubordinateList {
	_result := &_BACnetConstructedDataSubordinateList{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		SubordinateList:               subordinateList,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSubordinateListBuilder is a builder for BACnetConstructedDataSubordinateList
type BACnetConstructedDataSubordinateListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subordinateList []BACnetDeviceObjectReference) BACnetConstructedDataSubordinateListBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSubordinateListBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSubordinateListBuilder
	// WithSubordinateList adds SubordinateList (property field)
	WithSubordinateList(...BACnetDeviceObjectReference) BACnetConstructedDataSubordinateListBuilder
	// Build builds the BACnetConstructedDataSubordinateList or returns an error if something is wrong
	Build() (BACnetConstructedDataSubordinateList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSubordinateList
}

// NewBACnetConstructedDataSubordinateListBuilder() creates a BACnetConstructedDataSubordinateListBuilder
func NewBACnetConstructedDataSubordinateListBuilder() BACnetConstructedDataSubordinateListBuilder {
	return &_BACnetConstructedDataSubordinateListBuilder{_BACnetConstructedDataSubordinateList: new(_BACnetConstructedDataSubordinateList)}
}

type _BACnetConstructedDataSubordinateListBuilder struct {
	*_BACnetConstructedDataSubordinateList

	err *utils.MultiError
}

var _ (BACnetConstructedDataSubordinateListBuilder) = (*_BACnetConstructedDataSubordinateListBuilder)(nil)

func (m *_BACnetConstructedDataSubordinateListBuilder) WithMandatoryFields(subordinateList []BACnetDeviceObjectReference) BACnetConstructedDataSubordinateListBuilder {
	return m.WithSubordinateList(subordinateList...)
}

func (m *_BACnetConstructedDataSubordinateListBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSubordinateListBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataSubordinateListBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSubordinateListBuilder {
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

func (m *_BACnetConstructedDataSubordinateListBuilder) WithSubordinateList(subordinateList ...BACnetDeviceObjectReference) BACnetConstructedDataSubordinateListBuilder {
	m.SubordinateList = subordinateList
	return m
}

func (m *_BACnetConstructedDataSubordinateListBuilder) Build() (BACnetConstructedDataSubordinateList, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataSubordinateList.deepCopy(), nil
}

func (m *_BACnetConstructedDataSubordinateListBuilder) MustBuild() BACnetConstructedDataSubordinateList {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataSubordinateListBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataSubordinateListBuilder()
}

// CreateBACnetConstructedDataSubordinateListBuilder creates a BACnetConstructedDataSubordinateListBuilder
func (m *_BACnetConstructedDataSubordinateList) CreateBACnetConstructedDataSubordinateListBuilder() BACnetConstructedDataSubordinateListBuilder {
	if m == nil {
		return NewBACnetConstructedDataSubordinateListBuilder()
	}
	return &_BACnetConstructedDataSubordinateListBuilder{_BACnetConstructedDataSubordinateList: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubordinateList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSubordinateList) GetSubordinateList() []BACnetDeviceObjectReference {
	return m.SubordinateList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetZero() uint64 {
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
func CastBACnetConstructedDataSubordinateList(structType any) BACnetConstructedDataSubordinateList {
	if casted, ok := structType.(BACnetConstructedDataSubordinateList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubordinateList) GetTypeName() string {
	return "BACnetConstructedDataSubordinateList"
}

func (m *_BACnetConstructedDataSubordinateList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.SubordinateList) > 0 {
		for _, element := range m.SubordinateList {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubordinateList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSubordinateList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSubordinateList BACnetConstructedDataSubordinateList, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateList")
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

	subordinateList, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "subordinateList", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subordinateList' field"))
	}
	m.SubordinateList = subordinateList

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateList")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSubordinateList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSubordinateList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateList")
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

		if err := WriteComplexTypeArrayField(ctx, "subordinateList", m.GetSubordinateList(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'subordinateList' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateList")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubordinateList) IsBACnetConstructedDataSubordinateList() {}

func (m *_BACnetConstructedDataSubordinateList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSubordinateList) deepCopy() *_BACnetConstructedDataSubordinateList {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSubordinateListCopy := &_BACnetConstructedDataSubordinateList{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetDeviceObjectReference, BACnetDeviceObjectReference](m.SubordinateList),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSubordinateListCopy
}

func (m *_BACnetConstructedDataSubordinateList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
