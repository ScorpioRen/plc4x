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

// BACnetRecipientProcess is the corresponding interface of BACnetRecipientProcess
type BACnetRecipientProcess interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientEnclosed
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
	// IsBACnetRecipientProcess is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipientProcess()
	// CreateBuilder creates a BACnetRecipientProcessBuilder
	CreateBACnetRecipientProcessBuilder() BACnetRecipientProcessBuilder
}

// _BACnetRecipientProcess is the data-structure of this message
type _BACnetRecipientProcess struct {
	Recipient         BACnetRecipientEnclosed
	ProcessIdentifier BACnetContextTagUnsignedInteger
}

var _ BACnetRecipientProcess = (*_BACnetRecipientProcess)(nil)

// NewBACnetRecipientProcess factory function for _BACnetRecipientProcess
func NewBACnetRecipientProcess(recipient BACnetRecipientEnclosed, processIdentifier BACnetContextTagUnsignedInteger) *_BACnetRecipientProcess {
	if recipient == nil {
		panic("recipient of type BACnetRecipientEnclosed for BACnetRecipientProcess must not be nil")
	}
	return &_BACnetRecipientProcess{Recipient: recipient, ProcessIdentifier: processIdentifier}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRecipientProcessBuilder is a builder for BACnetRecipientProcess
type BACnetRecipientProcessBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recipient BACnetRecipientEnclosed) BACnetRecipientProcessBuilder
	// WithRecipient adds Recipient (property field)
	WithRecipient(BACnetRecipientEnclosed) BACnetRecipientProcessBuilder
	// WithRecipientBuilder adds Recipient (property field) which is build by the builder
	WithRecipientBuilder(func(BACnetRecipientEnclosedBuilder) BACnetRecipientEnclosedBuilder) BACnetRecipientProcessBuilder
	// WithProcessIdentifier adds ProcessIdentifier (property field)
	WithOptionalProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetRecipientProcessBuilder
	// WithOptionalProcessIdentifierBuilder adds ProcessIdentifier (property field) which is build by the builder
	WithOptionalProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetRecipientProcessBuilder
	// Build builds the BACnetRecipientProcess or returns an error if something is wrong
	Build() (BACnetRecipientProcess, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRecipientProcess
}

// NewBACnetRecipientProcessBuilder() creates a BACnetRecipientProcessBuilder
func NewBACnetRecipientProcessBuilder() BACnetRecipientProcessBuilder {
	return &_BACnetRecipientProcessBuilder{_BACnetRecipientProcess: new(_BACnetRecipientProcess)}
}

type _BACnetRecipientProcessBuilder struct {
	*_BACnetRecipientProcess

	err *utils.MultiError
}

var _ (BACnetRecipientProcessBuilder) = (*_BACnetRecipientProcessBuilder)(nil)

func (m *_BACnetRecipientProcessBuilder) WithMandatoryFields(recipient BACnetRecipientEnclosed) BACnetRecipientProcessBuilder {
	return m.WithRecipient(recipient)
}

func (m *_BACnetRecipientProcessBuilder) WithRecipient(recipient BACnetRecipientEnclosed) BACnetRecipientProcessBuilder {
	m.Recipient = recipient
	return m
}

func (m *_BACnetRecipientProcessBuilder) WithRecipientBuilder(builderSupplier func(BACnetRecipientEnclosedBuilder) BACnetRecipientEnclosedBuilder) BACnetRecipientProcessBuilder {
	builder := builderSupplier(m.Recipient.CreateBACnetRecipientEnclosedBuilder())
	var err error
	m.Recipient, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetRecipientEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetRecipientProcessBuilder) WithOptionalProcessIdentifier(processIdentifier BACnetContextTagUnsignedInteger) BACnetRecipientProcessBuilder {
	m.ProcessIdentifier = processIdentifier
	return m
}

func (m *_BACnetRecipientProcessBuilder) WithOptionalProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetRecipientProcessBuilder {
	builder := builderSupplier(m.ProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ProcessIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetRecipientProcessBuilder) Build() (BACnetRecipientProcess, error) {
	if m.Recipient == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'recipient' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetRecipientProcess.deepCopy(), nil
}

func (m *_BACnetRecipientProcessBuilder) MustBuild() BACnetRecipientProcess {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetRecipientProcessBuilder) DeepCopy() any {
	return m.CreateBACnetRecipientProcessBuilder()
}

// CreateBACnetRecipientProcessBuilder creates a BACnetRecipientProcessBuilder
func (m *_BACnetRecipientProcess) CreateBACnetRecipientProcessBuilder() BACnetRecipientProcessBuilder {
	if m == nil {
		return NewBACnetRecipientProcessBuilder()
	}
	return &_BACnetRecipientProcessBuilder{_BACnetRecipientProcess: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientProcess) GetRecipient() BACnetRecipientEnclosed {
	return m.Recipient
}

func (m *_BACnetRecipientProcess) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetRecipientProcess(structType any) BACnetRecipientProcess {
	if casted, ok := structType.(BACnetRecipientProcess); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientProcess); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientProcess) GetTypeName() string {
	return "BACnetRecipientProcess"
}

func (m *_BACnetRecipientProcess) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Optional Field (processIdentifier)
	if m.ProcessIdentifier != nil {
		lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetRecipientProcess) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientProcessParse(ctx context.Context, theBytes []byte) (BACnetRecipientProcess, error) {
	return BACnetRecipientProcessParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRecipientProcessParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientProcess, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientProcess, error) {
		return BACnetRecipientProcessParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetRecipientProcessParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientProcess, error) {
	v, err := (&_BACnetRecipientProcess{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetRecipientProcess) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetRecipientProcess BACnetRecipientProcess, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientProcess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientProcess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recipient, err := ReadSimpleField[BACnetRecipientEnclosed](ctx, "recipient", ReadComplex[BACnetRecipientEnclosed](BACnetRecipientEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	var processIdentifier BACnetContextTagUnsignedInteger
	_processIdentifier, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}
	if _processIdentifier != nil {
		processIdentifier = *_processIdentifier
		m.ProcessIdentifier = processIdentifier
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientProcess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientProcess")
	}

	return m, nil
}

func (m *_BACnetRecipientProcess) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientProcess) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRecipientProcess"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientProcess")
	}

	if err := WriteSimpleField[BACnetRecipientEnclosed](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipientEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "processIdentifier", GetRef(m.GetProcessIdentifier()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'processIdentifier' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientProcess"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientProcess")
	}
	return nil
}

func (m *_BACnetRecipientProcess) IsBACnetRecipientProcess() {}

func (m *_BACnetRecipientProcess) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRecipientProcess) deepCopy() *_BACnetRecipientProcess {
	if m == nil {
		return nil
	}
	_BACnetRecipientProcessCopy := &_BACnetRecipientProcess{
		m.Recipient.DeepCopy().(BACnetRecipientEnclosed),
		m.ProcessIdentifier.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	return _BACnetRecipientProcessCopy
}

func (m *_BACnetRecipientProcess) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
