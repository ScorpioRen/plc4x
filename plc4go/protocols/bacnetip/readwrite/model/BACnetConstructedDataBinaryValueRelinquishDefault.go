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

// BACnetConstructedDataBinaryValueRelinquishDefault is the corresponding interface of BACnetConstructedDataBinaryValueRelinquishDefault
type BACnetConstructedDataBinaryValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
	// IsBACnetConstructedDataBinaryValueRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryValueRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
	CreateBACnetConstructedDataBinaryValueRelinquishDefaultBuilder() BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
}

// _BACnetConstructedDataBinaryValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataBinaryValueRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetBinaryPVTagged
}

var _ BACnetConstructedDataBinaryValueRelinquishDefault = (*_BACnetConstructedDataBinaryValueRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryValueRelinquishDefault)(nil)

// NewBACnetConstructedDataBinaryValueRelinquishDefault factory function for _BACnetConstructedDataBinaryValueRelinquishDefault
func NewBACnetConstructedDataBinaryValueRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetBinaryPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryValueRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetBinaryPVTagged for BACnetConstructedDataBinaryValueRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryValueRelinquishDefault{
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

// BACnetConstructedDataBinaryValueRelinquishDefaultBuilder is a builder for BACnetConstructedDataBinaryValueRelinquishDefault
type BACnetConstructedDataBinaryValueRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetBinaryPVTagged) BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetBinaryPVTagged) BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetBinaryPVTaggedBuilder) BACnetBinaryPVTaggedBuilder) BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
	// Build builds the BACnetConstructedDataBinaryValueRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryValueRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryValueRelinquishDefault
}

// NewBACnetConstructedDataBinaryValueRelinquishDefaultBuilder() creates a BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
func NewBACnetConstructedDataBinaryValueRelinquishDefaultBuilder() BACnetConstructedDataBinaryValueRelinquishDefaultBuilder {
	return &_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder{_BACnetConstructedDataBinaryValueRelinquishDefault: new(_BACnetConstructedDataBinaryValueRelinquishDefault)}
}

type _BACnetConstructedDataBinaryValueRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataBinaryValueRelinquishDefault

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) = (*_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder)(nil)

func (m *_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetBinaryPVTagged) BACnetConstructedDataBinaryValueRelinquishDefaultBuilder {
	return m.WithRelinquishDefault(relinquishDefault)
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetBinaryPVTagged) BACnetConstructedDataBinaryValueRelinquishDefaultBuilder {
	m.RelinquishDefault = relinquishDefault
	return m
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetBinaryPVTaggedBuilder) BACnetBinaryPVTaggedBuilder) BACnetConstructedDataBinaryValueRelinquishDefaultBuilder {
	builder := builderSupplier(m.RelinquishDefault.CreateBACnetBinaryPVTaggedBuilder())
	var err error
	m.RelinquishDefault, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetBinaryPVTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) Build() (BACnetConstructedDataBinaryValueRelinquishDefault, error) {
	if m.RelinquishDefault == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataBinaryValueRelinquishDefault.deepCopy(), nil
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataBinaryValueRelinquishDefault {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataBinaryValueRelinquishDefaultBuilder()
}

// CreateBACnetConstructedDataBinaryValueRelinquishDefaultBuilder creates a BACnetConstructedDataBinaryValueRelinquishDefaultBuilder
func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) CreateBACnetConstructedDataBinaryValueRelinquishDefaultBuilder() BACnetConstructedDataBinaryValueRelinquishDefaultBuilder {
	if m == nil {
		return NewBACnetConstructedDataBinaryValueRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataBinaryValueRelinquishDefaultBuilder{_BACnetConstructedDataBinaryValueRelinquishDefault: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_VALUE
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetRelinquishDefault() BACnetBinaryPVTagged {
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

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetActualValue() BACnetBinaryPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryPVTagged(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryValueRelinquishDefault(structType any) BACnetConstructedDataBinaryValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataBinaryValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataBinaryValueRelinquishDefault"
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryValueRelinquishDefault BACnetConstructedDataBinaryValueRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetBinaryPVTagged](ctx, "relinquishDefault", ReadComplex[BACnetBinaryPVTagged](BACnetBinaryPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetBinaryPVTagged](ctx, "actualValue", (*BACnetBinaryPVTagged)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryValueRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryValueRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetBinaryPVTagged](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetBinaryPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryValueRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) IsBACnetConstructedDataBinaryValueRelinquishDefault() {
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) deepCopy() *_BACnetConstructedDataBinaryValueRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryValueRelinquishDefaultCopy := &_BACnetConstructedDataBinaryValueRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.RelinquishDefault.DeepCopy().(BACnetBinaryPVTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryValueRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataBinaryValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
