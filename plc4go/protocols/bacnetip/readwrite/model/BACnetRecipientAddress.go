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

// BACnetRecipientAddress is the corresponding interface of BACnetRecipientAddress
type BACnetRecipientAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetRecipient
	// GetAddressValue returns AddressValue (property field)
	GetAddressValue() BACnetAddressEnclosed
	// IsBACnetRecipientAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipientAddress()
	// CreateBuilder creates a BACnetRecipientAddressBuilder
	CreateBACnetRecipientAddressBuilder() BACnetRecipientAddressBuilder
}

// _BACnetRecipientAddress is the data-structure of this message
type _BACnetRecipientAddress struct {
	BACnetRecipientContract
	AddressValue BACnetAddressEnclosed
}

var _ BACnetRecipientAddress = (*_BACnetRecipientAddress)(nil)
var _ BACnetRecipientRequirements = (*_BACnetRecipientAddress)(nil)

// NewBACnetRecipientAddress factory function for _BACnetRecipientAddress
func NewBACnetRecipientAddress(peekedTagHeader BACnetTagHeader, addressValue BACnetAddressEnclosed) *_BACnetRecipientAddress {
	if addressValue == nil {
		panic("addressValue of type BACnetAddressEnclosed for BACnetRecipientAddress must not be nil")
	}
	_result := &_BACnetRecipientAddress{
		BACnetRecipientContract: NewBACnetRecipient(peekedTagHeader),
		AddressValue:            addressValue,
	}
	_result.BACnetRecipientContract.(*_BACnetRecipient)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRecipientAddressBuilder is a builder for BACnetRecipientAddress
type BACnetRecipientAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(addressValue BACnetAddressEnclosed) BACnetRecipientAddressBuilder
	// WithAddressValue adds AddressValue (property field)
	WithAddressValue(BACnetAddressEnclosed) BACnetRecipientAddressBuilder
	// WithAddressValueBuilder adds AddressValue (property field) which is build by the builder
	WithAddressValueBuilder(func(BACnetAddressEnclosedBuilder) BACnetAddressEnclosedBuilder) BACnetRecipientAddressBuilder
	// Build builds the BACnetRecipientAddress or returns an error if something is wrong
	Build() (BACnetRecipientAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRecipientAddress
}

// NewBACnetRecipientAddressBuilder() creates a BACnetRecipientAddressBuilder
func NewBACnetRecipientAddressBuilder() BACnetRecipientAddressBuilder {
	return &_BACnetRecipientAddressBuilder{_BACnetRecipientAddress: new(_BACnetRecipientAddress)}
}

type _BACnetRecipientAddressBuilder struct {
	*_BACnetRecipientAddress

	err *utils.MultiError
}

var _ (BACnetRecipientAddressBuilder) = (*_BACnetRecipientAddressBuilder)(nil)

func (m *_BACnetRecipientAddressBuilder) WithMandatoryFields(addressValue BACnetAddressEnclosed) BACnetRecipientAddressBuilder {
	return m.WithAddressValue(addressValue)
}

func (m *_BACnetRecipientAddressBuilder) WithAddressValue(addressValue BACnetAddressEnclosed) BACnetRecipientAddressBuilder {
	m.AddressValue = addressValue
	return m
}

func (m *_BACnetRecipientAddressBuilder) WithAddressValueBuilder(builderSupplier func(BACnetAddressEnclosedBuilder) BACnetAddressEnclosedBuilder) BACnetRecipientAddressBuilder {
	builder := builderSupplier(m.AddressValue.CreateBACnetAddressEnclosedBuilder())
	var err error
	m.AddressValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetAddressEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetRecipientAddressBuilder) Build() (BACnetRecipientAddress, error) {
	if m.AddressValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'addressValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetRecipientAddress.deepCopy(), nil
}

func (m *_BACnetRecipientAddressBuilder) MustBuild() BACnetRecipientAddress {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetRecipientAddressBuilder) DeepCopy() any {
	return m.CreateBACnetRecipientAddressBuilder()
}

// CreateBACnetRecipientAddressBuilder creates a BACnetRecipientAddressBuilder
func (m *_BACnetRecipientAddress) CreateBACnetRecipientAddressBuilder() BACnetRecipientAddressBuilder {
	if m == nil {
		return NewBACnetRecipientAddressBuilder()
	}
	return &_BACnetRecipientAddressBuilder{_BACnetRecipientAddress: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetRecipientAddress) GetParent() BACnetRecipientContract {
	return m.BACnetRecipientContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientAddress) GetAddressValue() BACnetAddressEnclosed {
	return m.AddressValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetRecipientAddress(structType any) BACnetRecipientAddress {
	if casted, ok := structType.(BACnetRecipientAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientAddress) GetTypeName() string {
	return "BACnetRecipientAddress"
}

func (m *_BACnetRecipientAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetRecipientContract.(*_BACnetRecipient).getLengthInBits(ctx))

	// Simple field (addressValue)
	lengthInBits += m.AddressValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetRecipientAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetRecipient) (__bACnetRecipientAddress BACnetRecipientAddress, err error) {
	m.BACnetRecipientContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	addressValue, err := ReadSimpleField[BACnetAddressEnclosed](ctx, "addressValue", ReadComplex[BACnetAddressEnclosed](BACnetAddressEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'addressValue' field"))
	}
	m.AddressValue = addressValue

	if closeErr := readBuffer.CloseContext("BACnetRecipientAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientAddress")
	}

	return m, nil
}

func (m *_BACnetRecipientAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetRecipientAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetRecipientAddress")
		}

		if err := WriteSimpleField[BACnetAddressEnclosed](ctx, "addressValue", m.GetAddressValue(), WriteComplex[BACnetAddressEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'addressValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetRecipientAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetRecipientAddress")
		}
		return nil
	}
	return m.BACnetRecipientContract.(*_BACnetRecipient).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetRecipientAddress) IsBACnetRecipientAddress() {}

func (m *_BACnetRecipientAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRecipientAddress) deepCopy() *_BACnetRecipientAddress {
	if m == nil {
		return nil
	}
	_BACnetRecipientAddressCopy := &_BACnetRecipientAddress{
		m.BACnetRecipientContract.(*_BACnetRecipient).deepCopy(),
		m.AddressValue.DeepCopy().(BACnetAddressEnclosed),
	}
	m.BACnetRecipientContract.(*_BACnetRecipient)._SubType = m
	return _BACnetRecipientAddressCopy
}

func (m *_BACnetRecipientAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
