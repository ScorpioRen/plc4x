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

// NLMSecurityResponse is the corresponding interface of NLMSecurityResponse
type NLMSecurityResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetResponseCode returns ResponseCode (property field)
	GetResponseCode() SecurityResponseCode
	// GetOriginalMessageId returns OriginalMessageId (property field)
	GetOriginalMessageId() uint32
	// GetOriginalTimestamp returns OriginalTimestamp (property field)
	GetOriginalTimestamp() uint32
	// GetVariableParameters returns VariableParameters (property field)
	GetVariableParameters() []byte
	// IsNLMSecurityResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMSecurityResponse()
	// CreateBuilder creates a NLMSecurityResponseBuilder
	CreateNLMSecurityResponseBuilder() NLMSecurityResponseBuilder
}

// _NLMSecurityResponse is the data-structure of this message
type _NLMSecurityResponse struct {
	NLMContract
	ResponseCode       SecurityResponseCode
	OriginalMessageId  uint32
	OriginalTimestamp  uint32
	VariableParameters []byte
}

var _ NLMSecurityResponse = (*_NLMSecurityResponse)(nil)
var _ NLMRequirements = (*_NLMSecurityResponse)(nil)

// NewNLMSecurityResponse factory function for _NLMSecurityResponse
func NewNLMSecurityResponse(responseCode SecurityResponseCode, originalMessageId uint32, originalTimestamp uint32, variableParameters []byte, apduLength uint16) *_NLMSecurityResponse {
	_result := &_NLMSecurityResponse{
		NLMContract:        NewNLM(apduLength),
		ResponseCode:       responseCode,
		OriginalMessageId:  originalMessageId,
		OriginalTimestamp:  originalTimestamp,
		VariableParameters: variableParameters,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMSecurityResponseBuilder is a builder for NLMSecurityResponse
type NLMSecurityResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseCode SecurityResponseCode, originalMessageId uint32, originalTimestamp uint32, variableParameters []byte) NLMSecurityResponseBuilder
	// WithResponseCode adds ResponseCode (property field)
	WithResponseCode(SecurityResponseCode) NLMSecurityResponseBuilder
	// WithOriginalMessageId adds OriginalMessageId (property field)
	WithOriginalMessageId(uint32) NLMSecurityResponseBuilder
	// WithOriginalTimestamp adds OriginalTimestamp (property field)
	WithOriginalTimestamp(uint32) NLMSecurityResponseBuilder
	// WithVariableParameters adds VariableParameters (property field)
	WithVariableParameters(...byte) NLMSecurityResponseBuilder
	// Build builds the NLMSecurityResponse or returns an error if something is wrong
	Build() (NLMSecurityResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMSecurityResponse
}

// NewNLMSecurityResponseBuilder() creates a NLMSecurityResponseBuilder
func NewNLMSecurityResponseBuilder() NLMSecurityResponseBuilder {
	return &_NLMSecurityResponseBuilder{_NLMSecurityResponse: new(_NLMSecurityResponse)}
}

type _NLMSecurityResponseBuilder struct {
	*_NLMSecurityResponse

	err *utils.MultiError
}

var _ (NLMSecurityResponseBuilder) = (*_NLMSecurityResponseBuilder)(nil)

func (m *_NLMSecurityResponseBuilder) WithMandatoryFields(responseCode SecurityResponseCode, originalMessageId uint32, originalTimestamp uint32, variableParameters []byte) NLMSecurityResponseBuilder {
	return m.WithResponseCode(responseCode).WithOriginalMessageId(originalMessageId).WithOriginalTimestamp(originalTimestamp).WithVariableParameters(variableParameters...)
}

func (m *_NLMSecurityResponseBuilder) WithResponseCode(responseCode SecurityResponseCode) NLMSecurityResponseBuilder {
	m.ResponseCode = responseCode
	return m
}

func (m *_NLMSecurityResponseBuilder) WithOriginalMessageId(originalMessageId uint32) NLMSecurityResponseBuilder {
	m.OriginalMessageId = originalMessageId
	return m
}

func (m *_NLMSecurityResponseBuilder) WithOriginalTimestamp(originalTimestamp uint32) NLMSecurityResponseBuilder {
	m.OriginalTimestamp = originalTimestamp
	return m
}

func (m *_NLMSecurityResponseBuilder) WithVariableParameters(variableParameters ...byte) NLMSecurityResponseBuilder {
	m.VariableParameters = variableParameters
	return m
}

func (m *_NLMSecurityResponseBuilder) Build() (NLMSecurityResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._NLMSecurityResponse.deepCopy(), nil
}

func (m *_NLMSecurityResponseBuilder) MustBuild() NLMSecurityResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_NLMSecurityResponseBuilder) DeepCopy() any {
	return m.CreateNLMSecurityResponseBuilder()
}

