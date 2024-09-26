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

// BACnetLightingCommand is the corresponding interface of BACnetLightingCommand
type BACnetLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetLightningOperation returns LightningOperation (property field)
	GetLightningOperation() BACnetLightingOperationTagged
	// GetTargetLevel returns TargetLevel (property field)
	GetTargetLevel() BACnetContextTagReal
	// GetRampRate returns RampRate (property field)
	GetRampRate() BACnetContextTagReal
	// GetStepIncrement returns StepIncrement (property field)
	GetStepIncrement() BACnetContextTagReal
	// GetFadeTime returns FadeTime (property field)
	GetFadeTime() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// IsBACnetLightingCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLightingCommand()
	// CreateBuilder creates a BACnetLightingCommandBuilder
	CreateBACnetLightingCommandBuilder() BACnetLightingCommandBuilder
}

// _BACnetLightingCommand is the data-structure of this message
type _BACnetLightingCommand struct {
	LightningOperation BACnetLightingOperationTagged
	TargetLevel        BACnetContextTagReal
	RampRate           BACnetContextTagReal
	StepIncrement      BACnetContextTagReal
	FadeTime           BACnetContextTagUnsignedInteger
	Priority           BACnetContextTagUnsignedInteger
}

var _ BACnetLightingCommand = (*_BACnetLightingCommand)(nil)

// NewBACnetLightingCommand factory function for _BACnetLightingCommand
func NewBACnetLightingCommand(lightningOperation BACnetLightingOperationTagged, targetLevel BACnetContextTagReal, rampRate BACnetContextTagReal, stepIncrement BACnetContextTagReal, fadeTime BACnetContextTagUnsignedInteger, priority BACnetContextTagUnsignedInteger) *_BACnetLightingCommand {
	if lightningOperation == nil {
		panic("lightningOperation of type BACnetLightingOperationTagged for BACnetLightingCommand must not be nil")
	}
	return &_BACnetLightingCommand{LightningOperation: lightningOperation, TargetLevel: targetLevel, RampRate: rampRate, StepIncrement: stepIncrement, FadeTime: fadeTime, Priority: priority}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLightingCommandBuilder is a builder for BACnetLightingCommand
type BACnetLightingCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lightningOperation BACnetLightingOperationTagged) BACnetLightingCommandBuilder
	// WithLightningOperation adds LightningOperation (property field)
	WithLightningOperation(BACnetLightingOperationTagged) BACnetLightingCommandBuilder
	// WithLightningOperationBuilder adds LightningOperation (property field) which is build by the builder
	WithLightningOperationBuilder(func(BACnetLightingOperationTaggedBuilder) BACnetLightingOperationTaggedBuilder) BACnetLightingCommandBuilder
	// WithTargetLevel adds TargetLevel (property field)
	WithOptionalTargetLevel(BACnetContextTagReal) BACnetLightingCommandBuilder
	// WithOptionalTargetLevelBuilder adds TargetLevel (property field) which is build by the builder
	WithOptionalTargetLevelBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLightingCommandBuilder
	// WithRampRate adds RampRate (property field)
	WithOptionalRampRate(BACnetContextTagReal) BACnetLightingCommandBuilder
	// WithOptionalRampRateBuilder adds RampRate (property field) which is build by the builder
	WithOptionalRampRateBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLightingCommandBuilder
	// WithStepIncrement adds StepIncrement (property field)
	WithOptionalStepIncrement(BACnetContextTagReal) BACnetLightingCommandBuilder
	// WithOptionalStepIncrementBuilder adds StepIncrement (property field) which is build by the builder
	WithOptionalStepIncrementBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLightingCommandBuilder
	// WithFadeTime adds FadeTime (property field)
	WithOptionalFadeTime(BACnetContextTagUnsignedInteger) BACnetLightingCommandBuilder
	// WithOptionalFadeTimeBuilder adds FadeTime (property field) which is build by the builder
	WithOptionalFadeTimeBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetLightingCommandBuilder
	// WithPriority adds Priority (property field)
	WithOptionalPriority(BACnetContextTagUnsignedInteger) BACnetLightingCommandBuilder
	// WithOptionalPriorityBuilder adds Priority (property field) which is build by the builder
	WithOptionalPriorityBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetLightingCommandBuilder
	// Build builds the BACnetLightingCommand or returns an error if something is wrong
	Build() (BACnetLightingCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLightingCommand
}

