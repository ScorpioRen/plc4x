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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesTimerState is the corresponding interface of BACnetPropertyStatesTimerState
type BACnetPropertyStatesTimerState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetTimerState returns TimerState (property field)
	GetTimerState() BACnetTimerStateTagged
	// IsBACnetPropertyStatesTimerState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesTimerState()
	// CreateBuilder creates a BACnetPropertyStatesTimerStateBuilder
	CreateBACnetPropertyStatesTimerStateBuilder() BACnetPropertyStatesTimerStateBuilder
}

// _BACnetPropertyStatesTimerState is the data-structure of this message
type _BACnetPropertyStatesTimerState struct {
	BACnetPropertyStatesContract
	TimerState BACnetTimerStateTagged
}

var _ BACnetPropertyStatesTimerState = (*_BACnetPropertyStatesTimerState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesTimerState)(nil)

// NewBACnetPropertyStatesTimerState factory function for _BACnetPropertyStatesTimerState
func NewBACnetPropertyStatesTimerState(peekedTagHeader BACnetTagHeader, timerState BACnetTimerStateTagged) *_BACnetPropertyStatesTimerState {
	if timerState == nil {
		panic("timerState of type BACnetTimerStateTagged for BACnetPropertyStatesTimerState must not be nil")
	}
	_result := &_BACnetPropertyStatesTimerState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		TimerState:                   timerState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesTimerStateBuilder is a builder for BACnetPropertyStatesTimerState
type BACnetPropertyStatesTimerStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timerState BACnetTimerStateTagged) BACnetPropertyStatesTimerStateBuilder
	// WithTimerState adds TimerState (property field)
	WithTimerState(BACnetTimerStateTagged) BACnetPropertyStatesTimerStateBuilder
	// WithTimerStateBuilder adds TimerState (property field) which is build by the builder
	WithTimerStateBuilder(func(BACnetTimerStateTaggedBuilder) BACnetTimerStateTaggedBuilder) BACnetPropertyStatesTimerStateBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesTimerState or returns an error if something is wrong
	Build() (BACnetPropertyStatesTimerState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesTimerState
}

// NewBACnetPropertyStatesTimerStateBuilder() creates a BACnetPropertyStatesTimerStateBuilder
func NewBACnetPropertyStatesTimerStateBuilder() BACnetPropertyStatesTimerStateBuilder {
	return &_BACnetPropertyStatesTimerStateBuilder{_BACnetPropertyStatesTimerState: new(_BACnetPropertyStatesTimerState)}
}

type _BACnetPropertyStatesTimerStateBuilder struct {
	*_BACnetPropertyStatesTimerState

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesTimerStateBuilder) = (*_BACnetPropertyStatesTimerStateBuilder)(nil)

func (b *_BACnetPropertyStatesTimerStateBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesTimerState
}

func (b *_BACnetPropertyStatesTimerStateBuilder) WithMandatoryFields(timerState BACnetTimerStateTagged) BACnetPropertyStatesTimerStateBuilder {
	return b.WithTimerState(timerState)
}

func (b *_BACnetPropertyStatesTimerStateBuilder) WithTimerState(timerState BACnetTimerStateTagged) BACnetPropertyStatesTimerStateBuilder {
	b.TimerState = timerState
	return b
}

func (b *_BACnetPropertyStatesTimerStateBuilder) WithTimerStateBuilder(builderSupplier func(BACnetTimerStateTaggedBuilder) BACnetTimerStateTaggedBuilder) BACnetPropertyStatesTimerStateBuilder {
	builder := builderSupplier(b.TimerState.CreateBACnetTimerStateTaggedBuilder())
	var err error
	b.TimerState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimerStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesTimerStateBuilder) Build() (BACnetPropertyStatesTimerState, error) {
	if b.TimerState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timerState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesTimerState.deepCopy(), nil
}

func (b *_BACnetPropertyStatesTimerStateBuilder) MustBuild() BACnetPropertyStatesTimerState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesTimerStateBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesTimerStateBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesTimerStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesTimerStateBuilder().(*_BACnetPropertyStatesTimerStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesTimerStateBuilder creates a BACnetPropertyStatesTimerStateBuilder
func (b *_BACnetPropertyStatesTimerState) CreateBACnetPropertyStatesTimerStateBuilder() BACnetPropertyStatesTimerStateBuilder {
	if b == nil {
		return NewBACnetPropertyStatesTimerStateBuilder()
	}
	return &_BACnetPropertyStatesTimerStateBuilder{_BACnetPropertyStatesTimerState: b.deepCopy()}
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

func (m *_BACnetPropertyStatesTimerState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesTimerState) GetTimerState() BACnetTimerStateTagged {
	return m.TimerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesTimerState(structType any) BACnetPropertyStatesTimerState {
	if casted, ok := structType.(BACnetPropertyStatesTimerState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesTimerState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesTimerState) GetTypeName() string {
	return "BACnetPropertyStatesTimerState"
}

func (m *_BACnetPropertyStatesTimerState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (timerState)
	lengthInBits += m.TimerState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesTimerState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesTimerState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesTimerState BACnetPropertyStatesTimerState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesTimerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesTimerState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timerState, err := ReadSimpleField[BACnetTimerStateTagged](ctx, "timerState", ReadComplex[BACnetTimerStateTagged](BACnetTimerStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timerState' field"))
	}
	m.TimerState = timerState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesTimerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesTimerState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesTimerState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesTimerState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesTimerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesTimerState")
		}

		if err := WriteSimpleField[BACnetTimerStateTagged](ctx, "timerState", m.GetTimerState(), WriteComplex[BACnetTimerStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timerState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesTimerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesTimerState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesTimerState) IsBACnetPropertyStatesTimerState() {}

func (m *_BACnetPropertyStatesTimerState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesTimerState) deepCopy() *_BACnetPropertyStatesTimerState {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesTimerStateCopy := &_BACnetPropertyStatesTimerState{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetTimerStateTagged](m.TimerState),
	}
	_BACnetPropertyStatesTimerStateCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesTimerStateCopy
}

func (m *_BACnetPropertyStatesTimerState) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
