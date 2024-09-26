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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionStateResponse is the corresponding interface of ConnectionStateResponse
type ConnectionStateResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
	// IsConnectionStateResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionStateResponse()
	// CreateBuilder creates a ConnectionStateResponseBuilder
	CreateConnectionStateResponseBuilder() ConnectionStateResponseBuilder
}

// _ConnectionStateResponse is the data-structure of this message
type _ConnectionStateResponse struct {
	KnxNetIpMessageContract
	CommunicationChannelId uint8
	Status                 Status
}

var _ ConnectionStateResponse = (*_ConnectionStateResponse)(nil)
var _ KnxNetIpMessageRequirements = (*_ConnectionStateResponse)(nil)

// NewConnectionStateResponse factory function for _ConnectionStateResponse
func NewConnectionStateResponse(communicationChannelId uint8, status Status) *_ConnectionStateResponse {
	_result := &_ConnectionStateResponse{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		CommunicationChannelId:  communicationChannelId,
		Status:                  status,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionStateResponseBuilder is a builder for ConnectionStateResponse
type ConnectionStateResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(communicationChannelId uint8, status Status) ConnectionStateResponseBuilder
	// WithCommunicationChannelId adds CommunicationChannelId (property field)
	WithCommunicationChannelId(uint8) ConnectionStateResponseBuilder
	// WithStatus adds Status (property field)
	WithStatus(Status) ConnectionStateResponseBuilder
	// Build builds the ConnectionStateResponse or returns an error if something is wrong
	Build() (ConnectionStateResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionStateResponse
}

// NewConnectionStateResponseBuilder() creates a ConnectionStateResponseBuilder
func NewConnectionStateResponseBuilder() ConnectionStateResponseBuilder {
	return &_ConnectionStateResponseBuilder{_ConnectionStateResponse: new(_ConnectionStateResponse)}
}

type _ConnectionStateResponseBuilder struct {
	*_ConnectionStateResponse

	err *utils.MultiError
}

var _ (ConnectionStateResponseBuilder) = (*_ConnectionStateResponseBuilder)(nil)

func (m *_ConnectionStateResponseBuilder) WithMandatoryFields(communicationChannelId uint8, status Status) ConnectionStateResponseBuilder {
	return m.WithCommunicationChannelId(communicationChannelId).WithStatus(status)
}

func (m *_ConnectionStateResponseBuilder) WithCommunicationChannelId(communicationChannelId uint8) ConnectionStateResponseBuilder {
	m.CommunicationChannelId = communicationChannelId
	return m
}

func (m *_ConnectionStateResponseBuilder) WithStatus(status Status) ConnectionStateResponseBuilder {
	m.Status = status
	return m
}

func (m *_ConnectionStateResponseBuilder) Build() (ConnectionStateResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ConnectionStateResponse.deepCopy(), nil
}

func (m *_ConnectionStateResponseBuilder) MustBuild() ConnectionStateResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ConnectionStateResponseBuilder) DeepCopy() any {
	return m.CreateConnectionStateResponseBuilder()
}

// CreateConnectionStateResponseBuilder creates a ConnectionStateResponseBuilder
func (m *_ConnectionStateResponse) CreateConnectionStateResponseBuilder() ConnectionStateResponseBuilder {
	if m == nil {
		return NewConnectionStateResponseBuilder()
	}
	return &_ConnectionStateResponseBuilder{_ConnectionStateResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionStateResponse) GetMsgType() uint16 {
	return 0x0208
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionStateResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionStateResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_ConnectionStateResponse) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectionStateResponse(structType any) ConnectionStateResponse {
	if casted, ok := structType.(ConnectionStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionStateResponse) GetTypeName() string {
	return "ConnectionStateResponse"
}

func (m *_ConnectionStateResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionStateResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionStateResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__connectionStateResponse ConnectionStateResponse, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	status, err := ReadEnumField[Status](ctx, "status", "Status", ReadEnum(StatusByValue, ReadUnsignedByte(readBuffer, uint8(8))), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	if closeErr := readBuffer.CloseContext("ConnectionStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionStateResponse")
	}

	return m, nil
}

func (m *_ConnectionStateResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionStateResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionStateResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
		}

		if err := WriteSimpleEnumField[Status](ctx, "status", "Status", m.GetStatus(), WriteEnum[Status, uint8](Status.GetValue, Status.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionStateResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionStateResponse) IsConnectionStateResponse() {}

func (m *_ConnectionStateResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionStateResponse) deepCopy() *_ConnectionStateResponse {
	if m == nil {
		return nil
	}
	_ConnectionStateResponseCopy := &_ConnectionStateResponse{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		m.CommunicationChannelId,
		m.Status,
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _ConnectionStateResponseCopy
}

func (m *_ConnectionStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
