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

// BACnetConstructedDataZoneFrom is the corresponding interface of BACnetConstructedDataZoneFrom
type BACnetConstructedDataZoneFrom interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetZoneFrom returns ZoneFrom (property field)
	GetZoneFrom() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataZoneFrom is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataZoneFrom()
	// CreateBuilder creates a BACnetConstructedDataZoneFromBuilder
	CreateBACnetConstructedDataZoneFromBuilder() BACnetConstructedDataZoneFromBuilder
}

// _BACnetConstructedDataZoneFrom is the data-structure of this message
type _BACnetConstructedDataZoneFrom struct {
	BACnetConstructedDataContract
	ZoneFrom BACnetDeviceObjectReference
}

var _ BACnetConstructedDataZoneFrom = (*_BACnetConstructedDataZoneFrom)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataZoneFrom)(nil)

// NewBACnetConstructedDataZoneFrom factory function for _BACnetConstructedDataZoneFrom
func NewBACnetConstructedDataZoneFrom(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, zoneFrom BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataZoneFrom {
	if zoneFrom == nil {
		panic("zoneFrom of type BACnetDeviceObjectReference for BACnetConstructedDataZoneFrom must not be nil")
	}
	_result := &_BACnetConstructedDataZoneFrom{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ZoneFrom:                      zoneFrom,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataZoneFromBuilder is a builder for BACnetConstructedDataZoneFrom
type BACnetConstructedDataZoneFromBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneFrom BACnetDeviceObjectReference) BACnetConstructedDataZoneFromBuilder
	// WithZoneFrom adds ZoneFrom (property field)
	WithZoneFrom(BACnetDeviceObjectReference) BACnetConstructedDataZoneFromBuilder
	// WithZoneFromBuilder adds ZoneFrom (property field) which is build by the builder
	WithZoneFromBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataZoneFromBuilder
	// Build builds the BACnetConstructedDataZoneFrom or returns an error if something is wrong
	Build() (BACnetConstructedDataZoneFrom, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataZoneFrom
}

// NewBACnetConstructedDataZoneFromBuilder() creates a BACnetConstructedDataZoneFromBuilder
func NewBACnetConstructedDataZoneFromBuilder() BACnetConstructedDataZoneFromBuilder {
	return &_BACnetConstructedDataZoneFromBuilder{_BACnetConstructedDataZoneFrom: new(_BACnetConstructedDataZoneFrom)}
}

type _BACnetConstructedDataZoneFromBuilder struct {
	*_BACnetConstructedDataZoneFrom

	err *utils.MultiError
}

var _ (BACnetConstructedDataZoneFromBuilder) = (*_BACnetConstructedDataZoneFromBuilder)(nil)

func (m *_BACnetConstructedDataZoneFromBuilder) WithMandatoryFields(zoneFrom BACnetDeviceObjectReference) BACnetConstructedDataZoneFromBuilder {
	return m.WithZoneFrom(zoneFrom)
}

func (m *_BACnetConstructedDataZoneFromBuilder) WithZoneFrom(zoneFrom BACnetDeviceObjectReference) BACnetConstructedDataZoneFromBuilder {
	m.ZoneFrom = zoneFrom
	return m
}

func (m *_BACnetConstructedDataZoneFromBuilder) WithZoneFromBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataZoneFromBuilder {
	builder := builderSupplier(m.ZoneFrom.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	m.ZoneFrom, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataZoneFromBuilder) Build() (BACnetConstructedDataZoneFrom, error) {
	if m.ZoneFrom == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'zoneFrom' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataZoneFrom.deepCopy(), nil
}

func (m *_BACnetConstructedDataZoneFromBuilder) MustBuild() BACnetConstructedDataZoneFrom {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataZoneFromBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataZoneFromBuilder()
}

// CreateBACnetConstructedDataZoneFromBuilder creates a BACnetConstructedDataZoneFromBuilder
func (m *_BACnetConstructedDataZoneFrom) CreateBACnetConstructedDataZoneFromBuilder() BACnetConstructedDataZoneFromBuilder {
	if m == nil {
		return NewBACnetConstructedDataZoneFromBuilder()
	}
	return &_BACnetConstructedDataZoneFromBuilder{_BACnetConstructedDataZoneFrom: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataZoneFrom) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataZoneFrom) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_FROM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataZoneFrom) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataZoneFrom) GetZoneFrom() BACnetDeviceObjectReference {
	return m.ZoneFrom
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataZoneFrom) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetZoneFrom())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataZoneFrom(structType any) BACnetConstructedDataZoneFrom {
	if casted, ok := structType.(BACnetConstructedDataZoneFrom); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneFrom); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataZoneFrom) GetTypeName() string {
	return "BACnetConstructedDataZoneFrom"
}

func (m *_BACnetConstructedDataZoneFrom) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (zoneFrom)
	lengthInBits += m.ZoneFrom.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataZoneFrom) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataZoneFrom) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataZoneFrom BACnetConstructedDataZoneFrom, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneFrom"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneFrom")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneFrom, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "zoneFrom", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneFrom' field"))
	}
	m.ZoneFrom = zoneFrom

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), zoneFrom)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneFrom"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneFrom")
	}

	return m, nil
}

func (m *_BACnetConstructedDataZoneFrom) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataZoneFrom) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneFrom"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneFrom")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "zoneFrom", m.GetZoneFrom(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneFrom' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneFrom"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneFrom")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataZoneFrom) IsBACnetConstructedDataZoneFrom() {}

func (m *_BACnetConstructedDataZoneFrom) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataZoneFrom) deepCopy() *_BACnetConstructedDataZoneFrom {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataZoneFromCopy := &_BACnetConstructedDataZoneFrom{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ZoneFrom.DeepCopy().(BACnetDeviceObjectReference),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataZoneFromCopy
}

func (m *_BACnetConstructedDataZoneFrom) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
