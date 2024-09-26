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

// S7PayloadUserDataItemCyclicServicesUnsubscribeRequest is the corresponding interface of S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
type S7PayloadUserDataItemCyclicServicesUnsubscribeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetFunction returns Function (property field)
	GetFunction() uint8
	// GetJobId returns JobId (property field)
	GetJobId() uint8
	// IsS7PayloadUserDataItemCyclicServicesUnsubscribeRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCyclicServicesUnsubscribeRequest()
	// CreateBuilder creates a S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
	CreateS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder() S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
}

// _S7PayloadUserDataItemCyclicServicesUnsubscribeRequest is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesUnsubscribeRequest struct {
	S7PayloadUserDataItemContract
	Function uint8
	JobId    uint8
}

var _ S7PayloadUserDataItemCyclicServicesUnsubscribeRequest = (*_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest)(nil)

// NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequest factory function for _S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
func NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, function uint8, jobId uint8) *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest {
	_result := &_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Function:                      function,
		JobId:                         jobId,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder is a builder for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
type S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(function uint8, jobId uint8) S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
	// WithFunction adds Function (property field)
	WithFunction(uint8) S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
	// WithJobId adds JobId (property field)
	WithJobId(uint8) S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
	// Build builds the S7PayloadUserDataItemCyclicServicesUnsubscribeRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCyclicServicesUnsubscribeRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
}

// NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder() creates a S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
func NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder() S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder {
	return &_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder{_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest: new(_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest)}
}

type _S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder struct {
	*_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) = (*_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder)(nil)

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) WithMandatoryFields(function uint8, jobId uint8) S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder {
	return m.WithFunction(function).WithJobId(jobId)
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) WithFunction(function uint8) S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder {
	m.Function = function
	return m
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) WithJobId(jobId uint8) S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder {
	m.JobId = jobId
	return m
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) Build() (S7PayloadUserDataItemCyclicServicesUnsubscribeRequest, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._S7PayloadUserDataItemCyclicServicesUnsubscribeRequest.deepCopy(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) MustBuild() S7PayloadUserDataItemCyclicServicesUnsubscribeRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder) DeepCopy() any {
	return m.CreateS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder()
}

// CreateS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder creates a S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder
func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) CreateS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder() S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder {
	if m == nil {
		return NewS7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder()
	}
	return &_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestBuilder{_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetCpuSubfunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetFunction() uint8 {
	return m.Function
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetJobId() uint8 {
	return m.JobId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesUnsubscribeRequest(structType any) S7PayloadUserDataItemCyclicServicesUnsubscribeRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesUnsubscribeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesUnsubscribeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (function)
	lengthInBits += 8

	// Simple field (jobId)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCyclicServicesUnsubscribeRequest S7PayloadUserDataItemCyclicServicesUnsubscribeRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	function, err := ReadSimpleField(ctx, "function", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'function' field"))
	}
	m.Function = function

	jobId, err := ReadSimpleField(ctx, "jobId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'jobId' field"))
	}
	m.JobId = jobId

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "function", m.GetFunction(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'function' field")
		}

		if err := WriteSimpleField[uint8](ctx, "jobId", m.GetJobId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'jobId' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesUnsubscribeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesUnsubscribeRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) IsS7PayloadUserDataItemCyclicServicesUnsubscribeRequest() {
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) deepCopy() *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCyclicServicesUnsubscribeRequestCopy := &_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.Function,
		m.JobId,
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCyclicServicesUnsubscribeRequestCopy
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
