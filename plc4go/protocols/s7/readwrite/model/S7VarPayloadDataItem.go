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
	"math"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7VarPayloadDataItem is the corresponding interface of S7VarPayloadDataItem
type S7VarPayloadDataItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetData returns Data (property field)
	GetData() []byte
}

// S7VarPayloadDataItemExactly can be used when we want exactly this type and not a type which fulfills S7VarPayloadDataItem.
// This is useful for switch cases.
type S7VarPayloadDataItemExactly interface {
	S7VarPayloadDataItem
	isS7VarPayloadDataItem() bool
}

// _S7VarPayloadDataItem is the data-structure of this message
type _S7VarPayloadDataItem struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	Data          []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7VarPayloadDataItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_S7VarPayloadDataItem) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_S7VarPayloadDataItem) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7VarPayloadDataItem factory function for _S7VarPayloadDataItem
func NewS7VarPayloadDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize, data []byte) *_S7VarPayloadDataItem {
	return &_S7VarPayloadDataItem{ReturnCode: returnCode, TransportSize: transportSize, Data: data}
}

// Deprecated: use the interface for direct cast
func CastS7VarPayloadDataItem(structType any) S7VarPayloadDataItem {
	if casted, ok := structType.(S7VarPayloadDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarPayloadDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarPayloadDataItem) GetTypeName() string {
	return "S7VarPayloadDataItem"
}

func (m *_S7VarPayloadDataItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Implicit Field (dataLength)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// Padding Field (padding)
	_timesPadding := uint8(utils.InlineIf((!(utils.GetLastItemFromContext(ctx))), func() any { return int32((int32(int32(len(m.GetData()))) % int32(int32(2)))) }, func() any { return int32(int32(0)) }).(int32))
	for ; _timesPadding > 0; _timesPadding-- {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_S7VarPayloadDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7VarPayloadDataItemParse(ctx context.Context, theBytes []byte) (S7VarPayloadDataItem, error) {
	return S7VarPayloadDataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7VarPayloadDataItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (S7VarPayloadDataItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7VarPayloadDataItem, error) {
		return S7VarPayloadDataItemParseWithBuffer(ctx, readBuffer)
	}
}

func S7VarPayloadDataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7VarPayloadDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarPayloadDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarPayloadDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	returnCode, err := ReadEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", ReadEnum(DataTransportErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnCode' field"))
	}

	transportSize, err := ReadEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", ReadEnum(DataTransportSizeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}

	dataLength, err := ReadImplicitField[uint16](ctx, "dataLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataLength' field"))
	}
	_ = dataLength

	data, err := readBuffer.ReadByteArray("data", int(utils.InlineIf(transportSize.SizeInBits(), func() any { return int32(math.Ceil(float64(dataLength) / float64(float64(8.0)))) }, func() any { return int32(dataLength) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if err := ReadPaddingField(ctx, ReadUnsignedByte(readBuffer, uint8(8)), (int)(utils.InlineIf((!(utils.GetLastItemFromContext(ctx))), func() any { return uint8((uint8(uint8(len(data))) % uint8(uint8(2)))) }, func() any { return uint8(uint8(0)) }).(uint8))); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing padding field"))
	}

	if closeErr := readBuffer.CloseContext("S7VarPayloadDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarPayloadDataItem")
	}

	// Create the instance
	return &_S7VarPayloadDataItem{
		ReturnCode:    returnCode,
		TransportSize: transportSize,
		Data:          data,
	}, nil
}

func (m *_S7VarPayloadDataItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7VarPayloadDataItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7VarPayloadDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7VarPayloadDataItem")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(ctx, m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transportSize")
	}
	_transportSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSize())
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transportSize")
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}
	dataLength := uint16(uint16(uint16(len(m.GetData()))) * uint16((utils.InlineIf((bool((m.GetTransportSize()) == (DataTransportSize_BIT))), func() any { return uint16(uint16(1)) }, func() any {
		return uint16((utils.InlineIf(m.GetTransportSize().SizeInBits(), func() any { return uint16(uint16(8)) }, func() any { return uint16(uint16(1)) }).(uint16)))
	}).(uint16))))
	if err := WriteImplicitField(ctx, "dataLength", dataLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataLength' field")
	}

	if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if err := WritePaddingField[uint8](ctx, "padding", int(utils.InlineIf((!(utils.GetLastItemFromContext(ctx))), func() any { return int32((int32(int32(len(m.GetData()))) % int32(int32(2)))) }, func() any { return int32(int32(0)) }).(int32)), uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'padding' field")
	}

	if popErr := writeBuffer.PopContext("S7VarPayloadDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7VarPayloadDataItem")
	}
	return nil
}

func (m *_S7VarPayloadDataItem) isS7VarPayloadDataItem() bool {
	return true
}

func (m *_S7VarPayloadDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
