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

// BACnetConstructedDataAuthenticationPolicyList is the corresponding interface of BACnetConstructedDataAuthenticationPolicyList
type BACnetConstructedDataAuthenticationPolicyList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationPolicyList returns AuthenticationPolicyList (property field)
	GetAuthenticationPolicyList() []BACnetAuthenticationPolicy
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataAuthenticationPolicyList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAuthenticationPolicyList()
	// CreateBuilder creates a BACnetConstructedDataAuthenticationPolicyListBuilder
	CreateBACnetConstructedDataAuthenticationPolicyListBuilder() BACnetConstructedDataAuthenticationPolicyListBuilder
}

// _BACnetConstructedDataAuthenticationPolicyList is the data-structure of this message
type _BACnetConstructedDataAuthenticationPolicyList struct {
	BACnetConstructedDataContract
	NumberOfDataElements     BACnetApplicationTagUnsignedInteger
	AuthenticationPolicyList []BACnetAuthenticationPolicy
}

var _ BACnetConstructedDataAuthenticationPolicyList = (*_BACnetConstructedDataAuthenticationPolicyList)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAuthenticationPolicyList)(nil)

// NewBACnetConstructedDataAuthenticationPolicyList factory function for _BACnetConstructedDataAuthenticationPolicyList
func NewBACnetConstructedDataAuthenticationPolicyList(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, authenticationPolicyList []BACnetAuthenticationPolicy, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationPolicyList {
	_result := &_BACnetConstructedDataAuthenticationPolicyList{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AuthenticationPolicyList:      authenticationPolicyList,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAuthenticationPolicyListBuilder is a builder for BACnetConstructedDataAuthenticationPolicyList
type BACnetConstructedDataAuthenticationPolicyListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(authenticationPolicyList []BACnetAuthenticationPolicy) BACnetConstructedDataAuthenticationPolicyListBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAuthenticationPolicyListBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAuthenticationPolicyListBuilder
	// WithAuthenticationPolicyList adds AuthenticationPolicyList (property field)
	WithAuthenticationPolicyList(...BACnetAuthenticationPolicy) BACnetConstructedDataAuthenticationPolicyListBuilder
	// Build builds the BACnetConstructedDataAuthenticationPolicyList or returns an error if something is wrong
	Build() (BACnetConstructedDataAuthenticationPolicyList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAuthenticationPolicyList
}

// NewBACnetConstructedDataAuthenticationPolicyListBuilder() creates a BACnetConstructedDataAuthenticationPolicyListBuilder
func NewBACnetConstructedDataAuthenticationPolicyListBuilder() BACnetConstructedDataAuthenticationPolicyListBuilder {
	return &_BACnetConstructedDataAuthenticationPolicyListBuilder{_BACnetConstructedDataAuthenticationPolicyList: new(_BACnetConstructedDataAuthenticationPolicyList)}
}

type _BACnetConstructedDataAuthenticationPolicyListBuilder struct {
	*_BACnetConstructedDataAuthenticationPolicyList

	err *utils.MultiError
}

var _ (BACnetConstructedDataAuthenticationPolicyListBuilder) = (*_BACnetConstructedDataAuthenticationPolicyListBuilder)(nil)

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) WithMandatoryFields(authenticationPolicyList []BACnetAuthenticationPolicy) BACnetConstructedDataAuthenticationPolicyListBuilder {
	return m.WithAuthenticationPolicyList(authenticationPolicyList...)
}

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAuthenticationPolicyListBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAuthenticationPolicyListBuilder {
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

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) WithAuthenticationPolicyList(authenticationPolicyList ...BACnetAuthenticationPolicy) BACnetConstructedDataAuthenticationPolicyListBuilder {
	m.AuthenticationPolicyList = authenticationPolicyList
	return m
}

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) Build() (BACnetConstructedDataAuthenticationPolicyList, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataAuthenticationPolicyList.deepCopy(), nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) MustBuild() BACnetConstructedDataAuthenticationPolicyList {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataAuthenticationPolicyListBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataAuthenticationPolicyListBuilder()
}

// CreateBACnetConstructedDataAuthenticationPolicyListBuilder creates a BACnetConstructedDataAuthenticationPolicyListBuilder
func (m *_BACnetConstructedDataAuthenticationPolicyList) CreateBACnetConstructedDataAuthenticationPolicyListBuilder() BACnetConstructedDataAuthenticationPolicyListBuilder {
	if m == nil {
		return NewBACnetConstructedDataAuthenticationPolicyListBuilder()
	}
	return &_BACnetConstructedDataAuthenticationPolicyListBuilder{_BACnetConstructedDataAuthenticationPolicyList: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_POLICY_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetAuthenticationPolicyList() []BACnetAuthenticationPolicy {
	return m.AuthenticationPolicyList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetZero() uint64 {
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
func CastBACnetConstructedDataAuthenticationPolicyList(structType any) BACnetConstructedDataAuthenticationPolicyList {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationPolicyList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationPolicyList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationPolicyList"
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AuthenticationPolicyList) > 0 {
		for _, element := range m.AuthenticationPolicyList {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAuthenticationPolicyList BACnetConstructedDataAuthenticationPolicyList, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationPolicyList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationPolicyList")
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

	authenticationPolicyList, err := ReadTerminatedArrayField[BACnetAuthenticationPolicy](ctx, "authenticationPolicyList", ReadComplex[BACnetAuthenticationPolicy](BACnetAuthenticationPolicyParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationPolicyList' field"))
	}
	m.AuthenticationPolicyList = authenticationPolicyList

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationPolicyList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationPolicyList")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationPolicyList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationPolicyList")
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

		if err := WriteComplexTypeArrayField(ctx, "authenticationPolicyList", m.GetAuthenticationPolicyList(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationPolicyList' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationPolicyList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationPolicyList")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) IsBACnetConstructedDataAuthenticationPolicyList() {
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) deepCopy() *_BACnetConstructedDataAuthenticationPolicyList {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAuthenticationPolicyListCopy := &_BACnetConstructedDataAuthenticationPolicyList{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetAuthenticationPolicy, BACnetAuthenticationPolicy](m.AuthenticationPolicyList),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAuthenticationPolicyListCopy
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
