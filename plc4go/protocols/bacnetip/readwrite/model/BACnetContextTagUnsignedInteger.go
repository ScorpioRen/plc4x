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

// BACnetContextTagUnsignedInteger is the corresponding interface of BACnetContextTagUnsignedInteger
type BACnetContextTagUnsignedInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
	// IsBACnetContextTagUnsignedInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagUnsignedInteger()
	// CreateBuilder creates a BACnetContextTagUnsignedIntegerBuilder
	CreateBACnetContextTagUnsignedIntegerBuilder() BACnetContextTagUnsignedIntegerBuilder
}

// _BACnetContextTagUnsignedInteger is the data-structure of this message
type _BACnetContextTagUnsignedInteger struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadUnsignedInteger
}

var _ BACnetContextTagUnsignedInteger = (*_BACnetContextTagUnsignedInteger)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagUnsignedInteger)(nil)

// NewBACnetContextTagUnsignedInteger factory function for _BACnetContextTagUnsignedInteger
func NewBACnetContextTagUnsignedInteger(header BACnetTagHeader, payload BACnetTagPayloadUnsignedInteger, tagNumberArgument uint8) *_BACnetContextTagUnsignedInteger {
	if payload == nil {
		panic("payload of type BACnetTagPayloadUnsignedInteger for BACnetContextTagUnsignedInteger must not be nil")
	}
	_result := &_BACnetContextTagUnsignedInteger{
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

// BACnetContextTagUnsignedIntegerBuilder is a builder for BACnetContextTagUnsignedInteger
type BACnetContextTagUnsignedIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadUnsignedInteger) BACnetContextTagUnsignedIntegerBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadUnsignedInteger) BACnetContextTagUnsignedIntegerBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadUnsignedIntegerBuilder) BACnetTagPayloadUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder
	// Build builds the BACnetContextTagUnsignedInteger or returns an error if something is wrong
	Build() (BACnetContextTagUnsignedInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagUnsignedInteger
}

// NewBACnetContextTagUnsignedIntegerBuilder() creates a BACnetContextTagUnsignedIntegerBuilder
func NewBACnetContextTagUnsignedIntegerBuilder() BACnetContextTagUnsignedIntegerBuilder {
	return &_BACnetContextTagUnsignedIntegerBuilder{_BACnetContextTagUnsignedInteger: new(_BACnetContextTagUnsignedInteger)}
}

type _BACnetContextTagUnsignedIntegerBuilder struct {
	*_BACnetContextTagUnsignedInteger

	err *utils.MultiError
}

var _ (BACnetContextTagUnsignedIntegerBuilder) = (*_BACnetContextTagUnsignedIntegerBuilder)(nil)

func (m *_BACnetContextTagUnsignedIntegerBuilder) WithMandatoryFields(payload BACnetTagPayloadUnsignedInteger) BACnetContextTagUnsignedIntegerBuilder {
	return m.WithPayload(payload)
}

func (m *_BACnetContextTagUnsignedIntegerBuilder) WithPayload(payload BACnetTagPayloadUnsignedInteger) BACnetContextTagUnsignedIntegerBuilder {
	m.Payload = payload
	return m
}

func (m *_BACnetContextTagUnsignedIntegerBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadUnsignedIntegerBuilder) BACnetTagPayloadUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder {
	builder := builderSupplier(m.Payload.CreateBACnetTagPayloadUnsignedIntegerBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagPayloadUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetContextTagUnsignedIntegerBuilder) Build() (BACnetContextTagUnsignedInteger, error) {
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetContextTagUnsignedInteger.deepCopy(), nil
}

func (m *_BACnetContextTagUnsignedIntegerBuilder) MustBuild() BACnetContextTagUnsignedInteger {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetContextTagUnsignedIntegerBuilder) DeepCopy() any {
	return m.CreateBACnetContextTagUnsignedIntegerBuilder()
}

// CreateBACnetContextTagUnsignedIntegerBuilder creates a BACnetContextTagUnsignedIntegerBuilder
func (m *_BACnetContextTagUnsignedInteger) CreateBACnetContextTagUnsignedIntegerBuilder() BACnetContextTagUnsignedIntegerBuilder {
	if m == nil {
		return NewBACnetContextTagUnsignedIntegerBuilder()
	}
	return &_BACnetContextTagUnsignedIntegerBuilder{_BACnetContextTagUnsignedInteger: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagUnsignedInteger) GetDataType() BACnetDataType {
	return BACnetDataType_UNSIGNED_INTEGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagUnsignedInteger) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagUnsignedInteger) GetPayload() BACnetTagPayloadUnsignedInteger {
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

func (m *_BACnetContextTagUnsignedInteger) GetActualValue() uint64 {
	ctx := context.Background()
	_ = ctx
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagUnsignedInteger(structType any) BACnetContextTagUnsignedInteger {
	if casted, ok := structType.(BACnetContextTagUnsignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagUnsignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagUnsignedInteger) GetTypeName() string {
	return "BACnetContextTagUnsignedInteger"
}

func (m *_BACnetContextTagUnsignedInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagUnsignedInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagUnsignedInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagUnsignedInteger BACnetContextTagUnsignedInteger, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagUnsignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagUnsignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadUnsignedInteger](ctx, "payload", ReadComplex[BACnetTagPayloadUnsignedInteger](BACnetTagPayloadUnsignedIntegerParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint64](ctx, "actualValue", (*uint64)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagUnsignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagUnsignedInteger")
	}

	return m, nil
}

func (m *_BACnetContextTagUnsignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagUnsignedInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagUnsignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagUnsignedInteger")
		}

		if err := WriteSimpleField[BACnetTagPayloadUnsignedInteger](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagUnsignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagUnsignedInteger")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagUnsignedInteger) IsBACnetContextTagUnsignedInteger() {}

func (m *_BACnetContextTagUnsignedInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagUnsignedInteger) deepCopy() *_BACnetContextTagUnsignedInteger {
	if m == nil {
		return nil
	}
	_BACnetContextTagUnsignedIntegerCopy := &_BACnetContextTagUnsignedInteger{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadUnsignedInteger),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagUnsignedIntegerCopy
}

func (m *_BACnetContextTagUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
