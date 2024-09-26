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

// CloseSecureChannelRequest is the corresponding interface of CloseSecureChannelRequest
type CloseSecureChannelRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// IsCloseSecureChannelRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCloseSecureChannelRequest()
	// CreateBuilder creates a CloseSecureChannelRequestBuilder
	CreateCloseSecureChannelRequestBuilder() CloseSecureChannelRequestBuilder
}

// _CloseSecureChannelRequest is the data-structure of this message
type _CloseSecureChannelRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader ExtensionObjectDefinition
}

var _ CloseSecureChannelRequest = (*_CloseSecureChannelRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSecureChannelRequest)(nil)

// NewCloseSecureChannelRequest factory function for _CloseSecureChannelRequest
func NewCloseSecureChannelRequest(requestHeader ExtensionObjectDefinition) *_CloseSecureChannelRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for CloseSecureChannelRequest must not be nil")
	}
	_result := &_CloseSecureChannelRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CloseSecureChannelRequestBuilder is a builder for CloseSecureChannelRequest
type CloseSecureChannelRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition) CloseSecureChannelRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) CloseSecureChannelRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CloseSecureChannelRequestBuilder
	// Build builds the CloseSecureChannelRequest or returns an error if something is wrong
	Build() (CloseSecureChannelRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CloseSecureChannelRequest
}

// NewCloseSecureChannelRequestBuilder() creates a CloseSecureChannelRequestBuilder
func NewCloseSecureChannelRequestBuilder() CloseSecureChannelRequestBuilder {
	return &_CloseSecureChannelRequestBuilder{_CloseSecureChannelRequest: new(_CloseSecureChannelRequest)}
}

type _CloseSecureChannelRequestBuilder struct {
	*_CloseSecureChannelRequest

	err *utils.MultiError
}

var _ (CloseSecureChannelRequestBuilder) = (*_CloseSecureChannelRequestBuilder)(nil)

func (m *_CloseSecureChannelRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition) CloseSecureChannelRequestBuilder {
	return m.WithRequestHeader(requestHeader)
}

func (m *_CloseSecureChannelRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) CloseSecureChannelRequestBuilder {
	m.RequestHeader = requestHeader
	return m
}

func (m *_CloseSecureChannelRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CloseSecureChannelRequestBuilder {
	builder := builderSupplier(m.RequestHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	m.RequestHeader, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return m
}

func (m *_CloseSecureChannelRequestBuilder) Build() (CloseSecureChannelRequest, error) {
	if m.RequestHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._CloseSecureChannelRequest.deepCopy(), nil
}

func (m *_CloseSecureChannelRequestBuilder) MustBuild() CloseSecureChannelRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_CloseSecureChannelRequestBuilder) DeepCopy() any {
	return m.CreateCloseSecureChannelRequestBuilder()
}

// CreateCloseSecureChannelRequestBuilder creates a CloseSecureChannelRequestBuilder
func (m *_CloseSecureChannelRequest) CreateCloseSecureChannelRequestBuilder() CloseSecureChannelRequestBuilder {
	if m == nil {
		return NewCloseSecureChannelRequestBuilder()
	}
	return &_CloseSecureChannelRequestBuilder{_CloseSecureChannelRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSecureChannelRequest) GetIdentifier() string {
	return "452"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSecureChannelRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSecureChannelRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCloseSecureChannelRequest(structType any) CloseSecureChannelRequest {
	if casted, ok := structType.(CloseSecureChannelRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSecureChannelRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSecureChannelRequest) GetTypeName() string {
	return "CloseSecureChannelRequest"
}

func (m *_CloseSecureChannelRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSecureChannelRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSecureChannelRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__closeSecureChannelRequest CloseSecureChannelRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSecureChannelRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSecureChannelRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	if closeErr := readBuffer.CloseContext("CloseSecureChannelRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSecureChannelRequest")
	}

	return m, nil
}

func (m *_CloseSecureChannelRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSecureChannelRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSecureChannelRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSecureChannelRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSecureChannelRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSecureChannelRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSecureChannelRequest) IsCloseSecureChannelRequest() {}

func (m *_CloseSecureChannelRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CloseSecureChannelRequest) deepCopy() *_CloseSecureChannelRequest {
	if m == nil {
		return nil
	}
	_CloseSecureChannelRequestCopy := &_CloseSecureChannelRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CloseSecureChannelRequestCopy
}

func (m *_CloseSecureChannelRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
