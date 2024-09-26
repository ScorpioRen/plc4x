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
const RequestDirectCommandAccess_AT byte = 0x40

// RequestDirectCommandAccess is the corresponding interface of RequestDirectCommandAccess
type RequestDirectCommandAccess interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Request
	// GetCalData returns CalData (property field)
	GetCalData() CALData
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetCalDataDecoded returns CalDataDecoded (virtual field)
	GetCalDataDecoded() CALData
	// IsRequestDirectCommandAccess is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequestDirectCommandAccess()
	// CreateBuilder creates a RequestDirectCommandAccessBuilder
	CreateRequestDirectCommandAccessBuilder() RequestDirectCommandAccessBuilder
}

// _RequestDirectCommandAccess is the data-structure of this message
type _RequestDirectCommandAccess struct {
	RequestContract
	CalData CALData
	Alpha   Alpha
}

var _ RequestDirectCommandAccess = (*_RequestDirectCommandAccess)(nil)
var _ RequestRequirements = (*_RequestDirectCommandAccess)(nil)

// NewRequestDirectCommandAccess factory function for _RequestDirectCommandAccess
func NewRequestDirectCommandAccess(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, calData CALData, alpha Alpha, cBusOptions CBusOptions) *_RequestDirectCommandAccess {
	_result := &_RequestDirectCommandAccess{
		RequestContract: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
		CalData:         calData,
		Alpha:           alpha,
	}
	_result.RequestContract.(*_Request)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RequestDirectCommandAccessBuilder is a builder for RequestDirectCommandAccess
type RequestDirectCommandAccessBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(calData CALData) RequestDirectCommandAccessBuilder
	// WithCalData adds CalData (property field)
	WithCalData(CALData) RequestDirectCommandAccessBuilder
	// WithCalDataBuilder adds CalData (property field) which is build by the builder
	WithCalDataBuilder(func(CALDataBuilder) CALDataBuilder) RequestDirectCommandAccessBuilder
	// WithAlpha adds Alpha (property field)
	WithOptionalAlpha(Alpha) RequestDirectCommandAccessBuilder
	// WithOptionalAlphaBuilder adds Alpha (property field) which is build by the builder
	WithOptionalAlphaBuilder(func(AlphaBuilder) AlphaBuilder) RequestDirectCommandAccessBuilder
	// Build builds the RequestDirectCommandAccess or returns an error if something is wrong
	Build() (RequestDirectCommandAccess, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RequestDirectCommandAccess
}

// NewRequestDirectCommandAccessBuilder() creates a RequestDirectCommandAccessBuilder
func NewRequestDirectCommandAccessBuilder() RequestDirectCommandAccessBuilder {
	return &_RequestDirectCommandAccessBuilder{_RequestDirectCommandAccess: new(_RequestDirectCommandAccess)}
}

type _RequestDirectCommandAccessBuilder struct {
	*_RequestDirectCommandAccess

	err *utils.MultiError
}

var _ (RequestDirectCommandAccessBuilder) = (*_RequestDirectCommandAccessBuilder)(nil)

func (m *_RequestDirectCommandAccessBuilder) WithMandatoryFields(calData CALData) RequestDirectCommandAccessBuilder {
	return m.WithCalData(calData)
}

func (m *_RequestDirectCommandAccessBuilder) WithCalData(calData CALData) RequestDirectCommandAccessBuilder {
	m.CalData = calData
	return m
}

func (m *_RequestDirectCommandAccessBuilder) WithCalDataBuilder(builderSupplier func(CALDataBuilder) CALDataBuilder) RequestDirectCommandAccessBuilder {
	builder := builderSupplier(m.CalData.CreateCALDataBuilder())
	var err error
	m.CalData, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "CALDataBuilder failed"))
	}
	return m
}

func (m *_RequestDirectCommandAccessBuilder) WithOptionalAlpha(alpha Alpha) RequestDirectCommandAccessBuilder {
	m.Alpha = alpha
	return m
}

func (m *_RequestDirectCommandAccessBuilder) WithOptionalAlphaBuilder(builderSupplier func(AlphaBuilder) AlphaBuilder) RequestDirectCommandAccessBuilder {
	builder := builderSupplier(m.Alpha.CreateAlphaBuilder())
	var err error
	m.Alpha, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "AlphaBuilder failed"))
	}
	return m
}

func (m *_RequestDirectCommandAccessBuilder) Build() (RequestDirectCommandAccess, error) {
	if m.CalData == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'calData' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._RequestDirectCommandAccess.deepCopy(), nil
}

func (m *_RequestDirectCommandAccessBuilder) MustBuild() RequestDirectCommandAccess {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_RequestDirectCommandAccessBuilder) DeepCopy() any {
	return m.CreateRequestDirectCommandAccessBuilder()
}