// NewBACnetLightingCommandBuilder() creates a BACnetLightingCommandBuilder
func NewBACnetLightingCommandBuilder() BACnetLightingCommandBuilder {
	return &_BACnetLightingCommandBuilder{_BACnetLightingCommand: new(_BACnetLightingCommand)}
}

type _BACnetLightingCommandBuilder struct {
	*_BACnetLightingCommand

	err *utils.MultiError
}

var _ (BACnetLightingCommandBuilder) = (*_BACnetLightingCommandBuilder)(nil)

func (m *_BACnetLightingCommandBuilder) WithMandatoryFields(lightningOperation BACnetLightingOperationTagged) BACnetLightingCommandBuilder {
	return m.WithLightningOperation(lightningOperation)
}

func (m *_BACnetLightingCommandBuilder) WithLightningOperation(lightningOperation BACnetLightingOperationTagged) BACnetLightingCommandBuilder {
	m.LightningOperation = lightningOperation
	return m
}

func (m *_BACnetLightingCommandBuilder) WithLightningOperationBuilder(builderSupplier func(BACnetLightingOperationTaggedBuilder) BACnetLightingOperationTaggedBuilder) BACnetLightingCommandBuilder {
	builder := builderSupplier(m.LightningOperation.CreateBACnetLightingOperationTaggedBuilder())
	var err error
	m.LightningOperation, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetLightingOperationTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalTargetLevel(targetLevel BACnetContextTagReal) BACnetLightingCommandBuilder {
	m.TargetLevel = targetLevel
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalTargetLevelBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLightingCommandBuilder {
	builder := builderSupplier(m.TargetLevel.CreateBACnetContextTagRealBuilder())
	var err error
	m.TargetLevel, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalRampRate(rampRate BACnetContextTagReal) BACnetLightingCommandBuilder {
	m.RampRate = rampRate
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalRampRateBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLightingCommandBuilder {
	builder := builderSupplier(m.RampRate.CreateBACnetContextTagRealBuilder())
	var err error
	m.RampRate, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalStepIncrement(stepIncrement BACnetContextTagReal) BACnetLightingCommandBuilder {
	m.StepIncrement = stepIncrement
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalStepIncrementBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetLightingCommandBuilder {
	builder := builderSupplier(m.StepIncrement.CreateBACnetContextTagRealBuilder())
	var err error
	m.StepIncrement, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalFadeTime(fadeTime BACnetContextTagUnsignedInteger) BACnetLightingCommandBuilder {
	m.FadeTime = fadeTime
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalFadeTimeBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetLightingCommandBuilder {
	builder := builderSupplier(m.FadeTime.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.FadeTime, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalPriority(priority BACnetContextTagUnsignedInteger) BACnetLightingCommandBuilder {
	m.Priority = priority
	return m
}

func (m *_BACnetLightingCommandBuilder) WithOptionalPriorityBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetLightingCommandBuilder {
	builder := builderSupplier(m.Priority.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.Priority, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetLightingCommandBuilder) Build() (BACnetLightingCommand, error) {
	if m.LightningOperation == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lightningOperation' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetLightingCommand.deepCopy(), nil
}

func (m *_BACnetLightingCommandBuilder) MustBuild() BACnetLightingCommand {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetLightingCommandBuilder) DeepCopy() any {
	return m.CreateBACnetLightingCommandBuilder()
}

// CreateBACnetLightingCommandBuilder creates a BACnetLightingCommandBuilder
func (m *_BACnetLightingCommand) CreateBACnetLightingCommandBuilder() BACnetLightingCommandBuilder {
	if m == nil {
		return NewBACnetLightingCommandBuilder()
	}
	return &_BACnetLightingCommandBuilder{_BACnetLightingCommand: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingCommand) GetLightningOperation() BACnetLightingOperationTagged {
	return m.LightningOperation
}

func (m *_BACnetLightingCommand) GetTargetLevel() BACnetContextTagReal {
	return m.TargetLevel
}

func (m *_BACnetLightingCommand) GetRampRate() BACnetContextTagReal {
	return m.RampRate
}

func (m *_BACnetLightingCommand) GetStepIncrement() BACnetContextTagReal {
	return m.StepIncrement
}

func (m *_BACnetLightingCommand) GetFadeTime() BACnetContextTagUnsignedInteger {
	return m.FadeTime
}

func (m *_BACnetLightingCommand) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLightingCommand(structType any) BACnetLightingCommand {
	if casted, ok := structType.(BACnetLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingCommand) GetTypeName() string {
	return "BACnetLightingCommand"
}

func (m *_BACnetLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lightningOperation)
	lengthInBits += m.LightningOperation.GetLengthInBits(ctx)

	// Optional Field (targetLevel)
	if m.TargetLevel != nil {
		lengthInBits += m.TargetLevel.GetLengthInBits(ctx)
	}

	// Optional Field (rampRate)
	if m.RampRate != nil {
		lengthInBits += m.RampRate.GetLengthInBits(ctx)
	}

	// Optional Field (stepIncrement)
	if m.StepIncrement != nil {
		lengthInBits += m.StepIncrement.GetLengthInBits(ctx)
	}

	// Optional Field (fadeTime)
	if m.FadeTime != nil {
		lengthInBits += m.FadeTime.GetLengthInBits(ctx)
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingCommandParse(ctx context.Context, theBytes []byte) (BACnetLightingCommand, error) {
	return BACnetLightingCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLightingCommandParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
		return BACnetLightingCommandParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLightingCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
	v, err := (&_BACnetLightingCommand{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetLightingCommand BACnetLightingCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningOperation, err := ReadSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", ReadComplex[BACnetLightingOperationTagged](BACnetLightingOperationTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningOperation' field"))
	}
	m.LightningOperation = lightningOperation

	var targetLevel BACnetContextTagReal
	_targetLevel, err := ReadOptionalField[BACnetContextTagReal](ctx, "targetLevel", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetLevel' field"))
	}
	if _targetLevel != nil {
		targetLevel = *_targetLevel
		m.TargetLevel = targetLevel
	}

	var rampRate BACnetContextTagReal
	_rampRate, err := ReadOptionalField[BACnetContextTagReal](ctx, "rampRate", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rampRate' field"))
	}
	if _rampRate != nil {
		rampRate = *_rampRate
		m.RampRate = rampRate
	}

	var stepIncrement BACnetContextTagReal
	_stepIncrement, err := ReadOptionalField[BACnetContextTagReal](ctx, "stepIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stepIncrement' field"))
	}
	if _stepIncrement != nil {
		stepIncrement = *_stepIncrement
		m.StepIncrement = stepIncrement
	}

	var fadeTime BACnetContextTagUnsignedInteger
	_fadeTime, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "fadeTime", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fadeTime' field"))
	}
	if _fadeTime != nil {
		fadeTime = *_fadeTime
		m.FadeTime = fadeTime
	}

	var priority BACnetContextTagUnsignedInteger
	_priority, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	if _priority != nil {
		priority = *_priority
		m.Priority = priority
	}

	if closeErr := readBuffer.CloseContext("BACnetLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingCommand")
	}

	return m, nil
}

func (m *_BACnetLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLightingCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingCommand")
	}

	if err := WriteSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", m.GetLightningOperation(), WriteComplex[BACnetLightingOperationTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'lightningOperation' field")
	}

	if err := WriteOptionalField[BACnetContextTagReal](ctx, "targetLevel", GetRef(m.GetTargetLevel()), WriteComplex[BACnetContextTagReal](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'targetLevel' field")
	}

	if err := WriteOptionalField[BACnetContextTagReal](ctx, "rampRate", GetRef(m.GetRampRate()), WriteComplex[BACnetContextTagReal](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'rampRate' field")
	}

	if err := WriteOptionalField[BACnetContextTagReal](ctx, "stepIncrement", GetRef(m.GetStepIncrement()), WriteComplex[BACnetContextTagReal](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'stepIncrement' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "fadeTime", GetRef(m.GetFadeTime()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'fadeTime' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", GetRef(m.GetPriority()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingCommand")
	}
	return nil
}

func (m *_BACnetLightingCommand) IsBACnetLightingCommand() {}

func (m *_BACnetLightingCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLightingCommand) deepCopy() *_BACnetLightingCommand {
	if m == nil {
		return nil
	}
	_BACnetLightingCommandCopy := &_BACnetLightingCommand{
		m.LightningOperation.DeepCopy().(BACnetLightingOperationTagged),
		m.TargetLevel.DeepCopy().(BACnetContextTagReal),
		m.RampRate.DeepCopy().(BACnetContextTagReal),
		m.StepIncrement.DeepCopy().(BACnetContextTagReal),
		m.FadeTime.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Priority.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	return _BACnetLightingCommandCopy
}

func (m *_BACnetLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
