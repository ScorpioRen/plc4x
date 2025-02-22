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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUUnconfirmedRequest is the corresponding interface of APDUUnconfirmedRequest
type APDUUnconfirmedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() BACnetUnconfirmedServiceRequest
	// IsAPDUUnconfirmedRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUUnconfirmedRequest()
	// CreateBuilder creates a APDUUnconfirmedRequestBuilder
	CreateAPDUUnconfirmedRequestBuilder() APDUUnconfirmedRequestBuilder
}

// _APDUUnconfirmedRequest is the data-structure of this message
type _APDUUnconfirmedRequest struct {
	APDUContract
	ServiceRequest BACnetUnconfirmedServiceRequest
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUUnconfirmedRequest = (*_APDUUnconfirmedRequest)(nil)
var _ APDURequirements = (*_APDUUnconfirmedRequest)(nil)

// NewAPDUUnconfirmedRequest factory function for _APDUUnconfirmedRequest
func NewAPDUUnconfirmedRequest(serviceRequest BACnetUnconfirmedServiceRequest, apduLength uint16) *_APDUUnconfirmedRequest {
	if serviceRequest == nil {
		panic("serviceRequest of type BACnetUnconfirmedServiceRequest for APDUUnconfirmedRequest must not be nil")
	}
	_result := &_APDUUnconfirmedRequest{
		APDUContract:   NewAPDU(apduLength),
		ServiceRequest: serviceRequest,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// APDUUnconfirmedRequestBuilder is a builder for APDUUnconfirmedRequest
type APDUUnconfirmedRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(serviceRequest BACnetUnconfirmedServiceRequest) APDUUnconfirmedRequestBuilder
	// WithServiceRequest adds ServiceRequest (property field)
	WithServiceRequest(BACnetUnconfirmedServiceRequest) APDUUnconfirmedRequestBuilder
	// WithServiceRequestBuilder adds ServiceRequest (property field) which is build by the builder
	WithServiceRequestBuilder(func(BACnetUnconfirmedServiceRequestBuilder) BACnetUnconfirmedServiceRequestBuilder) APDUUnconfirmedRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() APDUBuilder
	// Build builds the APDUUnconfirmedRequest or returns an error if something is wrong
	Build() (APDUUnconfirmedRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUUnconfirmedRequest
}

// NewAPDUUnconfirmedRequestBuilder() creates a APDUUnconfirmedRequestBuilder
func NewAPDUUnconfirmedRequestBuilder() APDUUnconfirmedRequestBuilder {
	return &_APDUUnconfirmedRequestBuilder{_APDUUnconfirmedRequest: new(_APDUUnconfirmedRequest)}
}

type _APDUUnconfirmedRequestBuilder struct {
	*_APDUUnconfirmedRequest

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDUUnconfirmedRequestBuilder) = (*_APDUUnconfirmedRequestBuilder)(nil)

func (b *_APDUUnconfirmedRequestBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
	contract.(*_APDU)._SubType = b._APDUUnconfirmedRequest
}

func (b *_APDUUnconfirmedRequestBuilder) WithMandatoryFields(serviceRequest BACnetUnconfirmedServiceRequest) APDUUnconfirmedRequestBuilder {
	return b.WithServiceRequest(serviceRequest)
}

func (b *_APDUUnconfirmedRequestBuilder) WithServiceRequest(serviceRequest BACnetUnconfirmedServiceRequest) APDUUnconfirmedRequestBuilder {
	b.ServiceRequest = serviceRequest
	return b
}

func (b *_APDUUnconfirmedRequestBuilder) WithServiceRequestBuilder(builderSupplier func(BACnetUnconfirmedServiceRequestBuilder) BACnetUnconfirmedServiceRequestBuilder) APDUUnconfirmedRequestBuilder {
	builder := builderSupplier(b.ServiceRequest.CreateBACnetUnconfirmedServiceRequestBuilder())
	var err error
	b.ServiceRequest, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetUnconfirmedServiceRequestBuilder failed"))
	}
	return b
}

func (b *_APDUUnconfirmedRequestBuilder) Build() (APDUUnconfirmedRequest, error) {
	if b.ServiceRequest == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'serviceRequest' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUUnconfirmedRequest.deepCopy(), nil
}

func (b *_APDUUnconfirmedRequestBuilder) MustBuild() APDUUnconfirmedRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_APDUUnconfirmedRequestBuilder) Done() APDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAPDUBuilder().(*_APDUBuilder)
	}
	return b.parentBuilder
}

func (b *_APDUUnconfirmedRequestBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDUUnconfirmedRequestBuilder) DeepCopy() any {
	_copy := b.CreateAPDUUnconfirmedRequestBuilder().(*_APDUUnconfirmedRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDUUnconfirmedRequestBuilder creates a APDUUnconfirmedRequestBuilder
func (b *_APDUUnconfirmedRequest) CreateAPDUUnconfirmedRequestBuilder() APDUUnconfirmedRequestBuilder {
	if b == nil {
		return NewAPDUUnconfirmedRequestBuilder()
	}
	return &_APDUUnconfirmedRequestBuilder{_APDUUnconfirmedRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUUnconfirmedRequest) GetApduType() ApduType {
	return ApduType_UNCONFIRMED_REQUEST_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUUnconfirmedRequest) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUUnconfirmedRequest) GetServiceRequest() BACnetUnconfirmedServiceRequest {
	return m.ServiceRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUUnconfirmedRequest(structType any) APDUUnconfirmedRequest {
	if casted, ok := structType.(APDUUnconfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*APDUUnconfirmedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_APDUUnconfirmedRequest) GetTypeName() string {
	return "APDUUnconfirmedRequest"
}

func (m *_APDUUnconfirmedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (serviceRequest)
	lengthInBits += m.ServiceRequest.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUUnconfirmedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUUnconfirmedRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUUnconfirmedRequest APDUUnconfirmedRequest, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUUnconfirmedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUUnconfirmedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	serviceRequest, err := ReadSimpleField[BACnetUnconfirmedServiceRequest](ctx, "serviceRequest", ReadComplex[BACnetUnconfirmedServiceRequest](BACnetUnconfirmedServiceRequestParseWithBufferProducer[BACnetUnconfirmedServiceRequest]((uint16)(uint16(apduLength)-uint16(uint16(1)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceRequest' field"))
	}
	m.ServiceRequest = serviceRequest

	if closeErr := readBuffer.CloseContext("APDUUnconfirmedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUUnconfirmedRequest")
	}

	return m, nil
}

func (m *_APDUUnconfirmedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUUnconfirmedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnconfirmedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUUnconfirmedRequest")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[BACnetUnconfirmedServiceRequest](ctx, "serviceRequest", m.GetServiceRequest(), WriteComplex[BACnetUnconfirmedServiceRequest](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceRequest' field")
		}

		if popErr := writeBuffer.PopContext("APDUUnconfirmedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUUnconfirmedRequest")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUUnconfirmedRequest) IsAPDUUnconfirmedRequest() {}

func (m *_APDUUnconfirmedRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUUnconfirmedRequest) deepCopy() *_APDUUnconfirmedRequest {
	if m == nil {
		return nil
	}
	_APDUUnconfirmedRequestCopy := &_APDUUnconfirmedRequest{
		m.APDUContract.(*_APDU).deepCopy(),
		utils.DeepCopy[BACnetUnconfirmedServiceRequest](m.ServiceRequest),
		m.reservedField0,
	}
	_APDUUnconfirmedRequestCopy.APDUContract.(*_APDU)._SubType = m
	return _APDUUnconfirmedRequestCopy
}

func (m *_APDUUnconfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
