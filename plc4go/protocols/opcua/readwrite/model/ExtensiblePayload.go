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

// ExtensiblePayload is the corresponding interface of ExtensiblePayload
type ExtensiblePayload interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Payload
	// GetPayload returns Payload (property field)
	GetPayload() ExtensionObject
	// IsExtensiblePayload is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensiblePayload()
	// CreateBuilder creates a ExtensiblePayloadBuilder
	CreateExtensiblePayloadBuilder() ExtensiblePayloadBuilder
}

// _ExtensiblePayload is the data-structure of this message
type _ExtensiblePayload struct {
	PayloadContract
	Payload ExtensionObject
}

var _ ExtensiblePayload = (*_ExtensiblePayload)(nil)
var _ PayloadRequirements = (*_ExtensiblePayload)(nil)

// NewExtensiblePayload factory function for _ExtensiblePayload
func NewExtensiblePayload(sequenceHeader SequenceHeader, payload ExtensionObject, byteCount uint32) *_ExtensiblePayload {
	if payload == nil {
		panic("payload of type ExtensionObject for ExtensiblePayload must not be nil")
	}
	_result := &_ExtensiblePayload{
		PayloadContract: NewPayload(sequenceHeader, byteCount),
		Payload:         payload,
	}
	_result.PayloadContract.(*_Payload)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ExtensiblePayloadBuilder is a builder for ExtensiblePayload
type ExtensiblePayloadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload ExtensionObject) ExtensiblePayloadBuilder
	// WithPayload adds Payload (property field)
	WithPayload(ExtensionObject) ExtensiblePayloadBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) ExtensiblePayloadBuilder
	// Build builds the ExtensiblePayload or returns an error if something is wrong
	Build() (ExtensiblePayload, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ExtensiblePayload
}

// NewExtensiblePayloadBuilder() creates a ExtensiblePayloadBuilder
func NewExtensiblePayloadBuilder() ExtensiblePayloadBuilder {
	return &_ExtensiblePayloadBuilder{_ExtensiblePayload: new(_ExtensiblePayload)}
}

type _ExtensiblePayloadBuilder struct {
	*_ExtensiblePayload

	err *utils.MultiError
}

var _ (ExtensiblePayloadBuilder) = (*_ExtensiblePayloadBuilder)(nil)

func (m *_ExtensiblePayloadBuilder) WithMandatoryFields(payload ExtensionObject) ExtensiblePayloadBuilder {
	return m.WithPayload(payload)
}

func (m *_ExtensiblePayloadBuilder) WithPayload(payload ExtensionObject) ExtensiblePayloadBuilder {
	m.Payload = payload
	return m
}

func (m *_ExtensiblePayloadBuilder) WithPayloadBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) ExtensiblePayloadBuilder {
	builder := builderSupplier(m.Payload.CreateExtensionObjectBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return m
}

func (m *_ExtensiblePayloadBuilder) Build() (ExtensiblePayload, error) {
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ExtensiblePayload.deepCopy(), nil
}

func (m *_ExtensiblePayloadBuilder) MustBuild() ExtensiblePayload {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ExtensiblePayloadBuilder) DeepCopy() any {
	return m.CreateExtensiblePayloadBuilder()
}

// CreateExtensiblePayloadBuilder creates a ExtensiblePayloadBuilder
func (m *_ExtensiblePayload) CreateExtensiblePayloadBuilder() ExtensiblePayloadBuilder {
	if m == nil {
		return NewExtensiblePayloadBuilder()
	}
	return &_ExtensiblePayloadBuilder{_ExtensiblePayload: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ExtensiblePayload) GetExtensible() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ExtensiblePayload) GetParent() PayloadContract {
	return m.PayloadContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensiblePayload) GetPayload() ExtensionObject {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastExtensiblePayload(structType any) ExtensiblePayload {
	if casted, ok := structType.(ExtensiblePayload); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensiblePayload); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensiblePayload) GetTypeName() string {
	return "ExtensiblePayload"
}

func (m *_ExtensiblePayload) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PayloadContract.(*_Payload).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ExtensiblePayload) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ExtensiblePayload) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Payload, extensible bool, byteCount uint32) (__extensiblePayload ExtensiblePayload, err error) {
	m.PayloadContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtensiblePayload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensiblePayload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[ExtensionObject](ctx, "payload", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(false))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("ExtensiblePayload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensiblePayload")
	}

	return m, nil
}

func (m *_ExtensiblePayload) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExtensiblePayload) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ExtensiblePayload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ExtensiblePayload")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "payload", m.GetPayload(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("ExtensiblePayload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ExtensiblePayload")
		}
		return nil
	}
	return m.PayloadContract.(*_Payload).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ExtensiblePayload) IsExtensiblePayload() {}

func (m *_ExtensiblePayload) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ExtensiblePayload) deepCopy() *_ExtensiblePayload {
	if m == nil {
		return nil
	}
	_ExtensiblePayloadCopy := &_ExtensiblePayload{
		m.PayloadContract.(*_Payload).deepCopy(),
		m.Payload.DeepCopy().(ExtensionObject),
	}
	m.PayloadContract.(*_Payload)._SubType = m
	return _ExtensiblePayloadCopy
}

func (m *_ExtensiblePayload) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
