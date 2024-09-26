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

// BACnetConstructedDataTimePatternValueRelinquishDefault is the corresponding interface of BACnetConstructedDataTimePatternValueRelinquishDefault
type BACnetConstructedDataTimePatternValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagTime
	// IsBACnetConstructedDataTimePatternValueRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimePatternValueRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
	CreateBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder() BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
}

// _BACnetConstructedDataTimePatternValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataTimePatternValueRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetApplicationTagTime
}

var _ BACnetConstructedDataTimePatternValueRelinquishDefault = (*_BACnetConstructedDataTimePatternValueRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimePatternValueRelinquishDefault)(nil)

// NewBACnetConstructedDataTimePatternValueRelinquishDefault factory function for _BACnetConstructedDataTimePatternValueRelinquishDefault
func NewBACnetConstructedDataTimePatternValueRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetApplicationTagTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimePatternValueRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetApplicationTagTime for BACnetConstructedDataTimePatternValueRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataTimePatternValueRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder is a builder for BACnetConstructedDataTimePatternValueRelinquishDefault
type BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetApplicationTagTime) BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetApplicationTagTime) BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
	// Build builds the BACnetConstructedDataTimePatternValueRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataTimePatternValueRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimePatternValueRelinquishDefault
}

// NewBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder() creates a BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
func NewBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder() BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder {
	return &_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder{_BACnetConstructedDataTimePatternValueRelinquishDefault: new(_BACnetConstructedDataTimePatternValueRelinquishDefault)}
}

type _BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataTimePatternValueRelinquishDefault

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) = (*_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder)(nil)

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetApplicationTagTime) BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder {
	return m.WithRelinquishDefault(relinquishDefault)
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetApplicationTagTime) BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder {
	m.RelinquishDefault = relinquishDefault
	return m
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder {
	builder := builderSupplier(m.RelinquishDefault.CreateBACnetApplicationTagTimeBuilder())
	var err error
	m.RelinquishDefault, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) Build() (BACnetConstructedDataTimePatternValueRelinquishDefault, error) {
	if m.RelinquishDefault == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataTimePatternValueRelinquishDefault.deepCopy(), nil
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataTimePatternValueRelinquishDefault {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder()
}

// CreateBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder creates a BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder
func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) CreateBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder() BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder {
	if m == nil {
		return NewBACnetConstructedDataTimePatternValueRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataTimePatternValueRelinquishDefaultBuilder{_BACnetConstructedDataTimePatternValueRelinquishDefault: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMEPATTERN_VALUE
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagTime {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetActualValue() BACnetApplicationTagTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagTime(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimePatternValueRelinquishDefault(structType any) BACnetConstructedDataTimePatternValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataTimePatternValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimePatternValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataTimePatternValueRelinquishDefault"
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimePatternValueRelinquishDefault BACnetConstructedDataTimePatternValueRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimePatternValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimePatternValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "relinquishDefault", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetApplicationTagTime](ctx, "actualValue", (*BACnetApplicationTagTime)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimePatternValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimePatternValueRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimePatternValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimePatternValueRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimePatternValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimePatternValueRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) IsBACnetConstructedDataTimePatternValueRelinquishDefault() {
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) deepCopy() *_BACnetConstructedDataTimePatternValueRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimePatternValueRelinquishDefaultCopy := &_BACnetConstructedDataTimePatternValueRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.RelinquishDefault.DeepCopy().(BACnetApplicationTagTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimePatternValueRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataTimePatternValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
