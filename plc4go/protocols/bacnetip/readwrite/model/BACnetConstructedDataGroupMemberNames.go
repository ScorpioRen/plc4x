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

// BACnetConstructedDataGroupMemberNames is the corresponding interface of BACnetConstructedDataGroupMemberNames
type BACnetConstructedDataGroupMemberNames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetGroupMemberNames returns GroupMemberNames (property field)
	GetGroupMemberNames() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataGroupMemberNames is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGroupMemberNames()
	// CreateBuilder creates a BACnetConstructedDataGroupMemberNamesBuilder
	CreateBACnetConstructedDataGroupMemberNamesBuilder() BACnetConstructedDataGroupMemberNamesBuilder
}

// _BACnetConstructedDataGroupMemberNames is the data-structure of this message
type _BACnetConstructedDataGroupMemberNames struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	GroupMemberNames     []BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataGroupMemberNames = (*_BACnetConstructedDataGroupMemberNames)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupMemberNames)(nil)

// NewBACnetConstructedDataGroupMemberNames factory function for _BACnetConstructedDataGroupMemberNames
func NewBACnetConstructedDataGroupMemberNames(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, groupMemberNames []BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupMemberNames {
	_result := &_BACnetConstructedDataGroupMemberNames{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		GroupMemberNames:              groupMemberNames,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGroupMemberNamesBuilder is a builder for BACnetConstructedDataGroupMemberNames
type BACnetConstructedDataGroupMemberNamesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(groupMemberNames []BACnetApplicationTagCharacterString) BACnetConstructedDataGroupMemberNamesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupMemberNamesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGroupMemberNamesBuilder
	// WithGroupMemberNames adds GroupMemberNames (property field)
	WithGroupMemberNames(...BACnetApplicationTagCharacterString) BACnetConstructedDataGroupMemberNamesBuilder
	// Build builds the BACnetConstructedDataGroupMemberNames or returns an error if something is wrong
	Build() (BACnetConstructedDataGroupMemberNames, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGroupMemberNames
}

// NewBACnetConstructedDataGroupMemberNamesBuilder() creates a BACnetConstructedDataGroupMemberNamesBuilder
func NewBACnetConstructedDataGroupMemberNamesBuilder() BACnetConstructedDataGroupMemberNamesBuilder {
	return &_BACnetConstructedDataGroupMemberNamesBuilder{_BACnetConstructedDataGroupMemberNames: new(_BACnetConstructedDataGroupMemberNames)}
}

type _BACnetConstructedDataGroupMemberNamesBuilder struct {
	*_BACnetConstructedDataGroupMemberNames

	err *utils.MultiError
}

var _ (BACnetConstructedDataGroupMemberNamesBuilder) = (*_BACnetConstructedDataGroupMemberNamesBuilder)(nil)

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) WithMandatoryFields(groupMemberNames []BACnetApplicationTagCharacterString) BACnetConstructedDataGroupMemberNamesBuilder {
	return m.WithGroupMemberNames(groupMemberNames...)
}

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupMemberNamesBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGroupMemberNamesBuilder {
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

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) WithGroupMemberNames(groupMemberNames ...BACnetApplicationTagCharacterString) BACnetConstructedDataGroupMemberNamesBuilder {
	m.GroupMemberNames = groupMemberNames
	return m
}

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) Build() (BACnetConstructedDataGroupMemberNames, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataGroupMemberNames.deepCopy(), nil
}

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) MustBuild() BACnetConstructedDataGroupMemberNames {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataGroupMemberNamesBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataGroupMemberNamesBuilder()
}

// CreateBACnetConstructedDataGroupMemberNamesBuilder creates a BACnetConstructedDataGroupMemberNamesBuilder
func (m *_BACnetConstructedDataGroupMemberNames) CreateBACnetConstructedDataGroupMemberNamesBuilder() BACnetConstructedDataGroupMemberNamesBuilder {
	if m == nil {
		return NewBACnetConstructedDataGroupMemberNamesBuilder()
	}
	return &_BACnetConstructedDataGroupMemberNamesBuilder{_BACnetConstructedDataGroupMemberNames: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupMemberNames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupMemberNames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MEMBER_NAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupMemberNames) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMemberNames) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataGroupMemberNames) GetGroupMemberNames() []BACnetApplicationTagCharacterString {
	return m.GroupMemberNames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMemberNames) GetZero() uint64 {
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
func CastBACnetConstructedDataGroupMemberNames(structType any) BACnetConstructedDataGroupMemberNames {
	if casted, ok := structType.(BACnetConstructedDataGroupMemberNames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupMemberNames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupMemberNames) GetTypeName() string {
	return "BACnetConstructedDataGroupMemberNames"
}

func (m *_BACnetConstructedDataGroupMemberNames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.GroupMemberNames) > 0 {
		for _, element := range m.GroupMemberNames {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupMemberNames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupMemberNames) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupMemberNames BACnetConstructedDataGroupMemberNames, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupMemberNames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupMemberNames")
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

	groupMemberNames, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "groupMemberNames", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupMemberNames' field"))
	}
	m.GroupMemberNames = groupMemberNames

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupMemberNames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupMemberNames")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupMemberNames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupMemberNames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupMemberNames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupMemberNames")
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

		if err := WriteComplexTypeArrayField(ctx, "groupMemberNames", m.GetGroupMemberNames(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'groupMemberNames' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupMemberNames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupMemberNames")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupMemberNames) IsBACnetConstructedDataGroupMemberNames() {}

func (m *_BACnetConstructedDataGroupMemberNames) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGroupMemberNames) deepCopy() *_BACnetConstructedDataGroupMemberNames {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGroupMemberNamesCopy := &_BACnetConstructedDataGroupMemberNames{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagCharacterString, BACnetApplicationTagCharacterString](m.GroupMemberNames),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGroupMemberNamesCopy
}

func (m *_BACnetConstructedDataGroupMemberNames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
