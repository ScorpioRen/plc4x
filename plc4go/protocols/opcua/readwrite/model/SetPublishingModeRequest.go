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

// SetPublishingModeRequest is the corresponding interface of SetPublishingModeRequest
type SetPublishingModeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetPublishingEnabled returns PublishingEnabled (property field)
	GetPublishingEnabled() bool
	// GetNoOfSubscriptionIds returns NoOfSubscriptionIds (property field)
	GetNoOfSubscriptionIds() int32
	// GetSubscriptionIds returns SubscriptionIds (property field)
	GetSubscriptionIds() []uint32
	// IsSetPublishingModeRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSetPublishingModeRequest()
	// CreateBuilder creates a SetPublishingModeRequestBuilder
	CreateSetPublishingModeRequestBuilder() SetPublishingModeRequestBuilder
}

// _SetPublishingModeRequest is the data-structure of this message
type _SetPublishingModeRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader       ExtensionObjectDefinition
	PublishingEnabled   bool
	NoOfSubscriptionIds int32
	SubscriptionIds     []uint32
	// Reserved Fields
	reservedField0 *uint8
}

var _ SetPublishingModeRequest = (*_SetPublishingModeRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SetPublishingModeRequest)(nil)

// NewSetPublishingModeRequest factory function for _SetPublishingModeRequest
func NewSetPublishingModeRequest(requestHeader ExtensionObjectDefinition, publishingEnabled bool, noOfSubscriptionIds int32, subscriptionIds []uint32) *_SetPublishingModeRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for SetPublishingModeRequest must not be nil")
	}
	_result := &_SetPublishingModeRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		PublishingEnabled:                 publishingEnabled,
		NoOfSubscriptionIds:               noOfSubscriptionIds,
		SubscriptionIds:                   subscriptionIds,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SetPublishingModeRequestBuilder is a builder for SetPublishingModeRequest
type SetPublishingModeRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, publishingEnabled bool, noOfSubscriptionIds int32, subscriptionIds []uint32) SetPublishingModeRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) SetPublishingModeRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) SetPublishingModeRequestBuilder
	// WithPublishingEnabled adds PublishingEnabled (property field)
	WithPublishingEnabled(bool) SetPublishingModeRequestBuilder
	// WithNoOfSubscriptionIds adds NoOfSubscriptionIds (property field)
	WithNoOfSubscriptionIds(int32) SetPublishingModeRequestBuilder
	// WithSubscriptionIds adds SubscriptionIds (property field)
	WithSubscriptionIds(...uint32) SetPublishingModeRequestBuilder
	// Build builds the SetPublishingModeRequest or returns an error if something is wrong
	Build() (SetPublishingModeRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SetPublishingModeRequest
}

// NewSetPublishingModeRequestBuilder() creates a SetPublishingModeRequestBuilder
func NewSetPublishingModeRequestBuilder() SetPublishingModeRequestBuilder {
	return &_SetPublishingModeRequestBuilder{_SetPublishingModeRequest: new(_SetPublishingModeRequest)}
}

type _SetPublishingModeRequestBuilder struct {
	*_SetPublishingModeRequest

	err *utils.MultiError
}

var _ (SetPublishingModeRequestBuilder) = (*_SetPublishingModeRequestBuilder)(nil)

func (m *_SetPublishingModeRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, publishingEnabled bool, noOfSubscriptionIds int32, subscriptionIds []uint32) SetPublishingModeRequestBuilder {
	return m.WithRequestHeader(requestHeader).WithPublishingEnabled(publishingEnabled).WithNoOfSubscriptionIds(noOfSubscriptionIds).WithSubscriptionIds(subscriptionIds...)
}

func (m *_SetPublishingModeRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) SetPublishingModeRequestBuilder {
	m.RequestHeader = requestHeader
	return m
}

