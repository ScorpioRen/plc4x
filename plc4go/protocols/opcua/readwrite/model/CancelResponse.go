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

// CancelResponse is the corresponding interface of CancelResponse
type CancelResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetCancelCount returns CancelCount (property field)
	GetCancelCount() uint32
	// IsCancelResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCancelResponse()
	// CreateBuilder creates a CancelResponseBuilder
	CreateCancelResponseBuilder() CancelResponseBuilder
}

// _CancelResponse is the data-structure of this message
type _CancelResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
	CancelCount    uint32
}

var _ CancelResponse = (*_CancelResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CancelResponse)(nil)

// NewCancelResponse factory function for _CancelResponse
func NewCancelResponse(responseHeader ExtensionObjectDefinition, cancelCount uint32) *_CancelResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for CancelResponse must not be nil")
	}
	_result := &_CancelResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		CancelCount:                       cancelCount,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CancelResponseBuilder is a builder for CancelResponse
type CancelResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ExtensionObjectDefinition, cancelCount uint32) CancelResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ExtensionObjectDefinition) CancelResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CancelResponseBuilder
	// WithCancelCount adds CancelCount (property field)
	WithCancelCount(uint32) CancelResponseBuilder
	// Build builds the CancelResponse or returns an error if something is wrong
	Build() (CancelResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CancelResponse
}

// NewCancelResponseBuilder() creates a CancelResponseBuilder
func NewCancelResponseBuilder() CancelResponseBuilder {
	return &_CancelResponseBuilder{_CancelResponse: new(_CancelResponse)}
}

type _CancelResponseBuilder struct {
	*_CancelResponse

	err *utils.MultiError
}

var _ (CancelResponseBuilder) = (*_CancelResponseBuilder)(nil)

func (m *_CancelResponseBuilder) WithMandatoryFields(responseHeader ExtensionObjectDefinition, cancelCount uint32) CancelResponseBuilder {
	return m.WithResponseHeader(responseHeader).WithCancelCount(cancelCount)
}

func (m *_CancelResponseBuilder) WithResponseHeader(responseHeader ExtensionObjectDefinition) CancelResponseBuilder {
	m.ResponseHeader = responseHeader
	return m
}

func (m *_CancelResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CancelResponseBuilder {
	builder := builderSupplier(m.ResponseHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	m.ResponseHeader, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return m
}

func (m *_CancelResponseBuilder) WithCancelCount(cancelCount uint32) CancelResponseBuilder {
	m.CancelCount = cancelCount
	return m
}

func (m *_CancelResponseBuilder) Build() (CancelResponse, error) {
	if m.ResponseHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._CancelResponse.deepCopy(), nil
}

func (m *_CancelResponseBuilder) MustBuild() CancelResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_CancelResponseBuilder) DeepCopy() any {
	return m.CreateCancelResponseBuilder()
}

// CreateCancelResponseBuilder creates a CancelResponseBuilder
func (m *_CancelResponse) CreateCancelResponseBuilder() CancelResponseBuilder {
	if m == nil {
		return NewCancelResponseBuilder()
	}
	return &_CancelResponseBuilder{_CancelResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CancelResponse) GetIdentifier() string {
	return "482"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CancelResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CancelResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_CancelResponse) GetCancelCount() uint32 {
	return m.CancelCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCancelResponse(structType any) CancelResponse {
	if casted, ok := structType.(CancelResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CancelResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CancelResponse) GetTypeName() string {
	return "CancelResponse"
}

func (m *_CancelResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (cancelCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CancelResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CancelResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__cancelResponse CancelResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CancelResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CancelResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	cancelCount, err := ReadSimpleField(ctx, "cancelCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cancelCount' field"))
	}
	m.CancelCount = cancelCount

	if closeErr := readBuffer.CloseContext("CancelResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CancelResponse")
	}

	return m, nil
}

func (m *_CancelResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CancelResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CancelResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CancelResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cancelCount", m.GetCancelCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cancelCount' field")
		}

		if popErr := writeBuffer.PopContext("CancelResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CancelResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CancelResponse) IsCancelResponse() {}

func (m *_CancelResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CancelResponse) deepCopy() *_CancelResponse {
	if m == nil {
		return nil
	}
	_CancelResponseCopy := &_CancelResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ExtensionObjectDefinition),
		m.CancelCount,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CancelResponseCopy
}

func (m *_CancelResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
