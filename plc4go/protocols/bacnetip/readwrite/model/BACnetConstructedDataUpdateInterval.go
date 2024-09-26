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

// BACnetConstructedDataUpdateInterval is the corresponding interface of BACnetConstructedDataUpdateInterval
type BACnetConstructedDataUpdateInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUpdateInterval returns UpdateInterval (property field)
	GetUpdateInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataUpdateInterval is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUpdateInterval()
	// CreateBuilder creates a BACnetConstructedDataUpdateIntervalBuilder
	CreateBACnetConstructedDataUpdateIntervalBuilder() BACnetConstructedDataUpdateIntervalBuilder
}

// _BACnetConstructedDataUpdateInterval is the data-structure of this message
type _BACnetConstructedDataUpdateInterval struct {
	BACnetConstructedDataContract
	UpdateInterval BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataUpdateInterval = (*_BACnetConstructedDataUpdateInterval)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUpdateInterval)(nil)

// NewBACnetConstructedDataUpdateInterval factory function for _BACnetConstructedDataUpdateInterval
func NewBACnetConstructedDataUpdateInterval(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, updateInterval BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUpdateInterval {
	if updateInterval == nil {
		panic("updateInterval of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataUpdateInterval must not be nil")
	}
	_result := &_BACnetConstructedDataUpdateInterval{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		UpdateInterval:                updateInterval,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataUpdateIntervalBuilder is a builder for BACnetConstructedDataUpdateInterval
type BACnetConstructedDataUpdateIntervalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(updateInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataUpdateIntervalBuilder
	// WithUpdateInterval adds UpdateInterval (property field)
	WithUpdateInterval(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataUpdateIntervalBuilder
	// WithUpdateIntervalBuilder adds UpdateInterval (property field) which is build by the builder
	WithUpdateIntervalBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataUpdateIntervalBuilder
	// Build builds the BACnetConstructedDataUpdateInterval or returns an error if something is wrong
	Build() (BACnetConstructedDataUpdateInterval, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataUpdateInterval
}

// NewBACnetConstructedDataUpdateIntervalBuilder() creates a BACnetConstructedDataUpdateIntervalBuilder
func NewBACnetConstructedDataUpdateIntervalBuilder() BACnetConstructedDataUpdateIntervalBuilder {
	return &_BACnetConstructedDataUpdateIntervalBuilder{_BACnetConstructedDataUpdateInterval: new(_BACnetConstructedDataUpdateInterval)}
}

type _BACnetConstructedDataUpdateIntervalBuilder struct {
	*_BACnetConstructedDataUpdateInterval

	err *utils.MultiError
}

var _ (BACnetConstructedDataUpdateIntervalBuilder) = (*_BACnetConstructedDataUpdateIntervalBuilder)(nil)

func (m *_BACnetConstructedDataUpdateIntervalBuilder) WithMandatoryFields(updateInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataUpdateIntervalBuilder {
	return m.WithUpdateInterval(updateInterval)
}

func (m *_BACnetConstructedDataUpdateIntervalBuilder) WithUpdateInterval(updateInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataUpdateIntervalBuilder {
	m.UpdateInterval = updateInterval
	return m
}

func (m *_BACnetConstructedDataUpdateIntervalBuilder) WithUpdateIntervalBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataUpdateIntervalBuilder {
	builder := builderSupplier(m.UpdateInterval.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.UpdateInterval, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataUpdateIntervalBuilder) Build() (BACnetConstructedDataUpdateInterval, error) {
	if m.UpdateInterval == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'updateInterval' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataUpdateInterval.deepCopy(), nil
}

func (m *_BACnetConstructedDataUpdateIntervalBuilder) MustBuild() BACnetConstructedDataUpdateInterval {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataUpdateIntervalBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataUpdateIntervalBuilder()
}

// CreateBACnetConstructedDataUpdateIntervalBuilder creates a BACnetConstructedDataUpdateIntervalBuilder
func (m *_BACnetConstructedDataUpdateInterval) CreateBACnetConstructedDataUpdateIntervalBuilder() BACnetConstructedDataUpdateIntervalBuilder {
	if m == nil {
		return NewBACnetConstructedDataUpdateIntervalBuilder()
	}
	return &_BACnetConstructedDataUpdateIntervalBuilder{_BACnetConstructedDataUpdateInterval: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUpdateInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetUpdateInterval() BACnetApplicationTagUnsignedInteger {
	return m.UpdateInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetUpdateInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUpdateInterval(structType any) BACnetConstructedDataUpdateInterval {
	if casted, ok := structType.(BACnetConstructedDataUpdateInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUpdateInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUpdateInterval) GetTypeName() string {
	return "BACnetConstructedDataUpdateInterval"
}

func (m *_BACnetConstructedDataUpdateInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (updateInterval)
	lengthInBits += m.UpdateInterval.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUpdateInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUpdateInterval) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUpdateInterval BACnetConstructedDataUpdateInterval, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUpdateInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	updateInterval, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "updateInterval", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateInterval' field"))
	}
	m.UpdateInterval = updateInterval

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), updateInterval)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUpdateInterval")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUpdateInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUpdateInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUpdateInterval")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "updateInterval", m.GetUpdateInterval(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'updateInterval' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUpdateInterval")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUpdateInterval) IsBACnetConstructedDataUpdateInterval() {}

func (m *_BACnetConstructedDataUpdateInterval) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUpdateInterval) deepCopy() *_BACnetConstructedDataUpdateInterval {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUpdateIntervalCopy := &_BACnetConstructedDataUpdateInterval{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.UpdateInterval.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUpdateIntervalCopy
}

func (m *_BACnetConstructedDataUpdateInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
