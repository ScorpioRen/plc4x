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

// DeviceConfigurationAckDataBlock is the corresponding interface of DeviceConfigurationAckDataBlock
type DeviceConfigurationAckDataBlock interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
	// IsDeviceConfigurationAckDataBlock is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeviceConfigurationAckDataBlock()
	// CreateBuilder creates a DeviceConfigurationAckDataBlockBuilder
	CreateDeviceConfigurationAckDataBlockBuilder() DeviceConfigurationAckDataBlockBuilder
}

// _DeviceConfigurationAckDataBlock is the data-structure of this message
type _DeviceConfigurationAckDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	Status                 Status
}

var _ DeviceConfigurationAckDataBlock = (*_DeviceConfigurationAckDataBlock)(nil)

// NewDeviceConfigurationAckDataBlock factory function for _DeviceConfigurationAckDataBlock
func NewDeviceConfigurationAckDataBlock(communicationChannelId uint8, sequenceCounter uint8, status Status) *_DeviceConfigurationAckDataBlock {
	return &_DeviceConfigurationAckDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter, Status: status}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeviceConfigurationAckDataBlockBuilder is a builder for DeviceConfigurationAckDataBlock
type DeviceConfigurationAckDataBlockBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(communicationChannelId uint8, sequenceCounter uint8, status Status) DeviceConfigurationAckDataBlockBuilder
	// WithCommunicationChannelId adds CommunicationChannelId (property field)
	WithCommunicationChannelId(uint8) DeviceConfigurationAckDataBlockBuilder
	// WithSequenceCounter adds SequenceCounter (property field)
	WithSequenceCounter(uint8) DeviceConfigurationAckDataBlockBuilder
	// WithStatus adds Status (property field)
	WithStatus(Status) DeviceConfigurationAckDataBlockBuilder
	// Build builds the DeviceConfigurationAckDataBlock or returns an error if something is wrong
	Build() (DeviceConfigurationAckDataBlock, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeviceConfigurationAckDataBlock
}

// NewDeviceConfigurationAckDataBlockBuilder() creates a DeviceConfigurationAckDataBlockBuilder
func NewDeviceConfigurationAckDataBlockBuilder() DeviceConfigurationAckDataBlockBuilder {
	return &_DeviceConfigurationAckDataBlockBuilder{_DeviceConfigurationAckDataBlock: new(_DeviceConfigurationAckDataBlock)}
}

type _DeviceConfigurationAckDataBlockBuilder struct {
	*_DeviceConfigurationAckDataBlock

	err *utils.MultiError
}

var _ (DeviceConfigurationAckDataBlockBuilder) = (*_DeviceConfigurationAckDataBlockBuilder)(nil)

func (m *_DeviceConfigurationAckDataBlockBuilder) WithMandatoryFields(communicationChannelId uint8, sequenceCounter uint8, status Status) DeviceConfigurationAckDataBlockBuilder {
	return m.WithCommunicationChannelId(communicationChannelId).WithSequenceCounter(sequenceCounter).WithStatus(status)
}

func (m *_DeviceConfigurationAckDataBlockBuilder) WithCommunicationChannelId(communicationChannelId uint8) DeviceConfigurationAckDataBlockBuilder {
	m.CommunicationChannelId = communicationChannelId
	return m
}

func (m *_DeviceConfigurationAckDataBlockBuilder) WithSequenceCounter(sequenceCounter uint8) DeviceConfigurationAckDataBlockBuilder {
	m.SequenceCounter = sequenceCounter
	return m
}

func (m *_DeviceConfigurationAckDataBlockBuilder) WithStatus(status Status) DeviceConfigurationAckDataBlockBuilder {
	m.Status = status
	return m
}

func (m *_DeviceConfigurationAckDataBlockBuilder) Build() (DeviceConfigurationAckDataBlock, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._DeviceConfigurationAckDataBlock.deepCopy(), nil
}

func (m *_DeviceConfigurationAckDataBlockBuilder) MustBuild() DeviceConfigurationAckDataBlock {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_DeviceConfigurationAckDataBlockBuilder) DeepCopy() any {
	return m.CreateDeviceConfigurationAckDataBlockBuilder()
}

// CreateDeviceConfigurationAckDataBlockBuilder creates a DeviceConfigurationAckDataBlockBuilder
func (m *_DeviceConfigurationAckDataBlock) CreateDeviceConfigurationAckDataBlockBuilder() DeviceConfigurationAckDataBlockBuilder {
	if m == nil {
		return NewDeviceConfigurationAckDataBlockBuilder()
	}
	return &_DeviceConfigurationAckDataBlockBuilder{_DeviceConfigurationAckDataBlock: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationAckDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DeviceConfigurationAckDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

func (m *_DeviceConfigurationAckDataBlock) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationAckDataBlock(structType any) DeviceConfigurationAckDataBlock {
	if casted, ok := structType.(DeviceConfigurationAckDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationAckDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationAckDataBlock) GetTypeName() string {
	return "DeviceConfigurationAckDataBlock"
}

func (m *_DeviceConfigurationAckDataBlock) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DeviceConfigurationAckDataBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceConfigurationAckDataBlockParse(ctx context.Context, theBytes []byte) (DeviceConfigurationAckDataBlock, error) {
	return DeviceConfigurationAckDataBlockParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeviceConfigurationAckDataBlockParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationAckDataBlock, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationAckDataBlock, error) {
		return DeviceConfigurationAckDataBlockParseWithBuffer(ctx, readBuffer)
	}
}

func DeviceConfigurationAckDataBlockParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationAckDataBlock, error) {
	v, err := (&_DeviceConfigurationAckDataBlock{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_DeviceConfigurationAckDataBlock) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationAckDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationAckDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	communicationChannelId, err := ReadSimpleField(ctx, "communicationChannelId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationChannelId' field"))
	}
	m.CommunicationChannelId = communicationChannelId

	sequenceCounter, err := ReadSimpleField(ctx, "sequenceCounter", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceCounter' field"))
	}
	m.SequenceCounter = sequenceCounter

	status, err := ReadEnumField[Status](ctx, "status", "Status", ReadEnum(StatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	if closeErr := readBuffer.CloseContext("DeviceConfigurationAckDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationAckDataBlock")
	}

	return m, nil
}

func (m *_DeviceConfigurationAckDataBlock) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationAckDataBlock) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DeviceConfigurationAckDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationAckDataBlock")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "communicationChannelId", m.GetCommunicationChannelId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationChannelId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sequenceCounter", m.GetSequenceCounter(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'sequenceCounter' field")
	}

	if err := WriteSimpleEnumField[Status](ctx, "status", "Status", m.GetStatus(), WriteEnum[Status, uint8](Status.GetValue, Status.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'status' field")
	}

	if popErr := writeBuffer.PopContext("DeviceConfigurationAckDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceConfigurationAckDataBlock")
	}
	return nil
}

func (m *_DeviceConfigurationAckDataBlock) IsDeviceConfigurationAckDataBlock() {}

func (m *_DeviceConfigurationAckDataBlock) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeviceConfigurationAckDataBlock) deepCopy() *_DeviceConfigurationAckDataBlock {
	if m == nil {
		return nil
	}
	_DeviceConfigurationAckDataBlockCopy := &_DeviceConfigurationAckDataBlock{
		m.CommunicationChannelId,
		m.SequenceCounter,
		m.Status,
	}
	return _DeviceConfigurationAckDataBlockCopy
}

func (m *_DeviceConfigurationAckDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
