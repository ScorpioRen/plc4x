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

// BrowseNextRequest is the corresponding interface of BrowseNextRequest
type BrowseNextRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetReleaseContinuationPoints returns ReleaseContinuationPoints (property field)
	GetReleaseContinuationPoints() bool
	// GetNoOfContinuationPoints returns NoOfContinuationPoints (property field)
	GetNoOfContinuationPoints() int32
	// GetContinuationPoints returns ContinuationPoints (property field)
	GetContinuationPoints() []PascalByteString
	// IsBrowseNextRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrowseNextRequest()
	// CreateBuilder creates a BrowseNextRequestBuilder
	CreateBrowseNextRequestBuilder() BrowseNextRequestBuilder
}

// _BrowseNextRequest is the data-structure of this message
type _BrowseNextRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader             ExtensionObjectDefinition
	ReleaseContinuationPoints bool
	NoOfContinuationPoints    int32
	ContinuationPoints        []PascalByteString
	// Reserved Fields
	reservedField0 *uint8
}

var _ BrowseNextRequest = (*_BrowseNextRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowseNextRequest)(nil)

// NewBrowseNextRequest factory function for _BrowseNextRequest
func NewBrowseNextRequest(requestHeader ExtensionObjectDefinition, releaseContinuationPoints bool, noOfContinuationPoints int32, continuationPoints []PascalByteString) *_BrowseNextRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for BrowseNextRequest must not be nil")
	}
	_result := &_BrowseNextRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ReleaseContinuationPoints:         releaseContinuationPoints,
		NoOfContinuationPoints:            noOfContinuationPoints,
		ContinuationPoints:                continuationPoints,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrowseNextRequestBuilder is a builder for BrowseNextRequest
type BrowseNextRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, releaseContinuationPoints bool, noOfContinuationPoints int32, continuationPoints []PascalByteString) BrowseNextRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) BrowseNextRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) BrowseNextRequestBuilder
	// WithReleaseContinuationPoints adds ReleaseContinuationPoints (property field)
	WithReleaseContinuationPoints(bool) BrowseNextRequestBuilder
	// WithNoOfContinuationPoints adds NoOfContinuationPoints (property field)
	WithNoOfContinuationPoints(int32) BrowseNextRequestBuilder
	// WithContinuationPoints adds ContinuationPoints (property field)
	WithContinuationPoints(...PascalByteString) BrowseNextRequestBuilder
	// Build builds the BrowseNextRequest or returns an error if something is wrong
	Build() (BrowseNextRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrowseNextRequest
}

// NewBrowseNextRequestBuilder() creates a BrowseNextRequestBuilder
func NewBrowseNextRequestBuilder() BrowseNextRequestBuilder {
	return &_BrowseNextRequestBuilder{_BrowseNextRequest: new(_BrowseNextRequest)}
}

type _BrowseNextRequestBuilder struct {
	*_BrowseNextRequest

	err *utils.MultiError
}

var _ (BrowseNextRequestBuilder) = (*_BrowseNextRequestBuilder)(nil)

func (m *_BrowseNextRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, releaseContinuationPoints bool, noOfContinuationPoints int32, continuationPoints []PascalByteString) BrowseNextRequestBuilder {
	return m.WithRequestHeader(requestHeader).WithReleaseContinuationPoints(releaseContinuationPoints).WithNoOfContinuationPoints(noOfContinuationPoints).WithContinuationPoints(continuationPoints...)
}

func (m *_BrowseNextRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) BrowseNextRequestBuilder {
	m.RequestHeader = requestHeader
	return m
}

