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

// S7PayloadUserDataItemCpuFunctionReadSzlRequest is the corresponding interface of S7PayloadUserDataItemCpuFunctionReadSzlRequest
type S7PayloadUserDataItemCpuFunctionReadSzlRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetSzlId returns SzlId (property field)
	GetSzlId() SzlId
	// GetSzlIndex returns SzlIndex (property field)
	GetSzlIndex() uint16
	// IsS7PayloadUserDataItemCpuFunctionReadSzlRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionReadSzlRequest()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
	CreateS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder() S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
}

// _S7PayloadUserDataItemCpuFunctionReadSzlRequest is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionReadSzlRequest struct {
	S7PayloadUserDataItemContract
	SzlId    SzlId
	SzlIndex uint16
}

var _ S7PayloadUserDataItemCpuFunctionReadSzlRequest = (*_S7PayloadUserDataItemCpuFunctionReadSzlRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionReadSzlRequest)(nil)

// NewS7PayloadUserDataItemCpuFunctionReadSzlRequest factory function for _S7PayloadUserDataItemCpuFunctionReadSzlRequest
func NewS7PayloadUserDataItemCpuFunctionReadSzlRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, szlId SzlId, szlIndex uint16) *_S7PayloadUserDataItemCpuFunctionReadSzlRequest {
	if szlId == nil {
		panic("szlId of type SzlId for S7PayloadUserDataItemCpuFunctionReadSzlRequest must not be nil")
	}
	_result := &_S7PayloadUserDataItemCpuFunctionReadSzlRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		SzlId:                         szlId,
		SzlIndex:                      szlIndex,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder is a builder for S7PayloadUserDataItemCpuFunctionReadSzlRequest
type S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(szlId SzlId, szlIndex uint16) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
	// WithSzlId adds SzlId (property field)
	WithSzlId(SzlId) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
	// WithSzlIdBuilder adds SzlId (property field) which is build by the builder
	WithSzlIdBuilder(func(SzlIdBuilder) SzlIdBuilder) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
	// WithSzlIndex adds SzlIndex (property field)
	WithSzlIndex(uint16) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionReadSzlRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionReadSzlRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionReadSzlRequest
}

// NewS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder() creates a S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
func NewS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder() S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder {
	return &_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder{_S7PayloadUserDataItemCpuFunctionReadSzlRequest: new(_S7PayloadUserDataItemCpuFunctionReadSzlRequest)}
}

type _S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionReadSzlRequest

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) = (*_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder)(nil)

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) WithMandatoryFields(szlId SzlId, szlIndex uint16) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder {
	return m.WithSzlId(szlId).WithSzlIndex(szlIndex)
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) WithSzlId(szlId SzlId) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder {
	m.SzlId = szlId
	return m
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) WithSzlIdBuilder(builderSupplier func(SzlIdBuilder) SzlIdBuilder) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder {
	builder := builderSupplier(m.SzlId.CreateSzlIdBuilder())
	var err error
	m.SzlId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "SzlIdBuilder failed"))
	}
	return m
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) WithSzlIndex(szlIndex uint16) S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder {
	m.SzlIndex = szlIndex
	return m
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) Build() (S7PayloadUserDataItemCpuFunctionReadSzlRequest, error) {
	if m.SzlId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'szlId' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._S7PayloadUserDataItemCpuFunctionReadSzlRequest.deepCopy(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionReadSzlRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder) DeepCopy() any {
	return m.CreateS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder()
}

// CreateS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder creates a S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder
func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) CreateS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder() S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder {
	if m == nil {
		return NewS7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilder{_S7PayloadUserDataItemCpuFunctionReadSzlRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetCpuSubfunction() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetSzlId() SzlId {
	return m.SzlId
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetSzlIndex() uint16 {
	return m.SzlIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionReadSzlRequest(structType any) S7PayloadUserDataItemCpuFunctionReadSzlRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionReadSzlRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionReadSzlRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlRequest"
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (szlId)
	lengthInBits += m.SzlId.GetLengthInBits(ctx)

	// Simple field (szlIndex)
	lengthInBits += 16

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionReadSzlRequest S7PayloadUserDataItemCpuFunctionReadSzlRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionReadSzlRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	szlId, err := ReadSimpleField[SzlId](ctx, "szlId", ReadComplex[SzlId](SzlIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'szlId' field"))
	}
	m.SzlId = szlId

	szlIndex, err := ReadSimpleField(ctx, "szlIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'szlIndex' field"))
	}
	m.SzlIndex = szlIndex

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionReadSzlRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionReadSzlRequest")
		}

		if err := WriteSimpleField[SzlId](ctx, "szlId", m.GetSzlId(), WriteComplex[SzlId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'szlId' field")
		}

		if err := WriteSimpleField[uint16](ctx, "szlIndex", m.GetSzlIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'szlIndex' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionReadSzlRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) IsS7PayloadUserDataItemCpuFunctionReadSzlRequest() {
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) deepCopy() *_S7PayloadUserDataItemCpuFunctionReadSzlRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionReadSzlRequestCopy := &_S7PayloadUserDataItemCpuFunctionReadSzlRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.SzlId.DeepCopy().(SzlId),
		m.SzlIndex,
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionReadSzlRequestCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
