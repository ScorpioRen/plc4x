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

// Constant values.
const MultipleServiceRequest_REQUESTPATHSIZE uint8 = 0x02
const MultipleServiceRequest_REQUESTPATH uint32 = 0x01240220

// MultipleServiceRequest is the corresponding interface of MultipleServiceRequest
type MultipleServiceRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetData returns Data (property field)
	GetData() Services
	// IsMultipleServiceRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMultipleServiceRequest()
	// CreateBuilder creates a MultipleServiceRequestBuilder
	CreateMultipleServiceRequestBuilder() MultipleServiceRequestBuilder
}

// _MultipleServiceRequest is the data-structure of this message
type _MultipleServiceRequest struct {
	CipServiceContract
	Data Services
}

var _ MultipleServiceRequest = (*_MultipleServiceRequest)(nil)
var _ CipServiceRequirements = (*_MultipleServiceRequest)(nil)

// NewMultipleServiceRequest factory function for _MultipleServiceRequest
func NewMultipleServiceRequest(data Services, serviceLen uint16) *_MultipleServiceRequest {
	if data == nil {
		panic("data of type Services for MultipleServiceRequest must not be nil")
	}
	_result := &_MultipleServiceRequest{
		CipServiceContract: NewCipService(serviceLen),
		Data:               data,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MultipleServiceRequestBuilder is a builder for MultipleServiceRequest
type MultipleServiceRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(data Services) MultipleServiceRequestBuilder
	// WithData adds Data (property field)
	WithData(Services) MultipleServiceRequestBuilder
	// WithDataBuilder adds Data (property field) which is build by the builder
	WithDataBuilder(func(ServicesBuilder) ServicesBuilder) MultipleServiceRequestBuilder
	// Build builds the MultipleServiceRequest or returns an error if something is wrong
	Build() (MultipleServiceRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MultipleServiceRequest
}

// NewMultipleServiceRequestBuilder() creates a MultipleServiceRequestBuilder
func NewMultipleServiceRequestBuilder() MultipleServiceRequestBuilder {
	return &_MultipleServiceRequestBuilder{_MultipleServiceRequest: new(_MultipleServiceRequest)}
}

type _MultipleServiceRequestBuilder struct {
	*_MultipleServiceRequest

	err *utils.MultiError
}

var _ (MultipleServiceRequestBuilder) = (*_MultipleServiceRequestBuilder)(nil)

func (m *_MultipleServiceRequestBuilder) WithMandatoryFields(data Services) MultipleServiceRequestBuilder {
	return m.WithData(data)
}

func (m *_MultipleServiceRequestBuilder) WithData(data Services) MultipleServiceRequestBuilder {
	m.Data = data
	return m
}

func (m *_MultipleServiceRequestBuilder) WithDataBuilder(builderSupplier func(ServicesBuilder) ServicesBuilder) MultipleServiceRequestBuilder {
	builder := builderSupplier(m.Data.CreateServicesBuilder())
	var err error
	m.Data, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "ServicesBuilder failed"))
	}
	return m
}

func (m *_MultipleServiceRequestBuilder) Build() (MultipleServiceRequest, error) {
	if m.Data == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'data' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._MultipleServiceRequest.deepCopy(), nil
}

func (m *_MultipleServiceRequestBuilder) MustBuild() MultipleServiceRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_MultipleServiceRequestBuilder) DeepCopy() any {
	return m.CreateMultipleServiceRequestBuilder()
}

// CreateMultipleServiceRequestBuilder creates a MultipleServiceRequestBuilder
func (m *_MultipleServiceRequest) CreateMultipleServiceRequestBuilder() MultipleServiceRequestBuilder {
	if m == nil {
		return NewMultipleServiceRequestBuilder()
	}
	return &_MultipleServiceRequestBuilder{_MultipleServiceRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MultipleServiceRequest) GetService() uint8 {
	return 0x0A
}

func (m *_MultipleServiceRequest) GetResponse() bool {
	return bool(false)
}

func (m *_MultipleServiceRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MultipleServiceRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MultipleServiceRequest) GetData() Services {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_MultipleServiceRequest) GetRequestPathSize() uint8 {
	return MultipleServiceRequest_REQUESTPATHSIZE
}

func (m *_MultipleServiceRequest) GetRequestPath() uint32 {
	return MultipleServiceRequest_REQUESTPATH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMultipleServiceRequest(structType any) MultipleServiceRequest {
	if casted, ok := structType.(MultipleServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MultipleServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MultipleServiceRequest) GetTypeName() string {
	return "MultipleServiceRequest"
}

func (m *_MultipleServiceRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Const Field (requestPathSize)
	lengthInBits += 8

	// Const Field (requestPath)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MultipleServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MultipleServiceRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__multipleServiceRequest MultipleServiceRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MultipleServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MultipleServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadConstField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)), MultipleServiceRequest_REQUESTPATHSIZE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	requestPath, err := ReadConstField[uint32](ctx, "requestPath", ReadUnsignedInt(readBuffer, uint8(32)), MultipleServiceRequest_REQUESTPATH)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPath' field"))
	}
	_ = requestPath

	data, err := ReadSimpleField[Services](ctx, "data", ReadComplex[Services](ServicesParseWithBufferProducer((uint16)(uint16(serviceLen)-uint16(uint16(6)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("MultipleServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MultipleServiceRequest")
	}

	return m, nil
}

func (m *_MultipleServiceRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MultipleServiceRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MultipleServiceRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MultipleServiceRequest")
		}

		if err := WriteConstField(ctx, "requestPathSize", MultipleServiceRequest_REQUESTPATHSIZE, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteConstField(ctx, "requestPath", MultipleServiceRequest_REQUESTPATH, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPath' field")
		}

		if err := WriteSimpleField[Services](ctx, "data", m.GetData(), WriteComplex[Services](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MultipleServiceRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MultipleServiceRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MultipleServiceRequest) IsMultipleServiceRequest() {}

func (m *_MultipleServiceRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MultipleServiceRequest) deepCopy() *_MultipleServiceRequest {
	if m == nil {
		return nil
	}
	_MultipleServiceRequestCopy := &_MultipleServiceRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
		m.Data.DeepCopy().(Services),
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _MultipleServiceRequestCopy
}

func (m *_MultipleServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
