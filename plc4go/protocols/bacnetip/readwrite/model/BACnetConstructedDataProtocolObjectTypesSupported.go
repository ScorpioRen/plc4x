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

// BACnetConstructedDataProtocolObjectTypesSupported is the corresponding interface of BACnetConstructedDataProtocolObjectTypesSupported
type BACnetConstructedDataProtocolObjectTypesSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetProtocolObjectTypesSupported returns ProtocolObjectTypesSupported (property field)
	GetProtocolObjectTypesSupported() BACnetObjectTypesSupportedTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectTypesSupportedTagged
	// IsBACnetConstructedDataProtocolObjectTypesSupported is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProtocolObjectTypesSupported()
	// CreateBuilder creates a BACnetConstructedDataProtocolObjectTypesSupportedBuilder
	CreateBACnetConstructedDataProtocolObjectTypesSupportedBuilder() BACnetConstructedDataProtocolObjectTypesSupportedBuilder
}

// _BACnetConstructedDataProtocolObjectTypesSupported is the data-structure of this message
type _BACnetConstructedDataProtocolObjectTypesSupported struct {
	BACnetConstructedDataContract
	ProtocolObjectTypesSupported BACnetObjectTypesSupportedTagged
}

var _ BACnetConstructedDataProtocolObjectTypesSupported = (*_BACnetConstructedDataProtocolObjectTypesSupported)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProtocolObjectTypesSupported)(nil)

