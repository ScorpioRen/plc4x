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

// BACnetContextTagDouble is the corresponding interface of BACnetContextTagDouble
type BACnetContextTagDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float64
	// IsBACnetContextTagDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagDouble()
	// CreateBuilder creates a BACnetContextTagDoubleBuilder
	CreateBACnetContextTagDoubleBuilder() BACnetContextTagDoubleBuilder
}

// _BACnetContextTagDouble is the data-structure of this message
type _BACnetContextTagDouble struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadDouble
}

var _ BACnetContextTagDouble = (*_BACnetContextTagDouble)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagDouble)(nil)

// NewBACnetContextTagDouble factory function for _BACnetContextTagDouble
func NewBACnetContextTagDouble(header BACnetTagHeader, payload BACnetTagPayloadDouble, tagNumberArgument uint8) *_BACnetContextTagDouble {
	if payload == nil {
		panic("payload of type BACnetTagPayloadDouble for BACnetContextTagDouble must not be nil")
	}
	_result := &_BACnetContextTagDouble{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagDoubleBuilder is a builder for BACnetContextTagDouble
type BACnetContextTagDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadDouble) BACnetContextTagDoubleBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadDouble) BACnetContextTagDoubleBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadDoubleBuilder) BACnetTagPayloadDoubleBuilder) BACnetContextTagDoubleBuilder
	// Build builds the BACnetContextTagDouble or returns an error if something is wrong
	Build() (BACnetContextTagDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagDouble
}

// NewBACnetContextTagDoubleBuilder() creates a BACnetContextTagDoubleBuilder
func NewBACnetContextTagDoubleBuilder() BACnetContextTagDoubleBuilder {
	return &_BACnetContextTagDoubleBuilder{_BACnetContextTagDouble: new(_BACnetContextTagDouble)}
}

type _BACnetContextTagDoubleBuilder struct {
	*_BACnetContextTagDouble

	err *utils.MultiError
}

var _ (BACnetContextTagDoubleBuilder) = (*_BACnetContextTagDoubleBuilder)(nil)

func (m *_BACnetContextTagDoubleBuilder) WithMandatoryFields(payload BACnetTagPayloadDouble) BACnetContextTagDoubleBuilder {
	return m.WithPayload(payload)
}

func (m *_BACnetContextTagDoubleBuilder) WithPayload(payload BACnetTagPayloadDouble) BACnetContextTagDoubleBuilder {
	m.Payload = payload
	return m
}

func (m *_BACnetContextTagDoubleBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadDoubleBuilder) BACnetTagPayloadDoubleBuilder) BACnetContextTagDoubleBuilder {
	builder := builderSupplier(m.Payload.CreateBACnetTagPayloadDoubleBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagPayloadDoubleBuilder failed"))
	}
	return m
}

func (m *_BACnetContextTagDoubleBuilder) Build() (BACnetContextTagDouble, error) {
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetContextTagDouble.deepCopy(), nil
}

func (m *_BACnetContextTagDoubleBuilder) MustBuild() BACnetContextTagDouble {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetContextTagDoubleBuilder) DeepCopy() any {
	return m.CreateBACnetContextTagDoubleBuilder()
}

// CreateBACnetContextTagDoubleBuilder creates a BACnetContextTagDoubleBuilder
func (m *_BACnetContextTagDouble) CreateBACnetContextTagDoubleBuilder() BACnetContextTagDoubleBuilder {
	if m == nil {
		return NewBACnetContextTagDoubleBuilder()
	}
	return &_BACnetContextTagDoubleBuilder{_BACnetContextTagDouble: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagDouble) GetDataType() BACnetDataType {
	return BACnetDataType_DOUBLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagDouble) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagDouble) GetPayload() BACnetTagPayloadDouble {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagDouble) GetActualValue() float64 {
	ctx := context.Background()
	_ = ctx
	return float64(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagDouble(structType any) BACnetContextTagDouble {
	if casted, ok := structType.(BACnetContextTagDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagDouble) GetTypeName() string {
	return "BACnetContextTagDouble"
}

func (m *_BACnetContextTagDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagDouble BACnetContextTagDouble, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadDouble](ctx, "payload", ReadComplex[BACnetTagPayloadDouble](BACnetTagPayloadDoubleParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[float64](ctx, "actualValue", (*float64)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagDouble")
	}

	return m, nil
}

func (m *_BACnetContextTagDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagDouble")
		}

		if err := WriteSimpleField[BACnetTagPayloadDouble](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagDouble")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagDouble) IsBACnetContextTagDouble() {}

func (m *_BACnetContextTagDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagDouble) deepCopy() *_BACnetContextTagDouble {
	if m == nil {
		return nil
	}
	_BACnetContextTagDoubleCopy := &_BACnetContextTagDouble{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadDouble),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagDoubleCopy
}

func (m *_BACnetContextTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
