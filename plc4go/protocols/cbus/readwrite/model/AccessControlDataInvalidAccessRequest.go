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

// AccessControlDataInvalidAccessRequest is the corresponding interface of AccessControlDataInvalidAccessRequest
type AccessControlDataInvalidAccessRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AccessControlData
	// GetAccessControlDirection returns AccessControlDirection (property field)
	GetAccessControlDirection() AccessControlDirection
	// GetData returns Data (property field)
	GetData() []byte
	// IsAccessControlDataInvalidAccessRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataInvalidAccessRequest()
	// CreateBuilder creates a AccessControlDataInvalidAccessRequestBuilder
	CreateAccessControlDataInvalidAccessRequestBuilder() AccessControlDataInvalidAccessRequestBuilder
}

// _AccessControlDataInvalidAccessRequest is the data-structure of this message
type _AccessControlDataInvalidAccessRequest struct {
	AccessControlDataContract
	AccessControlDirection AccessControlDirection
	Data                   []byte
}

var _ AccessControlDataInvalidAccessRequest = (*_AccessControlDataInvalidAccessRequest)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataInvalidAccessRequest)(nil)

// NewAccessControlDataInvalidAccessRequest factory function for _AccessControlDataInvalidAccessRequest
func NewAccessControlDataInvalidAccessRequest(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte, accessControlDirection AccessControlDirection, data []byte) *_AccessControlDataInvalidAccessRequest {
	_result := &_AccessControlDataInvalidAccessRequest{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
		AccessControlDirection:    accessControlDirection,
		Data:                      data,
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AccessControlDataInvalidAccessRequestBuilder is a builder for AccessControlDataInvalidAccessRequest
type AccessControlDataInvalidAccessRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(accessControlDirection AccessControlDirection, data []byte) AccessControlDataInvalidAccessRequestBuilder
	// WithAccessControlDirection adds AccessControlDirection (property field)
	WithAccessControlDirection(AccessControlDirection) AccessControlDataInvalidAccessRequestBuilder
	// WithData adds Data (property field)
	WithData(...byte) AccessControlDataInvalidAccessRequestBuilder
	// Build builds the AccessControlDataInvalidAccessRequest or returns an error if something is wrong
	Build() (AccessControlDataInvalidAccessRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AccessControlDataInvalidAccessRequest
}

// NewAccessControlDataInvalidAccessRequestBuilder() creates a AccessControlDataInvalidAccessRequestBuilder
func NewAccessControlDataInvalidAccessRequestBuilder() AccessControlDataInvalidAccessRequestBuilder {
	return &_AccessControlDataInvalidAccessRequestBuilder{_AccessControlDataInvalidAccessRequest: new(_AccessControlDataInvalidAccessRequest)}
}

type _AccessControlDataInvalidAccessRequestBuilder struct {
	*_AccessControlDataInvalidAccessRequest

	err *utils.MultiError
}

var _ (AccessControlDataInvalidAccessRequestBuilder) = (*_AccessControlDataInvalidAccessRequestBuilder)(nil)

func (m *_AccessControlDataInvalidAccessRequestBuilder) WithMandatoryFields(accessControlDirection AccessControlDirection, data []byte) AccessControlDataInvalidAccessRequestBuilder {
	return m.WithAccessControlDirection(accessControlDirection).WithData(data...)
}

func (m *_AccessControlDataInvalidAccessRequestBuilder) WithAccessControlDirection(accessControlDirection AccessControlDirection) AccessControlDataInvalidAccessRequestBuilder {
	m.AccessControlDirection = accessControlDirection
	return m
}

func (m *_AccessControlDataInvalidAccessRequestBuilder) WithData(data ...byte) AccessControlDataInvalidAccessRequestBuilder {
	m.Data = data
	return m
}

func (m *_AccessControlDataInvalidAccessRequestBuilder) Build() (AccessControlDataInvalidAccessRequest, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AccessControlDataInvalidAccessRequest.deepCopy(), nil
}

func (m *_AccessControlDataInvalidAccessRequestBuilder) MustBuild() AccessControlDataInvalidAccessRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AccessControlDataInvalidAccessRequestBuilder) DeepCopy() any {
	return m.CreateAccessControlDataInvalidAccessRequestBuilder()
}

// CreateAccessControlDataInvalidAccessRequestBuilder creates a AccessControlDataInvalidAccessRequestBuilder
func (m *_AccessControlDataInvalidAccessRequest) CreateAccessControlDataInvalidAccessRequestBuilder() AccessControlDataInvalidAccessRequestBuilder {
	if m == nil {
		return NewAccessControlDataInvalidAccessRequestBuilder()
	}
	return &_AccessControlDataInvalidAccessRequestBuilder{_AccessControlDataInvalidAccessRequest: m.deepCopy()}
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

func (m *_AccessControlDataInvalidAccessRequest) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AccessControlDataInvalidAccessRequest) GetAccessControlDirection() AccessControlDirection {
	return m.AccessControlDirection
}

func (m *_AccessControlDataInvalidAccessRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAccessControlDataInvalidAccessRequest(structType any) AccessControlDataInvalidAccessRequest {
	if casted, ok := structType.(AccessControlDataInvalidAccessRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataInvalidAccessRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataInvalidAccessRequest) GetTypeName() string {
	return "AccessControlDataInvalidAccessRequest"
}

func (m *_AccessControlDataInvalidAccessRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	// Simple field (accessControlDirection)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AccessControlDataInvalidAccessRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataInvalidAccessRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData, commandTypeContainer AccessControlCommandTypeContainer) (__accessControlDataInvalidAccessRequest AccessControlDataInvalidAccessRequest, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataInvalidAccessRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataInvalidAccessRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessControlDirection, err := ReadEnumField[AccessControlDirection](ctx, "accessControlDirection", "AccessControlDirection", ReadEnum(AccessControlDirectionByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessControlDirection' field"))
	}
	m.AccessControlDirection = accessControlDirection

	data, err := readBuffer.ReadByteArray("data", int(int32(commandTypeContainer.NumBytes())-int32(int32(3))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AccessControlDataInvalidAccessRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataInvalidAccessRequest")
	}

	return m, nil
}

func (m *_AccessControlDataInvalidAccessRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataInvalidAccessRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataInvalidAccessRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataInvalidAccessRequest")
		}

		if err := WriteSimpleEnumField[AccessControlDirection](ctx, "accessControlDirection", "AccessControlDirection", m.GetAccessControlDirection(), WriteEnum[AccessControlDirection, uint8](AccessControlDirection.GetValue, AccessControlDirection.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'accessControlDirection' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataInvalidAccessRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataInvalidAccessRequest")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataInvalidAccessRequest) IsAccessControlDataInvalidAccessRequest() {}

func (m *_AccessControlDataInvalidAccessRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AccessControlDataInvalidAccessRequest) deepCopy() *_AccessControlDataInvalidAccessRequest {
	if m == nil {
		return nil
	}
	_AccessControlDataInvalidAccessRequestCopy := &_AccessControlDataInvalidAccessRequest{
		m.AccessControlDataContract.(*_AccessControlData).deepCopy(),
		m.AccessControlDirection,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.AccessControlDataContract.(*_AccessControlData)._SubType = m
	return _AccessControlDataInvalidAccessRequestCopy
}

func (m *_AccessControlDataInvalidAccessRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