func (m *_BrowseNextRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) BrowseNextRequestBuilder {
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

func (m *_BrowseNextRequestBuilder) WithReleaseContinuationPoints(releaseContinuationPoints bool) BrowseNextRequestBuilder {
	m.ReleaseContinuationPoints = releaseContinuationPoints
	return m
}

func (m *_BrowseNextRequestBuilder) WithNoOfContinuationPoints(noOfContinuationPoints int32) BrowseNextRequestBuilder {
	m.NoOfContinuationPoints = noOfContinuationPoints
	return m
}

func (m *_BrowseNextRequestBuilder) WithContinuationPoints(continuationPoints ...PascalByteString) BrowseNextRequestBuilder {
	m.ContinuationPoints = continuationPoints
	return m
}

func (m *_BrowseNextRequestBuilder) Build() (BrowseNextRequest, error) {
	if m.RequestHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BrowseNextRequest.deepCopy(), nil
}

func (m *_BrowseNextRequestBuilder) MustBuild() BrowseNextRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BrowseNextRequestBuilder) DeepCopy() any {
	return m.CreateBrowseNextRequestBuilder()
}

// CreateBrowseNextRequestBuilder creates a BrowseNextRequestBuilder
func (m *_BrowseNextRequest) CreateBrowseNextRequestBuilder() BrowseNextRequestBuilder {
	if m == nil {
		return NewBrowseNextRequestBuilder()
	}
	return &_BrowseNextRequestBuilder{_BrowseNextRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseNextRequest) GetIdentifier() string {
	return "533"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseNextRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseNextRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_BrowseNextRequest) GetReleaseContinuationPoints() bool {
	return m.ReleaseContinuationPoints
}

func (m *_BrowseNextRequest) GetNoOfContinuationPoints() int32 {
	return m.NoOfContinuationPoints
}

func (m *_BrowseNextRequest) GetContinuationPoints() []PascalByteString {
	return m.ContinuationPoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrowseNextRequest(structType any) BrowseNextRequest {
	if casted, ok := structType.(BrowseNextRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseNextRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseNextRequest) GetTypeName() string {
	return "BrowseNextRequest"
}

func (m *_BrowseNextRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (releaseContinuationPoints)
	lengthInBits += 1

	// Simple field (noOfContinuationPoints)
	lengthInBits += 32

	// Array field
	if len(m.ContinuationPoints) > 0 {
		for _curItem, element := range m.ContinuationPoints {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ContinuationPoints), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowseNextRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowseNextRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__browseNextRequest BrowseNextRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowseNextRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseNextRequest")
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

	releaseContinuationPoints, err := ReadSimpleField(ctx, "releaseContinuationPoints", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'releaseContinuationPoints' field"))
	}
	m.ReleaseContinuationPoints = releaseContinuationPoints

	noOfContinuationPoints, err := ReadSimpleField(ctx, "noOfContinuationPoints", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfContinuationPoints' field"))
	}
	m.NoOfContinuationPoints = noOfContinuationPoints

	continuationPoints, err := ReadCountArrayField[PascalByteString](ctx, "continuationPoints", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer), uint64(noOfContinuationPoints))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'continuationPoints' field"))
	}
	m.ContinuationPoints = continuationPoints

	if closeErr := readBuffer.CloseContext("BrowseNextRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseNextRequest")
	}

	return m, nil
}

func (m *_BrowseNextRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseNextRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseNextRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseNextRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "releaseContinuationPoints", m.GetReleaseContinuationPoints(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'releaseContinuationPoints' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfContinuationPoints", m.GetNoOfContinuationPoints(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfContinuationPoints' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "continuationPoints", m.GetContinuationPoints(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'continuationPoints' field")
		}

		if popErr := writeBuffer.PopContext("BrowseNextRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseNextRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseNextRequest) IsBrowseNextRequest() {}

func (m *_BrowseNextRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrowseNextRequest) deepCopy() *_BrowseNextRequest {
	if m == nil {
		return nil
	}
	_BrowseNextRequestCopy := &_BrowseNextRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.ReleaseContinuationPoints,
		m.NoOfContinuationPoints,
		utils.DeepCopySlice[PascalByteString, PascalByteString](m.ContinuationPoints),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrowseNextRequestCopy
}

func (m *_BrowseNextRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
