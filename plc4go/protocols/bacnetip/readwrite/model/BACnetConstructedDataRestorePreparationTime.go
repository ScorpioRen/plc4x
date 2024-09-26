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

// BACnetConstructedDataRestorePreparationTime is the corresponding interface of BACnetConstructedDataRestorePreparationTime
type BACnetConstructedDataRestorePreparationTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRestorePreparationTime returns RestorePreparationTime (property field)
	GetRestorePreparationTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataRestorePreparationTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRestorePreparationTime()
	// CreateBuilder creates a BACnetConstructedDataRestorePreparationTimeBuilder
	CreateBACnetConstructedDataRestorePreparationTimeBuilder() BACnetConstructedDataRestorePreparationTimeBuilder
}

// _BACnetConstructedDataRestorePreparationTime is the data-structure of this message
type _BACnetConstructedDataRestorePreparationTime struct {
	BACnetConstructedDataContract
	RestorePreparationTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataRestorePreparationTime = (*_BACnetConstructedDataRestorePreparationTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRestorePreparationTime)(nil)

// NewBACnetConstructedDataRestorePreparationTime factory function for _BACnetConstructedDataRestorePreparationTime
func NewBACnetConstructedDataRestorePreparationTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, restorePreparationTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRestorePreparationTime {
	if restorePreparationTime == nil {
		panic("restorePreparationTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataRestorePreparationTime must not be nil")
	}
	_result := &_BACnetConstructedDataRestorePreparationTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RestorePreparationTime:        restorePreparationTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRestorePreparationTimeBuilder is a builder for BACnetConstructedDataRestorePreparationTime
type BACnetConstructedDataRestorePreparationTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(restorePreparationTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRestorePreparationTimeBuilder
	// WithRestorePreparationTime adds RestorePreparationTime (property field)
	WithRestorePreparationTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRestorePreparationTimeBuilder
	// WithRestorePreparationTimeBuilder adds RestorePreparationTime (property field) which is build by the builder
	WithRestorePreparationTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRestorePreparationTimeBuilder
	// Build builds the BACnetConstructedDataRestorePreparationTime or returns an error if something is wrong
	Build() (BACnetConstructedDataRestorePreparationTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRestorePreparationTime
}

// NewBACnetConstructedDataRestorePreparationTimeBuilder() creates a BACnetConstructedDataRestorePreparationTimeBuilder
func NewBACnetConstructedDataRestorePreparationTimeBuilder() BACnetConstructedDataRestorePreparationTimeBuilder {
	return &_BACnetConstructedDataRestorePreparationTimeBuilder{_BACnetConstructedDataRestorePreparationTime: new(_BACnetConstructedDataRestorePreparationTime)}
}

type _BACnetConstructedDataRestorePreparationTimeBuilder struct {
	*_BACnetConstructedDataRestorePreparationTime

	err *utils.MultiError
}

var _ (BACnetConstructedDataRestorePreparationTimeBuilder) = (*_BACnetConstructedDataRestorePreparationTimeBuilder)(nil)

func (m *_BACnetConstructedDataRestorePreparationTimeBuilder) WithMandatoryFields(restorePreparationTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRestorePreparationTimeBuilder {
	return m.WithRestorePreparationTime(restorePreparationTime)
}

func (m *_BACnetConstructedDataRestorePreparationTimeBuilder) WithRestorePreparationTime(restorePreparationTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRestorePreparationTimeBuilder {
	m.RestorePreparationTime = restorePreparationTime
	return m
}

func (m *_BACnetConstructedDataRestorePreparationTimeBuilder) WithRestorePreparationTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRestorePreparationTimeBuilder {
	builder := builderSupplier(m.RestorePreparationTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.RestorePreparationTime, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataRestorePreparationTimeBuilder) Build() (BACnetConstructedDataRestorePreparationTime, error) {
	if m.RestorePreparationTime == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'restorePreparationTime' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataRestorePreparationTime.deepCopy(), nil
}

func (m *_BACnetConstructedDataRestorePreparationTimeBuilder) MustBuild() BACnetConstructedDataRestorePreparationTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataRestorePreparationTimeBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataRestorePreparationTimeBuilder()
}

// CreateBACnetConstructedDataRestorePreparationTimeBuilder creates a BACnetConstructedDataRestorePreparationTimeBuilder
func (m *_BACnetConstructedDataRestorePreparationTime) CreateBACnetConstructedDataRestorePreparationTimeBuilder() BACnetConstructedDataRestorePreparationTimeBuilder {
	if m == nil {
		return NewBACnetConstructedDataRestorePreparationTimeBuilder()
	}
	return &_BACnetConstructedDataRestorePreparationTimeBuilder{_BACnetConstructedDataRestorePreparationTime: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESTORE_PREPARATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) GetRestorePreparationTime() BACnetApplicationTagUnsignedInteger {
	return m.RestorePreparationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRestorePreparationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRestorePreparationTime(structType any) BACnetConstructedDataRestorePreparationTime {
	if casted, ok := structType.(BACnetConstructedDataRestorePreparationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestorePreparationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetTypeName() string {
	return "BACnetConstructedDataRestorePreparationTime"
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (restorePreparationTime)
	lengthInBits += m.RestorePreparationTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRestorePreparationTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRestorePreparationTime BACnetConstructedDataRestorePreparationTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRestorePreparationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRestorePreparationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	restorePreparationTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "restorePreparationTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'restorePreparationTime' field"))
	}
	m.RestorePreparationTime = restorePreparationTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), restorePreparationTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestorePreparationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRestorePreparationTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRestorePreparationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRestorePreparationTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRestorePreparationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRestorePreparationTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "restorePreparationTime", m.GetRestorePreparationTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'restorePreparationTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestorePreparationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRestorePreparationTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRestorePreparationTime) IsBACnetConstructedDataRestorePreparationTime() {
}

func (m *_BACnetConstructedDataRestorePreparationTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRestorePreparationTime) deepCopy() *_BACnetConstructedDataRestorePreparationTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRestorePreparationTimeCopy := &_BACnetConstructedDataRestorePreparationTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.RestorePreparationTime.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRestorePreparationTimeCopy
}

func (m *_BACnetConstructedDataRestorePreparationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
