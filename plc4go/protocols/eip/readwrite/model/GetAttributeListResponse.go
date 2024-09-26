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

// GetAttributeListResponse is the corresponding interface of GetAttributeListResponse
type GetAttributeListResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// IsGetAttributeListResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetAttributeListResponse()
	// CreateBuilder creates a GetAttributeListResponseBuilder
	CreateGetAttributeListResponseBuilder() GetAttributeListResponseBuilder
}

// _GetAttributeListResponse is the data-structure of this message
type _GetAttributeListResponse struct {
	CipServiceContract
}

var _ GetAttributeListResponse = (*_GetAttributeListResponse)(nil)
var _ CipServiceRequirements = (*_GetAttributeListResponse)(nil)

// NewGetAttributeListResponse factory function for _GetAttributeListResponse
func NewGetAttributeListResponse(serviceLen uint16) *_GetAttributeListResponse {
	_result := &_GetAttributeListResponse{
		CipServiceContract: NewCipService(serviceLen),
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GetAttributeListResponseBuilder is a builder for GetAttributeListResponse
type GetAttributeListResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() GetAttributeListResponseBuilder
	// Build builds the GetAttributeListResponse or returns an error if something is wrong
	Build() (GetAttributeListResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GetAttributeListResponse
}

// NewGetAttributeListResponseBuilder() creates a GetAttributeListResponseBuilder
func NewGetAttributeListResponseBuilder() GetAttributeListResponseBuilder {
	return &_GetAttributeListResponseBuilder{_GetAttributeListResponse: new(_GetAttributeListResponse)}
}

type _GetAttributeListResponseBuilder struct {
	*_GetAttributeListResponse

	err *utils.MultiError
}

var _ (GetAttributeListResponseBuilder) = (*_GetAttributeListResponseBuilder)(nil)

func (m *_GetAttributeListResponseBuilder) WithMandatoryFields() GetAttributeListResponseBuilder {
	return m
}

func (m *_GetAttributeListResponseBuilder) Build() (GetAttributeListResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._GetAttributeListResponse.deepCopy(), nil
}

func (m *_GetAttributeListResponseBuilder) MustBuild() GetAttributeListResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_GetAttributeListResponseBuilder) DeepCopy() any {
	return m.CreateGetAttributeListResponseBuilder()
}

// CreateGetAttributeListResponseBuilder creates a GetAttributeListResponseBuilder
func (m *_GetAttributeListResponse) CreateGetAttributeListResponseBuilder() GetAttributeListResponseBuilder {
	if m == nil {
		return NewGetAttributeListResponseBuilder()
	}
	return &_GetAttributeListResponseBuilder{_GetAttributeListResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeListResponse) GetService() uint8 {
	return 0x03
}

func (m *_GetAttributeListResponse) GetResponse() bool {
	return bool(true)
}

func (m *_GetAttributeListResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeListResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// Deprecated: use the interface for direct cast
func CastGetAttributeListResponse(structType any) GetAttributeListResponse {
	if casted, ok := structType.(GetAttributeListResponse); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeListResponse); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeListResponse) GetTypeName() string {
	return "GetAttributeListResponse"
}

func (m *_GetAttributeListResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_GetAttributeListResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeListResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__getAttributeListResponse GetAttributeListResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeListResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeListResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GetAttributeListResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeListResponse")
	}

	return m, nil
}

func (m *_GetAttributeListResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeListResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeListResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeListResponse")
		}

		if popErr := writeBuffer.PopContext("GetAttributeListResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeListResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeListResponse) IsGetAttributeListResponse() {}

func (m *_GetAttributeListResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GetAttributeListResponse) deepCopy() *_GetAttributeListResponse {
	if m == nil {
		return nil
	}
	_GetAttributeListResponseCopy := &_GetAttributeListResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _GetAttributeListResponseCopy
}

func (m *_GetAttributeListResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