// CreateNLMSecurityResponseBuilder creates a NLMSecurityResponseBuilder
func (m *_NLMSecurityResponse) CreateNLMSecurityResponseBuilder() NLMSecurityResponseBuilder {
	if m == nil {
		return NewNLMSecurityResponseBuilder()
	}
	return &_NLMSecurityResponseBuilder{_NLMSecurityResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMSecurityResponse) GetMessageType() uint8 {
	return 0x0C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMSecurityResponse) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMSecurityResponse) GetResponseCode() SecurityResponseCode {
	return m.ResponseCode
}

func (m *_NLMSecurityResponse) GetOriginalMessageId() uint32 {
	return m.OriginalMessageId
}

func (m *_NLMSecurityResponse) GetOriginalTimestamp() uint32 {
	return m.OriginalTimestamp
}

func (m *_NLMSecurityResponse) GetVariableParameters() []byte {
	return m.VariableParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMSecurityResponse(structType any) NLMSecurityResponse {
	if casted, ok := structType.(NLMSecurityResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NLMSecurityResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NLMSecurityResponse) GetTypeName() string {
	return "NLMSecurityResponse"
}

func (m *_NLMSecurityResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (responseCode)
	lengthInBits += 8

	// Simple field (originalMessageId)
	lengthInBits += 32

	// Simple field (originalTimestamp)
	lengthInBits += 32

	// Array field
	if len(m.VariableParameters) > 0 {
		lengthInBits += 8 * uint16(len(m.VariableParameters))
	}

	return lengthInBits
}

func (m *_NLMSecurityResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMSecurityResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMSecurityResponse NLMSecurityResponse, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMSecurityResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMSecurityResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseCode, err := ReadEnumField[SecurityResponseCode](ctx, "responseCode", "SecurityResponseCode", ReadEnum(SecurityResponseCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseCode' field"))
	}
	m.ResponseCode = responseCode

	originalMessageId, err := ReadSimpleField(ctx, "originalMessageId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalMessageId' field"))
	}
	m.OriginalMessageId = originalMessageId

	originalTimestamp, err := ReadSimpleField(ctx, "originalTimestamp", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalTimestamp' field"))
	}
	m.OriginalTimestamp = originalTimestamp

	variableParameters, err := readBuffer.ReadByteArray("variableParameters", int(int32(apduLength)-int32((int32(int32(int32(int32(1))+int32(int32(1)))+int32(int32(4)))+int32(int32(4))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableParameters' field"))
	}
	m.VariableParameters = variableParameters

	if closeErr := readBuffer.CloseContext("NLMSecurityResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMSecurityResponse")
	}

	return m, nil
}

func (m *_NLMSecurityResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMSecurityResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMSecurityResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMSecurityResponse")
		}

		if err := WriteSimpleEnumField[SecurityResponseCode](ctx, "responseCode", "SecurityResponseCode", m.GetResponseCode(), WriteEnum[SecurityResponseCode, uint8](SecurityResponseCode.GetValue, SecurityResponseCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'responseCode' field")
		}

		if err := WriteSimpleField[uint32](ctx, "originalMessageId", m.GetOriginalMessageId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalMessageId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "originalTimestamp", m.GetOriginalTimestamp(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalTimestamp' field")
		}

		if err := WriteByteArrayField(ctx, "variableParameters", m.GetVariableParameters(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'variableParameters' field")
		}

		if popErr := writeBuffer.PopContext("NLMSecurityResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMSecurityResponse")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMSecurityResponse) IsNLMSecurityResponse() {}

func (m *_NLMSecurityResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMSecurityResponse) deepCopy() *_NLMSecurityResponse {
	if m == nil {
		return nil
	}
	_NLMSecurityResponseCopy := &_NLMSecurityResponse{
		m.NLMContract.(*_NLM).deepCopy(),
		m.ResponseCode,
		m.OriginalMessageId,
		m.OriginalTimestamp,
		utils.DeepCopySlice[byte, byte](m.VariableParameters),
	}
	m.NLMContract.(*_NLM)._SubType = m
	return _NLMSecurityResponseCopy
}

func (m *_NLMSecurityResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
