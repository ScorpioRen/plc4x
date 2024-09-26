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

// AdsReadWriteResponse is the corresponding interface of AdsReadWriteResponse
type AdsReadWriteResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetData returns Data (property field)
	GetData() []byte
	// IsAdsReadWriteResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsReadWriteResponse()
	// CreateBuilder creates a AdsReadWriteResponseBuilder
	CreateAdsReadWriteResponseBuilder() AdsReadWriteResponseBuilder
}

// _AdsReadWriteResponse is the data-structure of this message
type _AdsReadWriteResponse struct {
	AmsPacketContract
	Result ReturnCode
	Data   []byte
}

var _ AdsReadWriteResponse = (*_AdsReadWriteResponse)(nil)
var _ AmsPacketRequirements = (*_AdsReadWriteResponse)(nil)

// NewAdsReadWriteResponse factory function for _AdsReadWriteResponse
func NewAdsReadWriteResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, result ReturnCode, data []byte) *_AdsReadWriteResponse {
	_result := &_AdsReadWriteResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:            result,
		Data:              data,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsReadWriteResponseBuilder is a builder for AdsReadWriteResponse
type AdsReadWriteResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(result ReturnCode, data []byte) AdsReadWriteResponseBuilder
	// WithResult adds Result (property field)
	WithResult(ReturnCode) AdsReadWriteResponseBuilder
	// WithData adds Data (property field)
	WithData(...byte) AdsReadWriteResponseBuilder
	// Build builds the AdsReadWriteResponse or returns an error if something is wrong
	Build() (AdsReadWriteResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsReadWriteResponse
}

// NewAdsReadWriteResponseBuilder() creates a AdsReadWriteResponseBuilder
func NewAdsReadWriteResponseBuilder() AdsReadWriteResponseBuilder {
	return &_AdsReadWriteResponseBuilder{_AdsReadWriteResponse: new(_AdsReadWriteResponse)}
}

type _AdsReadWriteResponseBuilder struct {
	*_AdsReadWriteResponse

	err *utils.MultiError
}

var _ (AdsReadWriteResponseBuilder) = (*_AdsReadWriteResponseBuilder)(nil)

func (m *_AdsReadWriteResponseBuilder) WithMandatoryFields(result ReturnCode, data []byte) AdsReadWriteResponseBuilder {
	return m.WithResult(result).WithData(data...)
}

func (m *_AdsReadWriteResponseBuilder) WithResult(result ReturnCode) AdsReadWriteResponseBuilder {
	m.Result = result
	return m
}

func (m *_AdsReadWriteResponseBuilder) WithData(data ...byte) AdsReadWriteResponseBuilder {
	m.Data = data
	return m
}

func (m *_AdsReadWriteResponseBuilder) Build() (AdsReadWriteResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AdsReadWriteResponse.deepCopy(), nil
}

func (m *_AdsReadWriteResponseBuilder) MustBuild() AdsReadWriteResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AdsReadWriteResponseBuilder) DeepCopy() any {
	return m.CreateAdsReadWriteResponseBuilder()
}

// CreateAdsReadWriteResponseBuilder creates a AdsReadWriteResponseBuilder
func (m *_AdsReadWriteResponse) CreateAdsReadWriteResponseBuilder() AdsReadWriteResponseBuilder {
	if m == nil {
		return NewAdsReadWriteResponseBuilder()
	}
	return &_AdsReadWriteResponseBuilder{_AdsReadWriteResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadWriteResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *_AdsReadWriteResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadWriteResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadWriteResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadWriteResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsReadWriteResponse(structType any) AdsReadWriteResponse {
	if casted, ok := structType.(AdsReadWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadWriteResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadWriteResponse) GetTypeName() string {
	return "AdsReadWriteResponse"
}

func (m *_AdsReadWriteResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsReadWriteResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadWriteResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadWriteResponse AdsReadWriteResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadWriteResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadWriteResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	length, err := ReadImplicitField[uint32](ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	data, err := readBuffer.ReadByteArray("data", int(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AdsReadWriteResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadWriteResponse")
	}

	return m, nil
}

func (m *_AdsReadWriteResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadWriteResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadWriteResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadWriteResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}
		length := uint32(uint32(len(m.GetData())))
		if err := WriteImplicitField(ctx, "length", length, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadWriteResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadWriteResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadWriteResponse) IsAdsReadWriteResponse() {}

func (m *_AdsReadWriteResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsReadWriteResponse) deepCopy() *_AdsReadWriteResponse {
	if m == nil {
		return nil
	}
	_AdsReadWriteResponseCopy := &_AdsReadWriteResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
		m.Result,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsReadWriteResponseCopy
}

func (m *_AdsReadWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
