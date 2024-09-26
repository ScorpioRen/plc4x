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

// DF1ResponseMessage is the corresponding interface of DF1ResponseMessage
type DF1ResponseMessage interface {
	DF1ResponseMessageContract
	DF1ResponseMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsDF1ResponseMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1ResponseMessage()
}

// DF1ResponseMessageContract provides a set of functions which can be overwritten by a sub struct
type DF1ResponseMessageContract interface {
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetTransactionCounter returns TransactionCounter (property field)
	GetTransactionCounter() uint16
	// GetPayloadLength() returns a parser argument
	GetPayloadLength() uint16
	// IsDF1ResponseMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1ResponseMessage()
}

// DF1ResponseMessageRequirements provides a set of functions which need to be implemented by a sub struct
type DF1ResponseMessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
}

// _DF1ResponseMessage is the data-structure of this message
type _DF1ResponseMessage struct {
	_SubType           DF1ResponseMessage
	DestinationAddress uint8
	SourceAddress      uint8
	Status             uint8
	TransactionCounter uint16

	// Arguments.
	PayloadLength uint16
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ DF1ResponseMessageContract = (*_DF1ResponseMessage)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1ResponseMessage) GetDestinationAddress() uint8 {
	return m.DestinationAddress
}

func (m *_DF1ResponseMessage) GetSourceAddress() uint8 {
	return m.SourceAddress
}

func (m *_DF1ResponseMessage) GetStatus() uint8 {
	return m.Status
}

func (m *_DF1ResponseMessage) GetTransactionCounter() uint16 {
	return m.TransactionCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1ResponseMessage factory function for _DF1ResponseMessage
func NewDF1ResponseMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, payloadLength uint16) *_DF1ResponseMessage {
	return &_DF1ResponseMessage{DestinationAddress: destinationAddress, SourceAddress: sourceAddress, Status: status, TransactionCounter: transactionCounter, PayloadLength: payloadLength}
}

// Deprecated: use the interface for direct cast
func CastDF1ResponseMessage(structType any) DF1ResponseMessage {
	if casted, ok := structType.(DF1ResponseMessage); ok {
		return casted
	}
	if casted, ok := structType.(*DF1ResponseMessage); ok {
		return *casted
	}
	return nil
}

func (m *_DF1ResponseMessage) GetTypeName() string {
	return "DF1ResponseMessage"
}

func (m *_DF1ResponseMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1ResponseMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func DF1ResponseMessageParse[T DF1ResponseMessage](ctx context.Context, theBytes []byte, payloadLength uint16) (T, error) {
	return DF1ResponseMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), payloadLength)
}

func DF1ResponseMessageParseWithBufferProducer[T DF1ResponseMessage](payloadLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := DF1ResponseMessageParseWithBuffer[T](ctx, readBuffer, payloadLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func DF1ResponseMessageParseWithBuffer[T DF1ResponseMessage](ctx context.Context, readBuffer utils.ReadBuffer, payloadLength uint16) (T, error) {
	v, err := (&_DF1ResponseMessage{PayloadLength: payloadLength}).parse(ctx, readBuffer, payloadLength)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_DF1ResponseMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, payloadLength uint16) (__dF1ResponseMessage DF1ResponseMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1ResponseMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1ResponseMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	destinationAddress, err := ReadSimpleField(ctx, "destinationAddress", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationAddress' field"))
	}
	m.DestinationAddress = destinationAddress

	sourceAddress, err := ReadSimpleField(ctx, "sourceAddress", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAddress' field"))
	}
	m.SourceAddress = sourceAddress

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	commandCode, err := ReadDiscriminatorField[uint8](ctx, "commandCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandCode' field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	transactionCounter, err := ReadSimpleField(ctx, "transactionCounter", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionCounter' field"))
	}
	m.TransactionCounter = transactionCounter

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child DF1ResponseMessage
	switch {
	case commandCode == 0x4F: // DF1CommandResponseMessageProtectedTypedLogicalRead
		if _child, err = new(_DF1CommandResponseMessageProtectedTypedLogicalRead).parse(ctx, readBuffer, m, payloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1CommandResponseMessageProtectedTypedLogicalRead for type-switch of DF1ResponseMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandCode=%v]", commandCode)
	}

	if closeErr := readBuffer.CloseContext("DF1ResponseMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1ResponseMessage")
	}

	return _child, nil
}

func (pm *_DF1ResponseMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1ResponseMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DF1ResponseMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1ResponseMessage")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[uint8](ctx, "destinationAddress", m.GetDestinationAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'destinationAddress' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sourceAddress", m.GetSourceAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sourceAddress' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteDiscriminatorField(ctx, "commandCode", m.GetCommandCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'commandCode' field")
	}

	if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if err := WriteSimpleField[uint16](ctx, "transactionCounter", m.GetTransactionCounter(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1ResponseMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1ResponseMessage")
	}
	return nil
}

////
// Arguments Getter

func (m *_DF1ResponseMessage) GetPayloadLength() uint16 {
	return m.PayloadLength
}

//
////

func (m *_DF1ResponseMessage) IsDF1ResponseMessage() {}
