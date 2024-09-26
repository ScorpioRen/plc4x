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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// GetAttributeListRequest is the corresponding interface of GetAttributeListRequest
type GetAttributeListRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// IsGetAttributeListRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetAttributeListRequest()
	// CreateBuilder creates a GetAttributeListRequestBuilder
	CreateGetAttributeListRequestBuilder() GetAttributeListRequestBuilder
}

// _GetAttributeListRequest is the data-structure of this message
type _GetAttributeListRequest struct {
	CipServiceContract
}

var _ GetAttributeListRequest = (*_GetAttributeListRequest)(nil)
var _ CipServiceRequirements = (*_GetAttributeListRequest)(nil)

// NewGetAttributeListRequest factory function for _GetAttributeListRequest
func NewGetAttributeListRequest(serviceLen uint16) *_GetAttributeListRequest {
	_result := &_GetAttributeListRequest{
		CipServiceContract: NewCipService(serviceLen),
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GetAttributeListRequestBuilder is a builder for GetAttributeListRequest
type GetAttributeListRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() GetAttributeListRequestBuilder
	// Build builds the GetAttributeListRequest or returns an error if something is wrong
	Build() (GetAttributeListRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GetAttributeListRequest
}

// NewGetAttributeListRequestBuilder() creates a GetAttributeListRequestBuilder
func NewGetAttributeListRequestBuilder() GetAttributeListRequestBuilder {
	return &_GetAttributeListRequestBuilder{_GetAttributeListRequest: new(_GetAttributeListRequest)}
}

type _GetAttributeListRequestBuilder struct {
	*_GetAttributeListRequest

	err *utils.MultiError
}

var _ (GetAttributeListRequestBuilder) = (*_GetAttributeListRequestBuilder)(nil)

func (m *_GetAttributeListRequestBuilder) WithMandatoryFields() GetAttributeListRequestBuilder {
	return m
}

func (m *_GetAttributeListRequestBuilder) Build() (GetAttributeListRequest, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._GetAttributeListRequest.deepCopy(), nil
}

func (m *_GetAttributeListRequestBuilder) MustBuild() GetAttributeListRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_GetAttributeListRequestBuilder) DeepCopy() any {
	return m.CreateGetAttributeListRequestBuilder()
}

// CreateGetAttributeListRequestBuilder creates a GetAttributeListRequestBuilder
func (m *_GetAttributeListRequest) CreateGetAttributeListRequestBuilder() GetAttributeListRequestBuilder {
	if m == nil {
		return NewGetAttributeListRequestBuilder()
	}
	return &_GetAttributeListRequestBuilder{_GetAttributeListRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeListRequest) GetService() uint8 {
	return 0x03
}

func (m *_GetAttributeListRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeListRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeListRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// Deprecated: use the interface for direct cast
func CastGetAttributeListRequest(structType any) GetAttributeListRequest {
	if casted, ok := structType.(GetAttributeListRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeListRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeListRequest) GetTypeName() string {
	return "GetAttributeListRequest"
}

func (m *_GetAttributeListRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_GetAttributeListRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeListRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__getAttributeListRequest GetAttributeListRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeListRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeListRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GetAttributeListRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeListRequest")
	}

	return m, nil
}

func (m *_GetAttributeListRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeListRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeListRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeListRequest")
		}

		if popErr := writeBuffer.PopContext("GetAttributeListRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeListRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeListRequest) IsGetAttributeListRequest() {}

func (m *_GetAttributeListRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GetAttributeListRequest) deepCopy() *_GetAttributeListRequest {
	if m == nil {
		return nil
	}
	_GetAttributeListRequestCopy := &_GetAttributeListRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _GetAttributeListRequestCopy
}

func (m *_GetAttributeListRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
