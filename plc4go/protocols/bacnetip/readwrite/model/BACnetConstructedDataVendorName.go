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

// BACnetConstructedDataVendorName is the corresponding interface of BACnetConstructedDataVendorName
type BACnetConstructedDataVendorName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetVendorName returns VendorName (property field)
	GetVendorName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataVendorName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataVendorName()
	// CreateBuilder creates a BACnetConstructedDataVendorNameBuilder
	CreateBACnetConstructedDataVendorNameBuilder() BACnetConstructedDataVendorNameBuilder
}

// _BACnetConstructedDataVendorName is the data-structure of this message
type _BACnetConstructedDataVendorName struct {
	BACnetConstructedDataContract
	VendorName BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataVendorName = (*_BACnetConstructedDataVendorName)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataVendorName)(nil)

// NewBACnetConstructedDataVendorName factory function for _BACnetConstructedDataVendorName
func NewBACnetConstructedDataVendorName(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, vendorName BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVendorName {
	if vendorName == nil {
		panic("vendorName of type BACnetApplicationTagCharacterString for BACnetConstructedDataVendorName must not be nil")
	}
	_result := &_BACnetConstructedDataVendorName{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		VendorName:                    vendorName,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataVendorNameBuilder is a builder for BACnetConstructedDataVendorName
type BACnetConstructedDataVendorNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(vendorName BACnetApplicationTagCharacterString) BACnetConstructedDataVendorNameBuilder
	// WithVendorName adds VendorName (property field)
	WithVendorName(BACnetApplicationTagCharacterString) BACnetConstructedDataVendorNameBuilder
	// WithVendorNameBuilder adds VendorName (property field) which is build by the builder
	WithVendorNameBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataVendorNameBuilder
	// Build builds the BACnetConstructedDataVendorName or returns an error if something is wrong
	Build() (BACnetConstructedDataVendorName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataVendorName
}

// NewBACnetConstructedDataVendorNameBuilder() creates a BACnetConstructedDataVendorNameBuilder
func NewBACnetConstructedDataVendorNameBuilder() BACnetConstructedDataVendorNameBuilder {
	return &_BACnetConstructedDataVendorNameBuilder{_BACnetConstructedDataVendorName: new(_BACnetConstructedDataVendorName)}
}

type _BACnetConstructedDataVendorNameBuilder struct {
	*_BACnetConstructedDataVendorName

	err *utils.MultiError
}

var _ (BACnetConstructedDataVendorNameBuilder) = (*_BACnetConstructedDataVendorNameBuilder)(nil)

func (m *_BACnetConstructedDataVendorNameBuilder) WithMandatoryFields(vendorName BACnetApplicationTagCharacterString) BACnetConstructedDataVendorNameBuilder {
	return m.WithVendorName(vendorName)
}

func (m *_BACnetConstructedDataVendorNameBuilder) WithVendorName(vendorName BACnetApplicationTagCharacterString) BACnetConstructedDataVendorNameBuilder {
	m.VendorName = vendorName
	return m
}

func (m *_BACnetConstructedDataVendorNameBuilder) WithVendorNameBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataVendorNameBuilder {
	builder := builderSupplier(m.VendorName.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	m.VendorName, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataVendorNameBuilder) Build() (BACnetConstructedDataVendorName, error) {
	if m.VendorName == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'vendorName' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataVendorName.deepCopy(), nil
}

func (m *_BACnetConstructedDataVendorNameBuilder) MustBuild() BACnetConstructedDataVendorName {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataVendorNameBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataVendorNameBuilder()
}

// CreateBACnetConstructedDataVendorNameBuilder creates a BACnetConstructedDataVendorNameBuilder
func (m *_BACnetConstructedDataVendorName) CreateBACnetConstructedDataVendorNameBuilder() BACnetConstructedDataVendorNameBuilder {
	if m == nil {
		return NewBACnetConstructedDataVendorNameBuilder()
	}
	return &_BACnetConstructedDataVendorNameBuilder{_BACnetConstructedDataVendorName: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVendorName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVendorName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VENDOR_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVendorName) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVendorName) GetVendorName() BACnetApplicationTagCharacterString {
	return m.VendorName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVendorName) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetVendorName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVendorName(structType any) BACnetConstructedDataVendorName {
	if casted, ok := structType.(BACnetConstructedDataVendorName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVendorName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVendorName) GetTypeName() string {
	return "BACnetConstructedDataVendorName"
}

func (m *_BACnetConstructedDataVendorName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (vendorName)
	lengthInBits += m.VendorName.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataVendorName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataVendorName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataVendorName BACnetConstructedDataVendorName, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVendorName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVendorName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vendorName, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "vendorName", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorName' field"))
	}
	m.VendorName = vendorName

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), vendorName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVendorName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVendorName")
	}

	return m, nil
}

func (m *_BACnetConstructedDataVendorName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVendorName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVendorName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVendorName")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "vendorName", m.GetVendorName(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorName' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVendorName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVendorName")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVendorName) IsBACnetConstructedDataVendorName() {}

func (m *_BACnetConstructedDataVendorName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataVendorName) deepCopy() *_BACnetConstructedDataVendorName {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataVendorNameCopy := &_BACnetConstructedDataVendorName{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.VendorName.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataVendorNameCopy
}

func (m *_BACnetConstructedDataVendorName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
