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

// BACnetConstructedDataOccupancyCountEnable is the corresponding interface of BACnetConstructedDataOccupancyCountEnable
type BACnetConstructedDataOccupancyCountEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOccupancyCountEnable returns OccupancyCountEnable (property field)
	GetOccupancyCountEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataOccupancyCountEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyCountEnable()
	// CreateBuilder creates a BACnetConstructedDataOccupancyCountEnableBuilder
	CreateBACnetConstructedDataOccupancyCountEnableBuilder() BACnetConstructedDataOccupancyCountEnableBuilder
}

// _BACnetConstructedDataOccupancyCountEnable is the data-structure of this message
type _BACnetConstructedDataOccupancyCountEnable struct {
	BACnetConstructedDataContract
	OccupancyCountEnable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataOccupancyCountEnable = (*_BACnetConstructedDataOccupancyCountEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyCountEnable)(nil)

// NewBACnetConstructedDataOccupancyCountEnable factory function for _BACnetConstructedDataOccupancyCountEnable
func NewBACnetConstructedDataOccupancyCountEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyCountEnable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyCountEnable {
	if occupancyCountEnable == nil {
		panic("occupancyCountEnable of type BACnetApplicationTagBoolean for BACnetConstructedDataOccupancyCountEnable must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyCountEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyCountEnable:          occupancyCountEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOccupancyCountEnableBuilder is a builder for BACnetConstructedDataOccupancyCountEnable
type BACnetConstructedDataOccupancyCountEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(occupancyCountEnable BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyCountEnableBuilder
	// WithOccupancyCountEnable adds OccupancyCountEnable (property field)
	WithOccupancyCountEnable(BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyCountEnableBuilder
	// WithOccupancyCountEnableBuilder adds OccupancyCountEnable (property field) which is build by the builder
	WithOccupancyCountEnableBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataOccupancyCountEnableBuilder
	// Build builds the BACnetConstructedDataOccupancyCountEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataOccupancyCountEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOccupancyCountEnable
}

// NewBACnetConstructedDataOccupancyCountEnableBuilder() creates a BACnetConstructedDataOccupancyCountEnableBuilder
func NewBACnetConstructedDataOccupancyCountEnableBuilder() BACnetConstructedDataOccupancyCountEnableBuilder {
	return &_BACnetConstructedDataOccupancyCountEnableBuilder{_BACnetConstructedDataOccupancyCountEnable: new(_BACnetConstructedDataOccupancyCountEnable)}
}

type _BACnetConstructedDataOccupancyCountEnableBuilder struct {
	*_BACnetConstructedDataOccupancyCountEnable

	err *utils.MultiError
}

var _ (BACnetConstructedDataOccupancyCountEnableBuilder) = (*_BACnetConstructedDataOccupancyCountEnableBuilder)(nil)

func (m *_BACnetConstructedDataOccupancyCountEnableBuilder) WithMandatoryFields(occupancyCountEnable BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyCountEnableBuilder {
	return m.WithOccupancyCountEnable(occupancyCountEnable)
}

func (m *_BACnetConstructedDataOccupancyCountEnableBuilder) WithOccupancyCountEnable(occupancyCountEnable BACnetApplicationTagBoolean) BACnetConstructedDataOccupancyCountEnableBuilder {
	m.OccupancyCountEnable = occupancyCountEnable
	return m
}

func (m *_BACnetConstructedDataOccupancyCountEnableBuilder) WithOccupancyCountEnableBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataOccupancyCountEnableBuilder {
	builder := builderSupplier(m.OccupancyCountEnable.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.OccupancyCountEnable, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataOccupancyCountEnableBuilder) Build() (BACnetConstructedDataOccupancyCountEnable, error) {
	if m.OccupancyCountEnable == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'occupancyCountEnable' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataOccupancyCountEnable.deepCopy(), nil
}

func (m *_BACnetConstructedDataOccupancyCountEnableBuilder) MustBuild() BACnetConstructedDataOccupancyCountEnable {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataOccupancyCountEnableBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataOccupancyCountEnableBuilder()
}

// CreateBACnetConstructedDataOccupancyCountEnableBuilder creates a BACnetConstructedDataOccupancyCountEnableBuilder
func (m *_BACnetConstructedDataOccupancyCountEnable) CreateBACnetConstructedDataOccupancyCountEnableBuilder() BACnetConstructedDataOccupancyCountEnableBuilder {
	if m == nil {
		return NewBACnetConstructedDataOccupancyCountEnableBuilder()
	}
	return &_BACnetConstructedDataOccupancyCountEnableBuilder{_BACnetConstructedDataOccupancyCountEnable: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyCountEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_COUNT_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyCountEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountEnable) GetOccupancyCountEnable() BACnetApplicationTagBoolean {
	return m.OccupancyCountEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetOccupancyCountEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyCountEnable(structType any) BACnetConstructedDataOccupancyCountEnable {
	if casted, ok := structType.(BACnetConstructedDataOccupancyCountEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyCountEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyCountEnable) GetTypeName() string {
	return "BACnetConstructedDataOccupancyCountEnable"
}

func (m *_BACnetConstructedDataOccupancyCountEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyCountEnable)
	lengthInBits += m.OccupancyCountEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyCountEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyCountEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyCountEnable BACnetConstructedDataOccupancyCountEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyCountEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyCountEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyCountEnable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyCountEnable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyCountEnable' field"))
	}
	m.OccupancyCountEnable = occupancyCountEnable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), occupancyCountEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyCountEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyCountEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyCountEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyCountEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyCountEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyCountEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyCountEnable", m.GetOccupancyCountEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyCountEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyCountEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyCountEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyCountEnable) IsBACnetConstructedDataOccupancyCountEnable() {}

func (m *_BACnetConstructedDataOccupancyCountEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOccupancyCountEnable) deepCopy() *_BACnetConstructedDataOccupancyCountEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOccupancyCountEnableCopy := &_BACnetConstructedDataOccupancyCountEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.OccupancyCountEnable.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOccupancyCountEnableCopy
}

func (m *_BACnetConstructedDataOccupancyCountEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