func (m *_SetPublishingModeRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) SetPublishingModeRequestBuilder {
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

func (m *_SetPublishingModeRequestBuilder) WithPublishingEnabled(publishingEnabled bool) SetPublishingModeRequestBuilder {
	m.PublishingEnabled = publishingEnabled
	return m
}

func (m *_SetPublishingModeRequestBuilder) WithNoOfSubscriptionIds(noOfSubscriptionIds int32) SetPublishingModeRequestBuilder {
	m.NoOfSubscriptionIds = noOfSubscriptionIds
	return m
}

func (m *_SetPublishingModeRequestBuilder) WithSubscriptionIds(subscriptionIds ...uint32) SetPublishingModeRequestBuilder {
	m.SubscriptionIds = subscriptionIds
	return m
}

func (m *_SetPublishingModeRequestBuilder) Build() (SetPublishingModeRequest, error) {
	if m.RequestHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SetPublishingModeRequest.deepCopy(), nil
}

func (m *_SetPublishingModeRequestBuilder) MustBuild() SetPublishingModeRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SetPublishingModeRequestBuilder) DeepCopy() any {
	return m.CreateSetPublishingModeRequestBuilder()
}

// CreateSetPublishingModeRequestBuilder creates a SetPublishingModeRequestBuilder
func (m *_SetPublishingModeRequest) CreateSetPublishingModeRequestBuilder() SetPublishingModeRequestBuilder {
	if m == nil {
		return NewSetPublishingModeRequestBuilder()
	}
	return &_SetPublishingModeRequestBuilder{_SetPublishingModeRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetPublishingModeRequest) GetIdentifier() string {
	return "799"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetPublishingModeRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SetPublishingModeRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_SetPublishingModeRequest) GetPublishingEnabled() bool {
	return m.PublishingEnabled
}

func (m *_SetPublishingModeRequest) GetNoOfSubscriptionIds() int32 {
	return m.NoOfSubscriptionIds
}

func (m *_SetPublishingModeRequest) GetSubscriptionIds() []uint32 {
	return m.SubscriptionIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSetPublishingModeRequest(structType any) SetPublishingModeRequest {
	if casted, ok := structType.(SetPublishingModeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetPublishingModeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetPublishingModeRequest) GetTypeName() string {
	return "SetPublishingModeRequest"
}

func (m *_SetPublishingModeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (publishingEnabled)
	lengthInBits += 1

	// Simple field (noOfSubscriptionIds)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionIds) > 0 {
		lengthInBits += 32 * uint16(len(m.SubscriptionIds))
	}

	return lengthInBits
}

func (m *_SetPublishingModeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SetPublishingModeRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__setPublishingModeRequest SetPublishingModeRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetPublishingModeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetPublishingModeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	publishingEnabled, err := ReadSimpleField(ctx, "publishingEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingEnabled' field"))
	}
	m.PublishingEnabled = publishingEnabled

	noOfSubscriptionIds, err := ReadSimpleField(ctx, "noOfSubscriptionIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSubscriptionIds' field"))
	}
	m.NoOfSubscriptionIds = noOfSubscriptionIds

	subscriptionIds, err := ReadCountArrayField[uint32](ctx, "subscriptionIds", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfSubscriptionIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionIds' field"))
	}
	m.SubscriptionIds = subscriptionIds

	if closeErr := readBuffer.CloseContext("SetPublishingModeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetPublishingModeRequest")
	}

	return m, nil
}

func (m *_SetPublishingModeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetPublishingModeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetPublishingModeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetPublishingModeRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "publishingEnabled", m.GetPublishingEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingEnabled' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfSubscriptionIds", m.GetNoOfSubscriptionIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSubscriptionIds' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "subscriptionIds", m.GetSubscriptionIds(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionIds' field")
		}

		if popErr := writeBuffer.PopContext("SetPublishingModeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetPublishingModeRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetPublishingModeRequest) IsSetPublishingModeRequest() {}

func (m *_SetPublishingModeRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SetPublishingModeRequest) deepCopy() *_SetPublishingModeRequest {
	if m == nil {
		return nil
	}
	_SetPublishingModeRequestCopy := &_SetPublishingModeRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.PublishingEnabled,
		m.NoOfSubscriptionIds,
		utils.DeepCopySlice[uint32, uint32](m.SubscriptionIds),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SetPublishingModeRequestCopy
}

func (m *_SetPublishingModeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
