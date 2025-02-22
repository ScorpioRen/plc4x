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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadInputRegistersRequest is the corresponding interface of ModbusPDUReadInputRegistersRequest
type ModbusPDUReadInputRegistersRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// IsModbusPDUReadInputRegistersRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadInputRegistersRequest()
	// CreateBuilder creates a ModbusPDUReadInputRegistersRequestBuilder
	CreateModbusPDUReadInputRegistersRequestBuilder() ModbusPDUReadInputRegistersRequestBuilder
}

// _ModbusPDUReadInputRegistersRequest is the data-structure of this message
type _ModbusPDUReadInputRegistersRequest struct {
	ModbusPDUContract
	StartingAddress uint16
	Quantity        uint16
}

var _ ModbusPDUReadInputRegistersRequest = (*_ModbusPDUReadInputRegistersRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadInputRegistersRequest)(nil)

// NewModbusPDUReadInputRegistersRequest factory function for _ModbusPDUReadInputRegistersRequest
func NewModbusPDUReadInputRegistersRequest(startingAddress uint16, quantity uint16) *_ModbusPDUReadInputRegistersRequest {
	_result := &_ModbusPDUReadInputRegistersRequest{
		ModbusPDUContract: NewModbusPDU(),
		StartingAddress:   startingAddress,
		Quantity:          quantity,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadInputRegistersRequestBuilder is a builder for ModbusPDUReadInputRegistersRequest
type ModbusPDUReadInputRegistersRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUReadInputRegistersRequestBuilder
	// WithStartingAddress adds StartingAddress (property field)
	WithStartingAddress(uint16) ModbusPDUReadInputRegistersRequestBuilder
	// WithQuantity adds Quantity (property field)
	WithQuantity(uint16) ModbusPDUReadInputRegistersRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUReadInputRegistersRequest or returns an error if something is wrong
	Build() (ModbusPDUReadInputRegistersRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadInputRegistersRequest
}

// NewModbusPDUReadInputRegistersRequestBuilder() creates a ModbusPDUReadInputRegistersRequestBuilder
func NewModbusPDUReadInputRegistersRequestBuilder() ModbusPDUReadInputRegistersRequestBuilder {
	return &_ModbusPDUReadInputRegistersRequestBuilder{_ModbusPDUReadInputRegistersRequest: new(_ModbusPDUReadInputRegistersRequest)}
}

type _ModbusPDUReadInputRegistersRequestBuilder struct {
	*_ModbusPDUReadInputRegistersRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadInputRegistersRequestBuilder) = (*_ModbusPDUReadInputRegistersRequestBuilder)(nil)

func (b *_ModbusPDUReadInputRegistersRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUReadInputRegistersRequest
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUReadInputRegistersRequestBuilder {
	return b.WithStartingAddress(startingAddress).WithQuantity(quantity)
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) WithStartingAddress(startingAddress uint16) ModbusPDUReadInputRegistersRequestBuilder {
	b.StartingAddress = startingAddress
	return b
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) WithQuantity(quantity uint16) ModbusPDUReadInputRegistersRequestBuilder {
	b.Quantity = quantity
	return b
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) Build() (ModbusPDUReadInputRegistersRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadInputRegistersRequest.deepCopy(), nil
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) MustBuild() ModbusPDUReadInputRegistersRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadInputRegistersRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadInputRegistersRequestBuilder().(*_ModbusPDUReadInputRegistersRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadInputRegistersRequestBuilder creates a ModbusPDUReadInputRegistersRequestBuilder
func (b *_ModbusPDUReadInputRegistersRequest) CreateModbusPDUReadInputRegistersRequestBuilder() ModbusPDUReadInputRegistersRequestBuilder {
	if b == nil {
		return NewModbusPDUReadInputRegistersRequestBuilder()
	}
	return &_ModbusPDUReadInputRegistersRequestBuilder{_ModbusPDUReadInputRegistersRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadInputRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadInputRegistersRequest) GetFunctionFlag() uint8 {
	return 0x04
}

func (m *_ModbusPDUReadInputRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadInputRegistersRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadInputRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUReadInputRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadInputRegistersRequest(structType any) ModbusPDUReadInputRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadInputRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadInputRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadInputRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadInputRegistersRequest"
}

func (m *_ModbusPDUReadInputRegistersRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadInputRegistersRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadInputRegistersRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadInputRegistersRequest ModbusPDUReadInputRegistersRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadInputRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadInputRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingAddress, err := ReadSimpleField(ctx, "startingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingAddress' field"))
	}
	m.StartingAddress = startingAddress

	quantity, err := ReadSimpleField(ctx, "quantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'quantity' field"))
	}
	m.Quantity = quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadInputRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadInputRegistersRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadInputRegistersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadInputRegistersRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadInputRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadInputRegistersRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "startingAddress", m.GetStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "quantity", m.GetQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadInputRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadInputRegistersRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadInputRegistersRequest) IsModbusPDUReadInputRegistersRequest() {}

func (m *_ModbusPDUReadInputRegistersRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadInputRegistersRequest) deepCopy() *_ModbusPDUReadInputRegistersRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReadInputRegistersRequestCopy := &_ModbusPDUReadInputRegistersRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.StartingAddress,
		m.Quantity,
	}
	_ModbusPDUReadInputRegistersRequestCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadInputRegistersRequestCopy
}

func (m *_ModbusPDUReadInputRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
