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

// BACnetApplicationTagEnumerated is the corresponding interface of BACnetApplicationTagEnumerated
type BACnetApplicationTagEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadEnumerated
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint32
	// IsBACnetApplicationTagEnumerated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagEnumerated()
	// CreateBuilder creates a BACnetApplicationTagEnumeratedBuilder
	CreateBACnetApplicationTagEnumeratedBuilder() BACnetApplicationTagEnumeratedBuilder
}

// _BACnetApplicationTagEnumerated is the data-structure of this message
type _BACnetApplicationTagEnumerated struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadEnumerated
}

var _ BACnetApplicationTagEnumerated = (*_BACnetApplicationTagEnumerated)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagEnumerated)(nil)

// NewBACnetApplicationTagEnumerated factory function for _BACnetApplicationTagEnumerated
func NewBACnetApplicationTagEnumerated(header BACnetTagHeader, payload BACnetTagPayloadEnumerated) *_BACnetApplicationTagEnumerated {
	if payload == nil {
		panic("payload of type BACnetTagPayloadEnumerated for BACnetApplicationTagEnumerated must not be nil")
	}
	_result := &_BACnetApplicationTagEnumerated{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagEnumeratedBuilder is a builder for BACnetApplicationTagEnumerated
type BACnetApplicationTagEnumeratedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadEnumerated) BACnetApplicationTagEnumeratedBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadEnumerated) BACnetApplicationTagEnumeratedBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadEnumeratedBuilder) BACnetTagPayloadEnumeratedBuilder) BACnetApplicationTagEnumeratedBuilder
	// Build builds the BACnetApplicationTagEnumerated or returns an error if something is wrong
	Build() (BACnetApplicationTagEnumerated, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagEnumerated
}

// NewBACnetApplicationTagEnumeratedBuilder() creates a BACnetApplicationTagEnumeratedBuilder
func NewBACnetApplicationTagEnumeratedBuilder() BACnetApplicationTagEnumeratedBuilder {
	return &_BACnetApplicationTagEnumeratedBuilder{_BACnetApplicationTagEnumerated: new(_BACnetApplicationTagEnumerated)}
}

type _BACnetApplicationTagEnumeratedBuilder struct {
	*_BACnetApplicationTagEnumerated

	err *utils.MultiError
}

var _ (BACnetApplicationTagEnumeratedBuilder) = (*_BACnetApplicationTagEnumeratedBuilder)(nil)

func (m *_BACnetApplicationTagEnumeratedBuilder) WithMandatoryFields(payload BACnetTagPayloadEnumerated) BACnetApplicationTagEnumeratedBuilder {
	return m.WithPayload(payload)
}

func (m *_BACnetApplicationTagEnumeratedBuilder) WithPayload(payload BACnetTagPayloadEnumerated) BACnetApplicationTagEnumeratedBuilder {
	m.Payload = payload
	return m
}

func (m *_BACnetApplicationTagEnumeratedBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadEnumeratedBuilder) BACnetTagPayloadEnumeratedBuilder) BACnetApplicationTagEnumeratedBuilder {
	builder := builderSupplier(m.Payload.CreateBACnetTagPayloadEnumeratedBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagPayloadEnumeratedBuilder failed"))
	}
	return m
}

func (m *_BACnetApplicationTagEnumeratedBuilder) Build() (BACnetApplicationTagEnumerated, error) {
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetApplicationTagEnumerated.deepCopy(), nil
}

func (m *_BACnetApplicationTagEnumeratedBuilder) MustBuild() BACnetApplicationTagEnumerated {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetApplicationTagEnumeratedBuilder) DeepCopy() any {
	return m.CreateBACnetApplicationTagEnumeratedBuilder()
}

// CreateBACnetApplicationTagEnumeratedBuilder creates a BACnetApplicationTagEnumeratedBuilder
func (m *_BACnetApplicationTagEnumerated) CreateBACnetApplicationTagEnumeratedBuilder() BACnetApplicationTagEnumeratedBuilder {
	if m == nil {
		return NewBACnetApplicationTagEnumeratedBuilder()
	}
	return &_BACnetApplicationTagEnumeratedBuilder{_BACnetApplicationTagEnumerated: m.deepCopy()}
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

func (m *_BACnetApplicationTagEnumerated) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagEnumerated) GetPayload() BACnetTagPayloadEnumerated {
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

func (m *_BACnetApplicationTagEnumerated) GetActualValue() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagEnumerated(structType any) BACnetApplicationTagEnumerated {
	if casted, ok := structType.(BACnetApplicationTagEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagEnumerated) GetTypeName() string {
	return "BACnetApplicationTagEnumerated"
}

func (m *_BACnetApplicationTagEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagEnumerated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagEnumerated BACnetApplicationTagEnumerated, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadEnumerated](ctx, "payload", ReadComplex[BACnetTagPayloadEnumerated](BACnetTagPayloadEnumeratedParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint32](ctx, "actualValue", (*uint32)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagEnumerated")
	}

	return m, nil
}

func (m *_BACnetApplicationTagEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagEnumerated")
		}

		if err := WriteSimpleField[BACnetTagPayloadEnumerated](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadEnumerated](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagEnumerated")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagEnumerated) IsBACnetApplicationTagEnumerated() {}

func (m *_BACnetApplicationTagEnumerated) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagEnumerated) deepCopy() *_BACnetApplicationTagEnumerated {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagEnumeratedCopy := &_BACnetApplicationTagEnumerated{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadEnumerated),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagEnumeratedCopy
}

func (m *_BACnetApplicationTagEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
