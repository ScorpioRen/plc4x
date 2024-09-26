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

// UnknownMessage is the corresponding interface of UnknownMessage
type UnknownMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetUnknownData returns UnknownData (property field)
	GetUnknownData() []byte
	// IsUnknownMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUnknownMessage()
	// CreateBuilder creates a UnknownMessageBuilder
	CreateUnknownMessageBuilder() UnknownMessageBuilder
}

// _UnknownMessage is the data-structure of this message
type _UnknownMessage struct {
	KnxNetIpMessageContract
	UnknownData []byte

	// Arguments.
	TotalLength uint16
}

var _ UnknownMessage = (*_UnknownMessage)(nil)
var _ KnxNetIpMessageRequirements = (*_UnknownMessage)(nil)

// NewUnknownMessage factory function for _UnknownMessage
func NewUnknownMessage(unknownData []byte, totalLength uint16) *_UnknownMessage {
	_result := &_UnknownMessage{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		UnknownData:             unknownData,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UnknownMessageBuilder is a builder for UnknownMessage
type UnknownMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unknownData []byte) UnknownMessageBuilder
	// WithUnknownData adds UnknownData (property field)
	WithUnknownData(...byte) UnknownMessageBuilder
	// Build builds the UnknownMessage or returns an error if something is wrong
	Build() (UnknownMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UnknownMessage
}

// NewUnknownMessageBuilder() creates a UnknownMessageBuilder
func NewUnknownMessageBuilder() UnknownMessageBuilder {
	return &_UnknownMessageBuilder{_UnknownMessage: new(_UnknownMessage)}
}

type _UnknownMessageBuilder struct {
	*_UnknownMessage

	err *utils.MultiError
}

var _ (UnknownMessageBuilder) = (*_UnknownMessageBuilder)(nil)

func (m *_UnknownMessageBuilder) WithMandatoryFields(unknownData []byte) UnknownMessageBuilder {
	return m.WithUnknownData(unknownData...)
}

func (m *_UnknownMessageBuilder) WithUnknownData(unknownData ...byte) UnknownMessageBuilder {
	m.UnknownData = unknownData
	return m
}

func (m *_UnknownMessageBuilder) Build() (UnknownMessage, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._UnknownMessage.deepCopy(), nil
}

func (m *_UnknownMessageBuilder) MustBuild() UnknownMessage {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_UnknownMessageBuilder) DeepCopy() any {
	return m.CreateUnknownMessageBuilder()
}

// CreateUnknownMessageBuilder creates a UnknownMessageBuilder
func (m *_UnknownMessage) CreateUnknownMessageBuilder() UnknownMessageBuilder {
	if m == nil {
		return NewUnknownMessageBuilder()
	}
	return &_UnknownMessageBuilder{_UnknownMessage: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnknownMessage) GetMsgType() uint16 {
	return 0x020B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnknownMessage) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnknownMessage) GetUnknownData() []byte {
	return m.UnknownData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUnknownMessage(structType any) UnknownMessage {
	if casted, ok := structType.(UnknownMessage); ok {
		return casted
	}
	if casted, ok := structType.(*UnknownMessage); ok {
		return *casted
	}
	return nil
}

func (m *_UnknownMessage) GetTypeName() string {
	return "UnknownMessage"
}

func (m *_UnknownMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Array field
	if len(m.UnknownData) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownData))
	}

	return lengthInBits
}

func (m *_UnknownMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UnknownMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage, totalLength uint16) (__unknownMessage UnknownMessage, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnknownMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnknownMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unknownData, err := readBuffer.ReadByteArray("unknownData", int(int32(totalLength)-int32(int32(6))), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownData' field"))
	}
	m.UnknownData = unknownData

	if closeErr := readBuffer.CloseContext("UnknownMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnknownMessage")
	}

	return m, nil
}

func (m *_UnknownMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnknownMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnknownMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnknownMessage")
		}

		if err := WriteByteArrayField(ctx, "unknownData", m.GetUnknownData(), WriteByteArray(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownData' field")
		}

		if popErr := writeBuffer.PopContext("UnknownMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnknownMessage")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_UnknownMessage) GetTotalLength() uint16 {
	return m.TotalLength
}

//
////

func (m *_UnknownMessage) IsUnknownMessage() {}

func (m *_UnknownMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UnknownMessage) deepCopy() *_UnknownMessage {
	if m == nil {
		return nil
	}
	_UnknownMessageCopy := &_UnknownMessage{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.UnknownData),
		m.TotalLength,
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _UnknownMessageCopy
}

func (m *_UnknownMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
