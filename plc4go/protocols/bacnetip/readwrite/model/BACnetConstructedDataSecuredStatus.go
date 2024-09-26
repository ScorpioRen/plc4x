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

// BACnetConstructedDataSecuredStatus is the corresponding interface of BACnetConstructedDataSecuredStatus
type BACnetConstructedDataSecuredStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSecuredStatus returns SecuredStatus (property field)
	GetSecuredStatus() BACnetDoorSecuredStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorSecuredStatusTagged
	// IsBACnetConstructedDataSecuredStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSecuredStatus()
	// CreateBuilder creates a BACnetConstructedDataSecuredStatusBuilder
	CreateBACnetConstructedDataSecuredStatusBuilder() BACnetConstructedDataSecuredStatusBuilder
}

// _BACnetConstructedDataSecuredStatus is the data-structure of this message
type _BACnetConstructedDataSecuredStatus struct {
	BACnetConstructedDataContract
	SecuredStatus BACnetDoorSecuredStatusTagged
}

var _ BACnetConstructedDataSecuredStatus = (*_BACnetConstructedDataSecuredStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSecuredStatus)(nil)

// NewBACnetConstructedDataSecuredStatus factory function for _BACnetConstructedDataSecuredStatus
func NewBACnetConstructedDataSecuredStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, securedStatus BACnetDoorSecuredStatusTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSecuredStatus {
	if securedStatus == nil {
		panic("securedStatus of type BACnetDoorSecuredStatusTagged for BACnetConstructedDataSecuredStatus must not be nil")
	}
	_result := &_BACnetConstructedDataSecuredStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SecuredStatus:                 securedStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSecuredStatusBuilder is a builder for BACnetConstructedDataSecuredStatus
type BACnetConstructedDataSecuredStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(securedStatus BACnetDoorSecuredStatusTagged) BACnetConstructedDataSecuredStatusBuilder
	// WithSecuredStatus adds SecuredStatus (property field)
	WithSecuredStatus(BACnetDoorSecuredStatusTagged) BACnetConstructedDataSecuredStatusBuilder
	// WithSecuredStatusBuilder adds SecuredStatus (property field) which is build by the builder
	WithSecuredStatusBuilder(func(BACnetDoorSecuredStatusTaggedBuilder) BACnetDoorSecuredStatusTaggedBuilder) BACnetConstructedDataSecuredStatusBuilder
	// Build builds the BACnetConstructedDataSecuredStatus or returns an error if something is wrong
	Build() (BACnetConstructedDataSecuredStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSecuredStatus
}

// NewBACnetConstructedDataSecuredStatusBuilder() creates a BACnetConstructedDataSecuredStatusBuilder
func NewBACnetConstructedDataSecuredStatusBuilder() BACnetConstructedDataSecuredStatusBuilder {
	return &_BACnetConstructedDataSecuredStatusBuilder{_BACnetConstructedDataSecuredStatus: new(_BACnetConstructedDataSecuredStatus)}
}

type _BACnetConstructedDataSecuredStatusBuilder struct {
	*_BACnetConstructedDataSecuredStatus

	err *utils.MultiError
}

var _ (BACnetConstructedDataSecuredStatusBuilder) = (*_BACnetConstructedDataSecuredStatusBuilder)(nil)

func (m *_BACnetConstructedDataSecuredStatusBuilder) WithMandatoryFields(securedStatus BACnetDoorSecuredStatusTagged) BACnetConstructedDataSecuredStatusBuilder {
	return m.WithSecuredStatus(securedStatus)
}

func (m *_BACnetConstructedDataSecuredStatusBuilder) WithSecuredStatus(securedStatus BACnetDoorSecuredStatusTagged) BACnetConstructedDataSecuredStatusBuilder {
	m.SecuredStatus = securedStatus
	return m
}

func (m *_BACnetConstructedDataSecuredStatusBuilder) WithSecuredStatusBuilder(builderSupplier func(BACnetDoorSecuredStatusTaggedBuilder) BACnetDoorSecuredStatusTaggedBuilder) BACnetConstructedDataSecuredStatusBuilder {
	builder := builderSupplier(m.SecuredStatus.CreateBACnetDoorSecuredStatusTaggedBuilder())
	var err error
	m.SecuredStatus, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDoorSecuredStatusTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataSecuredStatusBuilder) Build() (BACnetConstructedDataSecuredStatus, error) {
	if m.SecuredStatus == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'securedStatus' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataSecuredStatus.deepCopy(), nil
}

func (m *_BACnetConstructedDataSecuredStatusBuilder) MustBuild() BACnetConstructedDataSecuredStatus {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataSecuredStatusBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataSecuredStatusBuilder()
}

// CreateBACnetConstructedDataSecuredStatusBuilder creates a BACnetConstructedDataSecuredStatusBuilder
func (m *_BACnetConstructedDataSecuredStatus) CreateBACnetConstructedDataSecuredStatusBuilder() BACnetConstructedDataSecuredStatusBuilder {
	if m == nil {
		return NewBACnetConstructedDataSecuredStatusBuilder()
	}
	return &_BACnetConstructedDataSecuredStatusBuilder{_BACnetConstructedDataSecuredStatus: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSecuredStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURED_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetSecuredStatus() BACnetDoorSecuredStatusTagged {
	return m.SecuredStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetActualValue() BACnetDoorSecuredStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorSecuredStatusTagged(m.GetSecuredStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSecuredStatus(structType any) BACnetConstructedDataSecuredStatus {
	if casted, ok := structType.(BACnetConstructedDataSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecuredStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSecuredStatus) GetTypeName() string {
	return "BACnetConstructedDataSecuredStatus"
}

func (m *_BACnetConstructedDataSecuredStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (securedStatus)
	lengthInBits += m.SecuredStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSecuredStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSecuredStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSecuredStatus BACnetConstructedDataSecuredStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecuredStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securedStatus, err := ReadSimpleField[BACnetDoorSecuredStatusTagged](ctx, "securedStatus", ReadComplex[BACnetDoorSecuredStatusTagged](BACnetDoorSecuredStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securedStatus' field"))
	}
	m.SecuredStatus = securedStatus

	actualValue, err := ReadVirtualField[BACnetDoorSecuredStatusTagged](ctx, "actualValue", (*BACnetDoorSecuredStatusTagged)(nil), securedStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecuredStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSecuredStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSecuredStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecuredStatus")
		}

		if err := WriteSimpleField[BACnetDoorSecuredStatusTagged](ctx, "securedStatus", m.GetSecuredStatus(), WriteComplex[BACnetDoorSecuredStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securedStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecuredStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSecuredStatus) IsBACnetConstructedDataSecuredStatus() {}

func (m *_BACnetConstructedDataSecuredStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSecuredStatus) deepCopy() *_BACnetConstructedDataSecuredStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSecuredStatusCopy := &_BACnetConstructedDataSecuredStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.SecuredStatus.DeepCopy().(BACnetDoorSecuredStatusTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSecuredStatusCopy
}

func (m *_BACnetConstructedDataSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