// NewBACnetConstructedDataProtocolObjectTypesSupported factory function for _BACnetConstructedDataProtocolObjectTypesSupported
func NewBACnetConstructedDataProtocolObjectTypesSupported(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, protocolObjectTypesSupported BACnetObjectTypesSupportedTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolObjectTypesSupported {
	if protocolObjectTypesSupported == nil {
		panic("protocolObjectTypesSupported of type BACnetObjectTypesSupportedTagged for BACnetConstructedDataProtocolObjectTypesSupported must not be nil")
	}
	_result := &_BACnetConstructedDataProtocolObjectTypesSupported{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProtocolObjectTypesSupported:  protocolObjectTypesSupported,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProtocolObjectTypesSupportedBuilder is a builder for BACnetConstructedDataProtocolObjectTypesSupported
type BACnetConstructedDataProtocolObjectTypesSupportedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(protocolObjectTypesSupported BACnetObjectTypesSupportedTagged) BACnetConstructedDataProtocolObjectTypesSupportedBuilder
	// WithProtocolObjectTypesSupported adds ProtocolObjectTypesSupported (property field)
	WithProtocolObjectTypesSupported(BACnetObjectTypesSupportedTagged) BACnetConstructedDataProtocolObjectTypesSupportedBuilder
	// WithProtocolObjectTypesSupportedBuilder adds ProtocolObjectTypesSupported (property field) which is build by the builder
	WithProtocolObjectTypesSupportedBuilder(func(BACnetObjectTypesSupportedTaggedBuilder) BACnetObjectTypesSupportedTaggedBuilder) BACnetConstructedDataProtocolObjectTypesSupportedBuilder
	// Build builds the BACnetConstructedDataProtocolObjectTypesSupported or returns an error if something is wrong
	Build() (BACnetConstructedDataProtocolObjectTypesSupported, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProtocolObjectTypesSupported
}

// NewBACnetConstructedDataProtocolObjectTypesSupportedBuilder() creates a BACnetConstructedDataProtocolObjectTypesSupportedBuilder
func NewBACnetConstructedDataProtocolObjectTypesSupportedBuilder() BACnetConstructedDataProtocolObjectTypesSupportedBuilder {
	return &_BACnetConstructedDataProtocolObjectTypesSupportedBuilder{_BACnetConstructedDataProtocolObjectTypesSupported: new(_BACnetConstructedDataProtocolObjectTypesSupported)}
}

type _BACnetConstructedDataProtocolObjectTypesSupportedBuilder struct {
	*_BACnetConstructedDataProtocolObjectTypesSupported

	err *utils.MultiError
}

var _ (BACnetConstructedDataProtocolObjectTypesSupportedBuilder) = (*_BACnetConstructedDataProtocolObjectTypesSupportedBuilder)(nil)

func (m *_BACnetConstructedDataProtocolObjectTypesSupportedBuilder) WithMandatoryFields(protocolObjectTypesSupported BACnetObjectTypesSupportedTagged) BACnetConstructedDataProtocolObjectTypesSupportedBuilder {
	return m.WithProtocolObjectTypesSupported(protocolObjectTypesSupported)
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupportedBuilder) WithProtocolObjectTypesSupported(protocolObjectTypesSupported BACnetObjectTypesSupportedTagged) BACnetConstructedDataProtocolObjectTypesSupportedBuilder {
	m.ProtocolObjectTypesSupported = protocolObjectTypesSupported
	return m
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupportedBuilder) WithProtocolObjectTypesSupportedBuilder(builderSupplier func(BACnetObjectTypesSupportedTaggedBuilder) BACnetObjectTypesSupportedTaggedBuilder) BACnetConstructedDataProtocolObjectTypesSupportedBuilder {
	builder := builderSupplier(m.ProtocolObjectTypesSupported.CreateBACnetObjectTypesSupportedTaggedBuilder())
	var err error
	m.ProtocolObjectTypesSupported, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetObjectTypesSupportedTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupportedBuilder) Build() (BACnetConstructedDataProtocolObjectTypesSupported, error) {
	if m.ProtocolObjectTypesSupported == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'protocolObjectTypesSupported' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataProtocolObjectTypesSupported.deepCopy(), nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupportedBuilder) MustBuild() BACnetConstructedDataProtocolObjectTypesSupported {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupportedBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataProtocolObjectTypesSupportedBuilder()
}

// CreateBACnetConstructedDataProtocolObjectTypesSupportedBuilder creates a BACnetConstructedDataProtocolObjectTypesSupportedBuilder
func (m *_BACnetConstructedDataProtocolObjectTypesSupported) CreateBACnetConstructedDataProtocolObjectTypesSupportedBuilder() BACnetConstructedDataProtocolObjectTypesSupportedBuilder {
	if m == nil {
		return NewBACnetConstructedDataProtocolObjectTypesSupportedBuilder()
	}
	return &_BACnetConstructedDataProtocolObjectTypesSupportedBuilder{_BACnetConstructedDataProtocolObjectTypesSupported: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_OBJECT_TYPES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetProtocolObjectTypesSupported() BACnetObjectTypesSupportedTagged {
	return m.ProtocolObjectTypesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetActualValue() BACnetObjectTypesSupportedTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectTypesSupportedTagged(m.GetProtocolObjectTypesSupported())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolObjectTypesSupported(structType any) BACnetConstructedDataProtocolObjectTypesSupported {
	if casted, ok := structType.(BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetTypeName() string {
	return "BACnetConstructedDataProtocolObjectTypesSupported"
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (protocolObjectTypesSupported)
	lengthInBits += m.ProtocolObjectTypesSupported.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProtocolObjectTypesSupported BACnetConstructedDataProtocolObjectTypesSupported, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolObjectTypesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolObjectTypesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolObjectTypesSupported, err := ReadSimpleField[BACnetObjectTypesSupportedTagged](ctx, "protocolObjectTypesSupported", ReadComplex[BACnetObjectTypesSupportedTagged](BACnetObjectTypesSupportedTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolObjectTypesSupported' field"))
	}
	m.ProtocolObjectTypesSupported = protocolObjectTypesSupported

	actualValue, err := ReadVirtualField[BACnetObjectTypesSupportedTagged](ctx, "actualValue", (*BACnetObjectTypesSupportedTagged)(nil), protocolObjectTypesSupported)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolObjectTypesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolObjectTypesSupported")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolObjectTypesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolObjectTypesSupported")
		}

		if err := WriteSimpleField[BACnetObjectTypesSupportedTagged](ctx, "protocolObjectTypesSupported", m.GetProtocolObjectTypesSupported(), WriteComplex[BACnetObjectTypesSupportedTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolObjectTypesSupported' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolObjectTypesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolObjectTypesSupported")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) IsBACnetConstructedDataProtocolObjectTypesSupported() {
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) deepCopy() *_BACnetConstructedDataProtocolObjectTypesSupported {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProtocolObjectTypesSupportedCopy := &_BACnetConstructedDataProtocolObjectTypesSupported{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ProtocolObjectTypesSupported.DeepCopy().(BACnetObjectTypesSupportedTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProtocolObjectTypesSupportedCopy
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
