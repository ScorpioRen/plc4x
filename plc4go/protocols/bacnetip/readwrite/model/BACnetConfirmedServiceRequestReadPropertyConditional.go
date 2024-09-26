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

// BACnetConfirmedServiceRequestReadPropertyConditional is the corresponding interface of BACnetConfirmedServiceRequestReadPropertyConditional
type BACnetConfirmedServiceRequestReadPropertyConditional interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
	// IsBACnetConfirmedServiceRequestReadPropertyConditional is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadPropertyConditional()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReadPropertyConditionalBuilder
	CreateBACnetConfirmedServiceRequestReadPropertyConditionalBuilder() BACnetConfirmedServiceRequestReadPropertyConditionalBuilder
}

// _BACnetConfirmedServiceRequestReadPropertyConditional is the data-structure of this message
type _BACnetConfirmedServiceRequestReadPropertyConditional struct {
	BACnetConfirmedServiceRequestContract
	BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestReadPropertyConditional = (*_BACnetConfirmedServiceRequestReadPropertyConditional)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestReadPropertyConditional)(nil)

// NewBACnetConfirmedServiceRequestReadPropertyConditional factory function for _BACnetConfirmedServiceRequestReadPropertyConditional
func NewBACnetConfirmedServiceRequestReadPropertyConditional(bytesOfRemovedService []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReadPropertyConditional {
	_result := &_BACnetConfirmedServiceRequestReadPropertyConditional{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		BytesOfRemovedService:                 bytesOfRemovedService,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReadPropertyConditionalBuilder is a builder for BACnetConfirmedServiceRequestReadPropertyConditional
type BACnetConfirmedServiceRequestReadPropertyConditionalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bytesOfRemovedService []byte) BACnetConfirmedServiceRequestReadPropertyConditionalBuilder
	// WithBytesOfRemovedService adds BytesOfRemovedService (property field)
	WithBytesOfRemovedService(...byte) BACnetConfirmedServiceRequestReadPropertyConditionalBuilder
	// Build builds the BACnetConfirmedServiceRequestReadPropertyConditional or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReadPropertyConditional, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReadPropertyConditional
}

// NewBACnetConfirmedServiceRequestReadPropertyConditionalBuilder() creates a BACnetConfirmedServiceRequestReadPropertyConditionalBuilder
func NewBACnetConfirmedServiceRequestReadPropertyConditionalBuilder() BACnetConfirmedServiceRequestReadPropertyConditionalBuilder {
	return &_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder{_BACnetConfirmedServiceRequestReadPropertyConditional: new(_BACnetConfirmedServiceRequestReadPropertyConditional)}
}

type _BACnetConfirmedServiceRequestReadPropertyConditionalBuilder struct {
	*_BACnetConfirmedServiceRequestReadPropertyConditional

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReadPropertyConditionalBuilder) = (*_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder) WithMandatoryFields(bytesOfRemovedService []byte) BACnetConfirmedServiceRequestReadPropertyConditionalBuilder {
	return m.WithBytesOfRemovedService(bytesOfRemovedService...)
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder) WithBytesOfRemovedService(bytesOfRemovedService ...byte) BACnetConfirmedServiceRequestReadPropertyConditionalBuilder {
	m.BytesOfRemovedService = bytesOfRemovedService
	return m
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder) Build() (BACnetConfirmedServiceRequestReadPropertyConditional, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestReadPropertyConditional.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder) MustBuild() BACnetConfirmedServiceRequestReadPropertyConditional {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestReadPropertyConditionalBuilder()
}

// CreateBACnetConfirmedServiceRequestReadPropertyConditionalBuilder creates a BACnetConfirmedServiceRequestReadPropertyConditionalBuilder
func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) CreateBACnetConfirmedServiceRequestReadPropertyConditionalBuilder() BACnetConfirmedServiceRequestReadPropertyConditionalBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestReadPropertyConditionalBuilder()
	}
	return &_BACnetConfirmedServiceRequestReadPropertyConditionalBuilder{_BACnetConfirmedServiceRequestReadPropertyConditional: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadPropertyConditional(structType any) BACnetConfirmedServiceRequestReadPropertyConditional {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadPropertyConditional); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadPropertyConditional); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadPropertyConditional"
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestReadPropertyConditional BACnetConfirmedServiceRequestReadPropertyConditional, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadPropertyConditional"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadPropertyConditional")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bytesOfRemovedService, err := readBuffer.ReadByteArray("bytesOfRemovedService", int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bytesOfRemovedService' field"))
	}
	m.BytesOfRemovedService = bytesOfRemovedService

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadPropertyConditional"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadPropertyConditional")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadPropertyConditional"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadPropertyConditional")
		}

		if err := WriteByteArrayField(ctx, "bytesOfRemovedService", m.GetBytesOfRemovedService(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadPropertyConditional"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadPropertyConditional")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) IsBACnetConfirmedServiceRequestReadPropertyConditional() {
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) deepCopy() *_BACnetConfirmedServiceRequestReadPropertyConditional {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReadPropertyConditionalCopy := &_BACnetConfirmedServiceRequestReadPropertyConditional{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.BytesOfRemovedService),
		m.ServiceRequestPayloadLength,
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestReadPropertyConditionalCopy
}

func (m *_BACnetConfirmedServiceRequestReadPropertyConditional) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
