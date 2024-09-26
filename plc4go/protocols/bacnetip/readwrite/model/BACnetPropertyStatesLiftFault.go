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

// BACnetPropertyStatesLiftFault is the corresponding interface of BACnetPropertyStatesLiftFault
type BACnetPropertyStatesLiftFault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLiftFault returns LiftFault (property field)
	GetLiftFault() BACnetLiftFaultTagged
	// IsBACnetPropertyStatesLiftFault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLiftFault()
	// CreateBuilder creates a BACnetPropertyStatesLiftFaultBuilder
	CreateBACnetPropertyStatesLiftFaultBuilder() BACnetPropertyStatesLiftFaultBuilder
}

// _BACnetPropertyStatesLiftFault is the data-structure of this message
type _BACnetPropertyStatesLiftFault struct {
	BACnetPropertyStatesContract
	LiftFault BACnetLiftFaultTagged
}

var _ BACnetPropertyStatesLiftFault = (*_BACnetPropertyStatesLiftFault)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLiftFault)(nil)

// NewBACnetPropertyStatesLiftFault factory function for _BACnetPropertyStatesLiftFault
func NewBACnetPropertyStatesLiftFault(peekedTagHeader BACnetTagHeader, liftFault BACnetLiftFaultTagged) *_BACnetPropertyStatesLiftFault {
	if liftFault == nil {
		panic("liftFault of type BACnetLiftFaultTagged for BACnetPropertyStatesLiftFault must not be nil")
	}
	_result := &_BACnetPropertyStatesLiftFault{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LiftFault:                    liftFault,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLiftFaultBuilder is a builder for BACnetPropertyStatesLiftFault
type BACnetPropertyStatesLiftFaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(liftFault BACnetLiftFaultTagged) BACnetPropertyStatesLiftFaultBuilder
	// WithLiftFault adds LiftFault (property field)
	WithLiftFault(BACnetLiftFaultTagged) BACnetPropertyStatesLiftFaultBuilder
	// WithLiftFaultBuilder adds LiftFault (property field) which is build by the builder
	WithLiftFaultBuilder(func(BACnetLiftFaultTaggedBuilder) BACnetLiftFaultTaggedBuilder) BACnetPropertyStatesLiftFaultBuilder
	// Build builds the BACnetPropertyStatesLiftFault or returns an error if something is wrong
	Build() (BACnetPropertyStatesLiftFault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLiftFault
}

// NewBACnetPropertyStatesLiftFaultBuilder() creates a BACnetPropertyStatesLiftFaultBuilder
func NewBACnetPropertyStatesLiftFaultBuilder() BACnetPropertyStatesLiftFaultBuilder {
	return &_BACnetPropertyStatesLiftFaultBuilder{_BACnetPropertyStatesLiftFault: new(_BACnetPropertyStatesLiftFault)}
}

type _BACnetPropertyStatesLiftFaultBuilder struct {
	*_BACnetPropertyStatesLiftFault

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLiftFaultBuilder) = (*_BACnetPropertyStatesLiftFaultBuilder)(nil)

func (m *_BACnetPropertyStatesLiftFaultBuilder) WithMandatoryFields(liftFault BACnetLiftFaultTagged) BACnetPropertyStatesLiftFaultBuilder {
	return m.WithLiftFault(liftFault)
}

func (m *_BACnetPropertyStatesLiftFaultBuilder) WithLiftFault(liftFault BACnetLiftFaultTagged) BACnetPropertyStatesLiftFaultBuilder {
	m.LiftFault = liftFault
	return m
}

func (m *_BACnetPropertyStatesLiftFaultBuilder) WithLiftFaultBuilder(builderSupplier func(BACnetLiftFaultTaggedBuilder) BACnetLiftFaultTaggedBuilder) BACnetPropertyStatesLiftFaultBuilder {
	builder := builderSupplier(m.LiftFault.CreateBACnetLiftFaultTaggedBuilder())
	var err error
	m.LiftFault, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetLiftFaultTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyStatesLiftFaultBuilder) Build() (BACnetPropertyStatesLiftFault, error) {
	if m.LiftFault == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'liftFault' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyStatesLiftFault.deepCopy(), nil
}

func (m *_BACnetPropertyStatesLiftFaultBuilder) MustBuild() BACnetPropertyStatesLiftFault {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyStatesLiftFaultBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyStatesLiftFaultBuilder()
}

// CreateBACnetPropertyStatesLiftFaultBuilder creates a BACnetPropertyStatesLiftFaultBuilder
func (m *_BACnetPropertyStatesLiftFault) CreateBACnetPropertyStatesLiftFaultBuilder() BACnetPropertyStatesLiftFaultBuilder {
	if m == nil {
		return NewBACnetPropertyStatesLiftFaultBuilder()
	}
	return &_BACnetPropertyStatesLiftFaultBuilder{_BACnetPropertyStatesLiftFault: m.deepCopy()}
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

func (m *_BACnetPropertyStatesLiftFault) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftFault) GetLiftFault() BACnetLiftFaultTagged {
	return m.LiftFault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftFault(structType any) BACnetPropertyStatesLiftFault {
	if casted, ok := structType.(BACnetPropertyStatesLiftFault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftFault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftFault) GetTypeName() string {
	return "BACnetPropertyStatesLiftFault"
}

func (m *_BACnetPropertyStatesLiftFault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (liftFault)
	lengthInBits += m.LiftFault.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftFault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLiftFault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLiftFault BACnetPropertyStatesLiftFault, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftFault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	liftFault, err := ReadSimpleField[BACnetLiftFaultTagged](ctx, "liftFault", ReadComplex[BACnetLiftFaultTagged](BACnetLiftFaultTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'liftFault' field"))
	}
	m.LiftFault = liftFault

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftFault")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLiftFault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftFault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftFault")
		}

		if err := WriteSimpleField[BACnetLiftFaultTagged](ctx, "liftFault", m.GetLiftFault(), WriteComplex[BACnetLiftFaultTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'liftFault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftFault")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftFault) IsBACnetPropertyStatesLiftFault() {}

func (m *_BACnetPropertyStatesLiftFault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLiftFault) deepCopy() *_BACnetPropertyStatesLiftFault {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLiftFaultCopy := &_BACnetPropertyStatesLiftFault{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.LiftFault.DeepCopy().(BACnetLiftFaultTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLiftFaultCopy
}

func (m *_BACnetPropertyStatesLiftFault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
