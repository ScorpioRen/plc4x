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

// BACnetConstructedDataEventState is the corresponding interface of BACnetConstructedDataEventState
type BACnetConstructedDataEventState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventStateTagged
	// IsBACnetConstructedDataEventState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventState()
	// CreateBuilder creates a BACnetConstructedDataEventStateBuilder
	CreateBACnetConstructedDataEventStateBuilder() BACnetConstructedDataEventStateBuilder
}

// _BACnetConstructedDataEventState is the data-structure of this message
type _BACnetConstructedDataEventState struct {
	BACnetConstructedDataContract
	EventState BACnetEventStateTagged
}

var _ BACnetConstructedDataEventState = (*_BACnetConstructedDataEventState)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventState)(nil)

// NewBACnetConstructedDataEventState factory function for _BACnetConstructedDataEventState
func NewBACnetConstructedDataEventState(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, eventState BACnetEventStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventState {
	if eventState == nil {
		panic("eventState of type BACnetEventStateTagged for BACnetConstructedDataEventState must not be nil")
	}
	_result := &_BACnetConstructedDataEventState{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventState:                    eventState,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventStateBuilder is a builder for BACnetConstructedDataEventState
type BACnetConstructedDataEventStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventState BACnetEventStateTagged) BACnetConstructedDataEventStateBuilder
	// WithEventState adds EventState (property field)
	WithEventState(BACnetEventStateTagged) BACnetConstructedDataEventStateBuilder
	// WithEventStateBuilder adds EventState (property field) which is build by the builder
	WithEventStateBuilder(func(BACnetEventStateTaggedBuilder) BACnetEventStateTaggedBuilder) BACnetConstructedDataEventStateBuilder
	// Build builds the BACnetConstructedDataEventState or returns an error if something is wrong
	Build() (BACnetConstructedDataEventState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventState
}

// NewBACnetConstructedDataEventStateBuilder() creates a BACnetConstructedDataEventStateBuilder
func NewBACnetConstructedDataEventStateBuilder() BACnetConstructedDataEventStateBuilder {
	return &_BACnetConstructedDataEventStateBuilder{_BACnetConstructedDataEventState: new(_BACnetConstructedDataEventState)}
}

type _BACnetConstructedDataEventStateBuilder struct {
	*_BACnetConstructedDataEventState

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventStateBuilder) = (*_BACnetConstructedDataEventStateBuilder)(nil)

func (m *_BACnetConstructedDataEventStateBuilder) WithMandatoryFields(eventState BACnetEventStateTagged) BACnetConstructedDataEventStateBuilder {
	return m.WithEventState(eventState)
}

func (m *_BACnetConstructedDataEventStateBuilder) WithEventState(eventState BACnetEventStateTagged) BACnetConstructedDataEventStateBuilder {
	m.EventState = eventState
	return m
}

func (m *_BACnetConstructedDataEventStateBuilder) WithEventStateBuilder(builderSupplier func(BACnetEventStateTaggedBuilder) BACnetEventStateTaggedBuilder) BACnetConstructedDataEventStateBuilder {
	builder := builderSupplier(m.EventState.CreateBACnetEventStateTaggedBuilder())
	var err error
	m.EventState, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEventStateTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataEventStateBuilder) Build() (BACnetConstructedDataEventState, error) {
	if m.EventState == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'eventState' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataEventState.deepCopy(), nil
}

func (m *_BACnetConstructedDataEventStateBuilder) MustBuild() BACnetConstructedDataEventState {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataEventStateBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataEventStateBuilder()
}

// CreateBACnetConstructedDataEventStateBuilder creates a BACnetConstructedDataEventStateBuilder
func (m *_BACnetConstructedDataEventState) CreateBACnetConstructedDataEventStateBuilder() BACnetConstructedDataEventStateBuilder {
	if m == nil {
		return NewBACnetConstructedDataEventStateBuilder()
	}
	return &_BACnetConstructedDataEventStateBuilder{_BACnetConstructedDataEventState: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventState) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventState) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventState) GetActualValue() BACnetEventStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventStateTagged(m.GetEventState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventState(structType any) BACnetConstructedDataEventState {
	if casted, ok := structType.(BACnetConstructedDataEventState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventState) GetTypeName() string {
	return "BACnetConstructedDataEventState"
}

func (m *_BACnetConstructedDataEventState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventState BACnetConstructedDataEventState, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	actualValue, err := ReadVirtualField[BACnetEventStateTagged](ctx, "actualValue", (*BACnetEventStateTagged)(nil), eventState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventState")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventState")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "eventState", m.GetEventState(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventState")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventState) IsBACnetConstructedDataEventState() {}

func (m *_BACnetConstructedDataEventState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventState) deepCopy() *_BACnetConstructedDataEventState {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventStateCopy := &_BACnetConstructedDataEventState{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.EventState.DeepCopy().(BACnetEventStateTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventStateCopy
}

func (m *_BACnetConstructedDataEventState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