// CreateRequestDirectCommandAccessBuilder creates a RequestDirectCommandAccessBuilder
func (m *_RequestDirectCommandAccess) CreateRequestDirectCommandAccessBuilder() RequestDirectCommandAccessBuilder {
	if m == nil {
		return NewRequestDirectCommandAccessBuilder()
	}
	return &_RequestDirectCommandAccessBuilder{_RequestDirectCommandAccess: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestDirectCommandAccess) GetParent() RequestContract {
	return m.RequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetCalData() CALData {
	return m.CalData
}

func (m *_RequestDirectCommandAccess) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetCalDataDecoded() CALData {
	ctx := context.Background()
	_ = ctx
	alpha := m.GetAlpha()
	_ = alpha
	return CastCALData(m.GetCalData())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetAt() byte {
	return RequestDirectCommandAccess_AT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRequestDirectCommandAccess(structType any) RequestDirectCommandAccess {
	if casted, ok := structType.(RequestDirectCommandAccess); ok {
		return casted
	}
	if casted, ok := structType.(*RequestDirectCommandAccess); ok {
		return *casted
	}
	return nil
}

func (m *_RequestDirectCommandAccess) GetTypeName() string {
	return "RequestDirectCommandAccess"
}

func (m *_RequestDirectCommandAccess) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.RequestContract.(*_Request).getLengthInBits(ctx))

	// Const Field (at)
	lengthInBits += 8

	// Manual Field (calData)
	lengthInBits += uint16(int32((int32(m.GetCalData().GetLengthInBytes(ctx)) * int32(int32(2)))) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_RequestDirectCommandAccess) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RequestDirectCommandAccess) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Request, cBusOptions CBusOptions) (__requestDirectCommandAccess RequestDirectCommandAccess, err error) {
	m.RequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestDirectCommandAccess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestDirectCommandAccess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	at, err := ReadConstField[byte](ctx, "at", ReadByte(readBuffer, 8), RequestDirectCommandAccess_AT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'at' field"))
	}
	_ = at

	calData, err := ReadManualField[CALData](ctx, "calData", readBuffer, EnsureType[CALData](ReadCALData(ctx, readBuffer)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calData' field"))
	}
	m.CalData = calData

	calDataDecoded, err := ReadVirtualField[CALData](ctx, "calDataDecoded", (*CALData)(nil), calData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calDataDecoded' field"))
	}
	_ = calDataDecoded

	var alpha Alpha
	_alpha, err := ReadOptionalField[Alpha](ctx, "alpha", ReadComplex[Alpha](AlphaParseWithBuffer, readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alpha' field"))
	}
	if _alpha != nil {
		alpha = *_alpha
		m.Alpha = alpha
	}

	if closeErr := readBuffer.CloseContext("RequestDirectCommandAccess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestDirectCommandAccess")
	}

	return m, nil
}

func (m *_RequestDirectCommandAccess) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestDirectCommandAccess) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestDirectCommandAccess"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestDirectCommandAccess")
		}

		if err := WriteConstField(ctx, "at", RequestDirectCommandAccess_AT, WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'at' field")
		}

		if err := WriteManualField[CALData](ctx, "calData", func(ctx context.Context) error { return WriteCALData(ctx, writeBuffer, m.GetCalData()) }, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'calData' field")
		}
		// Virtual field
		calDataDecoded := m.GetCalDataDecoded()
		_ = calDataDecoded
		if _calDataDecodedErr := writeBuffer.WriteVirtual(ctx, "calDataDecoded", m.GetCalDataDecoded()); _calDataDecodedErr != nil {
			return errors.Wrap(_calDataDecodedErr, "Error serializing 'calDataDecoded' field")
		}

		if err := WriteOptionalField[Alpha](ctx, "alpha", GetRef(m.GetAlpha()), WriteComplex[Alpha](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'alpha' field")
		}

		if popErr := writeBuffer.PopContext("RequestDirectCommandAccess"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestDirectCommandAccess")
		}
		return nil
	}
	return m.RequestContract.(*_Request).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestDirectCommandAccess) IsRequestDirectCommandAccess() {}

func (m *_RequestDirectCommandAccess) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RequestDirectCommandAccess) deepCopy() *_RequestDirectCommandAccess {
	if m == nil {
		return nil
	}
	_RequestDirectCommandAccessCopy := &_RequestDirectCommandAccess{
		m.RequestContract.(*_Request).deepCopy(),
		m.CalData.DeepCopy().(CALData),
		m.Alpha.DeepCopy().(Alpha),
	}
	m.RequestContract.(*_Request)._SubType = m
	return _RequestDirectCommandAccessCopy
}

func (m *_RequestDirectCommandAccess) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
