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

// BACnetConstructedDataEnergyMeter is the corresponding interface of BACnetConstructedDataEnergyMeter
type BACnetConstructedDataEnergyMeter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEnergyMeter returns EnergyMeter (property field)
	GetEnergyMeter() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataEnergyMeter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEnergyMeter()
	// CreateBuilder creates a BACnetConstructedDataEnergyMeterBuilder
	CreateBACnetConstructedDataEnergyMeterBuilder() BACnetConstructedDataEnergyMeterBuilder
}

// _BACnetConstructedDataEnergyMeter is the data-structure of this message
type _BACnetConstructedDataEnergyMeter struct {
	BACnetConstructedDataContract
	EnergyMeter BACnetApplicationTagReal
}

var _ BACnetConstructedDataEnergyMeter = (*_BACnetConstructedDataEnergyMeter)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEnergyMeter)(nil)

// NewBACnetConstructedDataEnergyMeter factory function for _BACnetConstructedDataEnergyMeter
func NewBACnetConstructedDataEnergyMeter(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, energyMeter BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEnergyMeter {
	if energyMeter == nil {
		panic("energyMeter of type BACnetApplicationTagReal for BACnetConstructedDataEnergyMeter must not be nil")
	}
	_result := &_BACnetConstructedDataEnergyMeter{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EnergyMeter:                   energyMeter,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEnergyMeterBuilder is a builder for BACnetConstructedDataEnergyMeter
type BACnetConstructedDataEnergyMeterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(energyMeter BACnetApplicationTagReal) BACnetConstructedDataEnergyMeterBuilder
	// WithEnergyMeter adds EnergyMeter (property field)
	WithEnergyMeter(BACnetApplicationTagReal) BACnetConstructedDataEnergyMeterBuilder
	// WithEnergyMeterBuilder adds EnergyMeter (property field) which is build by the builder
	WithEnergyMeterBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataEnergyMeterBuilder
	// Build builds the BACnetConstructedDataEnergyMeter or returns an error if something is wrong
	Build() (BACnetConstructedDataEnergyMeter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEnergyMeter
}

// NewBACnetConstructedDataEnergyMeterBuilder() creates a BACnetConstructedDataEnergyMeterBuilder
func NewBACnetConstructedDataEnergyMeterBuilder() BACnetConstructedDataEnergyMeterBuilder {
	return &_BACnetConstructedDataEnergyMeterBuilder{_BACnetConstructedDataEnergyMeter: new(_BACnetConstructedDataEnergyMeter)}
}

type _BACnetConstructedDataEnergyMeterBuilder struct {
	*_BACnetConstructedDataEnergyMeter

	err *utils.MultiError
}

var _ (BACnetConstructedDataEnergyMeterBuilder) = (*_BACnetConstructedDataEnergyMeterBuilder)(nil)

func (m *_BACnetConstructedDataEnergyMeterBuilder) WithMandatoryFields(energyMeter BACnetApplicationTagReal) BACnetConstructedDataEnergyMeterBuilder {
	return m.WithEnergyMeter(energyMeter)
}

func (m *_BACnetConstructedDataEnergyMeterBuilder) WithEnergyMeter(energyMeter BACnetApplicationTagReal) BACnetConstructedDataEnergyMeterBuilder {
	m.EnergyMeter = energyMeter
	return m
}

func (m *_BACnetConstructedDataEnergyMeterBuilder) WithEnergyMeterBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataEnergyMeterBuilder {
	builder := builderSupplier(m.EnergyMeter.CreateBACnetApplicationTagRealBuilder())
	var err error
	m.EnergyMeter, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataEnergyMeterBuilder) Build() (BACnetConstructedDataEnergyMeter, error) {
	if m.EnergyMeter == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'energyMeter' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataEnergyMeter.deepCopy(), nil
}

func (m *_BACnetConstructedDataEnergyMeterBuilder) MustBuild() BACnetConstructedDataEnergyMeter {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataEnergyMeterBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataEnergyMeterBuilder()
}

// CreateBACnetConstructedDataEnergyMeterBuilder creates a BACnetConstructedDataEnergyMeterBuilder
func (m *_BACnetConstructedDataEnergyMeter) CreateBACnetConstructedDataEnergyMeterBuilder() BACnetConstructedDataEnergyMeterBuilder {
	if m == nil {
		return NewBACnetConstructedDataEnergyMeterBuilder()
	}
	return &_BACnetConstructedDataEnergyMeterBuilder{_BACnetConstructedDataEnergyMeter: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEnergyMeter) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENERGY_METER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetEnergyMeter() BACnetApplicationTagReal {
	return m.EnergyMeter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetEnergyMeter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEnergyMeter(structType any) BACnetConstructedDataEnergyMeter {
	if casted, ok := structType.(BACnetConstructedDataEnergyMeter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEnergyMeter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEnergyMeter) GetTypeName() string {
	return "BACnetConstructedDataEnergyMeter"
}

func (m *_BACnetConstructedDataEnergyMeter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (energyMeter)
	lengthInBits += m.EnergyMeter.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEnergyMeter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEnergyMeter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEnergyMeter BACnetConstructedDataEnergyMeter, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEnergyMeter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEnergyMeter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	energyMeter, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "energyMeter", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'energyMeter' field"))
	}
	m.EnergyMeter = energyMeter

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), energyMeter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEnergyMeter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEnergyMeter")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEnergyMeter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEnergyMeter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEnergyMeter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEnergyMeter")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "energyMeter", m.GetEnergyMeter(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'energyMeter' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEnergyMeter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEnergyMeter")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEnergyMeter) IsBACnetConstructedDataEnergyMeter() {}

func (m *_BACnetConstructedDataEnergyMeter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEnergyMeter) deepCopy() *_BACnetConstructedDataEnergyMeter {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEnergyMeterCopy := &_BACnetConstructedDataEnergyMeter{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.EnergyMeter.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEnergyMeterCopy
}

func (m *_BACnetConstructedDataEnergyMeter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
